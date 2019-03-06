package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/errwrap"
	"github.com/landoop/tableprinter"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

//ALLORGS represents the API endpoint to manage organizations
const ALLORGS = "/orgs"

//ONEORD get one org goven an orgID
const ONEORD = "/orgs/{ORG-ID}"

//ALLORGPROJECTS Get all projects inside an org
const ALLORGPROJECTS = "/orgs/{ORG-ID}/groups"

//OrganizationAPI hold all API related to handling Organizations
type OrganizationAPI struct {
	//API OpsManagerAPI
}

//Organization hold infromation about an organization
type Organization struct {
	OrgName string `json:"name"`
	OrgID   string `json:"id"`
}

//OrganizationResponse hold the reponse from an API call
type OrganizationResponse struct {
	Orgs []Organization `json:"results"`
}

type OrganizationOut struct {
	OrganizationID   string `header:"id"`
	OrganizationName string `header:"name"`
}

//NewOrganizationAPI return a new instance
func NewOrganizationAPI(mops OpsManagerAPI) OrganizationAPI {
	//API = mops
	return OrganizationAPI{}
}

//GetAllOrgs return all organizations
func (o OrganizationAPI) GetAllOrgs(mops *OpsManagerAPI, pageNum int) (organizationResponse *OrganizationResponse, err error) {
	if mops.apiVersion >= 4 {
		fmt.Println("Current version only support MongoDB Ops Manager 3.6")
		os.Exit(1)
	}
	dr := dac.NewRequest(mops.username, mops.apiKey, "GET", mops.url+ALLORGS, "")
	resp, err := dr.Execute()
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("404")
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrapf("Failed reading HTTP response body: {{err}}", err)
	}

	fmt.Println(err)

	fmt.Println(string(resBody))
	orgResp := &OrganizationResponse{}
	err = json.Unmarshal(resBody, orgResp)
	if err != nil {
		return nil, errwrap.Wrapf("Failed unmarshalling organization response: {{err}}", err)
	}
	fmt.Println(resp)
	return orgResp, nil
}

//func (broker *OpenServiceBroker) Catalog() (catalogResp *brokerapi.CatalogResponse, err error) {
//GetOneOrg return one organization
func (o OrganizationAPI) getOneOrg(orgID string) Organization {
	org := Organization{OrgName: "", OrgID: ""}
	return org
}

//PrintResultTable output the result in a nice table
func (s SnapshotAPI) PrintResultTable(printer *tableprinter.Printer, orgs []Organization) {
	outs := []OrganizationOut{}
	for _, snap := range orgs {
		outs = append(outs, OrganizationOut{snap.ClusterID, snap.ClusterID, snap.IsCompleted})
	}
	printer.Print(outs)

}
