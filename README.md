<h1 align="center">
  Letronic App
</h1>
<h2 align="center">
  Golang Restfull Api With Gorilla/Mux & GORM
</h2>

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/aldiramdan/go-rentvehicle.git
```

2. Install dependencies

```bash
go get -u ./...
# or
go mod tidy
```

3. Add Env

```sh

# Database
DB_HOST = Your DB User
DB_PORT = Your DB Port
DB_HOST = Your DB Host
DB_NAME = Your DB Name
DB_PASS = Your DB Password

# App
PORT = Your Port App
BASE_URL = http://localhost:PORT //example
```

4. Run the app

```bash
go run *.go serve
```

5. Database Migration and Seeder

```bash
# Migration
go run *.go migrate --up //db migration up
# or
go run *.go migrate --down //db migration down

# Seeder
go run *.go seeder --up //db seeder up
# or
go run *.go seeder --down /db seeder down
```

🌟 You are all set!

## 💻 Built with

- [Golang](https://go.dev/): programming language
- [gorilla/mux](https://github.com/gorilla/mux): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS
