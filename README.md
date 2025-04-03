# Loan Approval System with ONNX, FastAPI, and Go

## Overview

This project is an educational demonstration of building and deploying a machine learning system using **PyTorch (ONNX export)**, **FastAPI**, and a frontend interface built in **Go** using **HTMX and `html/template`**. It walks through data preprocessing, model training, API design, and interface development for a real-world binary classification task (loan approval prediction).

## Features

- **Machine Learning Model**: A PyTorch-based MLP with input normalization, exported to **ONNX** for inference.
- **Data Processing**: Custom preprocessing pipeline written in Python using **pandas**.
- **API Deployment**: A **FastAPI** backend serves predictions from the ONNX model.
- **Form-Based UI**: A simple HTML form using **HTMX** and **Bootstrap**, served from a **Go-based** HTTP server.
- **Rate Limiting**: Configurable request throttling implemented in the Go gateway using token buckets.
- **Modular Structure**: Clear separation of training, inference, and presentation layers.

## Project Structure

```
.
├── gateway                  # Go-based frontend
│   ├── cmd/gateway          # Entry point
│   ├── internal/
│   │   ├── app              # Form handlers, routing, middleware
│   │   ├── client           # REST client to inference API
│   │   ├── config           # Viper-based configuration
│   │   └── templates        # HTML templates (HTMX compatible)
├── inference_service        # Python backend (FastAPI)
│   ├── api/routes.py        # Predict endpoint
│   ├── core/                # Models, validators, services
│   ├── repository/          # ONNX runtime wrapper
│   └── config.py
├── training                 # Offline training scripts
│   ├── config/config.yaml
│   ├── data/                # CSV dataset
│   ├── model/               # PyTorch + ONNX pipeline
│   └── src/                 # Training utilities
```

---

## **Installation & Setup**

### **Option 1: Manual Setup (No Docker Yet)**

#### **Step 1: Install Dependencies**

- **Python 3.10+**
- **Go 1.21+**

Install Python packages:

```sh
cd training
pip install -r requirements.txt
pip install onnxruntime fastapi uvicorn pydantic-settings
```

Install Go modules (from the `gateway` directory):

```sh
go mod tidy
```

---

#### **Step 2: Run the Inference Server**

```sh
cd inference_service
uvicorn main:app --host 0.0.0.0 --port 8000
```

This launches the **FastAPI** backend which serves predictions from the ONNX model.

---

#### **Step 3: Run the Gateway UI**

```sh
cd gateway/cmd/gateway
go run main.go
```

Access the UI at `http://localhost:8020`. The form will submit to the API and display results dynamically using HTMX.

---

## **API Endpoints**

### **POST /predict**

**Request Body**:
```json
{
  "annual_income": 800000,
  "loan_amount": 1000000,
  "loan_term": 15,
  "credit_score": 720
}
```

**Response**:
```json
{
  "approval_status": "approved",
  "confidence": 0.87
}
```

Validation is enforced via Pydantic validators with field-level constraints (non-negative values, valid score range, etc.).

---

## Key Technologies

- **PyTorch / ONNX** – Offline model training + runtime-optimized inference
- **FastAPI** – REST API serving the ONNX model
- **Go** – Lightweight web server and client using standard lib + `net/http`
- **HTMX + Bootstrap** – Dynamic, responsive HTML UI
- **ONNX Runtime** – High-speed inference execution engine
- **Viper** – Environment configuration and rate limiter setup

---

## Educational Goals

This project is designed as a **learning resource** for:

- Training and exporting PyTorch models with ONNX
- Building REST APIs with FastAPI and Pydantic validation
- Using Go to implement basic web services and form interfaces
- Integrating Python inference with Go-based UI
- Applying rate-limiting middleware in Go
- Designing modular and maintainable ML services

---

## Missing Features

- **No Test Coverage**: This project currently lacks unit or integration tests. If deployed in production, test coverage for both the Go gateway and Python API should be added.
- **No Containerization**: There are no Dockerfiles or Compose configurations. Containerization is planned for future versions.

---

## Suggested Extensions

- **Add Testing**: Write unit tests for:
  - FastAPI models and routes
  - ONNX inference logic
  - Go middleware and service validation
- **Containerize the Project**:
  - Add Dockerfiles for both services
  - Define `docker-compose.yml` for orchestration
- **Deploy to a Cloud Provider**:
  - Deploy the FastAPI service on **AWS**, **GCP**, or **Render**
  - Host the Go UI on **VPS**, **Cloud Run**, or **Netlify (static wrapper)**
- **Add Logging and Metrics**:
  - Use **Prometheus** or **OpenTelemetry** for monitoring
  - Add structured logging for Go and Python