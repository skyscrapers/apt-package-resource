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

// ByVersion allows to compare the package version using the Debian versioning policy
// type ByVersion []string

// func (s ByVersion) Len() int {
// 	return len(s)
// }
// func (s ByVersion) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }
// func (s ByVersion) Less(i, j int) bool {
// 	return deb.CompareVersions(s[i], s[j]) < 0
// }

// // Filter an array of strings given a predicate function
// func Filter(vs []string, f func(string) bool) []string {
// 	vsf := make([]string, 0)
// 	for _, v := range vs {
// 		if f(v) {
// 			vsf = append(vsf, v)
// 		}
// 	}
// 	return vsf
// }

// // SameOrHigher creates a predicate function testing versions to a reference version
// func SameOrHigher(referenceVersion string) func(string) bool {
// 	return func(versionToCompare string) bool {
// 		return deb.CompareVersions(referenceVersion, versionToCompare) <= 0
// 	}
// }

func aptResourceIn(cmd *commander.Command, args []string) error {
	var err error
	var context = cli.GetContext()

	if len(args) != 3 {
		cmd.Usage()
		return commander.ErrCommandError
	}

	packageName := args[0]
	packageVersion := args[1]
	resourcePath := args[2]

	packageQuery := fmt.Sprintf("%v (= %s)", packageName, packageVersion)
	q, err := query.Parse(packageQuery)
	if err != nil {
		return fmt.Errorf("Unable to create package query '%s': %s", packageQuery, err)
	}

	result := q.Query(context.CollectionFactory().PackageCollection())
	if result.Len() != 1 {
		return fmt.Errorf("Expecting just 1 package (%s) matching version %s, but got %d", packageName, packageVersion, result.Len())
	}

	metadataFile := fmt.Sprintf("%s/metadata", resourcePath)
	f, err := os.Create(metadataFile)
	if err != nil {
		return fmt.Errorf("Unable to open file '%s': %s", metadataFile, err)
	}

	defer f.Close()

	wf := bufio.NewWriter(f)
	ws := bufio.NewWriter(os.Stdout)
	err = result.ForEach(func(p *deb.Package) error {
		err = printInJSON(packageVersion, p, wf)
		if err != nil {
			return fmt.Errorf("Unable to write to file '%s': %s", metadataFile, err)
		}

		err = printInJSON(packageVersion, p, ws)
		if err != nil {
			return fmt.Errorf("Unable to write to stdout: %s", err)
		}
		return nil
	})
	return err
}

// printInJson prints the list of versions in a format Concourse expects as output
// http://concourse.ci/implementing-resources.html#in
func printInJSON(version string, packageInfo *deb.Package, w *bufio.Writer) error {
	_, err := fmt.Fprintf(w, `{
	"version": { "ref": "%s" },
	"metadata": [
		{ "name": "name", "value": "%s" },
		{ "name": "version", "value": "%s" },
		{ "name": "architecture", "value": "%s" }
	]
}`, version, packageInfo.Name, packageInfo.Version, packageInfo.Architecture)
	if err != nil {
		return err
	}

	w.Flush()
	return nil
}

func makeCmdIn() *commander.Command {
	cmd := &commander.Command{
		Run:       aptResourceIn,
		UsageLine: "in <package-name> <version> <resource-path>",
		Short:     "Prepares the package and metadata in the <resource-path>",
		Long: `
Copy the version of the given package and prepares the metadata for
further consumption

Example:

    $ apt-resource in aptly 0.9.1 ./aptly-releases
`,
		Flag: *flag.NewFlagSet("apt-resource-in", flag.ExitOnError),
	}

	return cmd
}
