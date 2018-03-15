// roger-fetch-vault-toke takes environment variables and returns an unwrapped Vault Token
// in a JSON to stdout for the vaultexec plugin to use or to stdout directly for
// wrapper scripts to use
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
	echoToken := flag.Bool("echo-token", false, "echos unwrapped Vault token to stdout for use by wrapper scripts")
	flag.Parse()
	token, err := gatekeeper.EnvRequestVaultToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not fetch token: %s\n", err)
		os.Exit(1)
	}
	if *echoToken == true {
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
