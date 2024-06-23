# LeetCode REST API Wrapper (LRAW) 🚀

## Description 📝
LRAW is a Go-based REST API that simplifies access to LeetCode user data. It wraps LeetCode's GraphQL API, providing an easy-to-use REST interface for developers.

## Features ✨
- Fetch user information by username
- Simple REST API endpoints
- Dockerized for easy deployment
- Modular codebase structure

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
