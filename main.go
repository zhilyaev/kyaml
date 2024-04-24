package main

import (
	"fmt"
	"log"
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"
)

var Cfg Config

func main() {
	file := "debug.yml"
	src, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read plan file %s: %w", file, err))
	}

	err = yaml.Unmarshal(src, &Cfg)
	if err != nil {
		log.Fatal(err)
	}

	for name, deployment := range Cfg.Deployments {
		fmt.Println(name)
		fmt.Println(deployment.Namespace)
		fmt.Println(deployment.Name)
	}

}
