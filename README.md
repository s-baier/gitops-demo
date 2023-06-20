# GitOps Demo

* [GitOps](https://www.gitops.tech/)
* [GitLab: What is GitOps?](https://about.gitlab.com/topics/gitops/)
* [Codefresh: Get Certified for GitOps with Argo](https://learning.codefresh.io/)
* [Weaveworks: Awesome GitOps](https://github.com/weaveworks/awesome-gitops)

## ArgoCD

* https://artifacthub.io/packages/helm/argo/argo-cd
* https://artifacthub.io/packages/helm/argo/argocd-apps

## Slideshow

* [markdown-slides](https://gitlab.com/da_doomer/markdown-slides)

```sh
## Markdown-slides works with Python >= 3.8.
python -m pip install git+https://gitlab.com/da_doomer/markdown-slides.git

## render slides
mdslides ./slides.md --include presentation

## start web server with slides on local port 8000
python -m http.server -d slides -p 8000
```

To change the theme go to `./slides/index.html` and set the href of `#theme` to one of the available ones under `./slides/dist/theme` (c.f. ).
