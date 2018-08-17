// roger-fetch-vault-token takes environment variables and returns an unwrapped Vault Token
// in a JSON to stdout for the vaultexec plugin to use or to stdout directly for
// wrapper scripts to use
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/channelmeter/vault-gatekeeper-mesos/gatekeeper"
)

type vaultExecConfig struct {
	Address string `json:"address"`
	Token   string `json:"token"`
	Path    string `json:"path"`
}

func main() {
	echoToken := flag.Bool(
		"echo-token",
		false,
		"echos unwrapped Vault token to stdout for use by wrapper scripts")
	flag.Parse()

	token, err := gatekeeper.EnvRequestVaultToken()
	if err != nil {
		log.Fatalf("could not fetch token: %s\n", err)
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
