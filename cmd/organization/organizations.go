package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/schabiyo/mops/apiclient"
)

// OrganizationsOpts represents the 'oerganizations' command
type OrganizationsOpts struct {
	Strict bool `long:"strict" description:"Validate " env:"EDEN_STRICT"`
}

// Execute is callback from go-flags.Commander interface
func (c OrganizationsOpts) Execute(_ []string) (err error) {
	broker := apiclient.NewOpsManagerInterface(
		Opts.Broker.URLOpt,
		Opts.Broker.ClientOpt,
		Opts.Broker.ClientSecretOpt,
		Opts.Broker.APIVersion,
	)

	if Opts.JSON {
		b, err := json.Marshal(catalogResp)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(b))
		os.Exit(0)
	}

	return
}
