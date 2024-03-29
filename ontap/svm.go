package ontap

import (
	"net/http"
)

type Aggregate struct {
	Resource
	AvailableSize int64 `json:"available_size"`
	Type string         `json:"type"`
	State string        `json:"state"`
	SnaplockType string `json:"snaplock_type,omitempty"`
}

type AdDomain struct {
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	Fqdn string               `json:"fqdn,omitempty"`
	Password string           `json:"password,omitempty"`
	User string               `json:"user,omitempty"`
}

type Dns struct {
	Resource
	Domains []string `json:"domains,omitempty"`
	Servers []string `json:"servers,omitempty"`
}

type Ldap struct {
	Resource
	AdDomain string  `json:"ad_domain,omitempty"`
	BaseDn string    `json:"base_dn,omitempty"`
	BindDn string    `json:"bind_dn,omitempty"`
	Servers []string `json:"servers,omitempty"`
	Enabled bool     `json:"enabled"`
}

type Nfs struct {
	Resource
	Enabled bool `json:"enabled"`
}

type Iscsi struct {
	Resource
	Enabled bool `json:"enabled"`
}

type Fcp struct {
	Resource
	Enabled bool `json:"enabled"`
}

type Nvme struct {
	Resource
	Enabled bool `json:"enabled"`
}

type Cifs struct {
	Resource
	Name string       `json:"name"`
	AdDomain AdDomain `json:"ad_domain,omitempty"`
	Enabled bool      `json:"enabled"`
}

type Nis struct {
	Resource
	Domain string    `json:"nis_domain,omitempty"`
	Servers []string `json:"nis_servers,omitempty"`
	Enabled bool     `json:"enabled"`
}

type NsSwitch struct {
	Resource
	Group []string    `json:"group,omitempty"`
	Hosts []string    `json:"hosts,omitempty"`
	NameMap []string  `json:"namemap,omitempty"`
	NetGroup []string `json:"netgroup,omitempty"`
	Passwd []string   `json:"passwd,omitempty"`
}

type Ip struct {
	Address string `json:"address,omitempty"`
	Netmask string `json:"netmask,omitempty"`
}

type IpInfo struct {
	Address string `json:"address,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Family  string `json:"family,omitempty"`
}

type IpInterfaceSvm struct {
	Resource
	Location struct {
		HomeNode Resource        `json:"home_node,omitempty"`
		BroadcastDomain Resource `json:"broadcast_domain,omitempty"`
	}                                `json:"location,omitempty"`
	Ip Ip                            `json:"ip,omitempty"`
	Services []string                `json:"services,omitempty"`
}

type FcPortReference struct {
	Resource
	Node string
}

type FcInterfaceSvm struct {
	Resource
	DataProtocol string              `json:"data_protocol,omitempty"`
	Location struct {
		port FcPortReference     `json:"port,omitempty"`
	}                                `json:"location,omitempty"`
}

type NetworkRoute struct {
	Gateway string     `json:"gateway,omitempty"`
	Destination IpInfo `json:"destination,omitempty"`
}

type S3Service struct {
	Resource
	Certificate Resource `json:"certificate,omitempty"`
	IsHttpEnabled bool   `json:"is_http_enabled"`
	IsHttpsEnabled bool  `json:"is_https_enabled"`
	Port int             `json:"port"`
	SecurePort int       `json:"secure_port"`
	Enabled bool         `json:"enabled"`
}

type Svm struct {
	Resource
	Aggregates []Aggregate            `json:"aggregates,omitempty"`
	AggregatesDelegated bool          `json:"aggregates_delegated"`
	Certificate Resource              `json:"certificate,omitempty"`
	Cifs Cifs                         `json:"cifs,omitempty"`
	Comment string                    `json:"comment,omitempty"`
	Dns Dns                           `json:"dns,omitempty"`
	Fcp Fcp                           `json:"fcp,omitempty"`
	FcInterfaces []FcInterfaceSvm     `json:"fc_interfaces,omitempty"`
	IpInterfaces []IpInterfaceSvm     `json:"ip_interfaces,omitempty"`
	IpSpace Resource                  `json:"ipspace,omitempty"`
	Iscsi Iscsi                       `json:"iscsi,omitempty"`
	Language string                   `json:"language,omitempty"`
	Ldap Ldap                         `json:"ldap,omitempty"`
	Nfs Nfs                           `json:"nfs,omitempty"`
	Nis Nis                           `json:"nis,omitempty"`
	NsSwitch NsSwitch                 `json:"nsswitch,omitempty"`
	Nvme Nvme                         `json:"nvme,omitempty"`
	Routes []NetworkRoute             `json:"routes,omitempty"`
	S3 S3Service                      `json:"s3,omitempty"`
	SnapMirror struct {
		IsProtected bool          `json:"is_protected"`
		ProtectedVolumesCount int `json:"protected_volumes_count"`
	}                                 `json:"snapmirror"`
	SnapshotPolicy Resource           `json:"snapshot_policy,omitempty"`
	State string                      `json:"state,omitempty"`
	Subtype string                    `json:"subtype,omitempty"`
	VolumeEfficiencyPolicy Resource   `json:"volume_efficiency_policy,omitempty"`
}

type SvmResponse struct {
	BaseResponse
	Svms []Svm `json:"records,omitempty"`
}

func (c *Client) SvmGetIter(parameters []string) (svms []Svm, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/svm/svms"
	reqParameters := parameters
	for {
		r := SvmResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, svm := range r.Svms {
			svms = append(svms, svm)
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

func (c *Client) SvmGet(href string, parameters []string) (*Svm, *RestResponse, error) {
	r := Svm{}
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
