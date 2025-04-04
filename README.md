# DigiRoboticApp-Backend

Robot Control API service built with Go.

## Tech Stack
- Go 1.19+
- Docker
- Gin Framework

## Development Setup

### Prerequisites
- Docker Desktop
- Go 1.19+
- Git

### Setup Instructions
```bash
# Clone the repository
git clone git@github.com:sonnyvedder/DigiRoboticApp-Backend.git
cd DigiRoboticApp-Backend

# Create environment file
copy .env.example .env

# Initialize Go modules
go mod init digiroboticapp
go mod tidy

# Start the service
docker-compose up -d