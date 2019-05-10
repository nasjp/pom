# pom

*pomodoro timer on cli*

### if you build ...

```shell
# this command execute tests before building
~/$ make build
```

### if you try testing ...
```shell
~/$ make test
```

### if you build in docker container ...
```shell
~/$ docker-compose up
```

### if you remove all docker images and containers ...
```shell
~/$ make clean-docker
```

### long message example
```go
Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
```

### Errors

#### One
```shell
build github.com/YukihiroTaniguchi/pom/cmd: cannot find module for path gopkg.in/mattn/go-colorable.v0
```
if you have this error, adding an override to go.mod
```go.mod
replace (
	gopkg.in/mattn/go-colorable.v0 => github.com/mattn/go-colorable v0.1.0
	gopkg.in/mattn/go-isatty.v0 => github.com/mattn/go-isatty v0.0.6
)
```
[mattn/go-colorable Go modules bug](https://github.com/mattn/go-colorable/issues/35)


#### Two
```shell
app_1  | 18:8:23 main        | Build Failed:
app_1  |  can't load package: package github.com/YukihiroTaniguchi/pom: unknown import path "github.com/YukihiroTaniguchi/pom": cannot find module providing package github.com/YukihiroTaniguchi/pom
pom_app_1 exited with code 1
```
if you use "fresh", you must use on cmd directory (there is "main.go")
