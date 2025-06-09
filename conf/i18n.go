package conf

import (
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Dictionary Dictionary
var Dictionary *map[interface{}]interface{}

// LoadLocales Load internationalization file
func LoadLocales(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}

	Dictionary = &m

	return nil
}

// T Translate
func T(key string) string {
	dic := *Dictionary
	keys := strings.Split(key, ".")
	for index, path := range keys {
		// If we've reached the last level, look for the target translation
		if len(keys) == (index + 1) {
			for k, v := range dic {
				if k, ok := k.(string); ok {
					if k == path {
						if value, ok := v.(string); ok {
							return value
						}
					}
				}
			}
			return path
		}
		// If there's another level, continue searching
		for k, v := range dic {
			if ks, ok := k.(string); ok {
				if ks == path {
					if dic, ok = v.(map[interface{}]interface{}); ok == false {
						return path
					}
				}
			} else {
				return ""
			}
		}
	}

	return ""
}
