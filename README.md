# Golang Web App
A simple web application built with Golang and PostgreSQL. The application is a simple CRUD application that allows users to create, read, update, and delete users from a PostgreSQL database.
A simple frontend is also included in the project to demonstrate the functionality of the application.

## What's needed to run the server
- [Golang](https://golang.org/)
- [Make](https://www.gnu.org/software/make/)
- [PostgreSQL](https://www.postgresql.org/)
- ``.env`` file with the following environment variables:
```text
export HOST_ADDR   = ":3000"
export DB_HOST     = "localhost"
export DB_PORT     = "PORT"
export DB_USER     = "USERNAME"
export DB_PASSWORD = "PASSWORD"
export DB_NAME     = "DATABASE_NAME"
export DB_SSLMODE  = "disable"
```
Replace the placeholders for the database connection with your own values.

## How to run the server
1. Clone the repository by running the following command:
```bash
git clone https://github.com/chrisdedman/Golang-Web-App.git
cd Golang-Web-App # Change directory to the project folder
```
2. Run the following command to start the server:
```bash
make run
```
3. Open your browser and navigate to `http://localhost:3000` (or any other port you specified in the `.env` file)
4. You should see the front page template of the web app.

## Technologies
- [Golang](https://golang.org/)
- [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)
- [PostgreSQL](https://www.postgresql.org/)

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.