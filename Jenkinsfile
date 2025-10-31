pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/rafmasloman/naro-app-backend.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t naro-app-be:latest .'
                }
            }
        }

        stage('Run Container') {
            steps {
                script {
                    // stop container lama kalau ada
                    sh 'docker stop naro-app-be || true'
                    sh 'docker rm naro-app-be || true'

                    // jalankan container baru
                    sh 'docker run -d -p 8001:8001 --name naro-app-be naro-app-be:latest'
                }
            }
        }
    }
}
