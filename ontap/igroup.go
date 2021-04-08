package ontap

import (
	"net/http"
)

type IgroupInitiator struct {
	Resource
	Igroup *Resource `json:"igroup,omitempty"`
	IgroupInitiators *[]Resource `json:"records,omitempty"`
}

type IgroupInitiatorResponse struct {
	BaseResponse
	IgroupInitiators []IgroupInitiator `json:"records,omitempty"`
}
	
type Igroup struct {
	Resource
	DeleteOnUnmap *bool          `json:"delete_on_unmap,omitempty"`
	Initiators []IgroupInitiator `json:"initiators,omitempty"`
	LunMaps []LunMap             `json:"lun_maps,omitempty"`
	OsType string                `json:"os_type,omitempty"`
	Protocol string              `json:"protocol,omitempty"`
	Svm *Resource                `json:"svm,omitempty"`
}

type IgroupResponse struct {
	BaseResponse
	Igroups []Igroup `json:"records,omitempty"`
}

func (c *Client) IgroupGetIter(parameters []string) (igroups []Igroup, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/san/igroups"
	reqParameters := parameters
	for {
		r := IgroupResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, igroup := range r.Igroups {
			igroups = append(igroups, igroup)
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

func (c *Client) IgroupGet(href string, parameters []string) (*Igroup, *RestResponse, error) {
	r := Igroup{}
	req, err := c.NewRequest("GET", href, parameters, nil)
	if err != nil {
		return nil, nil, err
	}
	res, err := c.Do(req, &r)
	if err != nil {
		return nil, nil, err
	}
	return &r, res, nil
}

func (c *Client) IgroupCreate(igroup *Igroup, parameters []string) (res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/san/igroups"
	if req, err = c.NewRequest("POST", path, parameters, igroup); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) IgroupModify(href string, igroup *Igroup) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("PATCH", href, []string{}, igroup); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) IgroupDelete(href string) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) IgroupInitiatorGetIter(href string, parameters []string) (initiators []IgroupInitiator, res *RestResponse, err error) {
	var req *http.Request
	path := href + "/initiators"
	reqParameters := parameters
	for {
		r := IgroupInitiatorResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, initiator := range r.IgroupInitiators {
			initiators = append(initiators, initiator)
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

func (c *Client) IgroupInitiatorCreate(href string, initiator *IgroupInitiator) (res *RestResponse, err error) {
	var req *http.Request
	path := href + "/initiators"
	if req, err = c.NewRequest("POST", path, []string{}, initiator); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) IgroupInitiatorDelete(href string) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}
