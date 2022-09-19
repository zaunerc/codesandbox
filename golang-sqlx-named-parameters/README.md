Code for [one of my blog posts](https://zauner.nllk.net/post/0044-golang-sqlx-named-parameters/).

# Run

1. `postgres$ docker-compose up --build`
2. `golang$ go run main.go`

# Example console output of the Go program

```
user@pc:~/repos/codesandbox/golang-sqlx-named-parameters/golang$ go run main.go 
2022/09/19 16:15:52 Creating DB conn host:localhost port:5432 username:postgres
2022/09/19 16:15:52 DB conn created
2022/09/19 16:15:52 Created book: &{ID:98c6f825-44e2-4e95-a629-f872c4759894 Title:alice in wonderland}
2022/09/19 16:15:52 Now going to search for books...
2022/09/19 16:15:52 Found books: [{ID:98c6f825-44e2-4e95-a629-f872c4759894 Title:alice in wonderland}]
2022/09/19 16:15:52 Finished successfully
user@pc:~/repos/codesandbox/golang-sqlx-named-parameters/golang$
```

# Notes about PostgreSQL

* Execute `$ docker-compose up --build` to build the image and start the container.
* You can connect to the database using the following connection details:
  * host: `localhost`
	* port: `5432`
	* database: `postgres`
	* username: `postgres`
	* password: `postgrespwd`
* The databse is stored in a volume. To delete the volume and the container
  execute the following command: `$ docker-compose rm -v`.
