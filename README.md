# Loan Approval System with ONNX, FastAPI, and Go

## Overview

Educational project demonstrating an end-to-end ML pipeline:
- Train and export a PyTorch model to ONNX
- Serve predictions via FastAPI
- Interface through a Go-based frontend with HTMX
- Deploy using Docker Compose to a VPS

## Architecture

```
User ↔ NGINX ↔ Go Gateway ↔ FastAPI (ONNX model)
```

## Key Technologies

- **PyTorch / ONNX** – Training and optimized inference
- **FastAPI** – REST API backend
- **Go** – HTTP UI and client interface
- **HTMX + Bootstrap** – Dynamic form-based UI
- **Docker Compose** – Multi-service deployment
- **Viper / Pydantic** – Configuration management

---

## Project Structure

```
.
├── gateway/              # Go-based frontend
├── inference_service/    # FastAPI backend
├── training/             # Model training + ONNX export
├── docker-compose.yml
├── Deploy.md             # VPS deployment instructions
└── README.md
```

---

## Training

Train the model using PyTorch and export to ONNX.

### Requirements

```bash
cd training
pip install -r requirements.txt
```

Ensure dataset exists:
```
training/data/loan_approval_dataset.csv
```

Update training config in:
```
training/config/config.yaml
```

### Run Training

```bash
python main.py
```

This will:
- Preprocess data
- Train the model
- Evaluate test accuracy
- Export ONNX model to `LoanApprovalPredictor.onnx`

---

## Running Locally (Manual)

### Python Inference API

```bash
cd inference_service
uvicorn main:app --host 0.0.0.0 --port 8000
```

### Go Gateway UI

```bash
cd gateway/cmd/gateway
go run main.go
```

Open: [http://localhost:8020](http://localhost:8020)

---

## API

### POST /predict

**Request**

```json
{
  "annual_income": 800000,
  "loan_amount": 1000000,
  "loan_term": 15,
  "credit_score": 720
}
```

**Response**

```json
{
  "approval_status": "approved",
  "confidence": 0.87
}
```

---

## Deployment (VPS)

Production deployment using Docker Compose.

### Steps

1. Provision VM (Ubuntu 22.04)
2. Install Docker and Compose
3. Transfer project files
4. Add `.env` files
5. Run:

```bash
docker-compose build
docker-compose up -d
```

See full details in [Deploy.md](./Deploy.md).

---

## Educational Objectives

This project is designed as a reference architecture for real-world ML model deployment.

### Core Topics

- **Model Training and Export**  
  Train a PyTorch model and export it to ONNX for portable, runtime-agnostic inference.

- **Inference API Design (FastAPI + ONNX Runtime)**  
  Implement a minimal Python wrapper over ONNX Runtime, exposing a REST endpoint for predictions. This simulates production ML deployment where Python is used purely for orchestration and compatibility with ONNX, TensorRT, or other backends.

- **Gateway Abstraction in Go**  
  Use a statically compiled gateway (in Go) to handle:
  - Form input validation
  - Rate limiting
  - UI rendering
  - External API communication (Python service)

  This reflects patterns commonly used in production where performance-critical services are isolated from Python inference backends.

- **Deployment via Docker Compose**  
  Coordinate multiple services (inference API, gateway, NGINX) using container orchestration. Demonstrates service wiring, `.env` config loading, and isolated runtime environments.

- **UI-Driven Interaction with HTMX**  
  Serve a client-facing form to submit inference requests and render results dynamically. No JavaScript bundling or frameworks; just HTML, HTMX, and Bootstrap.

### Deployment Model Emulated

This architecture reflects a common production deployment pattern:

[Client] → [NGINX] → [Gateway (Go)] → [Inference API (Python/ONNX)]


This separation of concerns:
- Decouples model execution from request handling
- Enables performant, scalable HTTP interfaces
- Isolates Python execution to the minimal required footprint
- Makes it easier to replace models without touching frontend or gateway logic

This project can serve as a foundational template or reference implementation for lightweight, containerized ML deployments.

---

## Suggested Extensions

- **Add Testing**:
  - Unit tests for FastAPI routes and models
  - ONNX inference logic validation
  - Gateway service logic and middleware
- **Add Observability**:
  - Structured logging (Python/Go)
  - Basic metrics via Prometheus or OpenTelemetry