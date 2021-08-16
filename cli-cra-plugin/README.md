# IBM Cloud Code Risk Analyzer plugin for ibmcloud
    
## Purpose
This repo is for the **cra** plug-in for IBM Cloud Code Risk Analyzer.

## Developing a plugin

The plugin development using the sdk: https://github.com/IBM-Cloud/ibm-cloud-cli-sdk
is documented here: https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started

Link to developer's guide: https://github.com/IBM-Cloud/ibm-cloud-cli-sdk/blob/master/docs/plugin_developer_guide.md

## Publishing a plugin
Follow the process documented here :https://github.ibm.com/Bluemix/bluemix-cli-repo#publish-a-new-plug-in

Create a PR to publish the plugin to YS1. Always publish a new plugin, unless we are fixing a bug in an existing plugin.
You should see a PR created in this repo - https://github.ibm.com/Bluemix/bluemix-cli-repo

once PR is merged. Download the plugin from stage and check it out.

Create a PR to push to prod - https://github.ibm.com/Bluemix/bluemix-cli-repo#publish-to-production-repository
again check the https://github.ibm.com/Bluemix/bluemix-cli-repo for the merge status of the PR

## Working on the plugin
1. Install GO lang. (go version go1.14.10 darwin/amd64)
2. clone https://github.ibm.com/oneibmcloud/cli-cra-plugin 
3. cd to cli-cra-plugin directory and run ./buildbinaries.sh
4. This should generate the few binaries in the ./binaries directory
5. On mac install the plugin using the command ./installlocally.sh
6. Verify that you see the *cra* plugin when you type the *ibmcloud plugin list* command
7. Run the test ./test.sh (after setting the require environment variables), the test.sh runs the tests for standalone environment only. Set the environment as appropriate to run other tests.
if you want to test using the dev toolchain then set the value of `IBM_CLOUD_DEVOPS_ENV` env variable to "dev"
(With `IBM_CLOUD_DEVOPS_ENV=dev`, only the test for **Standalone** env can be run, To run all the test use a production toolchain)

* **For Standalone env:**
    * unset PIPELINE_ENV
    * unset PIPELINE_BUILD_STAGE
    * unset DIFF_TOOLCHAIN
    * go test github.ibm.com/oneibmcloud/cli-doi-plugin/plugin/test -count=1 -v
* **For Pipeline env (not build stage):** 
    * export PIPELINE_ENV=true
    * export PIPELINE_BUILD_STAGE=false
    * unset DIFF_TOOLCHAIN
    * go test github.ibm.com/oneibmcloud/cli-doi-plugin/plugin/test -count=1 -v
* **For Pipeline env (build stage):** 
    * export PIPELINE_ENV=true
    * export PIPELINE_BUILD_STAGE=true
    * unset DIFF_TOOLCHAIN
    * go test github.ibm.com/oneibmcloud/cli-doi-plugin/plugin/test -count=1 -v
* **For Pipeline env sending data to different toolchain:** 
    * export PIPELINE_ENV=true
    * export PIPELINE_BUILD_STAGE=false
    * export DIFF_TOOLCHAIN=true
    * go test github.ibm.com/oneibmcloud/cli-doi-plugin/plugin/test -count=1 -v
* These tests can be ran against a specific cluster by exporting INSIGHTS_CLUSTER_URL. 

## Tekton pipeline to test the build 
https://cloud.ibm.com/devops/pipelines/tekton/b5933bbf-0c9f-4824-ae83-3b3a3dffbcc0?env_id=ibm:yp:us-south
