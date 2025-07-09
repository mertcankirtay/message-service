# Message Service

## Description

This service sends automated messages every 2 minutes.


## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/mertcankirtay/message-service.git
cd message-service
```

### 2. Create a .env File
You need to create a .env file in the root of the project with the following keys:

```bash
SERVICE_PORT=            # e.g., 8000
MONGO_USERNAME=          # your MongoDB username
MONGO_PASSWORD=          # your MongoDB password
REDIS_PASSWORD=          # your Redis password
GIN_MODE=release         # or "debug" for development
WEBHOOK_URL=             # your webhook endpoint
AUTH_KEY=                # your auth key or token
```

### 3. Pull the latest changes

```bash
git pull
```

### 4. Build Docker containers

```bash
docker compose build
```
### 5. Start the application

```bash
docker compose up
```

The services will start and be available on the ports defined in your .env file.

### Stopping the Application

To stop the containers, press Ctrl + C in the terminal where the app is running.
To stop and remove containers, use:

```bash
docker compose down
```