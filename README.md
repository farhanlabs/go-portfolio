<div align="center">

# 🚀 Go Portfolio — Cloud-Native Production-Grade Web Application

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-blue?style=for-the-badge&logo=go)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Orchestrated-326CE5?style=for-the-badge&logo=kubernetes)](https://kubernetes.io/)
[![Helm](https://img.shields.io/badge/Helm-Package%20Manager-0F1689?style=for-the-badge&logo=helm)](https://helm.sh/)
[![ArgoCD](https://img.shields.io/badge/ArgoCD-GitOps-EF7b4d?style=for-the-badge&logo=argo)](https://argoproj.github.io/)
[![GitHub Actions](https://img.shields.io/badge/CI%2FCD-GitHub%20Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)](https://github.com/features/actions)

*A robust, enterprise-grade cloud-native web application showcasing high-availability architecture, automated CI/CD pipelines, GitOps continuous delivery, and advanced Kubernetes orchestration.*

</div>

---

## 🏗️ Architecture & Tech Stack

| Layer | Technology / Tool | Implementation Details |
| :--- | :--- | :--- |
| **Backend** | **Go (Golang)** | High-performance HTTP server with automated unit testing (`main_test.go`) |
| **Frontend** | **HTML5 / CSS / JS** | Responsive multi-page layout (`static/` dashboard, projects, courses, about, contact, home) |
| **Containerization** | **Docker** | Multi-stage optimized container builds ensuring minimal image footprint |
| **Package Management**| **Helm** | Templated Kubernetes manifests (`go-portfolio-chart`) for dynamic releases |
| **Orchestration & Scale**| **Kubernetes & HPA** | Self-healing pods with resource governance and CPU-based autoscaling |
| **Continuous Delivery** | **ArgoCD & GitHub Actions** | Fully automated GitOps deployment loop triggered via secure CI/CD workflows |

---

## 🌟 Key Production Features

* **⚡ Horizontal Pod Autoscaler (HPA):** Dynamically scales application replica counts (min: 2, max: 5) based on real-time CPU utilization thresholds.
* **🛡️ Liveness & Readiness Probes:** Automated health checks (`/` endpoint monitoring) to safely manage traffic routing and automated crash-recovery restarts.
* **📦 Resource Governance:** Enforced strict CPU and Memory requests/limits to protect cluster nodes from resource exhaustion.
* **🔄 GitOps Workflow:** Zero-downtime automated synchronization where code commits trigger GitHub Actions container builds, dynamic tag updates, and ArgoCD cluster reconciliation.

---

## 📁 Project Directory Structure

```text
GO/
├── .github/
│   └── workflows/
│       └── ci-cd.yml              # Automated CI/CD build & image tag update pipeline
├── helm/
│   └── go-portfolio-chart/        # Helm package manager chart
│       ├── templates/
│       │   ├── deployment.yaml    # Core deployment spec with Probes & Resource limits
│       │   ├── hpa.yaml           # Horizontal Pod Autoscaler configuration
│       │   ├── ingress.yaml       # Ingress routing rules & load balancing
│       │   ├── issuer.yaml        # SSL/TLS Certificate issuer configuration
│       │   └── service.yaml       # Kubernetes internal service mapping
│       ├── .helmignore            # Helm packaging ignore rules
│       ├── Chart.yaml             # Chart metadata definitions
│       └── values.yaml            # Configurable parameters (replicas, resources, image tags)
├── k8s/                           # Raw Kubernetes manifests backup repository
├── static/                        # Frontend UI assets & pages
│   ├── images/                    # Asset media and graphics repository
│   ├── about.html                 # About page
│   ├── contact.html               # Contact page
│   ├── courses.html               # Courses catalog page
│   ├── dashboard.html             # Administrative dashboard UI
│   ├── home.html                  # Landing page
│   └── projects.html              # Portfolio projects showcase
├── Dockerfile                     # Multi-stage Docker packaging configuration
├── go.mod                         # Go module dependencies
├── LICENSE                        # Open-source licensing terms
├── main.go                        # Core web application server implementation
├── main_test.go                   # Backend unit testing suite
└── README.md                      # Comprehensive project documentation