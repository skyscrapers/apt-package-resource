package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPackageVersions(t *testing.T) {
	assert := assert.New(t)
	concourseInput :=
		`dir = "/Users/ringods/Downloads/apt-mirror"
		 
		 [mirror.package]
		 url = "http://archive.ubuntu.com/ubuntu"
		 suites = ["zesty"]
		 sections = ["main"]
		 mirror_source = false
	     architectures = ["amd64"]
		`

	result := checkPackageVersions(concourseInput)
	assert.Equal("", result, "should be equal")
}
