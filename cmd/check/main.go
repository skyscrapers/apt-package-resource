package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/cybozu-go/aptutil/mirror"
)

const (
	lockFilename = ".lock"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	concourseInput := string(bytes)
	fmt.Println(concourseInput)
	checkPackageVersions(concourseInput)
	// md, err := toml.DecodeFile(*configPath, config)
	// if err != nil {
	// 	log.ErrorExit(err)
	// }
	// if len(md.Undecoded()) > 0 {
	// 	log.Error("invalid config keys", map[string]interface{}{
	// 		"keys": fmt.Sprintf("%#v", md.Undecoded()),
	// 	})
	// 	os.Exit(1)
	// }

	// err = config.Log.Apply()
	// if err != nil {
	// 	log.ErrorExit(err)
	// }

	// err = mirror.Run(config, flag.Args())
	// if err != nil {
	// 	log.ErrorExit(err)
	// }
}

func checkPackageVersions(checkInput string) string {
	var mirrorConfig = mirror.NewConfig()
	toml.Decode(checkInput, &mirrorConfig)
	// fmt.Println(config.Package)
	// fmt.Println(mirrorConfig.Dir)
	// fmt.Println(mirrorConfig.MaxConns)
	// fmt.Println(mirrorConfig.Mirrors["package"].URL)
	// fmt.Println(mirrorConfig.Mirrors["package"].Architectures)
	mirror.Run(mirrorConfig, nil, OnlyMetadata)
	return ""
}

// OnlyMetadata runs a partial mirror process: release files, index files but no packages
func OnlyMetadata(m *mirror.Mirror) func(context.Context) error {
	return func(ctx context.Context) error {
		fileMap, err := m.UpdateMetadata(ctx)
		// _, err := m.UpdateMetadata(ctx)
		if err != nil {
			return err
		}
		for key, file := range fileMap {
			fmt.Printf("%-100s: %s\n", key, file.Path())
		}
		return nil
	}
}
