
# GOpportunities – Job Opportunities API

**GOpportunities** is a modern job opportunities API built with **Golang**, one of the highest-paying programming languages today.  
The project uses:

- **Gin** for routing,
- **GORM** for ORM-based database access,
- **SQLite** as the embedded database,
- **Swagger** for API documentation and testing.

It follows a modular and maintainable folder structure to ensure scalability and clarity in development.

---

## 🚀 Features

- Introduction to Golang and modern API development
- Complete environment setup for local development
- Route management with **Go-Gin**
- Lightweight database integration with **SQLite**
- ORM communication using **GORM**
- Swagger integration for API documentation and testing
- Clean and scalable project structure
- Fully functional job opportunities API: create, read, update, delete (CRUD)
- Automated tests for ensuring application quality

---

## 🧑‍💻 Installation

### Install dependencies:

```bash
go mod download
```

### Build the project:

```bash
go build
```

### Run the application:

```bash
./main
```

---

## 🛠️ Makefile Commands

This project includes a `Makefile` to simplify common tasks:

| Command               | Description                                               |
|-----------------------|-----------------------------------------------------------|
| `make run`            | Runs the app without generating Swagger docs              |
| `make run-with-docs`  | Generates Swagger docs and then runs the app              |
| `make build`          | Builds the binary (`gopportunities`)                      |
| `make test`           | Runs all unit tests                                       |
| `make docs`           | Generates Swagger documentation using Swag                |
| `make clean`          | Cleans the binary and removes the `./docs` directory      |

### Example:

```bash
make run
```

---

## 🐳 Docker and Docker Compose

### Docker

#### Build the image:

```bash
docker build -t your-image-name .
```

#### Run the container:

```bash
docker run -p 8080:8080 -e PORT=8080 your-image-name
```

### Docker Compose

#### Build the services:

```bash
docker compose build
```

#### Start the services:

```bash
docker compose up
```

#### Stop and remove services:

```bash
docker compose down
```

> For more information, refer to the official docs:  
> [Docker](https://docs.docker.com) | [Docker Compose](https://docs.docker.com/compose/)

---

## 🧰 Tech Stack

- [Golang](https://golang.org)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io)
- [SQLite](https://www.sqlite.org/index.html)
- [Swagger (Swaggo)](https://github.com/swaggo/swag)

---

## 📎 Usage

Once the app is running, you can access the Swagger UI to interact with the API at:

```bash
http://localhost:8080/swagger/index.html
```

> If you're using a different port, replace it accordingly.  
> Default `$PORT` is `8080`.

---

## 🤝 Contributing

We welcome contributions! Follow these steps:

1. Fork this repository
2. Create a new branch:  
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Commit your changes using [Conventional Commits](https://www.conventionalcommits.org)
4. Push your branch:  
   ```bash
   git push origin feature/your-feature-name
   ```
5. Open a Pull Request

---

## 📄 License

This project is licensed under the [MIT License](LICENSE.md).
