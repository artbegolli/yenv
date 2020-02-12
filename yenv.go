package yenv // import "github.com/artbegolli/yenv"

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
)

func ApplyEnvVariablesYAML(y []byte, o interface{}) error {

	yamlMap := map[string]interface{}{}
	if err := yaml.Unmarshal(y, &yamlMap); err != nil {
		return err
	}
	fmt.Println(yamlMap)
	findAndReplace(yamlMap)
	fmt.Println(yamlMap)

	return nil
}



func findAndReplace(m map[string]interface{}) error {

	for k, v := range m {
		switch v := v.(type) {
		case int, float64:
		case string:
			val, err := matchEnvVariable(v)
			m[k] = val
			if err != nil {
				return err
			}

		case []interface{}:
			for ak, nv := range v {
				s, ok := nv.(string)
				if !ok {
					if err := findAndReplace(nv.(map[string]interface{})); err != nil {
						return err
					}
				} else {
					val, err := matchEnvVariable(s)
					v[ak] = val
					if err != nil {
						return err
					}
				}
			}
		case map[string]interface{}:
			if err := findAndReplace(v); err != nil {
				return err
			}
		}
	}
	return nil
}

func matchEnvVariable(s string) (string, error) {
	pattern := `^\${[a-zA-Z0-9_.-]*}$`
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return "", err
	}

	if matched {
		trimmedVal := strings.Trim(s, "${}")
		return os.Getenv(trimmedVal), nil
	}

	return s, nil
}
