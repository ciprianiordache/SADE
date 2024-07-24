# SADE - Safe Media Delivery
![icon150.png](backend/static/img/icon150.png)

## Description

SADE (Safe Media Delivery) is a web application designed to securely deliver media content.
It utilizes a modern tech stack to provide a robust and scalable platform for media delivery.

## Features
- Secure media delivery
- User authentication and session management
- Magic link authentication system
- Stripe integration for payment
- Scalable backend using GO nad Gin framework
- Frontend built with Svelte and Vite

## Installation
###  Prerequisites
- Docker and Docker Compose
- Node.js (for local frontend development)
- Go (for local backend development)

### 1. Clone the repository:
    git clone
    cd SADE
### 2. Set up environment variable:
- Create a `.env` file in both the `frontend` and `backend` directory with the necessary environment variable
- Look at the `frontend/.env file` and `backend/.env` file.

### 3. Run the application using Docker Compose: 
    docker-compose up --build

## Usage 
Now you can access the frontend at `http://localgost:3000` and the backend at `http://localhost:8080`

### Frontend Development
To run the frontend development server:

    cd frontend
    npm install
    npm run dev

### Backend Development 
To run the backend locally:

    cd backend
    go mod download
    cd cmd
    go run .

## Contact Information 
For any inquiries or support, please contact [ciipriian5521@gmail.com].

## Contributing
Contribution are welcome!
