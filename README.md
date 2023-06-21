# GitOps Demo

* [GitOps](https://www.gitops.tech/)
* [GitLab: What is GitOps?](https://about.gitlab.com/topics/gitops/)
* [Weaveworks: Awesome GitOps](https://github.com/weaveworks/awesome-gitops)

## GitOps with Argo CD

* [Argo CD Documentation](https://argo-cd.readthedocs.io/en/stable/)
* [Argo CD Helm Chart](https://artifacthub.io/packages/helm/argo/argo-cd)
* [Argo CD Apps Helm Chart](https://artifacthub.io/packages/helm/argo/argocd-apps)

Idea: only require an inital `helm install` to install Argo CD on a new cluster.
Everything else, including Argo CD itself, will from this point on be managed through GitOps!
For this, we package the `argo-cd` chart and the `argocd-apps` chart in an [umbrella chart](https://helm.sh/docs/howto/charts_tips_and_tricks/#complex-charts-with-many-dependencies).

```sh
## change to argocd dir
cd argocd

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

## retrieve initial admin pw
## NOTE: the initial admin pw should be deleted
## NOTE: the default project permissions should be restricted as much as possible
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

## Slideshow

* [markdown-slides](https://gitlab.com/da_doomer/markdown-slides)

```sh
## Markdown-slides works with Python >= 3.8.
python -m pip install git+https://gitlab.com/da_doomer/markdown-slides.git

## render slides
mdslides ./presentation/slides.md --include presentation

## start web server with slides on local port 8000
python -m http.server -d slides -p 8000
```

To change the theme go to `./slides/index.html` and set the href of `#theme` to one of the available ones under `./slides/dist/theme` (c.f. ).
