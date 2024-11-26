
# Go Gin REST starter

![GitHub Actions workflow status badge](https://github.com/eugenius1/go-gin-rest/actions/workflows/go.yml/badge.svg)
[![codebeat badge](https://codebeat.co/badges/58a7af95-d6ed-45b0-a307-487f22ef9a67)](https://codebeat.co/projects/github-com-eugenius1-go-gin-rest-main)

A backend starter project for implementing REST APIs in Go and using Gin framework, including:

- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin) as web framework
- PostgreSQL as database
- [GORM](https://pkg.go.dev/gorm.io/gorm) for ORM (object–relational mapping)
- [ULID](https://pkg.go.dev/github.com/oklog/ulid/v2) (Universally Unique Lexicographically Sortable Identifier) for IDs
- A devcontainer with dependencies that can be used with VS Code or independently with [docker compose](.devcontainer/docker-compose.yml)

The starter example implements the following endpoints:

```
GET /v1/albums
POST /v1/albums
PUT /v1/albums
GET /v1/albums/:id
DELETE /v1/albums/:id
```

TODO:

- [ ] Logging
- [ ] JSON validation
- [ ] OpenAPI/Swagger
- [ ] API request examples in the repo with [Bruno](https://www.usebruno.com/)
- [ ] Mock generation from interfaces
- [ ] Add unit tests
