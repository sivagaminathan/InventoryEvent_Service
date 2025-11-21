# Inventory Event Service

A production-grade microservice built with Go that processes inventory-related events such as product creation, stock updates, and warehouse adjustments.  
This service exposes clean REST APIs, persists data in PostgreSQL, and publishes events to an internal queue (Kafka or a pluggable message broker).

---

## 🚀 Features

- **Golang 1.22+**
- **REST APIs (Chi or Gin)**
- **PostgreSQL for persistence**
- **Event publishing** (Kafka, NATS, or in-memory for local dev)
- **Structured logging (Zap or Zerolog)**
- **Config management using environment variables**
- **Graceful shutdown & context-aware requests**
- **Unit + Integration tests**
- **Dockerized service**
- **Makefile for local development**
- **GitHub Actions CI pipeline**
- **(Optional) Kubernetes manifests for deployment**

---

## 🧱 Architecture
