# roger-fetch-vault-token
The purpose of roger-fetch-vault-token is interact with vault-mesos-gatekeeper and fetch
a Vault token for use in a Mesos task.

# Configuration

## Required Envionment Variables
* `VAULT_ADDR` - the URL to the Vault instance - (ie. `https://vault.right.url:8200`)
* `GATEKEEPER_ADDR` - the URL to the Gatekeeper instance - (ie. `https://gatekeeper.url.example:19201`)

Note: `MESOS_TASK_ID` is also used but will be set via Mesos

# Usage
This can be used as a standalone binary in a wrapper script inside a Mesos container
with an `--echo-token` flag or as a plugin to [vaultexec](https://github.com/funnylookinhat/vaultexec)

# Builds and Releases

## Builds
In order to build the roger-fetch-vault-token binaries, do the following:
```
git clone git@github.com:seomoz/roger-fetch-vault-token.git
cd roger-fetch-vault-token
docker build -t roger-fetch-vault-token-build .
docker run -it --rm -v $(pwd)/output:/output roger-fetch-vault-token-build
```

The roger-fetch-vault-token binaries will be in the $(pwd)/output folder on your machine

