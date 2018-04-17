// roger-fetch-vault-token takes environment variables and returns an unwrapped Vault Token
// in a JSON to stdout for the vaultexec plugin to use or to stdout directly for
// wrapper scripts to use
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/channelmeter/vault-gatekeeper-mesos/gatekeeper"
)

type vaultExecConfig struct {
	Address string `json:"address"`
	Token   string `json:"token"`
	Path    string `json:"path"`
}

// Do a case insensitive search for MESOS_TASK_ID
func getTaskId(envVars []string) (string, error) {
	for _, envVar := range envVars {
		res := strings.SplitN(envVar, "=", 2)
		if strings.ToUpper(res[0]) == "MESOS_TASK_ID" {
			return res[1], nil
		}
	}
	return "", errors.New("mesos task id not found")
}

func main() {
	echoToken := flag.Bool(
		"echo-token",
		false,
		"echos unwrapped Vault token to stdout for use by wrapper scripts")
	flag.Parse()

	mesosTaskId, err := getTaskId(os.Environ())
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}

	token, err := gatekeeper.RequestVaultToken(mesosTaskId)
	if err != nil {
		log.Fatalf("ERROR: could not fetch token: %s\n", err)
	}

	if *echoToken == true {
		fmt.Printf(token)
	} else {
		vec := vaultExecConfig{Token: token}
		b, err := json.Marshal(vec)
		if err != nil {
			log.Fatalf("ERROR: conversion to JSON failed: %s\n", err)
		}
		fmt.Printf(string(b))
	}
}
