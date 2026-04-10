# GoArticle

Article management backend built with Go, Fiber, PostgreSQL, Redis, and Clean Architecture.

This project was created to explore how a content-oriented backend service can be structured with clear separation of concerns, local observability tooling, and production-style development practices.

## What this project does

GoArticle provides a REST API for managing articles, including article creation and paginated article retrieval.

## Highlights

- Clean Architecture structure
- REST API built with Go and Fiber
- PostgreSQL for persistent storage
- Redis for caching
- Swagger for API documentation
- Prometheus and Grafana for local observability
- Docker and Docker Compose for local setup

## Tech Stack

- Go
- Fiber
- PostgreSQL
- GORM
- Redis
- Docker
- Docker Compose
- Swagger
- Prometheus
- Grafana

## Architecture

This project follows a Clean Architecture approach, with responsibilities separated into:

- Models
- Repository
- Use Case
- Delivery

The goal is to keep business rules isolated from frameworks, storage details, and external services, making the codebase easier to test and evolve.

## Local Setup

### Requirements

- Go
- Docker
- Docker Compose
- golang-migrate CLI
- swag CLI (optional, for regenerating Swagger docs)
- mockery CLI (optional, for interface mocking)

### Run locally

```bash
git clone git@github.com:khalidalhabibie/GoArticle.git
cd GoArticle

mv env.example .env

make build
make run
make migrate.up
