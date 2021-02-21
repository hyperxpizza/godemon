package prepareProject

import (
	"encoding/json"
	"fmt"
	"go/build"
	"godemon/errors"
	"godemon/models"
	"io/ioutil"
	"os"
)

func LoadProjectInfo() models.Project {
	jsonFile, err := os.Open("project.json")
	errors.ErrorHandle(err)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	errors.ErrorHandle(err)
	var project models.Project
	err = json.Unmarshal(byteValue, &project)
	errors.ErrorHandle(err)
	if project.Name == "" {
		fmt.Println("Project name is empty")
		os.Exit(1)
	}
	if project.OS == "" && project.Arch == "" {
		project.OS = build.Default.GOOS
		project.Arch = build.Default.GOARCH
	} else if project.OS == "" {
		project.OS = build.Default.GOOS
	} else if project.Arch == "" {
		project.Arch = build.Default.GOARCH
	}
	return project
}
