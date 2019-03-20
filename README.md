# roger-fetch-vault-token

Binary to fetch Vault tokens for applications. Currently support two modes:
Kubernetes Vault auth
Vault-Gatekeeper-Mesos

This is primarily a plugin to [vaultexec](https://github.com/funnylookinhat/vaultexec)

## Configuration

### Required Envionment Variables

* `VAULT_ADDR` - the URL to the Vault instance - (ie. `https://vault.right.url:8200`)

#### Vault-Gatekeeper-Mesos Required Variables

* `GATEKEEPER_ADDR` - the URL to the Gatekeeper instance - (ie. `https://gatekeeper.url.example:19201`)

Note: `MESOS_TASK_ID` is also used but will be set via Mesos

#### Kubernetes Required Variables

* `VAULT_ROLE` - the role used by the application in Vault - (ie. `k8s-test-role`)

Note: `KUBERNETES_SERVICE_HOST` is also used to determine if Kubernetes is used but will be set by Kubnernetes

## Usage

This can be used as a standalone binary in a wrapper script inside a Mesos container
with an `--echo-token` flag or as a plugin to [vaultexec](https://github.com/funnylookinhat/vaultexec)

## Builds and Releases

### Dependencies

Note that this binary uses the vault-gatekeeper-mesos library at a very specific version.
BE CAREFUL RUNNING `go get -u` AS TO NOT UPGRADE UNEXPECTEDLY

The current version is `go get github.com/channelmeter/vault-gatekeeper-mesos/gatekeeper@0.6.0`

### Builds

In order to build the roger-fetch-vault-token binaries, do the following:

```text
git clone git@github.com:seomoz/roger-fetch-vault-token.git
cd roger-fetch-vault-token
docker build -t roger-fetch-vault-token-build .
docker run -it --rm -v $(pwd)/output:/output roger-fetch-vault-token-build
```

The roger-fetch-vault-token binaries will be in the $(pwd)/output folder on your machine
