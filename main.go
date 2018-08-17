// roger-fetch-vault-token takes environment variables and returns an unwrapped Vault Token
// in a JSON to stdout for the vaultexec plugin to use or to stdout directly for
// wrapper scripts to use
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/channelmeter/vault-gatekeeper-mesos/gatekeeper"
)

type vaultExecConfig struct {
	Address string `json:"address"`
	Token   string `json:"token"`
	Path    string `json:"path"`
}

func determineScheduler() (string, error) {
	if os.Getenv("MESOS_TASK_ID") != "" {
		return "mesos", nil
	}

	return "", fmt.Errorf("could not determine scheduler based on environment variables")
}

func fetchToken() (string, error) {
	token, err := gatekeeper.EnvRequestVaultToken()
	if err != nil {
		return "", fmt.Errorf("could not fetch token: %s", err)
	}
	return token, nil
}

func main() {
	echoToken := flag.Bool(
		"echo-token",
		false,
		"echos unwrapped Vault token to stdout for use by wrapper scripts")
	flag.Parse()

	scheduler, err := determineScheduler()
	if err != nil {
		log.Fatal(err)
	}

	var token string
	switch scheduler {
	case "mesos":
		token, err = fetchToken()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("no supported scheduler found")
	}

	if *echoToken == true {
		fmt.Printf(token)
	} else {
		vec := vaultExecConfig{Token: token}
		b, err := json.Marshal(vec)
		if err != nil {
			log.Fatalf("conversion to JSON failed: %s\n", err)
		}
		fmt.Printf(string(b))
	}
}
