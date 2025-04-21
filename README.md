
# 🚀 Task Management API (Backend Developer Assignment)

This project is a **Task Management RESTful API** built with **Golang**, using **PostgreSQL** as the database, **JWT** for authentication, and supporting advanced features like **rate limiting** and **caching**.

---

## ✨ Features

- **User Authentication**  
  - Login with username & password (JWT based).

- **Task Management**
  - Create, Read, Update, and Delete Tasks.
  - Filtering tasks by status and due dates.
  - Pagination and sorting support.
  - Task statuses: `Pending`, `In Progress`, `Completed`.

- **Advanced Functionalities**
  - **Rate Limiting**: 60 requests per minute per user.
  - **Caching**: Frequently accessed task lists are cached for 5 minutes.
  - **Secure Passwords**: Passwords are stored securely with bcrypt.
  - **Structured Logging**: Using logrus/zap for better logs.

---

## 📚 API Endpoints

| Method | Endpoint              | Description                      |
|:------:|:---------------------- |:-------------------------------- |
| POST   | `/login`                | User login (returns JWT Token)   |
| POST   | `/tasks`                | Create a new Task                |
| GET    | `/tasks`                | Get all Tasks (with filters)     |
| GET    | `/tasks/{id}`           | Get a Task by ID                 |
| PUT    | `/tasks/{id}`           | Update a Task by ID              |
| DELETE | `/tasks/{id}`           | Delete a Task by ID              |

🔐 **All task-related routes require Bearer Token authentication.**

---

## 🛠️ Tech Stack

- **Go** (Golang) — Backend Language
- **PostgreSQL** — Database
- **JWT** — Authentication
- **bcrypt** — Password Hashing
- **logrus** or **zap** — Logging
- **go-cache** / **Redis** — Caching
- **gorilla/mux** — Router
- **golang-migrate** — Database migrations
- **GORM / sqlx** — ORM/SQL

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/BhagatKaveri/Task-system.git
cd task-management-api
```

### 2. Configure Environment Variables

Create a `.env` file or set your environment variables manually:

```env
DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=Task_Application
JWT_SECRET=your_super_secret_key
PORT=33002
```

### 3. Install Go dependencies

```bash
go mod tidy
```

### 4. Setup the Database

Run database migrations:

```bash
migrate -path ./migrations -database "postgres://postgres:Sairam@123@localhost:5433/Task_Application?sslmode=disable" up
```

**Tables:**

- `users`
- `tasks`
- `task_status` enum type

Migrations are located at: `./migrations`

---

## 🚀 Running the Application

```bash
go run main.go
```

Server will start at:

```
http://localhost:33002
```

---

## 🔑 Authentication Instructions

- First, **Login** with `/login` API to receive a JWT token.
- For every **authorized request**, pass the JWT token in the headers:

```http
Authorization: Bearer your_token_here
```

Example login credentials:

```json
{
    "username": "admin",
    "password": "admin123"
}
```

*(Pre-create user in database with hashed password.)*

---

## 📄 Example API Usage

### 1. Create a Task

```bash
POST /tasks
Authorization: Bearer {your_token}

{
  "Title": "Complete backend API",
  "Description": "Finish backend REST API development",
  "Status": "Pending",
  "DueDate": "2025-04-30T23:59:59Z"
}
```

Response:

```json
{
  "ID": 1,
  "Title": "Complete backend API",
  "Description": "Finish backend REST API development",
  "Status": "Pending",
  "DueDate": "2025-04-30T23:59:59Z",
  "CreatedAt": "2025-04-20T10:00:00Z",
  "UpdatedAt": "2025-04-20T10:00:00Z"
}
```

---

### 2. Get All Tasks with Filters

```bash
GET /tasks?page=1&limit=5&status=Pending&due_date_after=2025-04-01&sort_by=due_date&sort_order=asc
Authorization: Bearer {your_token}
```

Response:

```json
{
  "tasks": [
    {
      "ID": 1,
      "Title": "Complete backend API",
      "Status": "Pending",
      "CreatedAt": "2025-04-20T10:00:00Z",
      "UpdatedAt": "2025-04-20T10:00:00Z"
    }
  ],
  "page": 1,
  "limit": 5,
  "total": 10
}
```

---

## 🧪 Testing

Run unit tests using:

```bash
go test ./...
```

Coverage includes:

- Task CRUD Operations
- JWT Authentication
- Error handling

---

## 📂 Project Structure

```
/task-management-api
│
├── cmd/              # Entry point (main.go)
├── config/           # Configuration loader
├── handler/          # HTTP request handlers
├── middleware/       # JWT middleware, Rate limiting
├── model/            # Database models
├── service/          # Business logic
├── db/               # Database interaction (SQL, ORM)
├── migrations/       # SQL migrations
├── utils/            # Helper functions (token, password hashing)
├── go.mod / go.sum   # Go modules
└── README.md         # Documentation
```

---

## 📦 Advanced Features (Planned/Future)

- Background Jobs for overdue task reminders (Emails).
- Full-text search on tasks.
- Audit Logs (Create, Update, Delete history).
- Dockerized Deployment.
- GitHub Actions CI/CD setup.

---

## 🛡️ Security Considerations

- Passwords are hashed with **bcrypt** before storing.
- JWT tokens are securely signed.
- Rate limiting is enabled to prevent brute-force attacks.
- Validation is performed for all user inputs.

---
## 📄 Swagger & Postman

- Swagger: [`swagger.yaml`](./swagger.yaml)
- Postman: [`Task system.postman_collection.json`](./Task%20system.postman_collection.json)

---

