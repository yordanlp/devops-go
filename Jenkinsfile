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
    }
}
