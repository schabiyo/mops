package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hashicorp/errwrap"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

//ALLCLUSTERS represents the API endpoint
const ALLCLUSTERS = "/groups/GROUP-ID/clusters"

//ONECLUSTER get one project and one cluster ID
const ONECLUSTER = "/groups/%s/clusters/%s"

//ClusterAPI hold all API related to handling Organizations
type ClusterAPI struct {
	//API OpsManagerAPI
}

//Cluster hold infromation about an organization
type Cluster struct {
	ClusterName     string `json:"clusterName"`
	ClusterTypeName string `json:"typeName"`
	ClusterID       string `json:"id"`
	LastHeartbeat   int    `json:"lastHeartbeat"`
}

//ClusterResponse hold the response from an API call
type ClusterResponse struct {
	Clusters   []Cluster `json:"results"`
	TotalCount int       `json:"totalCount"`
}

//NewClusterAPI return a new instance
func NewClusterAPI(mops OpsManagerAPI) ClusterAPI {
	//API = mops
	return ClusterAPI{}
}

//GetClusters return all projects for the user
func (c ClusterAPI) GetClusters(mops *OpsManagerAPI, pageNum int) (projectResponse *ProjectResponse, err error) {
	url := fmt.Sprintf(ALLPROJECTS, strconv.Itoa(pageNum))
	fmt.Println("url=" + url)
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
	projResp := &ProjectResponse{}
	err = json.Unmarshal(resBody, projResp)
	if err != nil {
		return nil, errwrap.Wrapf("Failed unmarshalling organization response: {{err}}", err)
	}
	return projResp, nil
}

//getOneProject return one organization
func (c ClusterAPI) getOneCluster(mops *OpsManagerAPI, projectID string) (cluster *Cluster, err error) {
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
	clusterResp := &Cluster{}
	err = json.Unmarshal(resBody, clusterResp)
	return clusterResp, nil
}
