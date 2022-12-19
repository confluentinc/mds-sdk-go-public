# MDS-SDK-GO-PUBLIC

This repository holds generated Go SDKs for the MDS API, including v1 MDS for on-prem. V1 MDS API is held in a versioned directory and is updated independently. Warning: the v2alpha1 endpoints are already deprecated, and should not be used under any circumstances.



## Contributing to mds-sdk-go-public

### with access to `metadata-service` repo:

1. Open a PR in `metadata-service` with your changes to the API spec file. This will trigger the Jenkins job to generate v1 MDS SDK in this public repo, compare that to the current version, and open a PR if there's any update.

2. Get the PR in `metadata-service` approved and merged. 

3. Get the PR in `mds-sdk-go-public` approved and merged.

4. After the PR is merged, create a release manually.


### for outside contributors:

Please contact @confluent/cli team or open a support ticket to put forward your suggestions.
