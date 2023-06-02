pipeline {
    agent any

    tools{
        go 'go 1.20.4'
    }

    stages {
        stage('test') {
            steps {
                sh 'go test .'
            }
        }
        
        stage('build') {
            steps {
                sh 'go build main.go'
                archiveArtifacts artifacts: 'main', followSymlinks: false
            }
        }

        stage('build and push docker image'){
            steps{
                sh 'docker build . --tag ttl.sh/yordan-main-go:1h'
                sh 'docker push ttl.sh/yordan-main-go:1h'
            }
        }
    }
}
