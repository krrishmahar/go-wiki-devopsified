
# Go Wiki – DevOpsified 🐳🚀

A simple Wiki web application written in Go, containerized and deployed using modern DevOps tooling including Docker, Kubernetes, GitHub Actions, Helm, and Argo CD.

---

## 🔥 What I Implemented

✅ Go backend using `net/http`, HTML templates & local file storage  
✅ Multi-stage Dockerfile with non-root user  
✅ Minikube setup with real K8s cluster  
✅ Kubernetes Deployment, Service, Ingress + NGINX Ingress Controller  
✅ Domain mapping via hosts file (`go-wiki.local`) + `minikube tunnel`  
✅ File permission fix using `securityContext`  
✅ Helm Chart to package the app  
✅ CI with GitHub Actions, CD via Argo CD (GitOps)

---

## 🧰 Tech Stack

<p align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="40" title="Go"/>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" width="40" title="Docker"/>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/kubernetes/kubernetes-plain.svg" width="40" title="Kubernetes"/>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/github/github-original.svg" width="40" title="GitHub Actions"/>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/helm/helm-original.svg" width="40" title="Helm"/>
  <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/nginx/nginx-original.svg" width="40" title="NGINX"/>
  <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linux/linux-original.svg" width="40" title="Linux"/>
  <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/bash/bash-original.svg" width="40" title="Bash"/>
</p>

---

## 🏗️ Architecture Overview

```
GitHub → GitHub Actions → Docker Build → Push to Repo
        ↓
       Argo CD ← GitOps Pull
        ↓
  Minikube or EC2 Cluster (K8s Deployment, Service, Ingress)
        ↓
NGINX Ingress Controller → go-wiki.local
```

---

## 📦 Setup Instructions

### Prerequisites

- Go 1.18+
- Docker
- Minikube + kubectl
- Helm 3.x
- Argo CD installed on Minikube

---

## 🚀 How to Run

### 1. Clone the Repository

```bash
git clone https://github.com/krrishmahar/go-wiki-devopsified.git
cd go-wiki-devopsified
```

### 2. Build & Run Locally (Dev Mode)

```bash
go run main.go
```

### 3. Build Docker Image

```bash
docker build -t gowiki .
```

### 4. Set Up Minikube

```bash
minikube start
minikube addons enable ingress
minikube tunnel
```

### 5. Apply Kubernetes Manifests

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml
```

### 6. (Optional) Install via Helm

```bash
helm install gowiki ./chart
```

### 7. Configure Hosts File

Add this to `/etc/hosts` or `C:\Windows\System32\drivers\etc\hosts`:

```
127.0.0.1 go-wiki.local
```

Now visit [http://go-wiki.local](http://go-wiki.local)

---

## 🤖 CI/CD Overview

- GitHub Actions handles:
  - Docker image build
  - Helm linting
  - Pushing manifests or Helm package
- Argo CD pulls from Git repo and deploys to Minikube automatically

---

## 📷 Screenshot

![App Screenshot](https://media.licdn.com/dms/image/v2/D4D22AQGa6bEfI8qfdg/feedshare-shrink_800/B4DZZd0xynGgAg-/0/1745330839512?e=1750896000&v=beta&t=n-AExE558UL3tXOCHJ8yHEahd9ioOmOWm71jyNjHVrU)

![App Scrrenshot 2](https://media.licdn.com/dms/image/v2/D4D22AQEtRr17xI0G3g/feedshare-shrink_800/B4DZZd0xzDHIAk-/0/1745330839289?e=1750896000&v=beta&t=cqERG_zNKkS194H0vKI_0ErxhLD-MGvA5pEcOdzQ3y8)

---

## 🤝 Contributing

If you learned something from this, consider leaving a ⭐!  
Contributions are welcome — feel free to fork and PR!

#HappyShipping 🚢
