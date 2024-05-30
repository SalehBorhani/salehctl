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

func loadYaml(yamlFile string) {
	var obj Object

	file, err := os.ReadFile(yamlFile)
	if err != nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("Failed to read the file...")
	}

	err = yaml.Unmarshal(file, &obj)
	if err != nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("Failed to load the file...")
	}
	fmt.Printf("%s is %s\n", obj.Name, obj.Kind)
}
