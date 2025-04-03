# DigiRoboticApp

Backend services for robotic control and automation system.

## Tech Stack

### Backend Services
- Go (Robot Control API)
- Python (Vision & Processing Services)

### Infrastructure
- Docker
- Redis
- RabbitMQ
- Kafka (planned)
- Elasticsearch (planned)

## Development Setup

### Prerequisites
- Docker Desktop
- Go 1.19+
- Python 3.9+
- Git

### Detailed Setup Instructions

```bash
# Clone the repository
git clone git@github.com:sonnyvedder/DigiRoboticApp.git
cd DigiRoboticApp

# Create necessary environment files
copy .env.example .env
copy go-backend\.env.example go-backend\.env
copy py-services\.env.example py-services\.env

# Initialize Go modules
cd go-backend
go mod init digiroboticapp
go mod tidy
cd ..

# Setup Python virtual environment
cd py-services
python -m venv venv
.\venv\Scripts\activate
pip install -r requirements.txt
cd ..
```