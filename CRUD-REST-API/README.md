# 🎬 Go REST API – Movie Management Service

This is a simple REST API built in Go for managing a collection of movies. It supports full CRUD operations and is a great beginner-level project for practicing web development in Go using the `gorilla/mux` router.

---

## 🚀 Features

- ✅ List all movies
- 📄 Get a single movie by ID
- ➕ Create a new movie
- 📝 Update an existing movie
- ❌ Delete a movie

---

## 📦 Technologies Used

- [Go (Golang)](https://golang.org/)
- [gorilla/mux](https://github.com/gorilla/mux) – HTTP router for request routing
- `encoding/json` – JSON marshaling/unmarshaling
- `net/http` – for the HTTP server

---

## 📁 Project Structure

```bash
.
├── main.go         # Main application file
├── README.md       # Project documentation
```
## ⚙️ Getting Started

# Prerequisites

- Go installed
- (Optional) Postman for testing

# Run the server

```bash
go run main.go
```
The server starts on:

[arduino](http://localhost:8080)

## 🔗 API Endpoints

| Method | Endpoint        | Description         |
|--------|------------------|---------------------|
| GET    | `/movies`        | List all movies     |
| GET    | `/movies/{id}`   | Get a movie by ID   |
| POST   | `/movies`        | Create a new movie  |
| PUT    | `/movies/{id}`   | Update a movie by ID|
| DELETE | `/movies/{id}`   | Delete a movie by ID|

## 📥 Sample JSON

# ➕ POST/PUT movie body

```bash
{
  "isbn": "12345",
  "title": "Example Movie",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

## 📝 Notes

-this project does not use a database --movies are sotred in memory (slices).
- Data will reset every time the server restarts.
- This is an educational project--not intended for production use

## 📚 Future Improvements

✅ Connect to a real database (e.g., PostgreSQL, MongoDB)

🛡 Add validation and error handling

🔒 Add authentication

🚀 Deploy on a cloud provider

## 👨‍🎓 Author

Made with ❤️ while learning Go.
Feel free to contribute, fork, or suggest improvements!