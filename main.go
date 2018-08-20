// roger-fetch-vault-token takes environment variables and returns an unwrapped Vault Token
// in a JSON to stdout for the vaultexec plugin to use or to stdout directly for
// wrapper scripts to use
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/channelmeter/vault-gatekeeper-mesos/gatekeeper"
)

type vaultExecConfig struct {
	Address string `json:"address"`
	Token   string `json:"token"`
	Path    string `json:"path"`
}

// Check for the presence of two environment variables
// Check Mesos first and then check for Kubernetes after
// These schedulers in our environment always set these variables
func determineScheduler() (string, error) {
	if os.Getenv("MESOS_TASK_ID") != "" {
		return "mesos", nil
	}
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		return "kubernetes", nil
	}
	return "", fmt.Errorf("could not determine scheduler based on environment variables")
}

// Read the Kubernetes service account JWT
// This is used to authenticate the container to Vault
func readJwtToken(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read jwt token")
	}

	return string(bytes.TrimSpace(data)), nil
}

// Uses the service account JWT and the VAULT_ROLE variable
// to login to Vault Kubernetes auth method and fetch a token
func k8sFetchToken() (string, error) {
	var vaultAddr string

	vaultAddr = os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		vaultAddr = "https://127.0.0.1:8200"
	}

	// This is the standard mount path for Vault Kubernetes auth method
	// This variable will need to be set in each container if different
	vaultK8SMountPath := os.Getenv("VAULT_K8S_MOUNT_PATH")
	if vaultK8SMountPath == "" {
		vaultK8SMountPath = "kubernetes"
	}

	role := os.Getenv("VAULT_ROLE")
	if role == "" {
		return "", fmt.Errorf("required environment variable missing: VAULT_ROLE")
	}

	// This is the default location of the service account JWT mount
	saPath := os.Getenv("SERVICE_ACCOUNT_PATH")
	if saPath == "" {
		saPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	}

	jwt, err := readJwtToken(saPath)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	url := vaultAddr + "/v1/auth/" + vaultK8SMountPath + "/login"
	postBody := strings.NewReader("{\"role\": \"" + role + "\", \"jwt\": \"" + jwt + "\"}")

	c := &http.Client{Timeout: 3 * time.Second}
	req, err := http.NewRequest("POST", url, postBody)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", fmt.Errorf("http request creation failed: %s", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", fmt.Errorf("http call failed: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http call failed: %s", resp.Status)
	}

	var s struct {
		Auth struct {
			ClientToken string `json:"client_token"`
		} `json:"auth"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return "", fmt.Errorf("failed to read body")
	}

	return s.Auth.ClientToken, nil
}

func mesosFetchToken() (string, error) {
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
		"echoes unwrapped Vault token to stdout for use by wrapper scripts")
	flag.Parse()

	scheduler, err := determineScheduler()
	if err != nil {
		log.Fatal(err)
	}

	var token string
	switch scheduler {
	case "mesos":
		token, err = mesosFetchToken()
		if err != nil {
			log.Fatal(err)
		}
	case "kubernetes":
		token, err = k8sFetchToken()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unsupported scheduler type: %s", scheduler)
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
