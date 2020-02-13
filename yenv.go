package yenv // import "github.com/artbegolli/yenv"

import (
	"os"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
)

func UnmarshallWithEnv(y []byte, o interface{}) error {

	yamlReplaced, err := ApplyEnvValues(y)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlReplaced, o); err != nil {
		return err
	}
	return nil
}

func ApplyEnvValues(y []byte) ([]byte, error) {
	yamlMap := map[string]interface{}{}
	if err := yaml.Unmarshal(y, &yamlMap); err != nil {
		return nil, err
	}
	if err := findAndReplace(yamlMap); err != nil {
		return nil, err
	}

	yamlReplaced, err := yaml.Marshal(yamlMap)
	if err != nil {
		return nil, err
	}
	return yamlReplaced, nil
}

func findAndReplace(m map[string]interface{}) error {

	for k, v := range m {
		switch v := v.(type) {
		case int, float64:
		case string:
			val := matchEnvVariable(v)
			m[k] = val

		case []interface{}:
			for ak, nv := range v {
				s, ok := nv.(string)
				if !ok {
					if err := findAndReplace(nv.(map[string]interface{})); err != nil {
						return err
					}
				} else {
					val := matchEnvVariable(s)
					v[ak] = val
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

func matchEnvVariable(s string) string {
	pattern := `\${[a-zA-Z0-9_.-]*}`

	re := regexp.MustCompile(pattern)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		trimmedVal := strings.Trim(match, "${}")
		return os.Getenv(trimmedVal)
	})

}
