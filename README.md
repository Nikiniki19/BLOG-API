# Blog CRUD API - Go Fiber

A simple Blog CRUD API built using **Go Fiber**. It allows users to create, read, update, and delete blog posts.

This project uses:
- **Fiber** as the web framework
- **GORM** for ORM
- **PostgreSQL** as the database
- **Swagger** for API documentation
- **Testify** and **mocking** for unit tests
- Deployed on **Render**


## 🚀 Features

- Create a blog post
- Retrieve all blog posts
- Retrieve a blog post by ID
- Update a blog post
- Delete a blog post
- Swagger documentation
- Unit testing with mocking


## 🛠️ Technologies Used

- Go
- Fiber
- GORM
- PostgreSQL
- Swagger
- Testify (unit testing)
- Mockery/Mockgen (for mocks)
- Render (deployment)


## 📦 API Endpoints

| Method | Endpoint                  | Description           |
|--------|---------------------------|-----------------------|
| POST   | `/api/blog-post/createPost` | Create a new blog post |
| GET    | `/api/blog-post/getAllPosts` | Retrieve all blog posts |
| GET    | `/api/blog-post/:id`        | Retrieve a blog post by ID |
| PATCH  | `/api/blog-post/:id`        | Update a blog post by ID |
| DELETE | `/api/blog-post/:id`        | Delete a blog post by ID |


## 🧪 Run Tests

```bash
go test ./... -cover
