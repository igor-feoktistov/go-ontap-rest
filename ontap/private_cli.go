package ontap

import (
	"net/http"
)

type PrivateCliAggregate struct {
	Name string          `json:"aggregate,omitempty"`
	State string         `json:"aggrstate,omitempty"`
	AvailableSize *int   `json:"availsize,omitempty"`
	Vserver string       `json:"vserver,omitempty"`
}

type PrivateCliAggregatesResponse struct {
	BaseResponse
	Aggregates []PrivateCliAggregate `json:"records,omitempty"`
}

type PrivateCliVolumeNode struct {
	Vserver string `json:"vserver,omitempty"`
	Volume string  `json:"volume,omitempty"`
	Node string    `json:"node,omitempty"`
}

type PrivateCliVolumeNodeResponse struct {
	BaseResponse
	Volumes []PrivateCliVolumeNode `json:"records,omitempty"`
}

type LunCreateFromFileRequest struct {
	LunPath string  `json:"path"`
	FilePath string `json:"file-path"`
	OsType string   `json:"ostype"`
}

type LunCopyStartRequest struct {
	LunSrcPath string `json:"source-path"`
	LunDstPath string `json:"destination-path"`
}

func (c *Client) PrivateCliAggregateGetIter(parameters []string) (aggregates []PrivateCliAggregate, res *http.Response, err error) {
	var req *http.Request
	path := "/api/private/cli/vserver/aggregates"
	reqParameters := []string{"fields=aggrstate,availsize"}
	for _, parameter := range parameters {
		reqParameters = append(reqParameters, parameter)
	}
	for {
		r := PrivateCliAggregatesResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, aggregate := range r.Aggregates {
			aggregates = append(aggregates, aggregate)
		}
		if r.IsPaginate() {
			path = r.GetNextRef()
			reqParameters = []string{}
		} else {
			break
		}
	}
	return
}

func (c *Client) PrivateCliVolumeGetNode(volumeName string) (node string, res *http.Response, err error) {
	var req *http.Request
	path := "/api/private/cli/volume"
	parameters := []string{"volume=" + volumeName, "fields=node"}
	r := PrivateCliVolumeNodeResponse{}
	if req, err = c.NewRequest("GET", path, parameters, nil); err != nil {
		return
	}
	if res, err = c.Do(req, &r); err == nil {
		node = r.Volumes[0].Node
	}
	return
}

func (c *Client) PrivateCliLunCreateFromFile(lunRequest *LunCreateFromFileRequest) (res *http.Response, err error) {
	var req *http.Request
	path := "/api/private/cli/lun"
	if req, err = c.NewRequest("POST", path, []string{}, lunRequest); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) PrivateCliLunCopyStart(lunRequest *LunCopyStartRequest) (res *http.Response, err error) {
	var req *http.Request
	path := "/api/private/cli/lun/copy/start"
	if req, err = c.NewRequest("POST", path, []string{}, lunRequest); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}
