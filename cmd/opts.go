package cmd

// OmOpts describes subset of flags/options for selecting target Ops Manager installation
type OmOpts struct {
	URLOpt    string `long:"url"           description:"MongoDB Op Manager URL"                env:"OM_URL" required:"true"`
	ClientOpt string `long:"client"        description:"Override username"        env:"OM_USERNAME" required:"true"`
	APIKeyOpt string `long:"client-secret" description:"Override API Key" env:"OM_API_KEY" required:"true"`
}

// MopsOpts describes the flags/options for the CLI
type MopsOpts struct {
	Version bool `short:"v" long:"version" description:"Show version"`

	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose []bool `long:"verbose" description:"Show verbose debug information" env:"MOPS_TRACE"`
	JSON    bool   `long:"json" description:"Print information in JSON format, for easier parsing" env:"MOPS_AS_JSON"`

	ConfigPathOpt string `long:"config" description:"Config file path" env:"MOPS_CONFIG" default:"~/.mops/config"`

	Broker OmOpts `group:"MongoDB Ops Manager Options"`
}

// Opts carries all the user provided options (from flags or env vars)
var Opts MopsOpts
