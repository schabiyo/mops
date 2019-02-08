package apiclient

// OpsManagerInterface is the client struct for connecting to remote Open Service Broker API
type OpsManagerInterface struct {
	url      string
	username string
	apiKey   string
	//catalog  *omsapi.OMResponse
}

// NewOpsManagerInterface constructs OpsManagerInterface
func NewOpsManagerInterface(url, username, apiKey string) *OpsManagerInterface {
	return &OpsManagerInterface{
		url:      url,
		username: username,
		apiKey:   apiKey,
	}
}
