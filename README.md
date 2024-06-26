# LeetCode REST API Wrapper (LRAW) 🚀

## Description 📝
LRAW is a Go-based REST API that simplifies access to LeetCode user data. It wraps LeetCode's GraphQL API, providing an easy-to-use REST interface for developers.

## Features ✨
- Fetch user information by username
- Simple REST API endpoints
- Dockerized for easy deployment
- Modular codebase structure

## Why this project? 🤓

### Problem 🤔
Another day I was making a platform for LeetCode users, but I realized there's no official API from LeetCode. Some people have made their own ways to do this, but I didn't really like them. LeetCode is nice, but it does not have a REST API (which is easy to use for the majority). It provides GraphQL, but not everyone is familiar with it, and LeetCode does not officially provide a `schema` even for that GraphQL. This lack of documentation and the complexity of GraphQL can be a significant hurdle for many developers who want to build applications using LeetCode data.

### Solution 💡
To fix this, I made this project. I like using Go, so I wrote it in Go. I spent some time figuring out how to get data from LeetCode, and I found the right GraphQL queries to use.

This project changes those hard-to-use GraphQL queries into a simple REST API. Now, developers can easily get all the info about any LeetCode user by just making a simple request with the username. This makes it much easier to use LeetCode data in other apps, even if you don't know how to use GraphQL.

## Installation 🛠️
1. Clone the repository
```bash
git clone https://github.com/mutasim/leetcode-api.git
```
2. Navigate to the project directory
```bash
cd leetcode-api
```
3. Install dependencies
```bash
go mod tidy
```

## Usage 🖥️

### Running locally 🐌
1. Start the server
```bash
go run cmd/server/main.go
```
2. Make a GET request to `http://localhost:8080/user/{username}`

> [!TIP]
> By default, the server runs on port `8080`. You can change this by setting the `PORT` environment variable:
> 
> `PORT=3000 go run cmd/server/main.go`
> 
> This would start the server on port `3000`.

### Using Docker 🐳
1. Build the Docker image
```bash
docker build -t leetcode-api .
```
2. Run the container
```bash
docker run -p 8080:8080 leetcode-api
```

## API Endpoints 🛣️
- GET `/user/{username}`: Fetch user information

## Development 💻
- To run tests:
```bash
go test ./test
```
- To build the binary:
```bash
go build -o leet-api cmd/server/main.go
```

## Deployment 🌐
LRAW can be easily deployed on platforms like:
- Google Cloud Run
- AWS ECS (Elastic Container Service)
- Heroku (with container runtime)
- DigitalOcean App Platform

Refer to each platform's documentation for specific deployment instructions.

## License 📄
This project is [MIT](./LICENCE) licensed.
