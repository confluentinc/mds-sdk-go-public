# MDS-SDK-GO-PUBLIC

This repository holds generated Go SDKs for the MDS API, including v1 MDS for on-prem. V1 MDS API is held in a versioned directory and is updated independently. Warning: the v2alpha1 endpoints are already deprecated, and should not be used under any circumstances.
See below on how to update a version of the MDS SDK.


## Updating mds-sdk-go-public after changing an API spec

1. Download/build/install this fork of the OpenAPI generator tool: https://github.com/confluentinc/openapi-generator

        mkdir -p "$HOME/code/confluentinc"
        cd "$HOME/code/confluentinc"
        git clone git@github.com:confluentinc/openapi-generator.git
        cd openapi-generator
        git switch ath-add-interfaces
        mvn clean install -DskipTests

1. Clone metadata-service if you haven't already

        cd "$HOME/code/confluentinc"
        git clone git@github.com:confluentinc/metadata-service.git

1. Install mocker if you haven't already

        cd "path/to/mds-sdk-go-public/mdsv1"
        go install github.com/travisjeffery/mocker/cmd/mocker

1. Generate the updated SDK from the mds-sdk-repo. Note that the path to the generator should be absolute.

        cd "path/to/mds-sdk-go-public"
        ./generate.sh <path/to/openapi-generator> ./mds<version> <path/to/metadata-service/../spec.yaml>

1. Run a `git diff` and ensure that the changes you see are only to the files
   you think should have been modified. If you see some minor changes across a
   bunch of files, likely something went wrong like specifying a wrong flag to
   openapi-generator. The `generate.sh` script may need to be updated.

1. Use `git` to commit and push the result.  You should likely do this as a PR.

1. Make sure downstream dependencies like CLI get the updated SDK, and that the new version causes no breakages.
