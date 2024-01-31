# Testapp

This repository contains a demo application in  `golang` and a `gitlab-ci.yml` pipeline file.


## Application code

The application code contains a basic demo in golang for showcasing.

## The `gitlab-ci.yml` file

This is a Gitlab pipeline file that contains pipeline code to do the following:

- `destroy_previous_deployment`
- `build_app`
- `build_container`
- `tag_container`
- `test`
- `generate_manifests_pipeline`
- `deploy_pipeline`
- `test_deployment_pipeline`
- `destroy_deployment_pipeline`
- `deploy_tst_environment`
- `deploy_acc_environment`
- `deploy_prd_environment`

The application will build and start it's own deployment from within Kubernetes. After that, it can be tagged and a pipeline will trigger the `gitops` pipeline from within the `gitops` repository. For more information, check out the `gitops` directory in this project.

