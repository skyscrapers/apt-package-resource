package command

import (
	"fmt"

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
	// resourcePath := args[2]

	packageQuery := fmt.Sprintf("%v (= %s)", packageName, packageVersion)
	q, err := query.Parse(packageQuery)
	if err != nil {
		return fmt.Errorf("unable to create package query '%s': %s", packageQuery, err)
	}

	result := q.Query(context.CollectionFactory().PackageCollection())
	if result.Len() > 1 {
		return fmt.Errorf("Not expecting more than 1 package %s with exact version %s", packageName, packageVersion)
	}
	err = result.ForEach(func(p *deb.Package) error {
		printInJSON(packageVersion, p)
		return nil
	})
	return err
}

// printInJson prints the list of versions in a format Concourse expects as output
// http://concourse.ci/implementing-resources.html#in
func printInJSON(version string, packageInfo *deb.Package) {
	fmt.Println("{")
	fmt.Printf("  \"version\": { \"ref\": \"%s\" },\n", version)
	fmt.Printf("  \"metadata\": [\n")
	fmt.Printf("    { \"name\": \"name\", \"value\": \"%s\" },\n", packageInfo.Name)
	fmt.Printf("    { \"name\": \"version\", \"value\": \"%s\" },\n", packageInfo.Version)
	fmt.Printf("    { \"name\": \"architecture\", \"value\": \"%s\" }\n", packageInfo.Architecture)
	fmt.Println("  ]")
	fmt.Println("}")
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
