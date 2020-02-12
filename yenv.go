package yenv // import "github.com/artbegolli/yenv"

import (
	"fmt"
	"os"
	"regexp"

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



func findAndReplace(m map[string]interface{}) {

	for _, v := range m {
		switch v := v.(type) {
		case int, float64:
		case string:
			fmt.Printf("%v\n", v)
		case []interface{}:
			for _, nv := range v {
				s, ok := nv.(string)
				fmt.Printf("%v\n", s)
				if !ok {
					findAndReplace(nv.(map[string]interface{}))
				}
			}
		case map[string]interface{}:
			findAndReplace(v)
		}
	}
	return
}

func matchEnvVariable(s string) (string, error) {
	pattern := `^\${[a-zA-Z0-9_.-]*}$`
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return "", err
	}

	if matched {
		return os.Getenv(s), nil
	}

	return s, nil
}
