/*
Copyright 2017 Samsung SDSA CNCT

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"time"

	"k8s.io/client-go/kubernetes"
)

var (
	resyncPeriod = 30 * time.Second
)

// API Server
type apiServer struct {
	// Command line / environment supplied configuration values
	cfg *config

	clientset *kubernetes.Clientset

	stopCh chan struct{}
}

func newAPIServer(clientset *kubernetes.Clientset, cfg *config) *apiServer {
	// Create and start the http router

	// create api server controller struct
	as := apiServer{
		clientset: clientset,
		stopCh:    make(chan struct{}),
		cfg:       cfg,
	}

	return &as
}

func (as *apiServer) run() {
	// run the controller and queue goroutines
	// go as.apiServer.Run(as.stopCh)
	// Allow time for the initial startup
	time.Sleep(5 * time.Second)
}