# dateTime-Server-FarahTharwat
## How to use 

### What to install first

1. Install Go. I am using version `1.22.5`. You can use this link as a guide: [install-go-in-vscode](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code#1-install-go)
2. Install Gin Framework using [intsall-gin](https://gin-gonic.com/docs/quickstart/)
3. Install Docker Desktop through this link [install-docker](https://docs.docker.com/get-docker/)
4. Install Docker compose using this [intsall-docker-compose](https://docs.docker.com/compose/install/)
5. Install docker in vscode if you are using it [docker-in-vscode](https://code.visualstudio.com/docs/containers/overview)
6. Clone the repo.

### How to run 

- If you want to run the tests of the APIs before using them, navigate to the `pkg` directory in your terminal using `cd pkg` and then run `go test ./... -v`.
- If you want to run the servers follow these steps: 
	- run in the terminal `docker compose up`. <br>
  - Then, you can navigate in your browser to `localhost:8080/dateTime` for server developed using Gin framework or to `localhost:8090/dateTime` for standard server.
