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

        stage('deploy to k8s cluster'){
            steps{
                withCredentials([sshUserPrivateKey(credentialsId: 'target-ssh-credentials', keyFileVariable: 'keyFile', usernameVariable: 'userName')]) {
                    sh "ssh-keyscan 192.168.105.4 > ~/.ssh/known_hosts"
                    
                    sh "ssh -l ${userName} -i ${keyFile} 192.168.105.4 -C kubectl delete deployment yordan-go-app-deployment"
                    sh "ssh -l ${userName} -i ${keyFile} 192.168.105.4 -C kubectl create deployment yordan-go-app-deployment --image=ttl.sh/yordan-main-go:1h --port 5555"
                    sh "ssh -l ${userName} -i ${keyFile} 192.168.105.4 -C kubectl scale --replicas=2 deployment yordan-go-app-deployment"
                    
                }
            }
        }

        // stage('health check'){
        //     steps{
        //         sh "curl -s http://192.168.105.4:5555/api"
        //     }
        // }
    }
}
