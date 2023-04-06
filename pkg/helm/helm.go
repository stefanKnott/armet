package helm

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	b64 "encoding/base64"
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var charts map[string]*HelmRelease
var chartLock *sync.Mutex

func GetReleaseSlice() []HelmRelease {
	cs := make([]HelmRelease, 0)
	chartLock.Lock()
	for _, hr := range charts {
		cs = append(cs, *hr)
	}
	chartLock.Unlock()

	return cs
}

func GetReleaseByNamespaceMap() map[string][]HelmRelease {
	cs := make(map[string][]HelmRelease)
	chartLock.Lock()
	for _, hr := range charts {
		if cs[hr.Namespace] == nil {
			cs[hr.Namespace] = make([]HelmRelease, 0)
		}
		cs[hr.Namespace] = append(cs[hr.Namespace], *hr)
	}
	chartLock.Unlock()

	return cs
}

//TODO: here we could take a list of kubeconfig files, create loop per config
func BeginLoop(kubeconfig string) {
	charts = make(map[string]*HelmRelease)
	chartLock = new(sync.Mutex)
	fmt.Printf("beginning loop using config at file: %s\n", kubeconfig)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		// get all helm secrets
		secrets, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{
			LabelSelector: "owner=helm",
		})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d helm secrets in the cluster\n", len(secrets.Items))

		for _, item := range secrets.Items {
			chartData := item.Data
			releaseEncodedCompressed := string(chartData["release"])
			releaseCompressed, _ := b64.StdEncoding.DecodeString(releaseEncodedCompressed)
			reader := bytes.NewReader([]byte(releaseCompressed))
			gzreader, e1 := gzip.NewReader(reader)
			if e1 != nil {
				fmt.Println(e1) // Maybe panic here, depends on your error handling.
			}

			output, e2 := ioutil.ReadAll(gzreader)
			if e2 != nil {
				fmt.Println(e2)
			}

			fmt.Printf("OUTPUT POST DECODE AND DECOMPRESS: %+v\n", string(output))
			hr := new(HelmRelease)
			err = json.Unmarshal(output, &hr)
			if err != nil {
				panic(err)
			}

			chartLock.Lock()
			charts[hr.Name] = hr
			chartLock.Unlock()
		}

		for _, chart := range charts {
			fmt.Println("iter thru charts")
			fmt.Printf("%+v\n", chart, chart.Info)
		}

		// parse helm secrets

		time.Sleep(10 * time.Second)
	}
}
