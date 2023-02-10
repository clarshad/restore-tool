package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	caas "github.com/clarshad/restore-tool/pkg/caas"
	apiclient "github.com/clarshad/restore-tool/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const caasDefaultNamespace = "caas-default"

func main() {
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

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("lenght: %v\n", len(pods.Items))

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	err = GetKubeCluster(dynamicClient)
	fmt.Println(err)

	dbCluster, err := apiclient.GetClusterDetails("dummyID", "dummySpaceID")
	if err != nil {
		fmt.Println(err)
	}

	// Build kubecluster resource
	_ = caas.CreateKubeCluster(dbCluster, caasDefaultNamespace)

	//TODO: get default management network- configmap from mgmt k3s cluster
	mgmtNetworks := map[string]string{}
	// Build kubemachine resource
	_ = caas.CreateKubeMachine(dbCluster, caasDefaultNamespace, mgmtNetworks)

	// Buid kubemachinetemplate resource
	_ = caas.CreateKubeMachineTemplate(dbCluster, caasDefaultNamespace)

	_ = caas.CreateKubeBootstrapConfig(dbCluster, caasDefaultNamespace)

}

func GetKubeCluster(client *dynamic.DynamicClient) error {

	kcRes := schema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha2", Resource: "kubeclusters"}
	list, err := client.Resource(kcRes).Namespace("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("failed to list kubecluster resource, error: %v\n", err)
		return err
	}

	for _, kc := range list.Items {
		fmt.Println(kc.Object)
	}

	return nil
}
