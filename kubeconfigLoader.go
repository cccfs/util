package util

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"os/exec"
	"path/filepath"
)

func Kubeconfig() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		if ok, _ := Exists(filepath.Join(home, ".kube", "config")); ok {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			outfile, err := os.Create("/tmp/.config")
			if err != nil {
				panic(err)
			}
			defer outfile.Close()
			cmd := exec.Command("bash", "-c", "kubectl config view --merge --flatten")
			cmd.Stdout = outfile
			err = cmd.Start(); if err != nil {
				panic(err)
			}
			cmd.Wait()
			kubeconfig = flag.String("kubeconfig", filepath.Join("/tmp/.config"), "using $KUBECONFIG variables to the kubeconfig file")
		}
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
	return clientset
}
