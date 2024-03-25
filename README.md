# Deep Focus App
A simple deep focus web app that helps users stay focused on their work. With this app, users can set goals for their deep work sessions. For example, if a user wants to have a 30-minute deep work session daily, they can input the goal and time. Successfully reaching the goal will increase their streak by one day, building on the previous streak unless it's broken by a missed day. In case of a missed day, the streak resets to zero. The app also stores all session history in a database, allowing users to review their past sessions. The objective of this app is to facilitate deep study sessions, reading periods, or any tasks users wish to accomplish, all recorded and accessible at any time.

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
3. Open your browser and navigate to `http://localhost:8080`
4. You should see `{"message": "Welcome to the server!"}`


## Technologies
- [Golang](https://golang.org/)
- [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)