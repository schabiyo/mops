package cmd

// ShowOrg represents the 'organization' command
type ShowOrg struct {
	Strict bool `long:"orgId" description:"The organization Id " env:""`
}

// Execute is callback from go-flags.Commander interface
func (c ShowOrg) Execute(_ []string) (err error) {

	return
}
