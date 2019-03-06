package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"

	"github.com/hashicorp/errwrap"
	"github.com/landoop/tableprinter"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

//ALLPROJECTS represents the API endpoint to manage organizations
const ALLPROJECTS = "/groups?pageNum=%s&itemsPerPage=100"

//ONEPROJECT get one org goven an orgID
const ONEPROJECT = "/groups/%s"

//ProjectAPI hold all API related to handling Organizations
type ProjectAPI struct {
	//API OpsManagerAPI
}

//Project hold infromation about an organization
type Project struct {
	ProjectName     string `json:"name"`
	ProjectOrgID    string `json:"orgId"`
	ProjectID       string `json:"id"`
	ReplicaSetCount int    `json:"replicaSetCount"`
	LastActiveAgent string `json:"lastActiveAgent"`
}

//ProjectResponse hold the response from an API call
type ProjectResponse struct {
	Projects   []Project `json:"results"`
	TotalCount int       `json:"totalCount"`
	NextPage   string
}

//ProjectOut is used to print out the result
type ProjectOut struct {
	ProjectID       string `header:"id"`
	ProjectName     string `header:"name"`
	ProjectOrgID    string `header:"orgId"`
	ReplicaSetCount int    `header:"replicaSet Count"`
	LastActiveAgent string `header:"last Active Agent"`
}

//NewProjectAPI return a new instance
func NewProjectAPI(mops OpsManagerAPI) ProjectAPI {
	return ProjectAPI{}
}

//GetProjects return all projects for the user
func (p ProjectAPI) GetProjects(mops *OpsManagerAPI, pageNum int) (projectResponse *ProjectResponse, err error) {
	url := fmt.Sprintf(ALLPROJECTS, strconv.Itoa(pageNum))
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
func (p ProjectAPI) getOneProject(mops *OpsManagerAPI, projectID string) (project *Project, err error) {
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

//PrintResultTable output the result in a nice table
func (p ProjectAPI) PrintResultTable(printer *tableprinter.Printer, projects []Project) {
	outs := []ProjectOut{}
	for _, proj := range projects {
		outs = append(outs, ProjectOut{proj.ProjectID, proj.ProjectName, proj.ProjectOrgID, proj.ReplicaSetCount, proj.LastActiveAgent})
	}

	sort.Slice(outs, func(i, j int) bool {
		return outs[j].LastActiveAgent < outs[i].LastActiveAgent
	})
	printer.Print(outs)

}
