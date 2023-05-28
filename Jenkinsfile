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

        stage('deploy'){
            steps{
                withCredentials([sshUserPrivateKey(credentialsId: 'target-ssh-credentials', keyFileVariable: 'keyFile', usernameVariable: 'userName')]) {
                    sh "ssh-keyscan 192.168.105.3 > ~/.ssh/known_hosts"
                    sh "scp -i ${keyFile} main ${userName}@192.168.105.3:"
                    sh "scp -i ${keyFile} myapp.service ${userName}@192.168.105.3:"
                    sh "ssh -l ${userName} -i ${keyFile} 192.168.105.3 -C sudo mv myapp.service /etc/systemd/system"
                    sh "ssh -l ${userName} -i ${keyFile} 192.168.105.3 -C sudo systemctl daemon-reload"
                    sh "ssh -l ${userName} -i ${keyFile} 192.168.105.3 -C sudo systemctl enable myapp"
                    sh "ssh -l ${userName} -i ${keyFile} 192.168.105.3 -C sudo systemctl start myapp"
                }
            }
        }
    }
}
