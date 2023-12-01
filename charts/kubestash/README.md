# KubeStash

[KubeStash by AppsCode](https://github.com/stashed/kubestash) - Backup your Kubernetes native applications

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubestash --version=v2023.12.1
$ helm upgrade -i kubestash appscode/kubestash -n kubestash --create-namespace --version=v2023.12.1
```

## Introduction

This chart deploys Backup operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.14+

## Installing the Chart

To install/upgrade the chart with the release name `kubestash`:

```bash
$ helm upgrade -i kubestash appscode/kubestash -n kubestash --create-namespace --version=v2023.12.1
```

The command deploys Backup operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubestash`:

```bash
$ helm uninstall kubestash -n kubestash
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubestash` chart and their default values.

|        Parameter        |                                                                                                                                                                                  Description                                                                                                                                                                                   |      Default       |
|-------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| global.license          | License for the product. Get a license by following the steps from [here](https://kubestash.com/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubestash \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubestash \` <br> `--set global.license=<license file content>` | <code>""</code>    |
| global.registry         | Docker registry used to pull Stash operator image                                                                                                                                                                                                                                                                                                                              | <code>""</code>    |
| global.registryFQDN     | Docker registry fqdn used to pull Stash related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                         | <code>""</code>    |
| global.imagePullSecrets | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubestash \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                           | <code>[]</code>    |
| global.skipCleaner      | Skip generating cleaner job YAML                                                                                                                                                                                                                                                                                                                                               | <code>false</code> |
| kubestash-operator      | Pass values to kubestash-operator chart                                                                                                                                                                                                                                                                                                                                        | <code>{}</code>    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubestash appscode/kubestash -n kubestash --create-namespace --version=v2023.12.1 --set global.registry=stashed
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubestash appscode/kubestash -n kubestash --create-namespace --version=v2023.12.1 --values values.yaml
```
