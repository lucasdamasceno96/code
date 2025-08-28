# Go URL Shortener

A robust and scalable URL shortening service built with Go, following Clean Architecture principles. This project was developed as a hands-on exercise to apply concepts of layered architecture, dependency injection, and testing in a real-world application.

## Table of Contents

  - [Introduction](https://www.google.com/search?q=%23introduction)
  - [Features](https://www.google.com/search?q=%23features)
  - [Tech Stack & Architecture](https://www.google.com/search?q=%23tech-stack--architecture)
  - [Project Structure](https://www.google.com/search?q=%23project-structure)
  - [Getting Started](https://www.google.com/search?q=%23getting-started)
      - [Prerequisites](https://www.google.com/search?q=%23prerequisites)
      - [Installation & Running](https://www.google.com/search?q=%23installation--running)
  - [API Usage](https://www.google.com/search?q=%23api-usage)
  - [My Learning Journey](https://www.google.com/search?q=%23my-learning-journey)
      - [From Tutorial to Working Application](https://www.google.com/search?q=%23from-tutorial-to-working-application)
      - [Key Challenges & Solutions](https://www.google.com/search?q=%23key-challenges--solutions)
  - [Testing Strategy](https://www.google.com/search?q=%23testing-strategy)
  - [Future Improvements](https://www.google.com/search?q=%23future-improvements)

## Introduction

This project is a fully functional URL shortener API. It takes a long URL and generates a unique, shorter alias that redirects to the original URL. The primary goal was not just to build the service, but to do so using best practices in software engineering, creating a codebase that is clean, maintainable, and easily testable.

## Features

  - **URL Shortening:** Converts long URLs into a compact format.
  - **URL Redirection:** Redirects short codes to their original URL.
  - **Persistent Storage:** Uses SQLite to store URL mappings.
  - **RESTful API:** Provides simple and clear API endpoints.
  - **Layered Architecture:** Ensures separation of concerns and modularity.

## Tech Stack & Architecture

  - **Language:** **Go (Golang)**
  - **API Framework:** **gorilla/mux** for HTTP routing.
  - **Database:** **SQLite 3** for lightweight and file-based storage.
  - **Driver:** `mattn/go-sqlite3`

The application follows a **Clean Architecture** approach, separating the code into four distinct layers:

1.  **Domain:** Contains the core data structures (e.g., `ShortURL` struct). It has no dependencies on other layers.
2.  **Repository:** Manages all communication with the database. It implements an interface defined by the service layer, abstracting away the database technology.
3.  **Service:** Contains the core business logic (e.g., generating a unique short code, validating data). It orchestrates the flow of data between the handlers and the repository.
4.  **Handler (HTTP):** Manages incoming HTTP requests, decodes JSON payloads, calls the appropriate service methods, and formats the HTTP response.

## Project Structure

```
url-shortener/
├── cmd/api/
│   └── main.go           # Application entry point
├── internal/
│   ├── config/           # Configuration structs
│   ├── domain/           # Core data models
│   ├── handler/          # HTTP request handlers
│   ├── repository/
│   │   └── sqlite/       # SQLite implementation of the repository
│   ├── service/          # Business logic
│   └── util/             # Utility functions (e.g., code generation)
├── data/                 # Directory for the SQLite database file
└── go.mod
```

## Getting Started

### Prerequisites

  - Go 1.19+
  - Git
  - A command-line tool like `curl` for testing.

### Installation & Running

1.  **Clone the repository:**

    ```bash
    git clone <your-repo-url>
    cd url-shortener
    ```

2.  **Install dependencies:**
    The project uses Go Modules, so dependencies will be downloaded automatically.

3.  **Create the data directory:**
    The SQLite driver can create the database file, but not the directory it resides in.

    ```bash
    mkdir data
    ```

4.  **Run the server:**

    ```bash
    go run cmd/api/main.go
    ```

    The server will start on `http://localhost:8080`.

## API Usage

### 1\. Shorten a URL

  - **Endpoint:** `POST /shorten`
  - **Request Body:**
    ```json
    {
      "url": "https://your-very-long-url.com/with/some/path"
    }
    ```
  - **Example with `curl`:**
    ```bash
    curl -X POST -H "Content-Type: application/json" \
    -d '{"url": "https://www.google.com/search?q=golang"}' \
    http://localhost:8080/shorten
    ```
  - **Success Response (201 Created):**
    ```json
    {
      "short_url": "http://localhost:8080/k8s-12345"
    }
    ```

### 2\. Redirect to Original URL

  - **Endpoint:** `GET /{short-code}`
  - **Example with `curl`:**
    ```bash
    curl -i http://localhost:8080/k8s-12345
    ```
  - **Success Response:** The server will respond with a `302 Found` redirect to the original URL.

## My Learning Journey

This project was a significant learning experience in practical software development. The initial phase involved following a tutorial, but the real learning began when I had to debug and extend the application.

### From Tutorial to Working Application

I started with a guide that laid out the layered architecture. Implementing each layer (Domain, Repository, Service, Handler) provided a clear understanding of **Separation of Concerns**. This structure made the code much easier to reason about compared to a monolithic approach.

### Key Challenges & Solutions

The path to a working application was filled with valuable debugging challenges:

1.  **Database Connection Issues:** My first major roadblock was that the application couldn't save data, returning a generic "Failed to create short URL" error.

      - **Solution:** I learned how to add **structured logging** across all layers. By placing `log.Printf` statements in the handler, service, and repository, I traced the error back to its source. The logs revealed specific SQL syntax errors, first a typo (`INSET` instead of `INSERT`) and then a column name mismatch (`short_url` instead of `short_code`). This taught me that good logging is essential for diagnostics.

2.  **HTTP Handler Bugs:** I encountered an `http: superfluous response.WriteHeader call` error.

      - **Solution:** I discovered that `http.Error()` implicitly writes a header. My mistake was not placing a `return` statement immediately after it, causing the code to continue execution and attempt a second `WriteHeader` call in the success path. This was a crucial lesson in understanding the Go `net/http` package's behavior.

3.  **Understanding Dependency Injection:** Initially, the concept was abstract. But when writing tests, it clicked. By having the service depend on a `repository.URLRepository` interface instead of a concrete `sqlite.Repo` struct, I was able to "mock" the repository during tests. This allowed me to test my business logic in complete isolation from the database, which was a powerful realization.

## Testing Strategy

The project includes unit tests to ensure the reliability of its core logic. The tests focus on individual packages to verify their functionality in isolation.

  - **Util Package:** A simple unit test verifies that the `GenerateShortCode` function produces a code with the correct prefix and length.
  - **Service Package:** This test is more advanced and uses a **mock repository**. By creating a `mockRepository` struct that satisfies the `URLRepository` interface, I can simulate database behaviors (like successfully saving data or reporting that a code doesn't exist) without needing a real database connection.

### How to Run Tests

From the root directory, run the following command to execute all tests in the project:

```bash
go test ./...
```

## Future Improvements

This project has a solid foundation that can be extended with many new features:

  - **Add Analytics:** Track the number of clicks for each short URL.
  - **URL Expiration:** Implement a feature to make short URLs expire after a certain time.
  - **Custom Short Codes:** Allow users to choose their own custom alias for a URL.
  - **Switch to a Robust Database:** For a production environment, migrate from SQLite to PostgreSQL or MySQL.