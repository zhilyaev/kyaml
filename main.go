package main

import (
	"fmt"
	"log"
	"os"

	"github.com/imdario/mergo"
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
		mergo.Merge(&deployment.ObjectMeta, Cfg.Metadata)

		fmt.Println(deployment.Annotations)
		fmt.Println(deployment.Name)

	}

}
