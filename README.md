# GoArticle

Article management backend built with Go, Fiber, PostgreSQL, Redis, and Clean Architecture.

This project was built to explore how a content-oriented backend service can be structured with clear separation of concerns, local observability tooling, and production-style development practices.

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

This project follows a Clean Architecture approach, separating responsibilities across delivery, use case, repository, and model layers to keep business logic isolated from infrastructure concerns.

![Clean Architecture](clean-arch.png)

## Local Setup

### Requirements

- Go
- Docker
- Docker Compose
- golang-migrate CLI
- swag CLI
- mockery CLI

### Run locally

    git clone git@github.com:khalidalhabibie/GoArticle.git
    cd GoArticle
    mv env.example .env
    make build
    make run
    make migrate.up

### Stop the application

    make stop

## Local Services

### Swagger Documentation

    http://localhost:8081/swagger/index.html

### Prometheus

    http://localhost:9090

### Grafana

    http://localhost:3000

Grafana default credentials:

    username: admin
    password: admin

## Why this repo matters

Although this is a smaller backend project, it reflects several engineering practices that matter in real services:

- layered application design
- cache integration
- local observability
- containerized development
- API documentation
- separation of business logic and infrastructure concerns

## Current Status

This repository is archived and kept as a reference project from an earlier stage of development.

## Possible Improvements

- add authentication and authorization
- improve test coverage
- add update and delete article flows
- add search and filtering support
- improve CI workflow
- introduce request validation and rate limiting

## About

Article management backend in Go with clean architecture, PostgreSQL, Redis caching, Swagger, and observability.
