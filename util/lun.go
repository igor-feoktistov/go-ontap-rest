package util

import (
	"fmt"
	"time"

	"go-ontap-rest/ontap"
)

const (
	MAX_WAIT_FOR_LUN = 300
)

func LunCreateFromFile(c *ontap.Client, lunPath string, filePath string, osType string) (err error) {
	lunRequest := ontap.LunCreateFromFileRequest{
		LunPath: lunPath,
		FilePath: filePath,
		OsType: osType,
	}
	if _, err = c.PrivateCliLunCreateFromFile(&lunRequest); err != nil {
		return
	}
	giveupTime := time.Now().Add(time.Second * MAX_WAIT_FOR_LUN)
	for time.Now().Before(giveupTime) {
		var luns []ontap.Lun
		if luns, _, err = c.LunGetIter([]string{"name=" + lunPath,"fields=status"}); err != nil {
			break
		}
		if len(luns) > 0 && luns[0].Status.State == "online" {
			return
		}
		time.Sleep(time.Second)
	}
	if err == nil {
		err = fmt.Errorf("LunCreateFromFile(): LUN is not available, maximum wait time exceeded")
	}
	return
}

func LunCopy(c *ontap.Client, lunSrcPath string, lunDstPath string) (err error) {
	lunRequest := ontap.LunCopyStartRequest{
		LunSrcPath: lunSrcPath,
		LunDstPath: lunDstPath,
	}
	if _, err = c.PrivateCliLunCopyStart(&lunRequest); err != nil {
		return
	}
	giveupTime := time.Now().Add(time.Second * MAX_WAIT_FOR_LUN)
	for time.Now().Before(giveupTime) {
		var luns []ontap.Lun
		if luns, _, err = c.LunGetIter([]string{"name=" + lunDstPath,"fields=status"}); err != nil {
			break
		}
		if len(luns) > 0 && luns[0].Status.State == "online" {
			return
		}
		time.Sleep(time.Second)
	}
	if err == nil {
		err = fmt.Errorf("LunCopy(): destination LUN is not available, maximum wait time exceeded")
	}
	return
}
