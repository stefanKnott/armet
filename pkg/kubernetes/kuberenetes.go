package kubernetes

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var CurrentContext string

func ParseContext(kubeconfig string) {
	yfile, err := ioutil.ReadFile(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})

	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}

	var ok bool
	CurrentContext, ok = data["current-context"].(string)
	if !ok {
		log.Fatal("string assert for current-context failed")
	}

	fmt.Printf("analyzing cluster: %s\n", CurrentContext)
}
