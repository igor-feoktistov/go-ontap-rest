package util

import (
	"io"
	"fmt"
	"path/filepath"
	"net"
	"net/http"

	"go-ontap-rest/ontap"
	"github.com/vmware/go-nfs-client/nfs"
	"github.com/vmware/go-nfs-client/nfs/rpc"
)

const (
    HTTP_MAX_WRITE_SIZE = 1024*1024
)

func UploadFileAPI(c *ontap.Client, volumeName string, filePath string, r io.Reader) (bytesUploaded int64, err error) {
	var volume *ontap.Volume
	if volume, err = createDirPath(c, volumeName, filePath); err != nil {
		return
	}
	inb := make([]byte, HTTP_MAX_WRITE_SIZE)
	for {
		var n int
		var read_err, write_err error
		n, read_err = r.Read(inb)
    		if n > 0 {
    			bytesWritten := 0
    			if bytesUploaded > 0 {
				bytesWritten, _, write_err = c.FileWrite("PATCH", volume.Uuid, filePath, []string{}, inb[:n])
    			} else {
				bytesWritten, _, write_err = c.FileWrite("POST", volume.Uuid, filePath, []string{}, inb[:n])
    			}
			if write_err != nil {
				err = write_err
				break
			}
			if bytesWritten != n {
    				err = fmt.Errorf("UploadFileAPI(): short write: requested %d, written %d", n, bytesWritten)
    				break
			}
			bytesUploaded += int64(n)
    		}
    		if read_err == io.EOF {
        		break
    		}
    		if read_err != nil {
    			err = read_err
        		break
    		}
	}
	return
}

func UploadFileNFS(c *ontap.Client, volumeName string, filePath string, r io.Reader) (bytesUploaded int64, err error) {
	var clientIP net.IP
	if clientIP, err = GetOutboundIP(); err != nil {
		return
	}
	var volume *ontap.Volume
	if volume, err = createDirPath(c, volumeName, filePath); err != nil {
		return
	}
	var lifs []ontap.IpInterface
	if lifs, err = DiscoverNfsLIFs(c, volumeName); err != nil {
		return
	}
	if err = createExportPolicyRule(c, volume.Nas.ExportPolicy.Name, clientIP.String()); err != nil {
		return
	}
	defer deleteExportPolicyRule(c, volume.Nas.ExportPolicy.Name, clientIP.String())
	mount, err := nfs.DialMount(lifs[0].Ip.Address)
	if err != nil {
		return
	}
	defer mount.Close()
	auth := rpc.NewAuthUnix("root", 0, 0)
	v, err := mount.Mount(volume.Nas.Path, auth.Auth())
	if err != nil {
		return
	}
	defer v.Close()
	w, err := v.OpenFile(filePath, 0644)
	if err != nil {
		return
	}
	defer w.Close()
	bytesUploaded, err = io.Copy(w, r)
	return
}

func DownloadFileAPI(c *ontap.Client, volumeName string, filePath string) (content []byte, err error) {
	var volumes []ontap.Volume
	if volumes, _, err = c.VolumeGetIter([]string{"name=" + volumeName}); err != nil {
		return
	}
	if len(volumes) == 0 {
		err = fmt.Errorf("DownloadFileAPI(): no volume \"%s\" found", volumeName)
		return
	}
	var files []ontap.FileInfo
	dirPath := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)
	if files, _, err = c.FileGetIter(volumes[0].Uuid, dirPath, []string{"type=file","name=" + fileName,"fields=size"}); err != nil {
		return
	}
	if len(files) == 0 {
		err = fmt.Errorf("DownloadFileAPI(): no file \"%s\" found", fileName)
		return
	}
	content, _, err = c.FileRead(volumes[0].Uuid, filePath, 0, *files[0].Size)
	return
}

func createDirPath(c *ontap.Client, volumeName string, filePath string) (volume *ontap.Volume, err error) {
	var response *http.Response
	var dirList []string
	var volumes []ontap.Volume
        if volumes, _, err = c.VolumeGetIter([]string{"name=" + volumeName,"fields=nas"}); err != nil {
    		return
    	}
    	if len(volumes) == 0 {
    		err = fmt.Errorf("createDirPath(): no volume \"%s\" found", volumeName)
    		return
    	}
    	volume = &volumes[0]
	for dir := filepath.Dir(filePath); dir != "/" && dir != "."; dir = filepath.Dir(dir) {
		dirList = append(dirList, dir)
	}
	for i := len(dirList) - 1; i >= 0; i-- {
		if _, response, err = c.FileGetIter(volume.Uuid, dirList[i], []string{"type=directory","return_metadata=true"}); err != nil {
    			if response.StatusCode == 404 {
    				unixPermissions := 755
            			fileInfo := ontap.FileInfo{
					Type: "directory",
					UnixPermissions: &unixPermissions,
				}
				if _, _, err = c.FileCreate(volumes[0].Uuid, dirList[i], &fileInfo); err != nil {
					break
				}
    			} else {
    				break
    			}
    		}
    	}
    	return
}

func createExportPolicyRule(c *ontap.Client, policyName string, clientIP string) (err error) {
	var exportPolicies []ontap.ExportPolicy
	if exportPolicies, _, err = c.ExportPolicyGetIter([]string{"name=" + policyName}); err != nil {
		return
	}
	if len(exportPolicies) > 0 {
		rule := ontap.ExportPolicyRule{
			AnonymousUser: "root",
			Protocols: []string{"any"},
			RoRule: []string{"any"},
			RwRule: []string{"any"},
			Superuser: []string{"any"},
			Clients: []ontap.ExportRuleClient{
				ontap.ExportRuleClient{
					Match: clientIP,
				},
			},
		}
		_, err = c.ExportPolicyRuleCreate(exportPolicies[0].GetRef(), &rule)
	} else {
    		err = fmt.Errorf("createExportPolicyRule(): no export-policy \"%s\" found", policyName)
	}
	return
}

func deleteExportPolicyRule(c *ontap.Client, policyName string, clientIP string) (err error) {
	var exportPolicies []ontap.ExportPolicy
	if exportPolicies, _, err = c.ExportPolicyGetIter([]string{"name=" + policyName}); err != nil {
		return
	}
	if len(exportPolicies) > 0 {
		var rules []ontap.ExportPolicyRule
		if rules, _, err = c.ExportPolicyRuleGetIter(exportPolicies[0].GetRef(), []string{"clients.match=" + clientIP}); err != nil {
			return
		}
		if len(rules) > 0 {
			_, err = c.ExportPolicyRuleDelete(rules[0].GetRef())
		} else {
    			err = fmt.Errorf("deleteExportPolicyRule(): no export-policy rule found")
		}
	} else {
    		err = fmt.Errorf("createExportPolicyRule(): no export-policy \"%s\" found", policyName)
	}
	return
}
