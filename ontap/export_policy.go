package ontap

import (
        "net/http"
)

type ExportRuleClient struct {
	Match string `json:"match,omitempty"`
}

type ExportPolicyRule struct {
	Resource
	AnonymousUser string       `json:"anonymous_user,omitempty"`
	Clients []ExportRuleClient `json:"clients,omitempty"`
	Index *int                 `json:"index,omitempty"`
	Protocols []string         `json:"protocols,omitempty"`
	RoRule []string            `json:"ro_rule,omitempty"`
	RwRule []string            `json:"rw_rule,omitempty"`
	Superuser []string         `json:"superuser,omitempty"`
}

type ExportPolicyRuleResponse struct {
	BaseResponse
	ExportPolicyRules []ExportPolicyRule `json:"records,omitempty"`
}

type ExportPolicyRef struct {
	Resource
	Id *int `json:"id,omitempty"`
}

type ExportPolicy struct {
	ExportPolicyRef
	Rules []ExportPolicyRule `json:"rules,omitempty"`
	Svm *Resource            `json:"svm,omitempty"`
}

type ExportPolicyResponse struct {
	BaseResponse
	ExportPolicies []ExportPolicy `json:"records,omitempty"`
}

func (c *Client) ExportPolicyGetIter(parameters []string) (expPolicies []ExportPolicy, res *http.Response, err error) {
	var req *http.Request
	path := "/api/protocols/nfs/export-policies"
	reqParameters := parameters
	for {
		r := ExportPolicyResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, expPolicy := range r.ExportPolicies {
			expPolicies = append(expPolicies, expPolicy)
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

func (c *Client) ExportPolicyGet(href string, parameters []string) (*ExportPolicy, *http.Response, error) {
	r := ExportPolicy{}
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

func (c *Client) ExportPolicyCreate(expPolicy *ExportPolicy, parameters []string) (res *http.Response, err error) {
	var req *http.Request
	path := "/api/protocols/nfs/export-policies"
	if req, err = c.NewRequest("POST", path, parameters, expPolicy); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) ExportPolicyModify(href string, expPolicy *ExportPolicy) (res *http.Response, err error) {
	var req *http.Request
	if req, err = c.NewRequest("PATCH", href, []string{}, expPolicy); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) ExportPolicyDelete(href string) (res *http.Response, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) ExportPolicyRuleGetIter(href string, parameters []string) (rules []ExportPolicyRule, res *http.Response, err error) {
	var req *http.Request
	path := href + "/rules"
	reqParameters := parameters
	for {
		r := ExportPolicyRuleResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, rule := range r.ExportPolicyRules {
			rules = append(rules, rule)
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

func (c *Client) ExportPolicyRuleCreate(href string, rule *ExportPolicyRule) (res *http.Response, err error) {
	var req *http.Request
	path := href + "/rules"
	if req, err = c.NewRequest("POST", path, []string{}, rule); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) ExportPolicyRuleModify(href string, rule *ExportPolicyRule) (res *http.Response, err error) {
	var req *http.Request
	if req, err = c.NewRequest("PATCH", href, []string{}, rule); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}

func (c *Client) ExportPolicyRuleDelete(href string) (res *http.Response, err error) {
	var req *http.Request
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	res, err = c.Do(req, nil)
	return
}
