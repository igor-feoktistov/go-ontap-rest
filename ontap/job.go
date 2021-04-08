package ontap

import (
	"net/http"
	"time"
)

type Job struct {
	Resource
	Code int           `json:"code"`
	Description string `json:"description,omitempty"`
	EndTime string     `json:"end_time,omitempty"`
	Message string     `json:"message,omitempty"`
	StartTime string   `json:"start_time,omitempty"`
	State string       `json:"state,omitempty"`
	Svm *Resource      `json:"svm,omitempty"`
}

type JobLinkResponse struct {
	JobLink struct {
		Resource
	} `json:"job"`
}

type JobResponse struct {
	BaseResponse
	Jobs []Job `json:"records,omitempty"`
}

func (c *Client) JobGetIter(parameters []string) (jobs []Job, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/cluster/jobs"
	reqParameters := parameters
	for {
		r := JobResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, job := range r.Jobs {
			jobs = append(jobs, job)
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

func (c *Client) JobGet(href string, parameters []string) (*Job, *RestResponse, error) {
	r := Job{}
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

func (c *Client) JobWaitUntilComplete(href string) (job *Job, err error) {
	if len(href) > 0 {
		for {
			if job, _, err = c.JobGet(href, []string{}); err != nil {
				break
			}
			if job.State == "success" || job.State == "failure" {
				break
			}
			time.Sleep(time.Second)
		}
	}
	return
}
