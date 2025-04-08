# Deployment Guide

This document describes a simple deployment method for the Loan Approval System on a Google Cloud Platform (GCP) virtual machine using Docker Compose.

---

## 1. Provision a GCP Virtual Machine

**Recommended VM configuration:**

| Setting           | Value                          |
|------------------|---------------------------------|
| OS               | Ubuntu 22.04 LTS                |
| Machine Type     | e2-medium (2 vCPU, 4 GB RAM)    |
| Boot Disk        | 20 GB (Standard SSD or higher)  |
| External IP      | Ephemeral (or reserve static)   |
| Firewall Rules   | Allow HTTP and HTTPS            |
| Network Tags     | `http-server`, `https-server`   |

---

## 2. Install Docker and Docker Compose

SSH into the VM and install Docker:

```bash
sudo apt update
sudo apt install -y docker.io docker-compose
sudo usermod -aG docker $USER
newgrp docker
```

---

## 3. Transfer Project Files

Transfer the project directory from your local system to the VM:

```bash
scp -r ./LoanApprovalSystem <username>@<VM_EXTERNAL_IP>:~/
```

Or use an alternative method such as VSCode Remote SSH or Git.

---

## 4. Add Environment Files

Create the following `.env` files inside the project directory.

**`gateway/.env`:**

```env
GATEWAY_INFERENCE_API_URL=http://inference:8000/predict
GATEWAY_PORT=8020
GATEWAY_RATE_LIMIT_INTERVAL=6s
GATEWAY_RATE_LIMIT_BURST=5
GATEWAY_RATE_LIMIT_TTL=10m
GATEWAY_RATE_LIMIT_CLEANUM=5m
```

**`inference_service/.env`:**

```env
model_path=models/loan_approval.onnx
prediction_threshold=0.5
```

Ensure that the ONNX model exists at the specified path.

---

## 5. Build and Start the Application

Navigate to the project root and run:

```bash
docker-compose build
docker-compose up -d
```

Check container status:

```bash
docker-compose ps
```

---

## 6. Access the Web Interface

Open a browser and navigate to:

```
http://<VM_EXTERNAL_IP>
```

If using a domain name, configure an `A` record pointing to the VM's external IP.

---

## Notes

- NGINX proxies external HTTP traffic (port 80) to the Go-based gateway (port 8020).
- Use a static external IP in GCP if DNS stability is required.
- TLS (HTTPS) is not configured by default. For production, integrate Certbot or use a managed certificate.

---
```