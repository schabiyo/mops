package organization

// OrgOpts represents the 'oerganization' command
type OrgOpts struct {
	Strict bool `long:"orgId" description:"The organization Id " env:""`
}

// Execute is callback from go-flags.Commander interface
func (c OrgOpts) Execute(_ []string) (err error) {

	return
}
