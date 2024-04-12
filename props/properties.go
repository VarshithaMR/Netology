package props

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type Properties struct {
	Server ServerProps `yaml:"server"`
}

// ServerProps has the properties needed to create Server
type ServerProps struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	ContextRoot string `yaml:"contextroot"`
}

func ReadProperties(filename string) (*Properties, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Properties
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
