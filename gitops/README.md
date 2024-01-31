# Gitops

## What does this repository contain?

### ArgoCD directory layout
This directory contains a base gitops directory layout that can be used by the ArgoCD's `app of all apps` deployment.

1. Deploy ArgoCD
2. Create a Gitops repository in your Gitlab instance.
3. [Create a Access token in Gitlab for this repository](https://docs.gitlab.com/ee/user/project/settings/project_access_tokens.html)
4. [Deploy the Gitlab agent for Kubernetes](https://docs.gitlab.com/ee/user/clusters/agent/install/)
5. Create a SSH key `ssh-keygen -t ed25519 -C "Gitlab CI cfgmgmtcamp2024" -f id_ed25519_gitlab`
6. Hash the private key with base64
7. [Create a environment variable called `SSH_PRIVATE_KEY_BASE64`](https://docs.gitlab.com/ee/ci/variables/)
8. Create a secret for the ArgoCD credentials. Check the file `argocd-configuration/credentials.yml` for a example. This secret should contain the URL for your gitops repo and the credentials you created earlier.
9. Apply the secret in de `argocd` namespace with `kubectl`
10. Create a repo configuration for the `app of all apps`. Check the file `argocd-configuration/repoconfig.yml` for a example. This repo configuration should contain the url of your gitops repo.
11. Deploy the repo configuration with `kubectl`
12. All the apps in the gitops folder should be deployed immediately after applying the repo configuration

### Gitlab CI file

The gitlab CI file contains a proof of concept to use a application repository and a Gitops repository together to automatically deploy applications. See the `gitlab-ci.yml` for more information.

### `Containerfile`

The `Containerfile` includes a container image for doing "gitops" stuff from within the pipeline.
