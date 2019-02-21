package apiclient

//ALLORGS represents the API endpoint to manage organizations
const ALLORGS = "/orgs"

//ONEORD get one org goven an orgID
const ONEORD = "/orgs/{ORG-ID}"

//ALLPROJECTS Get all projects inside an org
const ALLPROJECTS = "/orgs/{ORG-ID}/groups"

//OrganizationInterface hold all API related to handling Organizations
type OrganizationInterface struct {
}

//Organization hold infromation about an organization
type Organization struct {
	OrgName string `json:"name"`
	OrgID   string `json:"id"`
}

//OrganizationResponse hold the reponse from an API call
type OrganizationResponse struct {
	OrgName []Organization `json:"results"`
}

//NewOrganizationInterface return a new instance
func NewOrganizationInterface() *OrganizationInterface {
	return &OrganizationInterface{}
}

//GetAllOrgs return all organizations
func (o OrganizationInterface) getAllOrgs() []Organization {

	orgs := []Organization{Organization{"", ""}, Organization{"", ""}}
	return orgs
}

//GetOneOrg return one organization
func (o OrganizationInterface) getOneOrg(orgID string) Organization {
	org := Organization{OrgName: "", OrgID: ""}
	return org
}
