# roger-fetch-vault-token

# Configuration

## Required Envionment Variables
* `VAULT_ADDR` - the URL to the Vault instance - (ie. `https://vault.right.url:8200`)
* `GATEKEEPER_ADDR` - the URL to the Gatekeeper instance - (ie. `https://gatekeeper.url.example:19201`)

Note: `MESOS_TASK_ID` is also used but will be set via Mesos

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

## Releases
NOTE: Haven't done this yet so steps are not great
* create git tag
* edit release notes
* upload binaries in output directory to Github as part of the release
