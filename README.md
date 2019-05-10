# pom

*pomodoro timer on cli*


### test mode on docker-compose

```shell
~/$ docker-compose up
~/$ docker exec -it pom_app_1 go run main.go hoge
```

### if you build ...

```shell
# this command execute tests before building
~/$ make build
```

### if you try testing ...
```shell
~/$ make test
```

### long message example
```go
Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
```
