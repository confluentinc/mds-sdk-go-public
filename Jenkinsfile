#!/usr/bin/env groovy

def defaultConfig = [
        nodeLabel: 'docker-oraclejdk8',
        cron: '@weekly',
        testResultSpecs: ['junit': '**/*-reports/TEST-*.xml'],
        connectCveScan : true,
        downStreamValidate: true,
        upstreamProjects: 'confluentinc/metadata-service',
        slackChannel: '#mds-alerts'
]

def body = {
}

def config = jobConfig(body, defaultConfig)

def job = {
    stage('Validate mds-sdk-go') {
        configureGitSSH("github/confluent_jenkins", "private_key")
        withVaultEnv([["github/jenkins_gh_config", "user", "GITHUB_CLI_USER"], ["github/jenkins_gh_config", "oauth_token", "GITHUB_CLI_OAUTH_TOKEN"]]) {
            sh '''#!/bin/bash
                set -e
                mkdir -p ~/.config/gh
                echo "github.com:\n    user: $GITHUB_CLI_USER\n    oauth_token: $GITHUB_CLI_OAUTH_TOKEN" > ~/.config/gh/hosts.yml

                source ./jenkins-generate.sh
                diff=$(git status -s | grep -v '?? .ci/' || true)

                if [[ $CHANGE_BRANCH == mds-pr-* ]]; then
                    MDS_PR=${CHANGE_BRANCH#"mds-pr-"}
                    MDS_HASH=$(gh api repos/confluentinc/metadata-service/pulls/$MDS_PR/commits | jq -r '.[length-1] | .sha')
                    STATUS=$([[ -z ${diff} ]] && echo "success" || echo "failure")
                    STATUS_PAST=$([[ -z ${diff} ]] && echo "succeeded" || echo "failed")
                    echo "{ \\"state\\": \\"$STATUS\\", \\"target_url\\": \\"https://github.com/confluentinc/mds-sdk-go/pull/$CHANGE_ID\\", \\"description\\": \\"The mds-sdk-go build $STATUS_PAST\\", \\"context\\": \\"mds-sdk-go build status\\" }" > status
                    cat status
                    gh api repos/confluentinc/metadata-service/statuses/$MDS_HASH -X POST --input status
                    rm status
                fi

                [[ -z ${diff} ]] || (echo "\n\n======================================================================" \
                && echo "Error: ./generate.sh should not cause diffs. Displaying diffs:\n${diff}\n" && exit 1)

                '''
        }
    }
}
runJob config, job
