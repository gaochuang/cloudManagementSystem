package main

import (
	"kubernetes_management_system/pkg/client"
	"kubernetes_management_system/pkg/signals"
)

func main() {

	stopCh := signals.SetupSignalHandler()
	client := client.NewClient()
}
