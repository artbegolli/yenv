package yenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testObject = Object{
	ApiVersion: "",
	Kind:       "",
	Metadata:   Metadata{},
	Spec:       Spec{},
}

func TestApplyEnvVariablesYAML(t *testing.T) {

	err := os.Setenv("META_LABEL", "meta-label")
	assert.Equal(t, nil, err)
	err = os.Setenv("APP_LABEL", "arts-app")
	assert.Equal(t, nil, err)
	err = os.Setenv("PORT", "1231")
	assert.Equal(t, nil, err)
	err = os.Setenv("CONT_NAME", "arts-container")
	assert.Equal(t, nil, err)

	yaml, err := ioutil.ReadFile("./resources/test.yaml")
	assert.Equal(t, nil, err)

	emptyObj := Object{}
	if err := ApplyEnvVariablesYAML(yaml, emptyObj); err != nil {
		fmt.Println(err)
		return
	}
}

func TestApplyEnvVariablesMarshalled(t *testing.T) {


}

type Object struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec Spec `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
	Labels map[string]string
}

type Spec struct {
	Replicas string `json:"replicas"`
	Template Template `json:"template"`
	Spec ContainerSpec `json:"spec"`
}

type Template struct {
	Metadata Metadata `json:"metadata"`
}

type ContainerSpec struct {
	Containers []Container `json:"containers"`
}

type Container struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Ports []Port `json:"ports"`
}

type Port struct {
	ContainerPort string `json:"containerPort"`
}
