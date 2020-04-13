node {
    def app 
   

    stage('Clone repository') {
        /* Cloning the Repository to our Workspace */
            checkout scm

    }   

    stage('Lint Dockerfile') {
      
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
    // stage('Scan Dockerfile to find vulnerabilities') {
        
    //     aquaMicroscanner imageName: "omaroovee/inmemorydb:latest", notCompliesCmd: 'exit 4', onDisallowed: 'fail', outputFormat: 'html'
    
    // }
    stage('Testing The Application') {
        app.inside {
                  sh 'make test'
        }
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