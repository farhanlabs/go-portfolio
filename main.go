package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/projects.html")
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/dashboard.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

// DevOps Pipeline Simulation
func deployStreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	pipelineSteps := []string{
		"[STAGE 1/6] Building container image using optimized Dockerfile...",
		"[STAGE 2/6] GitHub Actions: Running CI tests, linting, and pushing to registry...",
		"[STAGE 3/6] Terraform: Provisioning cloud infrastructure, VPC, & EKS cluster nodes...",
		"[STAGE 4/6] Kubernetes (K8s): Applying manifests, deploying pods, services, & ingress...",
		"[STAGE 5/6] ArgoCD & Observability: Syncing GitOps state & setting up Prometheus/Grafana...",
		"SUCCESS: Pipeline completed! Website is live and serving traffic from the Kubernetes cluster. 🚀",
	}

	for _, step := range pipelineSteps {
		fmt.Fprintf(w, "data: %s\n\n", step)
		w.(http.Flusher).Flush()
		time.Sleep(1200 * time.Millisecond)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/api/deploy", deployStreamHandler)

	log.Println("CloudPulse Server running on http://0.0.0.0:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}