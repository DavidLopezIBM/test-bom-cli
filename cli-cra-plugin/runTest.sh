export TOOLCHAIN_ID=$1

ibmcloud login -a "https://test.cloud.ibm.com" --no-region --apikey $2

ibmcloud plugin install /tmp/cli-cra-plugin/doi-cli-linux-amd64-0.3.2

ibmcloud cra bom-generate --dir /tmp/cli-cra-plugin --report /tmp/test.json