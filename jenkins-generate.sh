#!/bin/bash -f

# Debug
# set -x
sudo apt-get update
sudo apt-get install make
export ORIGIN_DIR=`pwd`
cd ..

rm -rf metadata-service
git clone git@github.com:confluentinc/metadata-service.git

wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.2.1/openapi-generator-cli-5.2.1.jar -O openapi-generator-cli.jar
wget https://dl.google.com/go/go1.16.3.linux-amd64.tar.gz
wget https://github.com/cli/cli/releases/download/v0.11.0/gh_0.11.0_linux_amd64.tar.gz
tar -xzf go1.16.3.linux-amd64.tar.gz
tar -xzf gh_0.11.0_linux_amd64.tar.gz
export PATH=`pwd`/go/bin:`pwd`/gh_0.11.0_linux_amd64/bin:$PATH
export GOPATH=`pwd`/go

# If this jenkins script is run by a MDS PR, the cloned MDS repo should match that PR
# If this jenkins script is run by a mds-sdk-go-public PR, the cloned MDS repo should match the MDS PR that created the mds-sdk-go-public PR (if any)
cd metadata-service
if [[ $CHANGE_URL == *mds-sdk-go-public* ]] && [[ $CHANGE_BRANCH == mds-pr-* ]]; then
    MDS_PR=${CHANGE_BRANCH#"mds-pr-"}
    gh pr checkout $MDS_PR
fi
if [[ $CHANGE_URL == *metadata-service* ]]; then
    gh pr checkout $CHANGE_ID
fi
cd ..

cd $ORIGIN_DIR
go install github.com/travisjeffery/mocker/cmd/mocker@latest

./generate.sh `pwd`/../openapi-generator-cli.jar ./mdsv1 `pwd`/../metadata-service/rbac-api-server/src/main/resources/WEB-INF/openapi/mds-spec-1_0.yaml
cd $ORIGIN_DIR
git remote set-url origin git@github.com:confluentinc/mds-sdk-go-public.git

