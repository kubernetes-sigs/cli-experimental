---
title: "Documentation"
linkTitle: "Documentation"
type: docs
weight: 3
description: >
    How to make contributions to this site
---

This site uses [Docsy](https://www.docsy.dev) and was
forked from the [docsy-example](https://github.com/google/docsy-example)

## Prerequisites

- [Install hugo](https://gohugo.io/getting-started/installing/#fetch-from-github)
- Clone kustomize
  - `git clone git@github.com:kubernetes-sigs/cli-experimental && cd site/`

## Development

The doc input files are in the `site` directory.  The site can be hosted locally using
`hugo server`.

```shell script
cd site/
npm install
npm install -g postcss-cli
npm install autoprefixer
npm audit fix
hugo server
```

```shell script
...
Running in Fast Render Mode. For full rebuilds on change: hugo server --disableFastRender
Web Server is available at http://localhost:1313/ (bind address 127.0.0.1)
```

## Publishing

Hugo compiles the files under `site` Hugo into html which it puts in the `docs` folder:

```shell script
cd site/
hugo
```

```shell script
                   | EN  
-------------------+-----
  Pages            | 99  
  Paginator pages  |  0  
  Non-page files   |  0  
  Static files     | 47  
  Processed images |  0  
  Aliases          |  2  
  Sitemaps         |  1  
  Cleaned          |  0  
```

Add the `site/` and `docs/` folders to a commit, then create a PR.

## Publishing docs in forked repository

We use **Netlify** to publish changes in the site. You can also enable netlify on you're forked repo
by doing the following step.

- Log into Netlify using your Github Credentials.
- Click **New Site from Git** button in the Netlify Dashboard.
- The setup has 3 steps.
  - Connect to Git Provider - Select Github here and authenticate your Github account if not done earlier.
  - Pick a repository - Select the forked repository here.
  - Build options, and deploy! - Here set **Branch to deploy** to the branch that has the latest changes also set **Publish directory** to `./docs`.

![Netlify Setup Image][setup]

## Raising a PR for changes in the site

- Once deployed, you'll have a URL pointing to the newly deployed site. Submit the URL along with the PR.
- Make sure your changes are working as expected in the newly received netlify URL before PR.

![Netlify Deployed Image][deploy]


## Setting Custom Domain & DNS changes

{{< alert color="success" title="Note" >}}
This is applicable only for the site adminisrators on the event of site migration.
{{< /alert >}}
- Make sure you're a part of **Kubernetes Docs** Netlify team.
- Under **Site Settings** you'll find **Domain Management**, where in you can set the site's custom domain.
- Ideally, it should match up with the wild card `*.k8s.io`
- Once custom domains are set on Netlify, you can raise a PR in [k8s.io](https://github.com/kubernetes/k8s.io) github repository.
- You'll have to add this snippet in `dns/zone-configs/k8s.io._0_base.yaml` file:
```yaml
# <github repo url> (@maintainers)
<custom_name>:
  type: CNAME
  value: <current_netlify_url>.
```
{{< alert color="success" title="Useful Links" >}}
- Subproject Site Requests: https://github.com/kubernetes/community/blob/master/github-management/subproject-site-requests.md.
- Issue template for site request: https://github.com/kubernetes/org/issues/new/choose, select **Netlify site request**.
{{< /alert >}}



[setup]: /images/netlify_setup.png
[deploy]: /images/netlify_deployed.png