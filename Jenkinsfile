node {
    def app 
   

    stage('Clone repository') {
        /* Cloning the Repository to our Workspace */
            checkout scm

    }   

    stage('Linting') {
      
        docker.image('hadolint/hadolint:latest-debian').inside() {
                            sh 'hadolint ./Dockerfile | tee -a hadolint_lint.txt'
                            sh '''
                                lintErrors=$(stat --printf="%s"  hadolint_lint.txt)
                                if [ "$lintErrors" -gt "0" ]; then
                                    echo "Errors have been found, please see below"
                                    cat hadolint_lint.txt
                                    exit 1
                                else
                                    echo "There are no erros found on Dockerfile!!"
                                fi
                            '''
        }
        
    }
    stage('Build & Push to dockerhub') {
       
            app = docker.build("omaroovee/inmemorydb")
            docker.withRegistry('https://registry.hub.docker.com', 'docker-hub') {
                app.push("${env.BUILD_NUMBER}")
                app.push("latest")
            } 
        
    }

    stage('Build Docker Container') {

        sh 'docker run --name inmemorydb -d -p 80:80 omaroovee/inmemorydb:latest'
    }


    stage('Deploying to EKS') {
        app.inside {
                echo "Tests passed"
       }
    }
    stage("Cleaning Docker up") {
      
        sh "echo 'Cleaning Docker up'"
        sh "docker system prune"
    
        
    }
    
}