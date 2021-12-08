# Terrajet Exoscale Provider

`provider-jet-exoscale` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane-contrib/terrajet) code
generation tools and exposes XRM-conformant managed resources for the Exoscale API.

# Status

Warning: this is a WIP, currently for testing purposes and not yet in an alpha version :)

Only a limited number of resources are taken into account and not fully tested:
- instances
- sks cluster (without nodepool nor security groups)
- DBaas

## Getting Started

0. Pre-requisites

You need:
- a kubernetes cluster (easy to set up using [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/))
- kubectl
- helm

- Install crossplane
```
kubectl create namespace crossplane-system

helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
```

- Install crossplane cli
```
curl -sL https://raw.githubusercontent.com/crossplane/crossplane/release-1.5/install.sh | sh
```

1. Create a secret containing Exoscale creds

Note: envsubst is used in the following commands to replace env variables within the secret.yaml.tmpl
file. If you do not have envsubst, you can create the secret after having replaced the env var manually.

```
export EXOSCALE_API_KEY=...
export EXOSCALE_API_SECRET=...
cat examples/providerconfig/secret.yaml.tmpl| envsubst | kubectl apply -f -
```

2. Provider installation

Install the Exoscale provider by using the following command

```
kubectl crossplane install provider lucj/provider-jet-exoscale-amd64:v0.0.0-23.g16f4e86
```

Note: arm64 will be available soon

3. Create Exoscale resources:

- Compute instance

```
cat <<EOF | kubectl apply -f -
apiVersion: compute.exoscale.jet.crossplane.io/v1alpha1
kind: Instance
metadata:
  name: demo
spec:
  forProvider:
    displayName: "demo"
    zone: "de-fra-1"
    diskSize: 30
    template: "Linux Ubuntu 20.04 LTS 64-bit"
    securityGroups: ["default"]
  providerConfigRef:
    name: default
EOF
```

From your Exoscale portal you will see a new compute instance

![portal](./images/exoscale_compute.png)

- DBaas

```
cat <<EOF | kubectl apply -f -
apiVersion: database.exoscale.jet.crossplane.io/v1alpha1
kind: Database
metadata:
  name: db
spec:
  forProvider:
    zone: "de-fra-1"
    name: "test"
    type: "pg"
    plan: "startup-4"
    terminationProtection: false
  providerConfigRef:
    name: default
EOF
```

From your Exoscale portal you will see a new Postgres instance

![portal](./images/exoscale_database.png)

- SKS cluster

```
cat <<EOF | kubectl apply -f -
apiVersion: sks.exoscale.jet.crossplane.io/v1alpha1
kind: Cluster
metadata:
  name: demo
spec:
  forProvider:
    description: "Managed with Crossplane Exoscale Provider (generated with Terrajet)"
    zone: "de-fra-1"
    name: "demo"
  providerConfigRef:
    name: default
EOF
```

From your Exoscale portal you will see a new SKS cluster

![portal](./images/exoscale_sks_cluster.png)


## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/lucj/provider-jet-exoscale/issues).

## Governance and Owners

provider-jet-exoscale is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-exoscale adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-exoscale is under the Apache 2.0 license.
