# Terrajet Exoscale Provider

`provider-jet-exoscale` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane-contrib/terrajet) code
generation tools and exposes XRM-conformant managed resources for the Exoscale API.

# Status

Warning: this is a WIP, currently for testing purposes and not yet in an alpha version :)

Only a limited number of resources are taken into account and not fully tested:
- instances
- sks cluster (without related nodepool nor security groups)
- nodepool
- security group
- security group rule
- DBaas

## Getting Started

0. Pre-requisites

You need:
- a kubernetes cluster (easy to set up using [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/))
- kubectl
- helm

1. Install crossplane

Crossplane can be installed with Helm using the command below:

```
kubectl create namespace crossplane-system

helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
```

The crossplane cli can be installed with the following command:

```
curl -sL https://raw.githubusercontent.com/crossplane/crossplane/release-1.5/install.sh | sh
```

2. Create a secret containing Exoscale creds

Set Exoscale Key and Secret as env variables:

```
export EXOSCALE_API_KEY=...
export EXOSCALE_API_SECRET=...
```

Create a secret containing those credentials:

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: exoscale-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "key": "${EXOSCALE_API_KEY}",
      "secret": "${EXOSCALE_API_SECRET}"
    }
EOF
```

3. Provider installation

Install the Exoscale provider by using the following command replacing:
- ARCH with arm64 or amd64

```
kubectl crossplane install provider lucj/provider-jet-exoscale-ARCH:latest
```

4. Create a provider configuration

```
cat <<EOF | kubectl apply -f -
apiVersion: exoscale.jet.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: exoscale
spec:
  credentials:
    source: Secret
    secretRef:
      name: exoscale-creds
      namespace: crossplane-system
      key: credentials
EOF
```

5. Create Exoscale resources:

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
    name: exoscale
EOF
```

From your Exoscale portal you will see a new compute instance

![portal](./picts/exoscale_compute_instance.png)

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
    name: exoscale
EOF
```

From your Exoscale portal you will see a new Postgres instance

![portal](./picts/exoscale_database.png)

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
    name: exoscale
EOF
```

From your Exoscale portal you will see a new SKS cluster

![portal](./picts/exoscale_sks_cluster.png)


The 3 resources above appear as managed resources, in a couple of seconds they are in synced (meaning
that the specification is in sync with the acutal resource existing in the cluster):

```
$ watch kubectl get managed 

Every 2.0s: kubectl get managed                                                                                                                     jupiter.local: Wed Dec  8 16:47:03 2021

NAME                                               READY   SYNCED   EXTERNAL-NAME                          AGE
instance.compute.exoscale.jet.crossplane.io/demo   True    True     89a679c2-f5ef-4d6e-852e-757b0840c6b0   3m49s

NAME                                              READY   SYNCED   EXTERNAL-NAME   AGE
database.database.exoscale.jet.crossplane.io/db   True    True     test            3m39s

NAME                                          READY   SYNCED   EXTERNAL-NAME                          AGE
cluster.sks.exoscale.jet.crossplane.io/demo   True    True     a144b2f4-0091-4090-be69-3cdcd6a097de   3m26s
```

Removing the resources from the cluster only requires the usual `kubectl delete`:

```
kubectl delete instance demo
kubectl delete database db
kubectl delete cluster demo
```

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
