package ontap

import (
	"net/http"
)

type IpInterface struct {
	Resource
	Enabled bool                     `json:"enabled"`
	Ip IpInfo                        `json:"ip,omitempty"`
	IpSpace Resource                 `json:"ipspace,omitempty"`
	Location struct {
		AutoRevert bool          `json:"auto_revert"`
		BroadcastDomain Resource `json:"broadcast_domain,omitempty"`
		Failover string          `json:"failover,omitempty"`
		HomeNode Resource        `json:"home_node,omitempty"`
		HomePort struct {
			Resource
			Node Resource    `json:"node,omitempty"`
		}                        `json:"home_port,omitempty"`
		IsHome bool              `json:"is_home"`
		Node Resource            `json:"node"`
		Port struct {
			Resource
			Node Resource    `json:"node,omitempty"`
		}                        `json:"port,omitempty"`
	}                                `json:"location,omitempty"`
	Metrics struct {
		Resource
		Duration string          `json:"duration,omitempty"`
		Status string            `json:"status,omitempty"`
		Throughput struct {
			Read int         `json:"read"`
			Total int        `json:"total"`
			Write int        `json:"write"`
		}                        `json:"throughput,omitempty"`
		Timestamp string         `json:"timestamp,omitempty"`
	}                                `json:"metrics,omitempty"`
	Scope string                     `json:"scope,omitempty"`
	ServicePolicy Resource           `json:"service_policy,omitempty"`
	Services []string                `json:"services,omitempty"`
	State string                     `json:"state,omitempty"`
	Statistics struct {
		Status string            `json:"status,omitempty"`
		ThroughputRaw struct {
			Read int         `json:"read"`
			Total int        `json:"total"`
			Write int        `json:"write"`
		}                        `json:"throughput,omitempty"`
		Timestamp string         `json:"timestamp,omitempty"`
	}                                `json:"statistics,omitempty"`
	Svm Resource                     `json:"svm,omitempty"`
	Vip bool                         `json:"vip"`
}

type IpInterfaceResponse struct {
	BaseResponse
	IpInterfaces []IpInterface `json:"records,omitempty"`
}

func (c *Client) IpInterfaceGetIter(parameters []string) (ipInterfaces []IpInterface, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/network/ip/interfaces"
	reqParameters := parameters
	for {
		r := IpInterfaceResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, ipInterface := range r.IpInterfaces {
			ipInterfaces = append(ipInterfaces, ipInterface)
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

func (c *Client) IpInterfaceGet(href string, parameters []string) (*IpInterface, *RestResponse, error) {
	r := IpInterface{}
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
