# Bootstrap

* [Argo CD Documentation](https://argo-cd.readthedocs.io/en/stable/)
* [Argo CD Helm Chart](https://artifacthub.io/packages/helm/argo/argo-cd)
* [Argo CD Apps Helm Chart](https://artifacthub.io/packages/helm/argo/argocd-apps)

## Let Argo CD Manage Itself

Idea: only require an inital `helm install` to install Argo CD on a new cluster.
Everything else, including Argo CD itself, will from this point on be managed through GitOps!
For this, we package the `argo-cd` chart and the `argocd-apps` chart in an [umbrella chart](https://helm.sh/docs/howto/charts_tips_and_tricks/#complex-charts-with-many-dependencies).

```sh
## build dependencies
helm dependency build .

## install umbrella chart
helm upgrade -i argo-cd . \
  -f values.yaml \
  -f argocd-values.yaml \
  -f argocd-apps-values.yaml \
  -n argocd \
  --create-namespace

## expose dashboard (we will later enable ingress, but so far we got none)
kubectl port-forward service/argo-cd-argocd-server -n argocd 8080:443

## retrieve initial admin pw (should be deleted!!!)
## also: remove default project
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```
