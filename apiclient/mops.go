package apiclient

import (
	"fmt"
	"io/ioutil"
	"os"
)

// OpsManagerInterface is the client struct for connecting to remote Open Service Broker API
type OpsManagerInterface struct {
	url      string
	username string
	apiKey   string
	//catalog  *omsapi.OMResponse
	//Supported APIs
	OrganizationAPI OrganizationInterface
}

// NewOpsManagerInterface constructs OpsManagerInterface
func NewOpsManagerInterface() *OpsManagerInterface {
	mops := OpsManagerInterface{
		url:      os.Getenv("MOPS_URL"),
		username: os.Getenv("MOPS_USERNAME"),
		apiKey:   os.Getenv("MOPS_APIKEY"),
	}
	//Validate
	mops.validate()
	return &mops
}

//saveToFile save the result to a SCV file
func (o OpsManagerInterface) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte("null"), 0666)

}

func (o OpsManagerInterface) validate() {
	if o.apiKey == "" {
		fmt.Println("Please set your MOPS_APIKEY environment variable")
	}
	if o.username == "" {
		fmt.Println("Please set your MOPS_USERNAME environment variable")
	}
	if o.url == "" {
		fmt.Println("Please set your MOPS_URL environment variable")
	}
	os.Exit(1)
}
