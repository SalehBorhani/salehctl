package cmd

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

type Object struct {
	Kind string `yaml:"kind"`
	Name string `yaml:"name"`
}

func loadYaml(yaml_file string) {
	var obj Object

	yamlFile, err := os.ReadFile(yaml_file)
	if err != nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("Failed to read the file...")
	}

	err = yaml.Unmarshal(yamlFile, &obj)
	if err != nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("Failed to load the file...")
	}
	fmt.Printf("%s is %s", obj.Name, obj.Kind)
}
