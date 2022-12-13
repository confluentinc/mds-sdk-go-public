#!/bin/bash -f

function printhdr() {
  echo -e $1
  printf %${#1}s | tr ' ' '='
  echo
}

[[ $# -ne 3 ]] && echo "Usage: $0 <path/to/openapi-generator-cli.jar> <path/to/mds-sdk/mds/version> <path/to/metadata-service/../spec.yaml>" && exit 1

openapiPath=$1
tgtPath=$2
mdsPath=$3
version=$(basename ${tgtPath})

# Fail script on any errors
set -o errexit
# Debug
# set -x
cd $tgtPath

which mocker > /dev/null || (printhdr "Installing mocker"; go install github.com/travisjeffery/mocker/cmd/mocker; echo)

printhdr "Removing all files under $(pwd)"
find . ! '(' -name 'Makefile' -o -name '.' ')' -maxdepth 1 -print0 | xargs -0 rm -rf

printhdr "\nGenerating mds-sdk-go"
java -ea ${JAVA_OPTS} -Xms512M -Xmx1024M -server \
  -jar ${openapiPath} generate \
  --input-spec ${mdsPath} \
  --api-package ${version} --package-name ${version} --generator-name go-deprecated \
  --git-user-id confluentinc --git-repo-id mds-sdk-go \
  --additional-properties generateInterfaces=true,enumClassPrefix=true,isGoSubmodule=true

printhdr "\nGenerating mocks"
make mock

printhdr "\nFormatting go"
make fmt
