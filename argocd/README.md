# ArgoCD

> Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes.
>
> [Argo CD Documentation](https://argo-cd.readthedocs.io/en/stable/)

This [umbrella chart](https://helm.sh/docs/howto/charts_tips_and_tricks/#complex-charts-with-many-dependencies) packages the [Argo CD Helm Chart](https://artifacthub.io/packages/helm/argo/argo-cd) as well as the [Argo CD Apps Helm Chart](https://artifacthub.io/packages/helm/argo/argocd-apps) together, to enable an easy setup of a self-managed Argo CD instance, including additional AppProjects, Applications and ApplicationSets.

## TL;DR

```sh
## build dependencies
helm dependency build .

## install umbrella chart
helm install my-release .
```

## Prerequisites

* Kubernetes 1.19+
* Helm 3.2.0+
* PV provisioner support in the underlying infrastructure

## Installing the Chart

To install the chart with the release name `argo-cd`:

```sh
## build dependencies
helm dependency build .

## install umbrella chart
helm install argo-cd . \
  -f values.yaml \
  -f argocd-values.yaml \
  -f argocd-apps-values.yaml \
  -n argocd \
  --create-namespace
```

The command deploys argocd on the Kubernetes cluster with configuration for the chart itself as well as its subcharts.
For a full list of chart parameters of the subcharts, see:

* [argo-cd parameters](https://github.com/argoproj/argo-helm/blob/main/charts/argo-cd/values.yaml)
* [argocd-apps parameters](https://github.com/argoproj/argo-helm/blob/main/charts/argocd-apps/values.yaml)

After the installation, you can access the Argo CD dashboard via port-forwarding, if no Ingress is configured:

```sh
## expose dashboard (we will later enable ingress, but so far we got none)
kubectl port-forward service/argo-cd-argocd-server -n argocd 8080:443

## retrieve initial admin pw
## NOTE: the initial admin pw should be deleted
## NOTE: the default project permissions should be restricted as much as possible
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

## Uninstalling the Chart

To uninstall/delete the `argo-cd` deployment:

```sh
helm delete argo-cd
```

The command removes all the Kubernetes components associated with the chart and deletes the release.
