pipeline {
  agent any
  environment {
    APP_NAME = "ramdhany-portfolio-go"
    REGISTRY = credentials('docker-registry-url')     // e.g. registry.example.com
    REG_USER = credentials('docker-username')         // Jenkins creds
    REG_PASS = credentials('docker-password')         // Jenkins creds
    IMAGE = "${REGISTRY}/${APP_NAME}:${env.BUILD_NUMBER}"
    KUBE_CONTEXT = credentials('kubeconfig')          // Kubeconfig as secret file or string
    NAMESPACE = "portfolio"
  }
  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }
    stage('Go Build & Test') {
      steps {
        sh '''
          go version || true
          go mod download
          go vet ./...
          # No tests yet; placeholder to keep stage
          go build -o bin/app
        '''
      }
    }
    stage('Docker Build') {
      steps {
        sh '''
          echo "$REG_PASS" | docker login "$REGISTRY" -u "$REG_USER" --password-stdin
          docker build -t "$IMAGE" .
          docker push "$IMAGE"
        '''
      }
    }
    stage('Deploy to K8s') {
      steps {
        sh '''
          # Write kubeconfig and apply manifests
          mkdir -p $WORKSPACE/.kube
          echo "$KUBE_CONTEXT" > $WORKSPACE/.kube/config
          export KUBECONFIG=$WORKSPACE/.kube/config
          sed -e "s#{{IMAGE}}#$IMAGE#g" -e "s#{{APP_NAME}}#$APP_NAME#g" -e "s#{{NAMESPACE}}#$NAMESPACE#g" k8s/deployment.yaml | kubectl apply -f -
          sed -e "s#{{APP_NAME}}#$APP_NAME#g" -e "s#{{NAMESPACE}}#$NAMESPACE#g" k8s/service.yaml | kubectl apply -f -
          if [ -f k8s/ingress.yaml ]; then
            sed -e "s#{{APP_NAME}}#$APP_NAME#g" -e "s#{{NAMESPACE}}#$NAMESPACE#g" k8s/ingress.yaml | kubectl apply -f -
          fi
        '''
      }
    }
  }
  post {
    always {
      archiveArtifacts artifacts: 'bin/**', fingerprint: true, onlyIfSuccessful: false
    }
  }
}
