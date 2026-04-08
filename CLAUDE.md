# BadgerStatus - AI Context & Engineering Guidelines

## Project Overview
BadgerStatus is an uptime monitoring and alerting platform built specifically for University of Wisconsin-Madison students. It monitors critical university infrastructure (e.g., Canvas, MyUW) and sends email alerts to subscribed students when a service goes down.

## Architecture & Tech Stack
This is a monorepo consisting of a decoupled frontend client and a backend API, orchestrated via Docker.

* **Frontend:** Next.js (App Router), React, Tailwind CSS.
* **Backend:** Go (Golang), standard library `net/http` with `go-chi/chi` for routing.
* **Database:** PostgreSQL (for persisting users, subscriptions, and incident logs).
* **Cache/Queue:** Redis (for storing current uptime status and queuing email jobs).
* **Infrastructure:** Docker & Docker Compose.
* **Third-Party APIs:** Resend / SendGrid (for email delivery).

## Repository Structure
```text
badger-status/
├── docker-compose.yml       # Orchestrates DB, Redis, Backend, Frontend
├── backend/                 # Go API and background workers
│   ├── cmd/api/             # Entry point (main.go)
│   ├── internal/            # Core business logic, DB queries, pinger
│   ├── go.mod
│   └── Dockerfile           
└── frontend/                # Next.js Application
    ├── src/app/             # Next.js App Router pages
    ├── src/components/      # Reusable UI components
    ├── package.json
    └── Dockerfile

- Make SMALL, incremental changes. Do not rewrite architecture without  
  updating PROJECT.md.                   
  - Always add or update tests for new behavior. 
  - if a test fails fix them.
