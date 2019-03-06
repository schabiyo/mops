package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	strings "strings"

	"github.com/hashicorp/errwrap"
	"github.com/landoop/tableprinter"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

var jsonStr = `{
	"delivery" : {
       "methodName" : "AUTOMATED_RESTORE",
       "targetGroupId" : "PROJECT-ID",
       "targetClusterId" : "CLUSTER-ID"
     },
     "snapshotId" : "SNAPSHOT-ID"
}`

//ALLJOBS represents the API endpoint to manage Snapshots
const ALLJOBS = "/groups/%s/clusters/%s/restoreJobs?pageNum=%s&itemsPerPage=50"

//CREATEJOB endpoint for creating a Job
const CREATEJOB = "/groups/%s/clusters/%s/restoreJobs"

//ONEJOB get one org goven an orgID
const ONEJOB = "/groups/%s"

//RestoreAPI hold all API related to handling Snapshots
type RestoreAPI struct {
	//API OpsManagerAPI
}

//RestoreJob hold information about an organization
type RestoreJob struct {
	Status    string `json:"statusName"`
	ClusterID string `json:"clusterId"`
	JobID     string `json:"id"`
}

//RestoreJobResponse hold the response from an API call
type RestoreJobResponse struct {
	Jobs       []RestoreJob `json:"results"`
	TotalCount int          `json:"totalCount"`
	NextPage   string
}

//RestoreJobOut used for the output
type RestoreJobOut struct {
	JobID     string `header:"Job id"`
	ClusterID string `header:"cluster id"`
	Status    string `header:"status"`
}

//NewRestoreAPI return a new instance
func NewRestoreAPI(mops OpsManagerAPI) RestoreAPI {
	return RestoreAPI{}
}

//GetRestoreJobs return all restore jobs
func (r RestoreAPI) GetRestoreJobs(mops *OpsManagerAPI, groupID, clusterID string, pageNum int) (restoreJobResponse *RestoreJobResponse, err error) {
	url := fmt.Sprintf(ALLJOBS, groupID, clusterID, strconv.Itoa(pageNum))
	request := mops.url + url
	dr := dac.NewRequest(mops.username, mops.apiKey, "GET", request, "")
	resp, err := dr.Execute()
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrapf("Failed reading HTTP response body: {{err}}", err)
	}
	projResp := &RestoreJobResponse{}
	err = json.Unmarshal(resBody, projResp)
	if err != nil {
		return nil, errwrap.Wrapf("Failed unmarshalling organization response: {{err}}", err)
	}
	return projResp, nil
}

//GetOneJob return one snapshot details
func (r RestoreAPI) GetOneJob(mops *OpsManagerAPI, projectID string) (restoreJob *RestoreJob, err error) {
	dr := dac.NewRequest(mops.username, mops.apiKey, "GET", mops.url+ONEPROJECT+projectID, "")
	resp, err := dr.Execute()
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrapf("Failed reading HTTP response body: {{err}}", err)
	}
	projResp := &RestoreJob{}
	err = json.Unmarshal(resBody, projResp)
	return projResp, nil
}

//CreateRestoreJob return one snapshot details
func (r RestoreAPI) CreateRestoreJob(mops *OpsManagerAPI, projectID, clusterID, snapshotID string) (job *RestoreJobResponse, err error) {
	url := fmt.Sprintf(CREATEJOB, projectID, clusterID)
	request := mops.url + url
	body := strings.Replace(jsonStr, "PROJECT-ID", projectID, 1)
	body = strings.Replace(body, "CLUSTER-ID", clusterID, 1)
	body = strings.Replace(body, "SNAPSHOT-ID", snapshotID, 1)

	fmt.Println(body)

	dr := dac.NewRequest(mops.username, mops.apiKey, "POST", request, body)
	dr.Header.Set("Content-Type", "application/json")
	resp, err := dr.Execute()
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrapf("Failed reading HTTP response body: {{err}}", err)
	}
	fmt.Println(string(resBody))
	jobResp := &RestoreJobResponse{}
	err = json.Unmarshal(resBody, jobResp)
	return jobResp, nil
}

//GetRestoreJob return one snapshot details
func (r RestoreAPI) GetRestoreJob(mops *OpsManagerAPI, projectID string) (project *Project, err error) {
	dr := dac.NewRequest(mops.username, mops.apiKey, "POST", mops.url+ONEPROJECT+projectID, "")
	resp, err := dr.Execute()
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrapf("Failed reading HTTP response body: {{err}}", err)
	}
	projResp := &Project{}
	err = json.Unmarshal(resBody, projResp)
	return projResp, nil
}

//PrintResultTable output the result in a nice table
func (r RestoreAPI) PrintResultTable(printer *tableprinter.Printer, jobs []RestoreJob) {
	outs := []RestoreJobOut{}
	for _, job := range jobs {
		outs = append(outs, RestoreJobOut{job.JobID, job.ClusterID, job.Status})
	}
	printer.Print(outs)

}
