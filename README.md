<!-- Banner -->
<p align="center">
  <img src="static/images/banner.png" alt="Go Portfolio Banner" width="100%">
</p>

<h1 align="center">🚀 Go Portfolio — Cloud-Native Production-Grade Web Application</h1>

<p align="center">
A production-grade Go web application demonstrating end-to-end cloud-native delivery: Docker containerization, Kubernetes orchestration, Helm packaging, GitHub Actions CI/CD, and ArgoCD GitOps.
</p>

<p align="center">

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Container-2496ED?logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/Kubernetes-1.30-326CE5?logo=kubernetes&logoColor=white)
![Helm](https://img.shields.io/badge/Helm-Charts-0F1689?logo=helm&logoColor=white)
![ArgoCD](https://img.shields.io/badge/ArgoCD-GitOps-EF7B4D?logo=argo&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-CI/CD-2088FF?logo=githubactions&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)

</p>

<p align="center">
  <a href="#-overview">Overview</a> •
  <a href="#️-architecture">Architecture</a> •
  <a href="#-tech-stack">Tech Stack</a> •
  <a href="#-production-features">Features</a> •
  <a href="#-project-structure">Structure</a> •
  <a href="#-run-locally">Run Locally</a> •
  <a href="#️-deploy-using-helm">Deploy</a> •
  <a href="#-screenshots">Screenshots</a>
</p>

---

## 📖 Overview

This repository demonstrates a **production-style, cloud-native deployment** of a Go web application, built to reflect how real-world platform and DevOps teams ship software — not just how to run `go run main.go`.

Every layer of the stack was deliberately added to mirror a production environment:

- **Docker** → consistent, portable builds across environments
- **Kubernetes** → self-healing, declarative orchestration
- **Helm** → templated, versioned, environment-agnostic configuration
- **GitHub Actions** → automated build, test, and image publishing
- **ArgoCD (GitOps)** → the cluster's state is driven by Git, not manual `kubectl apply`
- **HPA + Probes + Resource Limits** → the app behaves predictably under load and fails safely

---

## 🏗️ Architecture

```
                        Developer
                            │
                            ▼
                     Git Push (GitHub)
                            │
                            ▼
                  GitHub Actions CI/CD
                            │
        ┌───────────────────┼────────────────────┐
        │                   │                    │
        ▼                   ▼                    ▼
    Build App          Run Tests         Docker Build
                                                │
                                                ▼
                                        Push Image to
                                          Docker Hub
                                                │
                                                ▼
                                     Update Helm Image Tag
                                                │
                                                ▼
                                           Git Repository
                                                │
                                                ▼
                                             ArgoCD
                                                │
                                                ▼
                                         Kubernetes Cluster
                                                │
        ┌────────────────────────────────────────┼────────────────────────┐
        ▼                                        ▼                        ▼
 Deployment                              Service                 Horizontal Pod Autoscaler
        │
        ▼
     Running Pods
```

**Why GitOps?** Instead of engineers running `kubectl apply` by hand, ArgoCD continuously watches the Git repo and syncs the cluster to match it. Git becomes the single source of truth — every deployment is versioned, auditable, and revertible with a `git revert`.

---

## ⚙️ Tech Stack

| Layer | Technology |
|---|---|
| Backend | Go (Golang) |
| Frontend | HTML • CSS • JavaScript |
| Containerization | Docker |
| CI/CD | GitHub Actions |
| GitOps | ArgoCD |
| Orchestration | Kubernetes |
| Package Manager | Helm |
| Scaling | Horizontal Pod Autoscaler (HPA) |
| Ingress | NGINX Ingress |
| TLS | cert-manager + Let's Encrypt |

---

## 🚀 Production Features

### ✅ GitHub Actions CI/CD
- Build the Go application
- Run unit tests
- Build the Docker image
- Push the image to Docker Hub
- Update the Helm chart's image tag
- Trigger automatic GitOps deployment via ArgoCD

### ☸️ Kubernetes
- Deployment — manages Pod lifecycle and rolling updates
- Service — stable internal networking for Pods
- Ingress — routes external traffic into the cluster
- Self-healing Pods — Kubernetes replaces failed containers automatically
- Rolling updates — zero-downtime deployments

### 📈 Horizontal Pod Autoscaler
Automatically scales Pods based on CPU utilization, so the app handles traffic spikes without manual intervention.

```yaml
minReplicas: 2
maxReplicas: 5
targetCPUUtilizationPercentage: 80
```

### 💻 Resource Management
Explicit requests and limits so the scheduler places Pods correctly and no single Pod can starve the node.

```yaml
resources:
  requests:
    cpu: 100m
    memory: 128Mi
  limits:
    cpu: 500m
    memory: 512Mi
```

This guarantees baseline resources, prevents noisy-neighbor issues, improves scheduling decisions, and keeps the deployment production-ready.

### ❤️ Health Checks — Liveness & Readiness Probes

Both probes are configured on the Deployment so Kubernetes knows exactly when a Pod is alive versus ready to serve traffic — the two are not the same thing, and conflating them is a common production mistake.

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 15
  timeoutSeconds: 3
  failureThreshold: 3

readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 3
  failureThreshold: 3
```

**Liveness Probe** — checks whether the application process is alive.
- If it fails → Kubernetes kills and restarts the container. This recovers apps stuck in a deadlock or unresponsive state.

**Readiness Probe** — checks whether the application is ready to receive traffic.
- If it fails → the Pod is pulled out of the Service's endpoint list. Traffic stops flowing to it, but the container itself is **not** restarted — this matters during startup, slow dependency checks, or temporary overload.

---

## 📂 Project Structure

```text
go-portfolio/
│
├── .github/
│   └── workflows/
│       └── ci-cd.yml              # GitHub Actions pipeline — build, test, Docker build/push, Helm tag update
│
├── helm/
│   └── go-portfolio-chart/
│       ├── templates/
│       │   ├── deployment.yaml    # Pod spec: container image, probes, resource requests/limits
│       │   ├── service.yaml       # Stable internal network endpoint for the Pods
│       │   ├── ingress.yaml       # External traffic routing + TLS termination
│       │   ├── hpa.yaml           # Autoscaling rules (min/max replicas, CPU target)
│       │   └── issuer.yaml        # cert-manager issuer for automated TLS certificates
│       │
│       ├── values.yaml            # Centralized, environment-agnostic configuration (image tag, replicas, resources)
│       ├── Chart.yaml             # Helm chart metadata (name, version, description)
│       └── .helmignore            # Files excluded when packaging the chart
│
├── k8s/                           # Raw Kubernetes manifests for quick local testing without Helm
│
├── static/
│   ├── images/                    # Banner and README screenshot assets
│   ├── home.html                  # Landing page
│   ├── about.html                 # About page
│   ├── contact.html               # Contact page
│   ├── dashboard.html             # Dashboard view
│   └── projects.html              # Projects showcase page
│
├── Dockerfile                     # Multi-stage build — compiles the Go binary, ships a minimal runtime image
├── go.mod                         # Go module definition and dependency versions
├── go.sum                         # Dependency checksums for reproducible builds
├── main.go                        # Application entrypoint — HTTP server and routes
├── main_test.go                   # Unit tests for the application
├── LICENSE                        # MIT License
└── README.md                      # Project documentation (this file)
```

**Why this structure?** Application code, infrastructure-as-code (Helm), and CI/CD pipeline definitions are kept in clearly separated directories. This mirrors how real engineering teams organize repos, so a reviewer — technical or non-technical — can immediately locate app logic vs. deployment config vs. automation.

---

## 🔄 CI/CD Pipeline

```
Developer
    │
    ▼
Git Push
    │
    ▼
GitHub Actions
    │
    ├── Checkout Code
    ├── Build
    ├── Unit Tests
    ├── Docker Build
    ├── Push Docker Image
    └── Update Helm Values
             │
             ▼
        Git Repository
             │
             ▼
           ArgoCD
             │
             ▼
     Kubernetes Cluster
```

---

## 📸 Screenshots

| ArgoCD | Kubernetes Pods |
|---|---|
| ![ArgoCD](static/images/argocd.png) | ![Pods](static/images/pods.png) |

| Horizontal Pod Autoscaler | Application |
|---|---|
| ![HPA](static/images/hpa.png) | ![Application](static/images/application.png) |

---

## 🚀 Run Locally

### Clone the Repository
```bash
git clone https://github.com/farhanlabs/go-portfolio.git
cd go-portfolio
```

### Run using Go
```bash
go mod tidy
go run main.go
```
App available at: `http://localhost:8080`

### Run using Docker
```bash
docker build -t go-portfolio .

docker run -d \
  -p 8080:8080 \
  go-portfolio
```

---

## ☸️ Deploy using Helm

**Install**
```bash
helm install go-portfolio ./helm/go-portfolio-chart
```

**Upgrade**
```bash
helm upgrade go-portfolio ./helm/go-portfolio-chart
```

**Uninstall**
```bash
helm uninstall go-portfolio
```

---

## 📊 Useful Commands

```bash
kubectl get pods
kubectl get svc
kubectl get ingress
kubectl get hpa
kubectl get deployments
```

---

## 👨‍💻 Author

**Md Farhan Rza**
DevOps Engineer • Full Stack Developer • Founder, Zevix Digital

[![GitHub](https://img.shields.io/badge/GitHub-farhanlabs-181717?logo=github&logoColor=white)](https://github.com/farhanlabs)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-0A66C2?logo=linkedin&logoColor=white)](#)

---

## 📜 License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

<p align="center">⭐ If this project helped you understand cloud-native deployment patterns, consider starring the repository.</p>
