[![Docker Compose CI](https://github.com/afteracademy/goserve-example-api-server-mongo/actions/workflows/docker_compose.yml/badge.svg)](https://github.com/afteracademy/goserve-example-api-server-mongo/actions/workflows/docker_compose.yml)
[![Architechture](https://img.shields.io/badge/Framework-blue?label=View&logo=go)](https://github.com/afteracademy/goserve)
[![Starter Project](https://img.shields.io/badge/Starter%20Project%20CLI-red?label=Get&logo=go)](https://github.com/afteracademy/goservegen)
[![Download](https://img.shields.io/badge/Download-Starter%20Project%20Mongo%20Zip-green.svg)](https://github.com/afteracademy/goservegen/raw/main/starter-project-mongo.zip)

<div align="center">

# MongoDB API Server Example

### Production-Ready Blog Service with GoServe Framework

![Banner](.extra/docs/goserve-banner.png)

**A complete, production-ready REST API service built with GoServe framework, MongoDB, Redis, JWT authentication, and role-based authorization.**

[![Documentation](https://img.shields.io/badge/📚_Read_Documentation-goserve.afteracademy.com-blue?style=for-the-badge)](http://goserve.afteracademy.com/mongo)

---
[![GoServe Framework](https://img.shields.io/badge/🚀_Framework-GoServe-blue?style=for-the-badge)](https://github.com/afteracademy/goserve)
[![API Documentation](https://img.shields.io/badge/📚_API_Docs-View_Here-blue?style=for-the-badge)](https://documenter.getpostman.com/view/1552895/2sA3XWdefu)
[![Download Starter](https://img.shields.io/badge/⬇️_Download-Starter_Project-green?style=for-the-badge)](https://github.com/afteracademy/goservegen/raw/main/starter-project-mongo.zip)
---
</div>

## Overview

This project is a fully production-ready blog service demonstrating best practices for building performant and secure backend REST API services with MongoDB. It showcases the application of the [GoServe framework](https://github.com/afteracademy/goserve) with clean architecture, feature separation, comprehensive testing, and production grade security.

## Features

- **GoServe Framework** - Built on the production-ready [GoServe v2](https://github.com/afteracademy/goserve) framework
- **Clean Architecture** - Well-structured, maintainable codebase with clear separation of concerns
- **MongoDB Integration** - Full MongoDB support with query builder and document validation
- **Redis Caching** - High-performance caching layer for frequently accessed data
- **JWT Authentication** - Secure token-based authentication with refresh tokens
- **Role-Based Authorization** - Fine-grained access control with role management
- **API Key Support** - Additional security layer for API access control
- **Request Validation** - Comprehensive input validation using validator v10
- **Testing Suite** - Extensive unit and integration test coverage
- **Docker Ready** - Complete Docker Compose setup for easy deployment
- **Auto-Generated APIs** - CLI tool for scaffolding new API endpoints
- **Type-Safe DTOs** - Structured data transfer objects for all requests/responses

## Technology Stack

- **Language**: Go 1.21+
- **Framework**: [GoServe v2](https://github.com/afteracademy/goserve)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: MongoDB ([mongo-driver](https://github.com/mongodb/mongo-go-driver))
- **Cache**: Redis ([go-redis](https://github.com/redis/go-redis))
- **Authentication**: JWT tokens
- **Validation**: [validator](https://github.com/go-playground/validator)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Testing**: [Testify](https://github.com/stretchr/testify)

## Quick Start

### Prerequisites

- Docker & Docker Compose ([Installation Guide](https://docs.docker.com/install/))
- Go 1.21+ (for local development)

### Installation

**1. Clone the Repository**

```bash
git clone https://github.com/afteracademy/goserve-example-api-server-mongo.git
cd goserve-example-api-server-mongo
```

**2. Generate RSA Keys**
```bash
go run .tools/rsa/keygen.go
```

**3. Create Environment Files**
```bash
go run .tools/copy/envs.go 
```

**4. Start with Docker Compose**
```bash
docker compose up --build
```

The API server will be available at: **http://localhost:8080**

**5. Health Check**
```bash
docker inspect --format='{{.State.Health.Status}}' goserver-mongo
```

**6. Run Tests**
```bash
docker exec -t goserver-mongo go test -v ./...
```

### Troubleshooting

If you encounter issues:
- Ensure port 8080 is available (change `SERVER_PORT` in `.env` if needed)
- Ensure port 27017 is available (change `DB_PORT` in `.env` if needed)
- Ensure port 6379 is available (change `REDIS_PORT` in `.env` if needed)

## Local Development

For local development without Docker:

```bash
go mod tidy
```

Keep Docker containers for `mongo` and `redis` running, but stop the `goserver-mongo` container.

Update the following in `.env` and `.test.env`:
```env
DB_HOST=localhost
REDIS_HOST=localhost
```

**Run the application:**
```bash
go run cmd/main.go
```

**Or use VS Code**: Use the `Run and Debug` panel for an enhanced development experience.

## Architecture

### Design Principles

The architecture is designed to make each API independent while sharing services among them. This promotes:
- **Code Reusability** - Shared services across multiple endpoints
- **Team Collaboration** - Reduced conflicts when working in teams
- **Feature Isolation** - Easier testing and maintenance

### Request Flow

![Request-Response-Design](.extra/docs/request-flow.svg)

**Startup Flow:**  
`cmd/main` → `startup/server` → `module, mongo, redis, router` → `api/[feature]/middlewares` → `api/[feature]/controller` → `api/[feature]/service` → `authentication, authorization` → `handlers` → `response`

### API Structure

```
Sample API
├── dto/
│   └── create_sample.go     # Data Transfer Objects
├── model/
│   └── sample.go            # MongoDB Collection Model
├── middleware/              # (Optional) Feature-specific middleware
│   └── custom.go
├── controller.go            # Route definitions & handlers
└── service.go              # Business logic & data operations
```

**Key Components:**
- **DTOs** - Request/response body definitions in `dto/` directory
- **Models** - MongoDB collection schemas in `model/` directory
- **Controller** - Defines endpoints and handles HTTP requests
- **Service** - Contains business logic and data operations
- **Middleware** - Authentication, authorization, and custom middleware

### Project Structure

| Directory | Purpose |
|-----------|---------|
| **api/** | Feature-based API implementations |
| **cmd/** | Application entry point (main.go) |
| **common/** | Shared code across all APIs |
| **config/** | Environment variable configuration |
| **keys/** | RSA keys for JWT token signing |
| **startup/** | Server initialization, DB, Redis, routing |
| **tests/** | Integration test suites |
| **utils/** | Utility functions |

**Helper Directories:**
- **.extra/** - MongoDB initialization scripts, assets, documentation
- **.github/** - CI/CD workflows
- **.tools/** - Code generators, key generation utilities
- **.vscode/** - Editor configuration and debug settings

## Generate New APIs

Scaffold a new API endpoint with a single command:

```bash
go run .tools/apigen.go sample
```

This creates the complete structure under `api/sample/` with:
- Model definitions
- DTO templates
- Controller skeleton
- Service interface

## API Documentation

<div align="center">

[![API Documentation](https://img.shields.io/badge/📚_View_Full_API_Documentation-blue?style=for-the-badge)](https://documenter.getpostman.com/view/1552895/2sA3XWdefu)

Complete API documentation with request/response examples and authentication details

</div>

## Code Examples

### Model

`api/sample/model/sample.go`

```go
package model

import (
  "context"
  "time"

  "github.com/go-playground/validator/v10"
  "github.com/afteracademy/goserve/v2/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  mongod "go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "samples"

type Sample struct {
  ID        primitive.ObjectID `bson:"_id,omitempty" validate:"-"`
  Field     string             `bson:"field" validate:"required"`
  Status    bool               `bson:"status" validate:"required"`
  CreatedAt time.Time          `bson:"createdAt" validate:"required"`
  UpdatedAt time.Time          `bson:"updatedAt" validate:"required"`
}

func NewSample(field string) (*Sample, error) {
  time := time.Now()
  doc := Sample{
    Field:     field,
    Status:    true,
    CreatedAt: time,
    UpdatedAt: time,
  }
  if err := doc.Validate(); err != nil {
    return nil, err
  }
  return &doc, nil
}
}

func (doc *Sample) GetValue() *Sample {
  return doc
}

func (doc *Sample) Validate() error {
  validate := validator.New()
  return validate.Struct(doc)
}

func (*Sample) EnsureIndexes(db mongo.Database) {
  indexes := []mongod.IndexModel{
    {
      Keys: bson.D{
        {Key: "_id", Value: 1},
        {Key: "status", Value: 1},
      },
    },
  }
  
  mongo.NewQueryBuilder[Sample](db, CollectionName).Query(context.Background()).CreateIndexes(indexes)
}
```

**Model Interface:** Implements `github.com/afteracademy/goserve/v2/mongo/database.Document[T]`

```golang
type Document[T any] interface {
  EnsureIndexes(Database)
  GetValue() *T
  Validate() error
}
``` 

### DTO

`api/sample/dto/create_sample.go`

```go
package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoSample struct {
	ID        primitive.ObjectID `json:"_id" binding:"required"`
	Field     string             `json:"field" binding:"required"`
	CreatedAt time.Time          `json:"createdAt" binding:"required"`
}
```

### Service

`api/sample/service.go`

```go
package sample

import (
  "github.com/afteracademy/goserve-example-api-server-mongo/api/sample/dto"
  "github.com/afteracademy/goserve-example-api-server-mongo/api/sample/model"
  "github.com/afteracademy/goserve/v2/mongo"
  "github.com/afteracademy/goserve/v2/redis"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
  FindSample(id primitive.ObjectID) (*model.Sample, error)
}

type service struct {
  sampleQueryBuilder mongo.QueryBuilder[model.Sample]
  infoSampleCache    redis.Cache[dto.InfoSample]
}

func NewService(db mongo.Database, store redis.Store) Service {
  return &service{
    sampleQueryBuilder: mongo.NewQueryBuilder[model.Sample](db, model.CollectionName),
    infoSampleCache: redis.NewCache[dto.InfoSample](store),
  }
}

func (s *service) FindSample(id primitive.ObjectID) (*model.Sample, error) {
  filter := bson.M{"_id": id}

  msg, err := s.sampleQueryBuilder.SingleQuery().FindOne(filter, nil)
  if err != nil {
    return nil, err
  }

  return msg, nil
}
```

**Key Features:**
- **Database Query:** `mongo.QueryBuilder[model.Sample]` provides type-safe MongoDB operations
- **Redis Cache:** `redis.Cache[dto.InfoSample]` provides type-safe caching operations

### Controller

`api/sample/controller.go`

```go
package sample

import (
  "github.com/gin-gonic/gin"
  "github.com/afteracademy/goserve-example-api-server-mongo/api/sample/dto"
  "github.com/afteracademy/goserve-example-api-server-mongo/common"
  coredto "github.com/afteracademy/goserve/v2/dto"
  "github.com/afteracademy/goserve/v2/network"
  "github.com/afteracademy/goserve-example-api-server-mongo/utils"
)

type controller struct {
  network.Controller
  common.ContextPayload
  service Service
}

func NewController(
  authMFunc network.AuthenticationProvider,
  authorizeMFunc network.AuthorizationProvider,
  service Service,
) network.Controller {
  return &controller{
    Controller: network.NewController("/sample", authMFunc, authorizeMFunc),
    ContextPayload: common.NewContextPayload(),
    service:  service,
  }
}

func (c *controller) MountRoutes(group *gin.RouterGroup) {
  group.GET("/id/:id", c.getSampleHandler)
}

func (c *controller) getSampleHandler(ctx *gin.Context) {
  mongoId, err := network.ReqParams[coredto.MongoId](ctx)
  if err != nil {
    network.SendBadRequestError(ctx, err.Error(), err)
    return
  }

  sample, err := c.service.FindSample(mongoId.ID)
  if err != nil {
    network.SendNotFoundError(ctx, "sample not found", err)
    return
  }

  data, err := utils.MapTo[dto.InfoSample](sample)
  if err != nil {
    network.SendInternalServerError(ctx, "something went wrong", err)
    return
  }

  network.SendSuccessDataResponse(ctx, "success", data)
}
```

**Controller Interface:** Implements `github.com/afteracademy/goserve/v2/network.Controller`

```golang
type Controller interface {
  BaseController
  MountRoutes(group *gin.RouterGroup)
}

type BaseController interface {
  ResponseSender
  Path() string
  Authentication() gin.HandlerFunc
  Authorization(role string) gin.HandlerFunc
}
``` 

### Register Controller

`startup/module.go`

```go
import (
  ...
  "github.com/afteracademy/goserve-example-api-server-mongo/api/sample"
)

...

func (m *module) Controllers() []network.Controller {
  return []network.Controller{
    ...
    sample.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), sample.NewService(m.DB, m.Store)),
  }
}
```

### Setup Indexes

`startup/indexes.go`

```go
import (
  ...
  sample "github.com/afteracademy/goserve-example-api-server-mongo/api/sample/model"
)

func EnsureDbIndexes(db mongo.Database) {
  go mongo.Document[sample.Sample](&sample.Sample{}).EnsureIndexes(db)
  ...
}
```

## Related Projects

Explore other GoServe example implementations:

1. **[GoServe Framework](https://github.com/afteracademy/goserve)**  
   Core framework with PostgreSQL, MongoDB, Redis, and NATS support

2. **[PostgreSQL API Server](https://github.com/afteracademy/goserve-example-api-server-postgres)**  
   Complete REST API with PostgreSQL and clean architecture

3. **[Microservices Example](https://github.com/afteracademy/gomicro)**  
   NATS-based microservices communication patterns

## Generate Starter Project

Use the [GoServeGen](https://github.com/afteracademy/goservegen) CLI to generate a starter project:

```bash
# Install GoServeGen CLI
go install github.com/afteracademy/goservegen@latest

# Generate a new project
goservegen create my-project --db=mongo
```

Or download the starter project directly:
- [MongoDB Starter Project (ZIP)](https://github.com/afteracademy/goservegen/raw/main/starter-project-mongo.zip)

## Articles & Tutorials

- [How to Architect Good Go Backend REST API Services](https://afteracademy.com/article/how-to-architect-good-go-backend-rest-api-services)
- [How to Create Microservices — A Practical Guide Using Go](https://afteracademy.com/article/how-to-create-microservices-a-practical-guide-using-go)
- [Implement JSON Web Token (JWT) Authentication using AccessToken and RefreshToken](https://afteracademy.com/article/implement-json-web-token-jwt-authentication-using-access-token-and-refresh-token)

## Contributing

We welcome contributions! Please feel free to:

- **Fork** the repository
- **Open** issues for bugs or feature requests
- **Submit** pull requests with improvements
- **Share** your feedback and suggestions

## Learn More

Subscribe to **AfterAcademy** on YouTube for in-depth tutorials and concept explanations:

[![YouTube](https://img.shields.io/badge/YouTube-Subscribe-red?style=for-the-badge&logo=youtube&logoColor=white)](https://www.youtube.com/@afteracad)

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support This Project

If you find this project useful, please consider:

- **Starring** ⭐ this repository
- **Sharing** with the community
- **Contributing** improvements
- **Reporting** bugs and issues

---

<div align="center">

**Built with love by [AfterAcademy](https://github.com/afteracademy)**

[GoServe Framework](https://github.com/afteracademy/goserve) • [API Documentation](https://documenter.getpostman.com/view/1552895/2sA3XWdefu) • [Articles](https://afteracademy.com) • [YouTube](https://www.youtube.com/@afteracad)

</div>