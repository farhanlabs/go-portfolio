# 🚀 Go Portfolio - Cloud-Native Full-Stack Application

A production-grade Go web application featuring custom dashboard and static frontend pages, containerized with Docker, orchestrated via Kubernetes manifests and Helm charts, and fully automated using a GitHub Actions CI/CD pipeline.

---

## 🛠️ Tech Stack & Architecture
- **Backend:** Go (Golang) with built-in unit testing (`main_test.go`)
- **Frontend / Static UI:** HTML templates and assets (`static/` including dashboard, projects, courses, about, contact, and home pages)
- **Containerization:** Multi-stage Docker build (`Dockerfile`)
- **Orchestration:** Kubernetes (`k8s/manifests/`) & Helm Package Manager (`helm/`)
- **CI/CD Automation:** GitHub Actions (`.github/workflows/ci.yaml`) handling build, test, linting, Docker image pushing, and dynamic Helm tag updates.

---

## 📁 Project Directory Structure
```text
GO/
├── .github/
│   └── workflows/
│       └── ci.yaml              # Automated CI/CD pipeline configuration
├── helm/
│   └── go-portfolio-chart/      # Helm chart for package management & releases
├── k8s/
│   └── manifests/               # Native Kubernetes manifests (Deployment, Ingress, Service, Issuer)
├── static/                      # Frontend HTML pages & images (Dashboard, Home, Projects, etc.)
├── Dockerfile                   # Container packaging configuration
├── go.mod                       # Go module dependencies
├── main.go                      # Main web server implementation
└── main_test.go                 # Go unit test suite