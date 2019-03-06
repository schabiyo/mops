package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hashicorp/errwrap"
	"github.com/landoop/tableprinter"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

//ALLSNAPSHOTS represents the API endpoint to manage Snapshots
const ALLSNAPSHOTS = "/groups/%s/clusters/%s/snapshots?pageNum=%s&itemsPerPage=50"

//ONESNAPSHOT get one org goven an orgID
const ONESNAPSHOT = "/groups/%s"

//SnapshotAPI hold all API related to handling Snapshots
type SnapshotAPI struct {
	//API OpsManagerAPI
}

//Snapshot hold infromation about an organization
type Snapshot struct {
	IsCompleted bool   `json:"complete"`
	ClusterID   string `json:"clusterId"`
	SnapshotID  string `json:"id"`
}

//SnapshotResponse hold the response from an API call
type SnapshotResponse struct {
	Snapshots  []Snapshot `json:"results"`
	TotalCount int        `json:"totalCount"`
	NextPage   string
}

type SnapshotOut struct {
	SnapshotID  string `header:"snapshot id"`
	ClusterID   string `header:"cluster id"`
	IsCompleted bool   `header:"completed"`
}

//NewSnapshotAPI return a new instance
func NewSnapshotAPI(mops OpsManagerAPI) SnapshotAPI {
	//API = mops
	return SnapshotAPI{}
}

//GetSnapshots return all snapshots for one cluster
func (s SnapshotAPI) GetSnapshots(mops *OpsManagerAPI, groupID, clusterID string, pageNum int) (snapshotResponse *SnapshotResponse, err error) {
	url := fmt.Sprintf(ALLSNAPSHOTS, groupID, clusterID, strconv.Itoa(pageNum))
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
	projResp := &SnapshotResponse{}
	err = json.Unmarshal(resBody, projResp)
	if err != nil {
		return nil, errwrap.Wrapf("Failed unmarshalling organization response: {{err}}", err)
	}
	return projResp, nil
}

//GetOneSnapshot return one snapshot details
func (s SnapshotAPI) GetOneSnapshot(mops *OpsManagerAPI, projectID string) (project *Project, err error) {
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
	projResp := &Project{}
	err = json.Unmarshal(resBody, projResp)
	return projResp, nil
}

//CreateRestoreJob return one snapshot details
func (s SnapshotAPI) CreateRestoreJob(mops *OpsManagerAPI, projectID string) (project *Project, err error) {
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

//GetRestoreJob return one snapshot details
func (s SnapshotAPI) GetRestoreJob(mops *OpsManagerAPI, projectID string) (project *Project, err error) {
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
func (s SnapshotAPI) PrintResultTable(printer *tableprinter.Printer, snaps []Snapshot) {
	outs := []SnapshotOut{}
	for _, snap := range snaps {
		outs = append(outs, SnapshotOut{snap.ClusterID, snap.ClusterID, snap.IsCompleted})
	}
	printer.Print(outs)

}
