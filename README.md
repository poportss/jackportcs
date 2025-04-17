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
DB_CONN="your_db_con"
DB_MAX_CONN="your_port"
API_PORT="your_port"
APP_PAGARME_API="pagarme_key"
STEAM_WEB_API_KEY="your_steam_api"
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
