# RSS Feed Aggregator

## Overview

This project is an RSS Feed Aggregator built with Go. It uses goroutines for concurrent scraping of RSS feeds from various blogs. It also features a user authentication middleware to ensure that users are logged in before accessing certain routes.

## Features

- Concurrent scraping of RSS feeds using goroutines
- User authentication middleware
- CRUD operations for users, feeds, and feed follows
- RESTful API with routes for managing users, feeds, and posts

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/feedaggregator.git
    cd feedaggregator
    ```

2. Configure environment variables at .env file:
    ```sh
    DB_URL="postgresql://user:password@localhost:5432/feedaggregator?sslmode=disable"
    PORT=8080
    ```

3. Run the application:
    ```sh
    go run .
    ```

## API Endpoints

- **Readiness and Error Check**
    - `GET /v1/ready`
    - `GET /v1/error`

- **User Management**
    - `POST /v1/users`
    - `GET /v1/users` (Authenticated)

- **Feed Management**
    - `POST /v1/feeds` (Authenticated)
    - `GET /v1/feeds`

- **Feed Follows Management**
    - `POST /v1/feeds-follows` (Authenticated)
    - `GET /v1/feeds-follows` (Authenticated)
    - `DELETE /v1/feeds-follows/{feedFollowID}` (Authenticated)

- **Posts Management**
    - `GET /v1/posts` (Authenticated)

## Code Overview

### Scraping Feeds

The `startScraping` function uses goroutines to scrape RSS feeds concurrently. It fetches the feeds from the database, marks them as fetched, and then processes each item in the feed, parsing the publication date and storing the post in the database.

### Middleware Authentication

The `middlewareAuth` function ensures that a valid API key is provided in the request header. It retrieves the user associated with the API key from the database and passes the user to the handler function.

### Router Setup

    v1Router := chi.NewRouter()
    
    v1Router.Get("/ready", handlerReadiness)
    v1Router.Get("/error", handlerError)
    
    v1Router.Post("/users", apiCfg.handlerCreateUser)
    v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
    
    v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
    v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
    
    v1Router.Post("/feeds-follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
    v1Router.Get("/feeds-follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
    v1Router.Delete("/feeds-follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
    
    v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))
    
    router.Mount("/v1", v1Router)

### Contributing
If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
