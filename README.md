# roger-gk-mesos

# Configuration

## Required Envionment Variables
* `VAULT_ADDR` - the URL to the Vault instance - (ie. `https://vault.right.url:8200`)
* `GATEKEEPER_ADDR` - the URL to the Gatekeeper instance - (ie. `https://gatekeeper.url.example:19201`)

Note: `MESOS_TASK_ID` is also used but will be set via Mesos

# Builds and Releases

## Builds
In order to build the roger-gk-mesos binaries, do the following:
```
git clone git@github.com:seomoz/roger-gk-mesos.git
cd roger-gk-mesos
docker build -t roger-gk-mesos-build .
docker run -it --rm -v $(pwd)/output:/output roger-gk-mesos-build
```

The roger-gk-mesos binaries will be in the $(pwd)/output folder on your machine
