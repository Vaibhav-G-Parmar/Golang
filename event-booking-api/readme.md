# Event Booking API - ** OnGoing Development **

A RESTful API for managing events and bookings built with Go. Provides user authentication, event management, and booking operations.

<img width="200" height="320" alt="golang-api" src="https://github.com/user-attachments/assets/4619df65-a4c7-4221-8cc1-eb852adaa1b4" />

## Features
- CRUD for events (create, read, update, delete)
- User registration and JWT-based authentication - * upcoming *

## Tech stack
- Language: Go
- Database: SQLite
- Authentication: JWT * upcoming *

## Prerequisites
- Go 1.20+

## Run locally (after cloning)

1. Clone the repo and enter the project directory
```bash
git clone <repo-url>
cd <repo-directory>/event-booking-api

# replace <repo-url> and <repo-directory> accordingly
```

2. Install dependencies
```bash
go mod download
```

3. Run the server
```bash
go run .
```

4. Open the API
```text
http://localhost:8080
```

## API Endpoints
- `POST /events` - Create a new event
- `GET /events` - Get all events
- `GET /events/:id` - Get event by ID
- `PUT /events/:id` - Update event by ID
- `DELETE /events/:id` - Delete event by ID

Please refer to `api-test` folder for sample requests.

## Testing the API
You can use tools like Postman, curl, or HTTP to test the API endpoints.

## Database
The application uses SQLite for data storage. The database file `events.db` will be created automatically in the project directory when you run the application for the first time.


