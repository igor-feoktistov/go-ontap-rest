package ontap

import (
	"fmt"
	"strings"
	"net/http"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"path"
)

type FileInfo struct {
	Resource
	AccessedTime string `json:"accessed_time,omitempty"`
	Analytics *struct {
		ByAccessedTime *struct {
			BytesUsed *struct {
				Labels []string       `json:"labels,omitempty"`
				NewestLabel []string  `json:"newest_label,omitempty"`
				OldestLabel []string  `json:"oldest_label,omitempty"`
				Percentages []float32 `json:"percentages,omitempty"`
				Values []int          `json:"values,omitempty"`
			}                             `json:"bytes_used ,omitempty"`
		}                                     `json:"by_accessed_time,omitempty"`
		ByModifiedTime *struct {
			BytesUsed *struct {
				Labels []string       `json:"labels,omitempty"`
				NewestLabel []string  `json:"newest_label,omitempty"`
				OldestLabel []string  `json:"oldest_label,omitempty"`
				Percentages []float32 `json:"percentages,omitempty"`
				Values []int          `json:"values,omitempty"`
			}                             `json:"bytes_used ,omitempty"`
		}                                     `json:"by_modified_time,omitempty"`
		BytesUsed int                         `json:"bytes_used"`
		FileCount int                         `json:"file_count"`
		SubdirCount int                       `json:"subdir_count"`
	}                                             `json:"analytics,omitempty"`
	BytesUsed *int                                `json:"bytes_used,omitempty"`
	ChangedTime string                            `json:"changed_time,omitempty"`
	CreationTime string                           `json:"creation_time,omitempty"`
	FillEnabled *bool                             `json:"fill_enabled,omitempty"`
	GroupId *int                                  `json:"group_id,omitempty"`
	HardLinksCount *int                           `json:"hard_links_count,omitempty"`
	InodeGeneration *int                          `json:"inode_generation,omitempty"`
	InodeNumber *int                              `json:"inode_number,omitempty"`
	IsEmpty *bool                                 `json:"is_empty,omitempty"`
	IsJunction *bool                              `json:"is_junction,omitempty"`
	IsSnapshot *bool                              `json:"is_snapshot,omitempty"`
	IsVmAligned *bool                             `json:"is_vm_aligned,omitempty"`
	ModifiedTime string                           `json:"modified_time,omitempty"`
	OverwriteEnabled *bool                        `json:"overwrite_enabled,omitempty"`
	OwnerId *int                                  `json:"owner_id,omitempty"`
	Path string                                   `json:"path,omitempty"`
	QosPolicy *Resource                           `json:"qos_policy,omitempty"`
	Size *int                                     `json:"size,omitempty"`
	Target string                                 `json:"target,omitempty"`
	Type string                                   `json:"type,omitempty"`
	UniqueBytes *int                              `json:"unique_bytes,omitempty"`
	UnixPermissions *int                          `json:"unix_permissions,omitempty"`
	Volume *Resource                              `json:"volume,omitempty"`
}

type FileInfoResponse struct {
	BaseResponse
	Files []FileInfo `json:"records,omitempty"`
}

type FileWriteResponse struct {
	BytesWritten int `json:"bytes_written"`
}

func (c *Client) FileGetIter(volumeUuid string, path string, parameters []string) (files []FileInfo, res *RestResponse, err error) {
	var req *http.Request
	getPath := fmt.Sprintf("/api/storage/volumes/%s/files/%s", volumeUuid, path)
	reqParameters := parameters
	for {
		r := FileInfoResponse{}
		req, err = c.NewRequest("GET", getPath, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, file := range r.Files {
			files = append(files, file)
		}
		if r.IsPaginate() {
			getPath = r.GetNextRef()
			reqParameters = []string{}
		} else {
			break
		}
	}
	return
}

func (c *Client) FileCreate(volumeUuid string, filePath string, fileInfo *FileInfo) (fileResponse *FileInfoResponse, res *RestResponse, err error) {
	var req *http.Request
	href := fmt.Sprintf("/api/storage/volumes/%s/files/%s", volumeUuid, strings.ReplaceAll(strings.ReplaceAll(filePath, "/", "%2F"), ".", "%2E"))
	if req, err = c.NewRequest("POST", href, []string{}, fileInfo); err != nil {
		return
	}
	fileResponse = &FileInfoResponse{}
	res, err = c.Do(req, fileResponse)
	return
}

func (c *Client) FileDelete(volumeUuid string, filePath string, parameters []string) (res *RestResponse, err error) {
	var req *http.Request
	jobLink := JobLinkResponse{}
	var job *Job
	href := fmt.Sprintf("/api/storage/volumes/%s/files/%s", volumeUuid, strings.ReplaceAll(strings.ReplaceAll(filePath, "/", "%2F"), ".", "%2E"))
	if req, err = c.NewRequest("DELETE", href, parameters, nil); err != nil {
		return
	}
	if res, err = c.Do(req, &jobLink); err != nil {
		return
	}
	if job, err = c.JobWaitUntilComplete(jobLink.JobLink.GetRef()); err == nil && job != nil {
		if job.State == "failure" {
			err = fmt.Errorf("Error: REST code=%d, REST message=\"%s\"", job.Code, job.Message)
		}
	}
	return
}

func (c *Client) FileWrite(method string, volumeUuid string, filePath string, parameters []string, body []byte) (bytesWritten int, res *RestResponse, err error) {
	var req *http.Request
	href := fmt.Sprintf("/api/storage/volumes/%s/files/%s", volumeUuid, strings.ReplaceAll(strings.ReplaceAll(filePath, "/", "%2F"), ".", "%2E"))
	if req, err = c.NewFormFileRequest(method, href, parameters, body); err != nil {
		return
	}
	r := FileWriteResponse{}
	if res, err = c.Do(req, &r); err == nil {
		bytesWritten = r.BytesWritten
	}		
	return
}

func (c *Client) FileRead(volumeUuid string, filePath string, contentOffset int, contentLength int) (content []byte, res *RestResponse, err error) {
	var req *http.Request
	href := fmt.Sprintf("/api/storage/volumes/%s/files/%s", volumeUuid, strings.ReplaceAll(strings.ReplaceAll(filePath, "/", "%2F"), ".", "%2E"))
	parameters := []string{fmt.Sprintf("byte_offset=%d", contentOffset), fmt.Sprintf("length=%d", contentLength)}
	if req, err = c.NewRequest("GET", href, parameters, nil); err != nil {
		return
	}
	req.Header.Set("Accept", "multipart/form-data")
	if res, err = c.Do(req, nil); err != nil {
		return
	}
	var headerParameters map[string]string
	if _, headerParameters, err = mime.ParseMediaType(res.HttpResponse.Header.Get("Content-Type")); err != nil {
		return
	}
	var mpartReader *multipart.Reader
	if boundary, ok := headerParameters["boundary"]; ok {
		mpartReader = multipart.NewReader(res.HttpResponse.Body, boundary)
	} else {
		err = fmt.Errorf("FileRead(): expected response in Mime format")
		return
	}
	fileName := path.Base(strings.ReplaceAll(filePath, "%2F", "/"))
	for {
		p, readErr := mpartReader.NextPart()
		if readErr == io.EOF {
			err = fmt.Errorf("FileRead(): no file \"%s\" found in response", fileName)
			break
		}
		if p.FileName() == fileName {
			content, err = ioutil.ReadAll(p)
			return
		}
	}
	return
}
