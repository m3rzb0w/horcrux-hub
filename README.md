# Horcrux-Hub

Welcome to Horcrux-Hub! This project allows you to manage your Horcruxes conveniently. To get started, follow the instructions below.

## Prerequisites

Before running this project locally, ensure that you have either [Docker](https://www.docker.com) and/or [Node](https://nodejs.org) and [Go](https://go.dev/) installed on your system.

Add these environment variables in the front folder in a .env file :

```bash
touch front/.env.local
echo "VITE_WS_URL=ws://localhost:8080/countdown" > front/.env.local
echo "VITE_WS_CHAT=ws://localhost:8080/ws" >> front/.env.local
```
---

### Using Go/Node

Launch the entire project with the Makefile:

```bash
make run
```
Launch only the frontend:

```bash
cd front/
npm run dev
```
---
### Using Docker

Build the Docker image:

```bash
docker build -t horcrux-hub .
```
Run the Docker container:

```bash
docker run -p 8080:8080 horcrux-hub:latest
```

---

Feel free to choose the method that suits your preference or existing development environment.

Enjoy managing your Horcruxes with Horcrux-Hub! 🧙