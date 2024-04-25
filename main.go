package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeyaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var magic map[string]any

var collections = []string{
	"deployments",
	"statefulsets",
}

func main() {
	// Read configuration
	file := "debug.yml"
	src, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read plan file %s: %w", file, err))
	}

	err = yaml.Unmarshal(src, &magic)
	if err != nil {
		log.Fatal(err)
	}

	// Merge magic
	for _, collectionName := range collections {
		collection, ok := magic[collectionName].(map[string]any)
		if ok {
			fmt.Println(collectionName, ":")
			for name, resource := range collection {
				fmt.Println("\t", name, "-", resource)
				merge(resource.(map[string]any))
			}
		} else {
			fmt.Println(collectionName, ": 404 not found")
		}
	}

	magicBytes, err := yaml.Marshal(&magic)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("=========    MAGIC   =============\n", string(magicBytes), "=========================\n")

	// Get true configuration
	err = kubeyaml.Unmarshal(magicBytes, &Cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Manipulation with configuration
	for name, deployment := range Cfg.Deployments {
		if deployment.ObjectMeta.Name == "" {
			deployment.ObjectMeta.Name = name.String()
		}
	}

	// Render configuration
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println("Creating deployment...")
	for _, deployment := range Cfg.Deployments {
		result, err := deploymentsClient.Create(context.TODO(), &deployment, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	}

}

func merge(fields map[string]any) {
	for name, value := range fields {
		v, found := value.(map[string]any)
		for key, _ := range magic {
			// TODO: extends here with aliases functionality
			if name == key && found {
				mergo.Map(&v, magic[key])
			}
		}

		if found {
			merge(value.(map[string]any))
		}
	}
}
