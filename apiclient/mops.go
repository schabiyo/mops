package apiclient

import (
	"fmt"
	"io/ioutil"
	"os"
)

// OpsManagerAPI is the client struct for connecting to remote Open Service Broker API
type OpsManagerAPI struct {
	url        string
	username   string
	apiKey     string
	apiVersion float32
	//catalog  *omsapi.OMResponse
	//Supported APIs
	OrgAPI      OrganizationAPI
	ProjectAPI  ProjectAPI
	SnapshotAPI SnapshotAPI
	RestoreAPI  RestoreAPI
}

// NewOpsManagerAPI constructs NewOpsManagerAPI
func NewOpsManagerAPI() *OpsManagerAPI {
	mops := OpsManagerAPI{
		url:        os.Getenv("MOPS_URL") + "/api/public/v1.0",
		username:   os.Getenv("MOPS_USERNAME"),
		apiKey:     os.Getenv("MOPS_APIKEY"),
		apiVersion: 3.6,
	}
	//Validate
	mops.validate()
	mops.OrgAPI = NewOrganizationAPI(mops)
	mops.ProjectAPI = NewProjectAPI(mops)
	return &mops
}

//saveToFile save the result to a SCV file
func (o OpsManagerAPI) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte("null"), 0666)

}

func (o OpsManagerAPI) validate() {
	if o.apiKey == "" {
		fmt.Println("Please set your MOPS_APIKEY environment variable")
		os.Exit(1)
	}
	if o.username == "" {
		fmt.Println("Please set your MOPS_USERNAME environment variable")
		os.Exit(1)
	}
	if os.Getenv("MOPS_URL") == "" {
		fmt.Println("Please set your MOPS_URL environment variable")
		os.Exit(1)
	}
	if o.apiVersion >= 4 {
		fmt.Println("Current version only support MongoDB Ops Manager 3.6")
		os.Exit(1)
	}
}
