# Messaging System 

This project demonstrates a simple but scalable architecture for managing message delivery using Go, PostgreSQL, Redis, and background workers.

This application runs on localhost:8080. You can change the port via docker files. Make sure your pgsql and redis ports are empty.

---

## Getting Started

To spin up everything:

```bash
make rebuild
```

What this does:

* Brings up PostgreSQL and Redis containers
* Applies any pending database migrations
* Starts the main application
* Kicks off the background worker automatically

---

## Test Data (Faker)

You can easily generate sample messages:

```bash
make faker
```

This will insert a bunch of fake unsent messages into the database, useful for testing the worker.

---

## Worker Management

You can control the background worker via HTTP:

```bash
POST /start   # Start the worker
POST /stop    # Stop the worker
```

The worker looks for unsent messages (`is_sent = false`), sends them to a webhook, and marks them as sent if delivery is successful.

---

## API Endpoints

| Method | Endpoint             | Description           |
| ------ |----------------------| --------------------- |
| GET    | `/messages`          | Returns sent messages |
| POST   | `/start` | Starts the worker     |
| POST   | `/stop`  | Stops the worker      |

---

## Technical Notes

* Written in Go (Golang)
* Docker + Makefile for easy local setup
* PostgreSQL is used for message storage
* Redis is used to log successfully sent messages
* Exponential backoff strategy for retrying failed webhook requests
* Worker logic driven by time.Ticker and context.Context
* Base repository pattern for clean DB access
* JSON response format is consistent across all endpoints

---

## Design Philosophy

* Clean separation of concerns: handler, service, repository
* Designed with testing and observability in mind
* Emphasis on simplicity: no frameworks, minimal dependencies
* Supports bulk message generation for load testing
* Retry logic simulates production-grade reliability

---

## Swagger

Swagger documentation is available at the `/docs/index.html` endpoint.

