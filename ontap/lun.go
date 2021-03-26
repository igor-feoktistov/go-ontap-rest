package ontap

import (
	"fmt"
	"net/http"
)

type LunClone struct {
	Source NameReference `json:"source"`
}

type QtreeResource struct {
	Resource
	Id int `json:"id"`
}

type LunLocation struct {
	LogicalUnit string   `json:"logical_unit,omitempty"`
	Qtree *QtreeResource `json:"qtree,omitempty"`
	Volume *Resource     `json:"volume,omitempty"`
}

type LunSpaceGuarantee struct {
	Requested bool `json:"requested"`
	Reserved bool  `json:"reserved"`
}

type LunSpace struct {
	Guarantee *LunSpaceGuarantee `json:"guarantee,omitempty"`
	Size *int `json:"size,omitempty"`
	Used *int `json:"used,omitempty"`
}

type Lun struct {
	Resource
	AutoDelete *bool                            `json:"auto_delete,omitempty"`
	Class string                                `json:"class,omitempty"`
	Clone *LunClone                             `json:"clone,omitempty"`
	Comment string                              `json:"comment,omitempty"`
	CreateTime string                           `json:"create_time,omitempty"`
	Enabled *bool                               `json:"enabled,omitempty"`
	Location *LunLocation                       `json:"location,omitempty"`
	Metric *struct {
		Resource
		Duration string                     `json:"duration,omitempty"`
		Iops *struct {
			Other int                   `json:"other"`
			Read int                    `json:"read"`
			Total int                   `json:"total"`
			Write int                   `json:"write"`
		}                                   `json:"iops,omitempty"`
		Latency *struct {
			Other int                   `json:"other"`
			Read int                    `json:"read"`
			Total int                   `json:"total"`
			Write int                   `json:"write"`
		}                                   `json:"latency,omitempty"`
		Status string                       `json:"status,omitempty"`
		Throughput *struct {
			Other int                   `json:"other"`
			Read int                    `json:"read"`
			Total int                   `json:"total"`
			Write int                   `json:"write"`
		}                                   `json:"throughput,omitempty"`
		Timestamp string                    `json:"timestamp,omitempty"`
	}                                           `json:"metric,omitempty"`
	Movement *struct {
		MaxThroughput string                `json:"max_throughput,omitempty"`
		Paths struct {
			Destination string          `json:"destination,omitempty"`
			Source string               `json:"source,omitempty"`
		}                                   `json:"paths,omitempty"`
		Progress *struct {
			Elapsed int                 `json:"elapsed"`
			Failure *struct {
				Code string         `json:"code,omitempty"`
				Message string      `json:"message,omitempty"`
			}                           `json:"failure,omitempty"`
			PercentComplete int         `json:"percent_complete"`
			State string                `json:"state,omitempty"`
			VolumeSnapshotBlocked bool  `json:"volume_snapshot_blocked"`
		}                                   `json:"progress,omitempty"`
	}                                           `json:"movement,omitempty"`
	Name string                                 `json:"name,omitempty"`
	OsType string                               `json:"os_type,omitempty"`
	QosPolicy *Resource                         `json:"qos_policy,omitempty"`
	SerialNumber string                         `json:"serial_number,omitempty"`
	Space *LunSpace                             `json:"space,omitempty"`
	Statistics *struct {
		IopsRaw *struct {
			Other int                   `json:"other"`
			Read int                    `json:"read"`
			Total int                   `json:"total"`
			Write int                   `json:"write"`
		}                                   `json:"iops_raw,omitempty"`
		LatencyRaw *struct {
			Other int                   `json:"other"`
			Read int                    `json:"read"`
			Total int                   `json:"total"`
			Write int                   `json:"write"`
		}                                   `json:"latency_raw,omitempty"`
		Status string                       `json:"status,omitempty"`
		ThroughputRaw *struct {
			Other int                   `json:"other"`
			Read int                    `json:"read"`
			Total int                   `json:"total"`
			Write int                   `json:"write"`
		}                                   `json:"throughput_raw,omitempty"`
		Timestamp string                    `json:"timestamp,omitempty"`
	}                                           `json:"statistics,omitempty"`
	Status *struct {
		ContainerState string               `json:"container_state,omitempty"`
		Mapped bool                         `json:"mapped"`
		ReadOnly bool                       `json:"read_only"`
		State string                        `json:"state,omitempty"`
	}                                           `json:"status,omitempty"`
	Svm *Resource                               `json:"svm,omitempty"`
}

type LunResponse struct {
	BaseResponse
	Luns []Lun `json:"records,omitempty"`
}

type LunRef struct {
	Resource
	Node *Resource `json:"node,omitempty"`
}

type LunMapResponse struct {
	BaseResponse
	LunMaps []LunMap `json:"records,omitempty"`
}

type LunMap struct {
	Resource
	Igroup *Igroup         `json:"igroup,omitempty"`
	LogicalUnitNumber *int `json:"logical_unit_number,omitempty"`
	Lun *LunRef            `json:"lun,omitempty"`
	Svm *Resource          `json:"svm,omitempty"`
}

func (c *Client) LunGetIter(parameters []string) (luns []Lun, res *http.Response, err error) {
	var req *http.Request
	path := "/api/storage/luns"
	reqParameters := parameters
	for {
		r := LunResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, lun := range r.Luns {
			luns = append(luns, lun)
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

func (c *Client) LunGet(href string, parameters []string) (*Lun, *http.Response, error) {
	r := Lun{}
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

func (c *Client) LunCreate(lun *Lun, parameters []string) (res *http.Response, err error) {
	var req *http.Request
	path := "/api/storage/luns"
	if req, err = c.NewRequest("POST", path, parameters, lun); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunModify(href string, lun *Lun) (res *http.Response, err error) {
	var req *http.Request
	if req, err = c.NewRequest("PATCH", href, []string{}, lun); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunDelete(href string) (res *http.Response, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunMapGetIter(parameters []string) (lunMaps []LunMap, res *http.Response, err error) {
	var req *http.Request
	path := "/api/protocols/san/lun-maps"
	reqParameters := parameters
	for {
		r := LunMapResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, lunMap := range r.LunMaps {
			lunMaps = append(lunMaps, lunMap)
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

func (c *Client) LunMapGet(href string, parameters []string) (*LunMap, *http.Response, error) {
	r := LunMap{}
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

func (c *Client) LunMapCreate(lunMap *LunMap, parameters []string) (res *http.Response, err error) {
	var req *http.Request
	path := "/api/protocols/san/lun-maps"
	if req, err = c.NewRequest("POST", path, parameters, lunMap); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunMapDelete(lunUuid string, igroupUuid string) (res *http.Response, err error) {
	var req *http.Request
	path := fmt.Sprintf("/api/protocols/san/lun-maps/%s/%s", lunUuid, igroupUuid)
	if req, err = c.NewRequest("DELETE", path, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}
