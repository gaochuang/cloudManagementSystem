package cluster

import (
	"context"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNodeEvent(client *kubernetes.Clientset, namespace string, field string) (*v1.EventList, error) {
	events, err := client.CoreV1().Events(namespace).List(context.TODO(), v12.ListOptions{
		FieldSelector: field,
	})

	if err != nil {
		return nil, err
	}
	return events, nil
}
