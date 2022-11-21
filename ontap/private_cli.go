package ontap

import (
	"net/http"
)

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
	LunPath string         `json:"path"`
	FilePath string        `json:"file-path"`
	OsType string          `json:"ostype"`
	SpaceReserve string    `json:"space-reserve,omitempty"`
	SpaceAllocation string `json:"space-allocation,omitempty"`
}

func (c *Client) PrivateCliVolumeGetNode(volumeName string) (node string, res *RestResponse, err error) {
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

func (c *Client) PrivateCliLunCreateFromFile(lunRequest *LunCreateFromFileRequest) (res *RestResponse, err error) {
	var req *http.Request
	path := "/api/private/cli/lun"
	if req, err = c.NewRequest("POST", path, []string{}, lunRequest); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}
