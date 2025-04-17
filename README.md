# JackPortCS

JackPortCS is a Go-based API for managing CS:GO skins, Steam integration, and payment processing.

## 🚀 Technologies

- Go 1.23.4
- Gin Framework
- GORM
- PostgreSQL
- JWT Authentication
- Stripe Payment Processing
- Steam API Integration

## 📋 Prerequisites

- Go 1.23.4 or higher
- PostgreSQL
- Docker (optional)

## 🔧 Installation

1. Clone the repository:
```bash
git clone https://github.com/poportss/jackportcs.git
cd jackportcs
```

2. Install dependencies:
```bash 
go mod download
```

3. Configure environment variables:
Create a `.env` file in the project root with the following variables:
```
DB_HOST=your_host
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database
DB_PORT=your_port
```

4. Run database migrations:
```bash
go run cmd/main.go migrate
```

5. Start the server:
```bash
go run cmd/main.go
```

## 🏗️ Project Structure

```
.
├── build/              # Dockerfiles for different environments
├── cmd/                # Application entry point
├── db/                 # SQL scripts
├── internal/           # Main source code
│   ├── baseservice/    # Base services
│   ├── database/       # Database configuration
│   ├── dto/            # Data Transfer Objects
│   ├── middleware/     # Middlewares
│   ├── migrations/     # Database migrations
│   ├── models/         # Data models
│   ├── pkg/            # Main packages
│   ├── rest/           # REST server configuration
│   ├── routes/         # Route definitions
│   └── utils/          # Utilities
```

## 🛠️ Features

- User authentication
- CS:GO skins management
- Steam integration
- Stripe payment processing
- Virtual wallet system
- Cases management

## 📦 Docker

The project includes Dockerfiles for different environments:

- `Dockerfile.dev`: Development environment
- `Dockerfile.stg`: Staging environment
- `Dockerfile.prod`: Production environment

To build and run with Docker:

```bash
docker build -f build/Dockerfile.dev -t jackportcs .
docker run -p 8080:8080 jackportcs
```

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details. 