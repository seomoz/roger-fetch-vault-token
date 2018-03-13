// roger-gk-mesos takes environment variables and returns JSON with an unwrapped Vault Token
// to stdout for the vaultexec plugin to use
package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/json"

	"github.com/channelmeter/vault-gatekeeper-mesos/gatekeeper"
)

type VaultExecConfig struct {
	Address string `json:"address"`
	Token	string `json:"token"`
	Path	string `json:"path"`
}

func main() {
	setVaultEnv := flag.Bool("env", false, "set VAULT_TOKEN environment variable instead of JSON output")
	flag.Parse()
	token, err := gatekeeper.EnvRequestVaultToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not fetch token: %s\n", err)
		os.Exit(1)
	}
	if *setVaultEnv == true {
		fmt.Printf(token)
	} else {
		vec := VaultExecConfig{Token: token}
		b, err := json.Marshal(vec)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: conversion to JSON failed: %s\n", err)
			os.Exit(1)
		}
		os.Stdout.Write(b)
	}
}
