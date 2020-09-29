package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config stores the settings read from YAML file
type Config struct {
	Sections []struct {
		Name       string   `yaml:"name"`
		Qualifiers []string `yaml:"qualifiers"`
	} `yaml:"sections"`
	Replace struct {
		Text []struct {
			From string `yaml:"from"`
			To   string `yaml:"to"`
		} `yaml:"text"`
		Link []struct {
			From string `yaml:"from"`
			To   string `yaml:"to"`
		} `yaml:"link"`
	} `yaml:"replace"`
}

// Validate the Config loaded from YAML
func (c *Config) Validate() {

}

func (c *Config) addDefault() {
	c.Sections = append(c.Sections,
		struct {
			Name       string   `yaml:"name"`
			Qualifiers []string `yaml:"qualifiers"`
		}{
			"Default",
			[]string{},
		})
}

// GenerateFromYML generates a config file from filePath provided
func GenerateFromYML(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
