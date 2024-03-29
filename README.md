# Deep Focus App
A simple deep focus web app that helps users stay focused on their work. With this app, users can set goals for their deep work sessions. For example, if a user wants to have a 30-minute deep work session daily, they can input the goal and time. Successfully reaching the goal will increase their streak by one day, building on the previous streak unless it's broken by a missed day. In case of a missed day, the streak resets to zero. The app also stores all session history in a database, allowing users to review their past sessions. The objective of this app is to facilitate deep study sessions, reading periods, or any tasks users wish to accomplish, all recorded and accessible at any time.

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
Replace the placeholders for the database [username], [password], [port],[database_name] with your actual values.

## How to run the server
1. Clone the repository by running the following command:
```bash
git clone https://github.com/sandbox-science/deep-focus.git
cd deep-focus # Change directory to the project folder
```
2. Run the following command to start the server:
```bash
make run
```
3. Open your browser and navigate to `http://localhost:8080` (or any other port you specified in the `.env` file)
4. You should see the front page of the app

## How to contribute
1. Clone the repository by following the instructions in the ["How to run the server"](#how-to-run-the-server) section 
2. Create a new branch ``git checkout -b choose-a-branch-name``
3. Make your changes
4. Commit your changes ``git commit -m "Your message"``
5. Push your changes ``git push origin choose-a-branch-name``
6. Create a pull request (from GitHub on the project page)

Always make sure that your local main branch is up-to-date with the remote main branch. You can do this by running the following commands:
```bash
git checkout main
git pull origin main
```
and to update your local branch with the latest changes from the main branch, go back to you own local branch by running:
```bash
git checkout choose-a-branch-name
git merge main
```

## Technologies
- [Golang](https://golang.org/)
- [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)
- [PostgreSQL](https://www.postgresql.org/)

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.