package ontap

import (
        "fmt"
	"net/http"
)

type NvmeNamespaceLocation struct {
	Namespace string     `json:"namespace,omitempty"`
	Node *Resource       `json:"node,omitempty"`
	Qtree *QtreeResource `json:"qtree,omitempty"`
	Volume *Resource     `json:"volume,omitempty"`
}

type NvmeNamespaceSpaceGuarantee struct {
	Requested bool `json:"requested"`
	Reserved bool  `json:"reserved"`
}

type NvmeNamespaceSpace struct {
        BlockSize *int                         `json:"blocksize,omitempty"`
	Guarantee *NvmeNamespaceSpaceGuarantee `json:"guarantee,omitempty"`
	Size *int64                            `json:"size,omitempty"`
	Used *int64                            `json:"used,omitempty"`
}

type NvmeNamespace struct {
	Resource
	AutoDelete *bool                            `json:"auto_delete,omitempty"`
	Comment string                              `json:"comment,omitempty"`
	CreateTime string                           `json:"create_time,omitempty"`
	Enabled *bool                               `json:"enabled,omitempty"`
	Location *NvmeNamespaceLocation             `json:"location,omitempty"`
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
	Space *NvmeNamespaceSpace                   `json:"space,omitempty"`
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
	SubsystemMap *NvmeSubsystemMap              `json:"subsystem_map,omitempty"`
	Svm *Resource                               `json:"svm,omitempty"`
	Uuid string                                 `json:"uuid,omitempty"`
}

type NvmeNamespaceResponse struct {
	BaseResponse
	NvmeNamespaces []NvmeNamespace `json:"records,omitempty"`
}

type NvmeHostIoQueue struct {
        Count int `json:"count,omitempty"`
        Depth int `json:"depth,omitempty"`
}

type NvmeHost struct {
	Resource
        DhHmacChap *DhHmacChapProtocol `json:"dh_hmac_chap,omitempty"`
	IoQueue *NvmeHostIoQueue       `json:"io_queue,omitempty"`
        Nqn string                     `json:"nqn,omitempty"`
        Subsystem *Resource            `json:"subsystem,omitempty"`
}

type NvmeHostResponse struct {
	BaseResponse
	NvmeHosts []NvmeHost `json:"records,omitempty"`
}

type NvmeSubsystemIoQueue struct {
        Default *NvmeHostIoQueue `json:"default,omitempty"`
}

type NvmeSubsystemMap struct {
        Resource
	Anagrpid string     `json:"anagrpid,omitempty"`
	Namespace *Resource `json:"namespace,omitempty"`
	Nsid string         `json:"nsid,omitempty"`
	Subsystem *Resource `json:"subsystem,omitempty"`
	Svm *Resource       `json:"svm,omitempty"`
}

type NvmeSubsystemMapResponse struct {
	BaseResponse
	NvmeSubsystemMaps []NvmeSubsystemMap `json:"records,omitempty"`
}

type NvmeSubsystem struct {
	Resource
	Comment string                              `json:"comment,omitempty"`
	DeleteOnUnmap *bool                         `json:"delete_on_unmap,omitempty"`
	Hosts []NvmeHost                            `json:"hosts,omitempty"`
	IoQueue *NvmeSubsystemIoQueue               `json:"io_queue,omitempty"`
	Name string                                 `json:"name,omitempty"`
	OsType string                               `json:"os_type,omitempty"`
	SerialNumber string                         `json:"serial_number,omitempty"`
	SubsystemMaps []NvmeSubsystemMap            `json:"subsystem_maps,omitempty"`
	Svm *Resource                               `json:"svm,omitempty"`
	TargetNqn string                            `json:"target_nqn,omitempty"`
	Uuid string                                 `json:"uuid,omitempty"`
	VendorUuids []string                        `json:"vendor_uuids,omitempty"`
}

type NvmeSubsystemResponse struct {
	BaseResponse
	NvmeSubsystems []NvmeSubsystem `json:"records,omitempty"`
}

type NvmeInterface struct {
	Resource
	Enabled *bool                               `json:"enabled,omitempty"`
	FcInterface *struct {
	        Resource
	        Port *struct {
	                Resource
	                Node *struct {
	                        Name string         `json:"name,omitempty"`
	                }                           `json:"node,omitempty"`
		}                                   `json:"port,omitempty"`
		Wwnn string                         `json:"wwnn,omitempty"`
		Wwpn string                         `json:"wwpn,omitempty"`
	}                                           `json:"fc_interface,omitempty"`
	InterfaceType string                        `json:"interface_type,omitempty"`
	IpInterface *struct {
	        Resource
	        Ip *struct {
	                Address string              `json:"address,omitempty"`
	        }                                   `json:"ip,omitempty"`
	        Location *struct {
	                Port *struct {
	                        Resource
	                        Node *struct {
	                                Name string `json:"name,omitempty"`
	                        }                   `json:"node,omitempty"`
		        }                           `json:"port,omitempty"`
		}                                   `json:"location,omitempty"`
	}                                           `json:"ip_interface,omitempty"`
	Name string                                 `json:"name,omitempty"`
	Node *Resource                              `json:"node,omitempty"`
	Svm *Resource                               `json:"svm,omitempty"`
	TransportAddress string                     `json:"transport_address,omitempty"`
	TransportProtocols []string                 `json:"transport_protocols,omitempty"`
	Uuid string                                 `json:"uuid,omitempty"`
}

type NvmeInterfaceResponse struct {
	BaseResponse
	NvmeInterfaces []NvmeInterface `json:"records,omitempty"`
}

func (c *Client) NvmeNamespaceGetIter(parameters []string) (namespaces []NvmeNamespace, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/namespaces"
	reqParameters := parameters
	for {
		r := NvmeNamespaceResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, namespace := range r.NvmeNamespaces {
			namespaces = append(namespaces, namespace)
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

func (c *Client) NvmeNamespaceGet(href string, parameters []string) (*NvmeNamespace, *RestResponse, error) {
	r := NvmeNamespace{}
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

func (c *Client) NvmeNamespaceGetByPath(svmName string, namespacePath string, parameters []string) (namespace *NvmeNamespace, res *RestResponse, err error) {
        var namespaces []NvmeNamespace
	var req *http.Request
	if namespaces, _, err = c.NvmeNamespaceGetIter([]string{"svm.name=" + svmName,"name=" + namespacePath}); err != nil {
	        return
	}
	if len(namespaces) > 0 {
	        namespace = &NvmeNamespace{}
	        if req, err = c.NewRequest("GET", namespaces[0].GetRef(), parameters, nil); err != nil {
		        return
	        }
	        res, err = c.Do(req, namespace)
	} else {
	        err = fmt.Errorf("no NVME namespace \"%s\" found", namespacePath)
	}
	return
}

func (c *Client) NvmeNamespaceCreate(namespace *NvmeNamespace, parameters []string) (namespaces []NvmeNamespace, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/namespaces"
	if req, err = c.NewRequest("POST", path, parameters, namespace); err != nil {
		return
	}
	r := NvmeNamespaceResponse{}
	if res, err = c.Do(req, &r); err == nil {
	        namespaces = r.NvmeNamespaces
	}
	return
}

func (c *Client) NvmeNamespaceModify(href string, namespace *NvmeNamespace) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("PATCH", href, []string{}, namespace); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) NvmeNamespaceDelete(href string) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) NvmeSubsystemGetIter(parameters []string) (subsystems []NvmeSubsystem, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/nvme/subsystems"
	reqParameters := parameters
	for {
		r := NvmeSubsystemResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, subsystem := range r.NvmeSubsystems {
			subsystems = append(subsystems, subsystem)
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

func (c *Client) NvmeSubsystemGet(href string, parameters []string) (*NvmeSubsystem, *RestResponse, error) {
	r := NvmeSubsystem{}
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

func (c *Client) NvmeSubsystemGetByPath(svmName string, namespacePath string) (subsystemHref string, res *RestResponse, err error) {
        var namespaces []NvmeNamespace
	if namespaces, res, err = c.NvmeNamespaceGetIter([]string{"svm.name=" + svmName,"name=" + namespacePath,"fields=subsystem_map"}); err == nil {
	        if len(namespaces) > 0 {
	                if namespaces[0].SubsystemMap != nil {
	                        subsystemHref = namespaces[0].SubsystemMap.Subsystem.GetRef()
	                }
	        } else {
	                err = fmt.Errorf("no NVME namespace \"%s\" found", namespacePath)
	        }
	}
	return
}

func (c *Client) NvmeSubsystemCreate(subsystem *NvmeSubsystem, parameters []string) (subsystems []NvmeSubsystem, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/nvme/subsystems"
	if req, err = c.NewRequest("POST", path, parameters, subsystem); err != nil {
		return
	}
	r := NvmeSubsystemResponse{}
	if res, err = c.Do(req, &r); err == nil {
	        subsystems = r.NvmeSubsystems
	}
	return
}

func (c *Client) NvmeSubsystemModify(href string, subsystem *NvmeSubsystem) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("PATCH", href, []string{}, subsystem); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) NvmeSubsystemDelete(href string, parameters []string) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, parameters, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) NvmeHostGetIter(subsystemHref string, parameters []string) (hosts []NvmeHost, res *RestResponse, err error) {
	var req *http.Request
	path := subsystemHref + "/hosts"
	reqParameters := parameters
	for {
		r := NvmeHostResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, host := range r.NvmeHosts {
			hosts = append(hosts, host)
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

func (c *Client) NvmeHostCreate(subsystemHref string, host *NvmeHost, parameters []string) (hosts []NvmeHost, res *RestResponse, err error) {
	var req *http.Request
	path := subsystemHref + "/hosts"
	if req, err = c.NewRequest("POST", path, parameters, host); err != nil {
		return
	}
	r := NvmeHostResponse{}
	if res, err = c.Do(req, &r); err == nil {
	        hosts = r.NvmeHosts
	}
	return
}

func (c *Client) NvmeHostDelete(subsystemHref string, nqn string) (res *RestResponse, err error) {
	var req *http.Request
	path := subsystemHref + "/hosts/" + nqn
	if req, err = c.NewRequest("DELETE", path, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) NvmeSubsystemMapGetIter(parameters []string) (subsystemMaps []NvmeSubsystemMap, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/nvme/subsystem-maps"
	reqParameters := parameters
	for {
		r := NvmeSubsystemMapResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, subsystemMap := range r.NvmeSubsystemMaps {
			subsystemMaps = append(subsystemMaps, subsystemMap)
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

func (c *Client) NvmeSubsystemMapGetByPath(svmName string, namespacePath string) (subsystemMapHref string, res *RestResponse, err error) {
        var namespaces []NvmeNamespace
	if namespaces, res, err = c.NvmeNamespaceGetIter([]string{"svm.name=" + svmName,"name=" + namespacePath,"fields=subsystem_map"}); err == nil {
	        if len(namespaces) > 0 {
	                if namespaces[0].SubsystemMap != nil {
	                        subsystemMapHref = namespaces[0].SubsystemMap.GetRef()
	                }
	        } else {
	                err = fmt.Errorf("no NVME namespace \"%s\" found", namespacePath)
	        }
	}
	return
}

func (c *Client) NvmeSubsystemMapCreate(subsystemMap *NvmeSubsystemMap, parameters []string) (subsystemMaps []NvmeSubsystemMap, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/nvme/subsystem-maps"
	if req, err = c.NewRequest("POST", path, parameters, subsystemMap); err != nil {
		return
	}
	r := NvmeSubsystemMapResponse{}
	if res, err = c.Do(req, &r); err == nil {
	        subsystemMaps = r.NvmeSubsystemMaps
	}
	return
}

func (c *Client) NvmeSubsystemMapDelete(href string) (res *RestResponse, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) NvmeInterfaceGetIter(parameters []string) (nvmeInterfaces []NvmeInterface, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/nvme/interfaces"
	reqParameters := parameters
	for {
		r := NvmeInterfaceResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, nvmeInterface := range r.NvmeInterfaces {
			nvmeInterfaces = append(nvmeInterfaces, nvmeInterface)
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
