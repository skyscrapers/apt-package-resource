package command

import (
	"fmt"
	"sort"
	"strings"

	cli "github.com/smira/aptly/cmd"
	"github.com/smira/aptly/deb"
	"github.com/smira/aptly/query"
	"github.com/smira/commander"
	"github.com/smira/flag"
)

// ByVersion allows to compare the package version using the Debian versioning policy
type ByVersion []string

func (s ByVersion) Len() int {
	return len(s)
}
func (s ByVersion) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByVersion) Less(i, j int) bool {
	return deb.CompareVersions(s[i], s[j]) < 0
}

// Filter an array of strings given a predicate function
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// SameOrHigher creates a predicate function testing versions to a reference version
func SameOrHigher(referenceVersion string) func(string) bool {
	return func(versionToCompare string) bool {
		return deb.CompareVersions(referenceVersion, versionToCompare) <= 0
	}
}

func aptResourceCheck(cmd *commander.Command, args []string) error {
	var err error
	var context = cli.GetContext()

	if len(args) != 2 {
		cmd.Usage()
		return commander.ErrCommandError
	}

	q, err := query.Parse(args[0])
	if err != nil {
		return fmt.Errorf("unable to show: %s", err)
	}

	result := q.Query(context.CollectionFactory().PackageCollection())
	versions := make([]string, 0, 5)

	err = result.ForEach(func(p *deb.Package) error {
		versions = append(versions, p.Version)
		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to show: %s", err)
	}

	sort.Sort(ByVersion(versions))
	matchingVersions := Filter(versions, SameOrHigher(getVersionToCompare(args[1], versions)))
	printCheckJSON(matchingVersions)

	return err
}

func getVersionToCompare(givenVersion string, foundVersions []string) string {
	if strings.Compare(givenVersion, "latest") == 0 {
		return foundVersions[len(foundVersions)-1]
	}
	return givenVersion
}

// printCheckJson prints the list of versions in a format Concourse expects as output
// http://concourse.ci/implementing-resources.html#resource-check
func printCheckJSON(versions []string) {
	fmt.Println("[")
	for i, version := range versions {
		fmt.Printf("  { \"ref\": \"%s\" }", version)
		if i+1 < len(versions) {
			fmt.Print(",")
		}
		fmt.Println()
	}
	fmt.Println("]")
}

func makeCmdCheck() *commander.Command {
	cmd := &commander.Command{
		Run:       aptResourceCheck,
		UsageLine: "check <package-name> <from-version>",
		Short:     "check for new versions of the package",
		Long: `
Command shows displays detailed meta-information about packages
matching query. Information from Debian control file is displayed.
Optionally information about package files and
inclusion into mirrors/snapshots/local repos is shown.

Example:

    $ apt-resource check aptly 0.9.1
`,
		Flag: *flag.NewFlagSet("apt-resource-check", flag.ExitOnError),
	}

	return cmd
}
