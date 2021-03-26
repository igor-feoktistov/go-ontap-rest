package ontap

import (
	"net/http"
	"fmt"
)

type Snapshot struct {
	Resource
	Comment string            `json:"comment,omitempty"`
	CreateTime string         `json:"create_time,omitempty"`
	ExpiryTime string         `json:"expiry_time,omitempty"`
	Owners []string           `json:"owners,omitempty"`
	SnaplockExpiryTime string `json:"snaplock_expiry_time,omitempty"`
	SnapmirrorLabel string    `json:"snapmirror_label,omitempty"`
	State string              `json:"state,omitempty"`
	Svm *Resource             `json:"svm,omitempty"`
	Volume *Resource          `json:"volume,omitempty"`
}

type SnapshotResponse struct {
	BaseResponse
	Snapshots []Snapshot `json:"records,omitempty"`
}

func (c *Client) SnapshotGetIter(volumeUuid string, parameters []string) (snapshots []Snapshot, res *http.Response, err error) {
	var req *http.Request
	getPath := fmt.Sprintf("/api/storage/volumes/%s/snapshots", volumeUuid)
	reqParameters := parameters
	for {
		r := SnapshotResponse{}
		req, err = c.NewRequest("GET", getPath, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, snapshot := range r.Snapshots {
			snapshots = append(snapshots, snapshot)
		}
		if r.IsPaginate() {
			getPath = r.GetNextRef()
			reqParameters = []string{}
		} else {
			break
		}
	}
	return
}

func (c *Client) SnapshotGet(href string, parameters []string) (*Snapshot, *http.Response, error) {
	r := Snapshot{}
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

func (c *Client) SnapshotCreate(volumeUuid string, snapshot *Snapshot) (res *http.Response, err error) {
	var req *http.Request
	var job *Job
	jobLink := JobLinkResponse{}
	if req, err = c.NewRequest("POST", fmt.Sprintf("/api/storage/volumes/%s/snapshots", volumeUuid), []string{}, snapshot); err != nil {
		return
	}
	if res, err = c.Do(req, &jobLink); err != nil {
		return
	}
	if job, err = c.JobWaitUntilComplete(jobLink.JobLink.GetRef()); err == nil {
		if job != nil && job.State == "failure" {
			err = fmt.Errorf("Error: REST code=%d, REST message=\"%s\"", job.Code, job.Message)
		}
	}
	return
}

func (c *Client) SnapshotModify(href string, snapshot *Snapshot) (res *http.Response, err error) {
	var req *http.Request
	var job *Job
	jobLink := JobLinkResponse{}
	if req, err = c.NewRequest("PATCH", href, []string{}, snapshot); err != nil {
		return
	}
	if res, err = c.Do(req, &jobLink); err != nil {
		return
	}
	if job, err = c.JobWaitUntilComplete(jobLink.JobLink.GetRef()); err == nil {
		if job != nil && job.State == "failure" {
			err = fmt.Errorf("Error: REST code=%d, REST message=\"%s\"", job.Code, job.Message)
		}
	}
	return
}

func (c *Client) SnapshotDelete(href string) (res *http.Response, err error) {
	var req *http.Request
	var job *Job
	jobLink := JobLinkResponse{}
	if req, err = c.NewRequest("DELETE", href, []string{}, nil); err != nil {
		return
	}
	if res, err = c.Do(req, &jobLink); err != nil {
		return
	}
	if job, err = c.JobWaitUntilComplete(jobLink.JobLink.GetRef()); err == nil {
		if job != nil && job.State == "failure" {
			err = fmt.Errorf("Error: REST code=%d, REST message=\"%s\"", job.Code, job.Message)
		}
	}
	return
}
