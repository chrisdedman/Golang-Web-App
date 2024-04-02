# Golang Template Web Application
A simple web application built with Golang and PostgreSQL. The application is a simple CRUD application that allows users to create, read, update, and delete users from a PostgreSQL database.
A simple frontend is also included in the project to demonstrate the functionality of the application.

## Project Guide
This is a guide for the Golang template CRUD web application project. You will find different guidelines to help you set up the project, and the database, and how to use Git and GitHub. You can find the wiki [here](https://github.com/chrisdedman/Golang-Web-App/wiki).

### Table of Content
1. [How to set up and use PostgreSQL](https://github.com/chrisdedman/Golang-Web-App/wiki/PostgresSQL-Setup-Guide)
2. [How to use Git and GitHub](https://github.com/chrisdedman/Golang-Web-App/wiki/Git-&-GitHub-Guide)
3. [How to set up the project](#whats-needed-to-run-the-server)
4. [How to run the server](#how-to-run-the-server)
5. [Technologies](#technologies)
6. [License](#license)
7. [Demo](#demo)

## What's needed to run the server
- [Golang](https://golang.org/)
- [Make](https://www.gnu.org/software/make/)
- [PostgreSQL](https://www.postgresql.org/)
- ``.env`` file

In the root directory, create a new file named ``.env``.<br>
Copy the content of ``.env.example`` into your ``.env`` file.

Replace the placeholders for the database connection with your own values (read the [PostgreSQL setup guide](#table-of-content) for more information).

## How to run the server
1. Run the following commands on your terminal to clone the repository and run the server:
```bash
git clone https://github.com/chrisdedman/Golang-Web-App.git
cd Golang-Web-App # Change directory to the project folder
make run          # Run the server using Makefile script (required Make)
```
2. Open your browser and navigate to `http://localhost:3000` (or any other port you specified in the `.env` file)
3. You should see the front page template of the web app (see the [demo](#demo) section below)

## Technologies
- [Golang](https://golang.org/)
- [GORM](https://gorm.io/)
- [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)
- [PostgreSQL](https://www.postgresql.org/)
- [Make](https://www.gnu.org/software/make/)

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Demo

![Golang Web App Homepage](/assets/homepage.png)
![Golang Web App Login Page](/assets/login.png)
![Golang Web App Register Page](/assets/register.png)
![Golang Web App Dashboard Page](/assets/dashboard.png)
