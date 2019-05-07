#!groovy

import groovy.transform.Field

@Field String email_to = 'kph@platinasystems.com'
@Field String email_from = 'jenkins-bot@platinasystems.com'
@Field String email_reply_to = 'no-reply@platinasystems.com'

pipeline {
    agent any
    environment {
	GOPATH = "$WORKSPACE/go-pkg"
	HOME = "$WORKSPACE"
    }
    stages {
	stage('Checkout') {
	    steps {
		echo "Running build #${env.BUILD_ID} on ${env.JENKINS_URL} GOPATH ${GOPATH}"
		dir('loopback') {
		    git([
			url: 'https://github.com/platinasystems/loopback.git',
			branch: 'master'
		    ])
		}
	    }
	}
	stage('Build') {
	    steps {
		dir('loopback') {
		    echo "Building loopback..."
		    sh 'set +x; export PATH=/usr/local/go/bin:${PATH}; go build -x; go test -x'
		}
	    }
	}
    }

    post {
	success {
	    mail body: "loopback build ok: ${env.BUILD_URL}",
		from: email_from,
		replyTo: email_reply_to,
		subject: 'loopback build ok',
		to: email_to
	}
	failure {
	    cleanWs()
	    mail body: "loopback build error: ${env.BUILD_URL}",
		from: email_from,
		replyTo: email_reply_to,
		subject: 'loopback BUILD FAILED',
		to: email_to
	}
    }
}
