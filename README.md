# Ramdhany Portfolio (Go)

A fast, minimal Go web app to host a personal portfolio/resume with a modern UI (Tailwind + Alpine).

## Run locally

```bash
go mod download
go run main.go
# open http://localhost:8080
```

## Docker

```bash
docker build -t ramdhany/portfolio:dev .
docker run -p 8080:8080 ramdhany/portfolio:dev
```

## Jenkins Pipeline (Kubernetes)

- Create Jenkins credentials:
  - `docker-registry-url` (Secret text) → e.g., `registry.example.com`
  - `docker-username` (Username/Password)
  - `docker-password` (Username/Password or Secret text)
  - `kubeconfig` (Secret text or file) → kubeconfig content for deploy cluster
- Update `NAMESPACE` in `Jenkinsfile` and the Ingress host in `k8s/ingress.yaml`.
- Push this repo to GitHub, then create a Jenkins multibranch or pipeline job and point to it.

The pipeline will:
1. Checkout
2. Build Go
3. Build & Push Docker image
4. Deploy to Kubernetes (apply Deployment/Service/Ingress)

> If your Jenkins agents cannot run Docker, switch to Kaniko or BuildKit-in-Kubernetes.

## Customize Content

Edit the `HomeHandler` in `main.go` or read data from JSON/ENV later. The UI uses Tailwind via CDN for simplicity.
