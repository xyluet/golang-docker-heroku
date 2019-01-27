# Testing `heroku.yml` with golang and docker

### Requirements.

You should have these tools installed in your system:

1. [Golang](https://golang.org/dl/)
2. [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)
3. [Git](https://git-scm.com)

Make sure you are logged in with heroku cli.

`Dockerfile` is annotated, be sure to check it out.

### Instructions at creating and deploying the app

1. `cd $GOPATH/src/github.com/your-username/`
2. `mkdir gin-heroku-api && cd gin-heroku-api` creates the package.
3. `go get -u github.com/gin-gonic/gin` installs gin.
4. `touch main.go` creates the executable go file.
5. Write the golang source code in `main.go`. Example can be seen in this repository.
6. Create `templates/pong.html` if you're using this `main.go` as an example application.
7. `go run main.go` tests the app. It should be working on http://localhost:8080.
8. `touch heroku.yml` creates the heroku file. Be sure to [populate it](https://devcenter.heroku.com/articles/build-docker-images-heroku-yml). Example is in the repo.
9. `git init` creates a git repository.
10. `git commit -am "Initial commit"` commits the code.
11. `heroku create your-golang-gin-app` creates the heroku app.
12. `heroku stack:set container` enables container in the app.
13. `git push heroku master` deploys the application.
14. `heroku open` should open the deployed application. Visit `/ping` or `/pong` to test the app.