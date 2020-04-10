pipeline {
    environment {
        dockerhubCredentials = 'dockerhubCredentials'
    }
    agent any
    stages {
        stage('Lint Dockerfile') {
            steps {
            }
        }
        stage('Build & Push to dockerhub') {
            steps {
            }
        }
        stage('Scan Dockerfile to find vulnerabilities') {
            steps{
            }
        }
        stage('Build Docker Container') {
      		steps {
            }
        }
        stage('Testing The Application') {
      		steps {
            }
        }
        stage('Deploying to EKS') {
            steps {
            }
        }
        stage("Cleaning Docker up") {
            steps {
                script {
                    sh "echo 'Cleaning Docker up'"
                    sh "docker system prune"
                }
            }
        }
    }
}