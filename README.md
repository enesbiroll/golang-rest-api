
# 🛠 Golang REST API with Fiber Framework

A simple, fast, and modern **REST API** built with **Golang** and the **Fiber** framework. This API allows you to perform **CRUD operations** for student data, with database interaction via **GORM** and structured logging using **Logrus**. The API also features **rate limiting** for controlling the number of requests per IP and supports **soft delete** functionality.

![Golang API](https://img.shields.io/badge/Go-Fiber-blue?style=for-the-badge&logo=go&logoColor=white)

## 🚀 Features

- **📝 CRUD Operations**: Create, read, update, and delete student data.
- **🗄️ Database Integration**: MySQL with **GORM** ORM for database interaction.
- **🖥️ Logging**: Structured logging with **Logrus** for event tracking.
- **⚠️ Error Handling**: Centralized handling of errors and success responses.
- **🗑️ Soft Delete**: Support for soft deletion of student data.
- **🔒 Rate Limiting**: Custom rate limits and IP banning for API security.

## 🛠 Tech Stack

- **Golang**: Backend programming language.
- **Fiber**: Lightweight web framework for Go.
- **MySQL**: Relational database for storing student data.
- **GORM**: ORM for Go that simplifies database interactions.
- **Logrus**: Structured and flexible logger for Go.
- **Dotenv**: Manages environment variables.

## ⚡ Installation

### 1. Clone the repository:

```bash
git clone https://github.com/yourusername/golang-fiber-api.git
cd golang-fiber-api
```

### 2. Install dependencies:

Make sure you have Go installed on your machine.

```bash
go mod tidy
```

### 3. Set up environment variables:

Create a `.env` file in the root directory and configure the following environment variables:

```plaintext
MYSQL_HOST=localhost
MYSQL_USER=root
MYSQL_PASS=yourpassword
MYSQL_DB_NAME=golangtest
```

### 4. Start the server:

Run the application using:

```bash
go run main.go
```

This will start the server at `http://localhost:3000`.

## 🔧 API Endpoints

### 1. Create a Student

**POST** `/students`

**Request Body**:

```json
{
  "name": "John Doe",
  "code": "S1234"
}
```

**Response**:

```json
{
  "status": "success",
  "message": "Student created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "code": "S1234",
    "created_at": "2025-04-03T14:59:59Z",
    "updated_at": "2025-04-03T14:59:59Z"
  }
}
```

### 2. Get All Students

**GET** `/students`

**Response**:

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "code": "S1234",
    "created_at": "2025-04-03T14:59:59Z",
    "updated_at": "2025-04-03T14:59:59Z"
  }
]
```

### 3. Get Student by ID

**GET** `/students/{id}`

**Response**:

```json
{
  "id": 1,
  "name": "John Doe",
  "code": "S1234",
  "created_at": "2025-04-03T14:59:59Z",
  "updated_at": "2025-04-03T14:59:59Z"
}
```

### 4. Update a Student

**PUT** `/students/{id}`

**Request Body**:

```json
{
  "name": "Johnathan Doe",
  "code": "S5678"
}
```

**Response**:

```json
{
  "status": "success",
  "message": "Student updated successfully",
  "data": {
    "id": 1,
    "name": "Johnathan Doe",
    "code": "S5678",
    "created_at": "2025-04-03T14:59:59Z",
    "updated_at": "2025-04-03T15:00:59Z"
  }
}
```

### 5. Delete a Student

**DELETE** `/students/{id}`

**Response**:

```json
{
  "status": "success",
  "message": "Student deleted successfully"
}
```

## 📝 Logging

The application uses **Logrus** for logging. Logs are outputted to the console by default, but you can also configure them to be written to a file.

### Log Levels:
- **Info**: General information.
- **Warning**: Non-critical issues.
- **Error**: Issues impacting functionality.
- **Fatal**: Critical errors causing the application to stop.

Logs are generated for database queries, server start-up, and errors.

## ⚠️ Error Handling

The API returns standardized responses for both success and error cases. For example, if a student is not found:

**Error Response**:

```json
{
  "status": "error",
  "message": "Student not found"
}
```

## 🔒 Rate Limiting

The application implements rate limiting to protect the API from overuse:

- **POST /students**: Limited to 3 requests per 15 minutes per IP.
- **PUT /students/{id}**: Limited to 5 requests per 10 minutes per IP.
- **DELETE /students/{id}**: Limited to 2 requests per 30 minutes per IP.

Exceeding the limit will result in a `429 Too Many Requests` response:

```json
{
  "status": "error",
  "message": "Too many requests, please try again later."
}
```

## 🤝 Contribution

Feel free to fork this repository, make improvements, and contribute!

### Steps to Contribute:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-xyz`.
3. Commit your changes: `git commit -am 'Add new feature'`.
4. Push the changes: `git push origin feature-xyz`.
5. Open a pull request.

Made with ❤️ by Enes Birol
