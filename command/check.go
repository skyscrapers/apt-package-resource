package command

import (
	"bufio"
	"fmt"
	"os"

	cli "github.com/smira/aptly/cmd"
	"github.com/smira/aptly/deb"
	"github.com/smira/aptly/query"
	"github.com/smira/commander"
	"github.com/smira/flag"
)

func aptResourceCheck(cmd *commander.Command, args []string) error {
	var err error
	var context = cli.GetContext()

	if len(args) != 1 {
		cmd.Usage()
		return commander.ErrCommandError
	}

	q, err := query.Parse(args[0])
	if err != nil {
		return fmt.Errorf("unable to show: %s", err)
	}

	w := bufio.NewWriter(os.Stdout)

	result := q.Query(context.CollectionFactory().PackageCollection())

	err = result.ForEach(func(p *deb.Package) error {
		p.Stanza().WriteTo(w, p.IsSource, false)
		w.Flush()
		fmt.Printf("\n")

		return nil
	})

	if err != nil {
		return fmt.Errorf("unable to show: %s", err)
	}

	return err
}

func makeCmdCheck() *commander.Command {
	cmd := &commander.Command{
		Run:       aptResourceCheck,
		UsageLine: "check <package-query>",
		Short:     "check for new versions of the package",
		Long: `
Command shows displays detailed meta-information about packages
matching query. Information from Debian control file is displayed.
Optionally information about package files and
inclusion into mirrors/snapshots/local repos is shown.

Example:

    $ apt-resource check aptly
`,
		Flag: *flag.NewFlagSet("apt-resource-check", flag.ExitOnError),
	}

	return cmd
}
