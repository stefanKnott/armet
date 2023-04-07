package helm

import (
	"bytes"
	"compress/gzip"
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	k "github.com/stefanKnott/armet/pkg/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// map[chartName]chart
var charts map[string]*HelmRelease

// map[cluster][namespace][chartname]chart
var chartsByCluster map[string]map[string]map[string]*HelmRelease

var chartLock *sync.Mutex

func GetReleases() map[string]map[string][]HelmRelease {
	ret := make(map[string]map[string][]HelmRelease)
	chartLock.Lock()
	for cluster, namespaces := range chartsByCluster {
		if ret[cluster] == nil {
			ret[cluster] = make(map[string][]HelmRelease)
		}
		for namespace, charts := range namespaces {
			if ret[cluster][namespace] == nil {
				ret[cluster][namespace] = make([]HelmRelease, 0)
			}
			for _, chart := range charts {
				ret[cluster][namespace] = append(ret[cluster][namespace], *chart)
			}
		}
	}
	defer chartLock.Unlock()

	return ret
}

func GetReleaseByNamespaceMap(cluster string) map[string][]HelmRelease {
	cs := make(map[string][]HelmRelease)
	chartLock.Lock()
	for _, ns := range chartsByCluster[cluster] {
		for _, ch := range ns {
			if cs[ch.Namespace] == nil {
				cs[ch.Namespace] = make([]HelmRelease, 0)
			}
			cs[ch.Namespace] = append(cs[ch.Namespace], *ch)
		}
	}
	chartLock.Unlock()

	return cs
}

func BeginLoop(pathToConfigs string) {
	charts = make(map[string]*HelmRelease)
	chartsByCluster = make(map[string]map[string]map[string]*HelmRelease)
	chartLock = new(sync.Mutex)

	filepath.Walk(pathToConfigs, func(pathToConfigs string, info os.FileInfo, err error) error {
		fmt.Printf("beginning loop using config at file: %s\n", info.Name())
		if info.IsDir() {
			return nil
		}
		path := pathToConfigs
		go func(path string) {
			clusterName := k.ParseContext(path)

			// use the current context in kubeconfig
			config, err := clientcmd.BuildConfigFromFlags("", path)
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

					hr := new(HelmRelease)
					err = json.Unmarshal(output, &hr)
					if err != nil {
						panic(err)
					}

					chartLock.Lock()
					charts[hr.Name] = hr

					if chartsByCluster[clusterName] == nil {
						chartsByCluster[clusterName] = make(map[string]map[string]*HelmRelease)
					}
					if chartsByCluster[clusterName][hr.Namespace] == nil {
						chartsByCluster[clusterName][hr.Namespace] = make(map[string]*HelmRelease)
					}
					chartsByCluster[clusterName][hr.Namespace][hr.Name] = hr
					chartLock.Unlock()

				}

				fmt.Printf("chartsByCluster: %+v\n", chartsByCluster)
				time.Sleep(10 * time.Second)
			}
		}(path)

		return nil
	})
}
