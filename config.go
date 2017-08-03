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
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"
	flag "github.com/spf13/pflag"
)

type config struct {
	flagSet     *flag.FlagSet
	kubeconfig  *string
	proxy       *string
	serviceName *string
	version     *bool
	healthCheck *bool
}

func newConfig() *config {
	return &config{
		kubeconfig:  flag.String("kubeconfig", "", "absolute path to the kubeconfig file"),
		proxy:       flag.String("proxy", "", "kubctl proxy server running at the given url"),
		serviceName: flag.String("service-name", "", "API Service name"),
		version:     flag.Bool("version", false, "display version info and exit"),
		healthCheck: flag.Bool("health-check", true, "enable health checking for API service"),
	}
}

func (cfg *config) String() string {
	return fmt.Sprintf("kubeconfig: %s, proxy: %s, service-name: %s, health-check: %t, version: %t",
		*cfg.kubeconfig, *cfg.proxy, *cfg.serviceName, *cfg.healthCheck, *cfg.version)
}

var envSupport = map[string]bool{
	"kubeconfig":   true,
	"proxy":        true,
	"service-name": true,
	"version":      false,
	"health-check": true,
}

func variableName(name string) string {
	return "KRAK8S_" + strings.ToUpper(strings.Replace(name, "-", "_", -1))
}

// Just like Flags.Parse() except we try to get recognized values for the valid
// set of flags from environment variables.  We choose to use the environment
// value if 1) the value hasen't already been set as command line flags and the
// flas is a member of the supported set (see map defined above).
func (cfg *config) envParse() error {
	var err error

	alreadySet := make(map[string]bool)
	cfg.flagSet.Visit(func(f *flag.Flag) {
		if envSupport[f.Name] {
			alreadySet[variableName(f.Name)] = true
		}
	})

	usedEnvKey := make(map[string]bool)
	cfg.flagSet.VisitAll(func(f *flag.Flag) {
		if envSupport[f.Name] {
			key := variableName(f.Name)
			if !alreadySet[key] {
				val := os.Getenv(key)
				if val != "" {
					usedEnvKey[key] = true
					if serr := cfg.flagSet.Set(f.Name, val); serr != nil {
						err = fmt.Errorf("invalid value %q for %s: %v", val, key, serr)
					}
					glog.V(3).Infof("recognized and used environment variable %s=%s", key, val)
				}
			}
		}
	})

	for _, env := range os.Environ() {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) != 2 {
			glog.Warningf("found invalid env %s", env)
		}
		if usedEnvKey[kv[0]] {
			continue
		}
		if alreadySet[kv[0]] {
			glog.V(3).Infof("recognized environment variable %s, but unused: superseeded by command line flag ", kv[0])
			continue
		}
		if strings.HasPrefix(env, "KRAK8S_") {
			glog.Warningf("unrecognized environment variable %s", env)
		}
	}

	return err
}