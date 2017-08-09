// Package command implements console commands
package command

import (
	"os"

	"github.com/smira/commander"
	"github.com/smira/flag"
)

// RootCommand creates root command in command tree
func RootCommand() *commander.Command {
	cmd := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "Concourse resource type to work with APT repository packages",
		Long: `
apt-resource is a tool to create partial and full mirrors of remote
repositories, manage local repositories, filter them, merge,
upgrade individual packages, take snapshots and publish them
back as Debian repositories.

aptly's goal is to establish repeatability and controlled changes
in a package-centric environment. aptly allows one to fix a set of packages
in a repository, so that package installation and upgrade becomes
deterministic. At the same time aptly allows one to perform controlled,
fine-grained changes in repository contents to transition your
package environment to new version.`,
		Flag: *flag.NewFlagSet("apt-resource", flag.ExitOnError),
		Subcommands: []*commander.Command{
			makeCmdCheck(),
			// makeCmdIn(),
		},
	}

	cmd.Flag.Int("db-open-attempts", 10, "number of attempts to open DB if it's locked by other instance")
	cmd.Flag.String("config", "", "location of configuration file (default locations are /etc/aptly.conf, ~/.aptly.conf)")

	return cmd
}
