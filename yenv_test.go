package yenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

var testObject = Object{
	ApiVersion: "",
	Kind:       "",
	Metadata:   Metadata{},
	Spec:       Spec{},
}

func TestUnmarshallWithEnv(t *testing.T) {

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
	if err := UnmarshallWithEnv(yaml, &emptyObj); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emptyObj)
}

func TestFindAndReplace(t *testing.T) {

	err := os.Setenv("META_LABEL", "meta-label")
	assert.Equal(t, nil, err)
	err = os.Setenv("APP_LABEL", "arts-app")
	assert.Equal(t, nil, err)
	err = os.Setenv("PORT", "1231")
	assert.Equal(t, nil, err)
	err = os.Setenv("CONT_NAME", "arts-container")
	assert.Equal(t, nil, err)

	yamlFile, err := ioutil.ReadFile("./resources/test.yaml")
	assert.Equal(t, nil, err)

	yamlMap := map[string]interface{}{}
	err = yaml.Unmarshal(yamlFile, &yamlMap)
	assert.Equal(t, nil, err)

	fmt.Println("BEFORE:: ", yamlMap)
	err = findAndReplace(yamlMap)
	assert.Equal(t, nil, err)

	fmt.Println("AFTER:: ", yamlMap)

}

func TestMatchEnvVariable(t *testing.T) {
	err := os.Setenv("TEST_ENV1", "replaced-test-1")
	assert.Equal(t, nil, err)
	err = os.Setenv("TEST_ENV2", "replaced-test-2")
	assert.Equal(t, nil, err)

	actual1 := matchEnvVariable("${TEST_ENV1}")
	assert.Equal(t, "replaced-test-1", actual1)

	actual2 := matchEnvVariable("FROM ${TEST_ENV1} ${TEST_ENV2}")
	assert.Equal(t, "FROM replaced-test-1 replaced-test-2", actual2)

	actual3 := matchEnvVariable("this is not an env")
	assert.Equal(t, "this is not an env", actual3)
}

type Object struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec Spec `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
	Labels []string
}

type Spec struct {
	Replicas string `json:"replicas"`
	Template Template `json:"template"`
	Spec ContainerSpec `json:"spec"`
}

type Template struct {
	Metadata Metadata2 `json:"metadata"`
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

type Metadata2 struct {
	Labels map[string]string `json:"labels"`
}
