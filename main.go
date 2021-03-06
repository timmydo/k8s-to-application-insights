package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
)

var (
	aikey     = flag.String("aikey", os.Getenv("AIKEY"), "application insights instrumentation key")
	cluster   = flag.String("cluster", os.Getenv("MONITOR_CLUSTER"), "deployment cluster name")
	namespace = flag.String("namespace", os.Getenv("MONITOR_NAMESPACE"), "deployment namespace")
	delay     = flag.Duration("delay", 10*time.Second, "delay between reporting")
)

func main() {
	flag.Parse()

	if *aikey == "" {
		log.Fatalln("No instrumentation key provided.")
	}

	if *namespace == "" {
		log.Fatalln("No namespace provided")
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	client := appinsights.NewTelemetryClient(*aikey)

	appinsights.NewDiagnosticsMessageListener(func(msg string) error {
		log.Printf("%s\n", msg)
		return nil
	})

	for {
		deploy, err := clientset.AppsV1().Deployments(*namespace).List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		for _, deployment := range deploy.Items {
			metricName := fmt.Sprintf("deploymentPercentAvailable_%s_%s_%s", *cluster, deployment.Namespace, deployment.Name)
			track(client, metricName, float64(deployment.Status.AvailableReplicas)/float64(deployment.Status.Replicas))
		}

		pods, err := clientset.CoreV1().Pods(*namespace).List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		for _, pod := range pods.Items {
			for idx, podStatus := range pod.Status.ContainerStatuses {
				metricName := fmt.Sprintf("podRestartCount_%s_%s_%s_%d", *cluster, pod.Namespace, pod.Name, idx)
				track(client, metricName, float64(podStatus.RestartCount))
			}
		}

		time.Sleep(*delay)
	}
}

func track(client appinsights.TelemetryClient, metricName string, val float64) {
	client.Track(appinsights.NewMetricTelemetry(metricName, val))
	log.Printf("%s=%f\n", metricName, val)
}
