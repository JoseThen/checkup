package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YamlExam struct {
	Name     string
	Endpoint string
	Tests    []struct {
		Code  int
		Paths []string
	}
}

func ReadYaml(file string) YamlExam {
	var exam YamlExam
	yamlFile, err := ioutil.ReadFile(file)
	ErrorCheck(err)
	err = yaml.Unmarshal(yamlFile, &exam)
	ErrorCheck(err)
	return exam
}
