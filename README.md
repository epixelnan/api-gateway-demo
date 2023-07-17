# API Gateway Demo Using nginx

## DIY

1. Create `helloapi` and `welcomeapi` ("upstream services"); test without Docker
2. Create `docker-compose.yml` without the nginx service; expose the ports of upstream services and test using `curl localhost:8091/hello` and `curl localhost:8092/welcome`
3. Add the nginx service to `docker-compose.yml`, comment out the port expositions of the upstream services and test using `curl localhost:8090/hello` and `curl localhost:8090/welcome`

## Output

```
$ curl localhost:8090/hello
Hello from helloapi!

$ curl localhost:8090/welcome
Welcome from welcomeapi!
```

## Notes

The directories `helloapi` and `welcomeapi` were initialized using the commands
`go mod init helloapi` and `go mod init welcomeapi` respectively.
