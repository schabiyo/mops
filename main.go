package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"github.com/pivotal-cf/jhanda"
	commands "github.com/schabiyo/mops/cmd"
	"gopkg.in/yaml.v2"
)

type options struct {
	Target   string `yaml:"target"                short:"t"  long:"target"              env:"OM_TARGET"                              description:"url of the Ops Manager instance"`
	Username string `yaml:"username"              short:"u"  long:"username"            env:"OM_USERNAME"                            description:"username to use to authenticate against the Ops Manager instance "`
	APIKey   string `yaml:"api-key"         	  short:"s"  long:"api-key"       env:"OM_API_KEY"                       description:"API Key  for the Ops Manager"`

	JSON bool `yaml:"json"                		  short:"j"  long:"json"              env:"MOPS_AS_JSON"                              description:"Print information in JSON format, for easier parsing"`

	Config bool `yaml:"json"                	short:"c"  long:"config"              env:"MOPS_CONFIG"                              description:"Config file path"`

	Help    bool   `                             short:"h"  long:"help"                                             default:"false" description:"prints this usage information"`
	Trace   bool   `yaml:"trace"                 short:"tr" long:"trace"               env:"OM_TRACE"                               description:"prints HTTP requests and response payloads"`
	Env     string `                             short:"e"  long:"env"                                                              description:"env file with login credentials"`
	Version bool   `                             short:"v"  long:"version"                                          default:"false" description:"prints the om release version"`
}

func main() {
	rand.Seed(5000)

	stdout := log.New(os.Stdout, "", 0)
	stderr := log.New(os.Stderr, "", 0)

	var global options

	args, err := jhanda.Parse(&global, os.Args[1:])
	if err != nil {
		stderr.Fatal(err)
	}

	err = setEnvFileProperties(&global)
	if err != nil {
		stderr.Fatal(err)
	}

	globalFlagsUsage, err := jhanda.PrintUsage(global)
	if err != nil {
		stderr.Fatal(err)
	}

	var command string
	if len(args) > 0 {
		command, args = args[0], args[1:]
	}

	if global.Version {
		command = "version"
	}

	if global.Help {
		command = "help"
	}

	if command == "" {
		command = "help"
	}

	commandSet := jhanda.CommandSet{}
	commandSet["credentials"] = commands.NewCredentials(api, presenter, stdout)
	commandSet["curl"] = commands.NewCurl(api, stdout, stderr)
	commandSet["errands"] = commands.NewErrands(presenter, api)
	commandSet["export-installation"] = commands.NewExportInstallation(api, stderr)
	commandSet["generate-certificate"] = commands.NewGenerateCertificate(api, stdout)
	commandSet["generate-certificate-authority"] = commands.NewGenerateCertificateAuthority(api, presenter)
	commandSet["help"] = commands.NewHelp(os.Stdout, globalFlagsUsage, commandSet)
	commandSet["upload-stemcell"] = commands.NewUploadStemcell(form, api, stdout)
	commandSet["version"] = commands.NewVersion(version, os.Stdout)

	err = commandSet.Execute(command, args)
	if err != nil {
		stderr.Fatal(err)
	}

}

func setEnvFileProperties(global *options) error {
	if global.Env == "" {
		return nil
	}

	var opts options
	file, err := os.Open(global.Env)
	if err != nil {
		return fmt.Errorf("env file does not exist: %s", err)
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("cannot read env file: %s", err)
	}

	err = yaml.Unmarshal(contents, &opts)
	if err != nil {
		return fmt.Errorf("could not parse env file: %s", err)
	}

	if global.Username == "" {
		global.Username = opts.Username
	}
	if global.APIKey == "" {
		global.APIKey = opts.APIKey
	}
	if global.Target == "" {
		global.Target = opts.Target
	}
	if global.Trace == false {
		global.Trace = opts.Trace
	}
	if global.Username == "" {
		global.Username = opts.Username
	}

	return nil
}
