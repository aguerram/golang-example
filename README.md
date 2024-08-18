# Structure

myapp/
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── app/
│   │   ├── routes/
│   │   │   └── routes.go
│   │   ├── handlers/
│   │   │   ├── api/
│   │   │   │   ├── product_handler.go
│   │   │   │   └── user_handler.go
│   │   │   └── web/
│   │   │       ├── home_handler.go
│   │   │       └── about_handler.go
│   │   ├── services/
│   │   │   ├── product_service.go
│   │   │   ├── user_service.go
│   │   │   └── cache_service.go
│   │   ├── repositories/
│   │   │   ├── product_repository.go
│   │   │   └── user_repository.go
│   │   ├── models/
│   │   │   ├── product.go
│   │   │   └── user.go
│   │   ├── dtos/
│   │   │   ├── request/
│   │   │   │   ├── product_request.go
│   │   │   │   └── user_request.go
│   │   │   └── response/
│   │   │       ├── product_response.go
│   │   │       └── user_response.go
│   │   ├── middleware/
│   │   │   └── auth_middleware.go
│   │   ├── utils/
│   │   │   ├── validator.go
│   │   │   └── json_utils.go
│   │   ├── server/
│   │   │   └── server.go
│   │   └── templates/
│   │       ├── home.html
│   │       └── about.html
└── pkg/
    ├── database/
    │   ├── connection.go
    │   └── migrations.go
    └── redis/
        └── client.go
