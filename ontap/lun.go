package ontap

import (
	"fmt"
	"net/http"
	"io"
        "mime"
        "mime/multipart"
)

type QtreeResource struct {
	Resource
	Id int `json:"id"`
}

type LunPaths struct {
        Destination string `json:"destination,omitempty"`
        Source string      `json:"source,omitempty"`
}

type LunClone struct {
	Source NameReference `json:"source"`
}

type LunCopy struct {
	Source NameReference `json:"source"`
}

type LunMovement struct {
        MaxThroughput string `json:"max_throughput,omitempty"`
        Paths LunPaths       `json:"paths,omitempty"`
}

type LunLocation struct {
	LogicalUnit string   `json:"logical_unit,omitempty"`
	Node *Resource       `json:"node,omitempty"`
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
	Copy *LunCopy                               `json:"copy,omitempty"`
	Movement *LunMovement                       `json:"movement,omitempty"`
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

type IgroupRef struct {
	Resource
	Initiators []string `json:"initiators,omitempty"`
	OsType string       `json:"os_type,omitempty"`
	Protocol string     `json:"protocol,omitempty"`
}

type LunMapResponse struct {
	BaseResponse
	LunMaps []LunMap `json:"records,omitempty"`
}

type LunMap struct {
	Resource
	Igroup *IgroupRef      `json:"igroup,omitempty"`
	LogicalUnitNumber *int `json:"logical_unit_number,omitempty"`
	Lun *LunRef            `json:"lun,omitempty"`
	Svm *Resource          `json:"svm,omitempty"`
}

func (c *Client) LunGetIter(parameters []string) (luns []Lun, res *RestResponse, err error) {
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

func (c *Client) LunGet(href string, parameters []string) (*Lun, *RestResponse, error) {
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

func (c *Client) LunGetByPath(lunPath string, parameters []string) (lun *Lun, res *RestResponse, err error) {
        var luns []Lun
	var req *http.Request
	if luns, _, err = c.LunGetIter([]string{"name=" + lunPath}); err != nil {
	        return
	}
	lun = &Lun{}
	if req, err = c.NewRequest("GET", luns[0].GetRef(), parameters, nil); err != nil {
		return
	}
	if res, err = c.Do(req, lun); err != nil {
		return
	}
	return
}

func (c *Client) LunCreate(lun *Lun, parameters []string) (res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/luns"
	if req, err = c.NewRequest("POST", path, parameters, lun); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunModify(href string, lun *Lun) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("PATCH", href, []string{}, lun); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunDelete(href string) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunMapGetIter(parameters []string) (lunMaps []LunMap, res *RestResponse, err error) {
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

func (c *Client) LunMapGet(href string, parameters []string) (*LunMap, *RestResponse, error) {
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

func (c *Client) LunMapCreate(lunMap *LunMap, parameters []string) (res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/san/lun-maps"
	if req, err = c.NewRequest("POST", path, parameters, lunMap); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunMapDelete(lunUuid string, igroupUuid string) (res *RestResponse, err error) {
	var req *http.Request
	path := fmt.Sprintf("/api/protocols/san/lun-maps/%s/%s", lunUuid, igroupUuid)
	if req, err = c.NewRequest("DELETE", path, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) LunRead(href string, dataOffset int64, dataSize int64) (data []byte, bytesRead int64, res *RestResponse, err error) {
	var req *http.Request
	data = make([]byte, dataSize)
	var bytesReadMax int64
        for {
                if (dataSize - bytesRead) > 1048576 {
                        bytesReadMax = 1048576
                } else {
                        bytesReadMax = dataSize - bytesRead
                }
	        parameters := []string{fmt.Sprintf("data.offset=%d", dataOffset + bytesRead), fmt.Sprintf("data.size=%d", bytesReadMax)}
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
		        err = fmt.Errorf("LunRead(): expected response in Mime format")
		        return
	        }
	        var p *multipart.Part
	        if p, err = mpartReader.NextPart(); err == nil {
	                n, readErr := p.Read(data[bytesRead:])
	                if readErr != nil {
                                if readErr != io.EOF {
		                        err = fmt.Errorf("LunRead(): mpart read error: %s", readErr)
		                        break
		                }
		        }
		        bytesRead += int64(n)
                        if bytesRead >= dataSize {
                                break
                        }
	        }
	}
	return
}

func (c *Client) LunWrite(href string, dataOffset int64, dataReader io.Reader) (bytesWritten int64, res *RestResponse, err error) {
	var req *http.Request
	var parameters []string
	writeBuffer := make([]byte, 1048576)
	for {
	        n, readErr := dataReader.Read(writeBuffer)
                if n > 0 {
	                parameters = []string{fmt.Sprintf("data.offset=%d", dataOffset + bytesWritten)}
	                if req, err = c.NewFormFileRequest("PATCH", href, parameters, writeBuffer[0:n]); err != nil {
		                return
	                }
	                if res, err = c.Do(req, nil); err == nil {
		                bytesWritten += int64(n)
		        }
		}
                if readErr != nil {
                        if readErr != io.EOF {
                                err = readErr
                        }
                        break
                }
	}
	return
}
