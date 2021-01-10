package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// ExamFile ... Struct to hold structure of exam files
type ExamFile struct {
	Name     string
	Endpoint string
	Tests    []struct {
		Code  int
		Paths []string
	}
}

// ReadExam ... Help read a given file and set to a struct
func ReadExam(file string) ExamFile {
	var exam ExamFile
	examFile, err := ioutil.ReadFile(file)
	ErrorCheck(err)
	extension := filepath.Ext(strings.TrimSpace(file))
	err = fmt.Errorf("Wrong Extension: %v", extension)
	if extension == ".yaml" || extension == ".yml" {
		err = yaml.Unmarshal(examFile, &exam)
	}
	if extension == ".json" {
		err = json.Unmarshal(examFile, &exam)
	}
	ErrorCheck(err)
	return exam
}
