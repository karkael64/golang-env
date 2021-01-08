# A reader environment variables

Read operating system environment variables, read project .env file variables, then mixes variables in a single list.

## Add in your project

Load module and create `.env` file in your working directory, by running:

```sh
go get github.com/karkael64/golang-env
touch .env
```

## Use in your project file

Here is an exemple for displaying available variable in your working directory.

```go
import (
  "fmt"
  env "github.com/karkael64/golang-env"
)

func init() {
  var envs map[string]string = env.LoadEnv()
  fmt.Printf("%+v\n", envs)
}
```

## Notes

* File `.env` should be in the current working directory
* Every keys are transformed to uppercase
* Each key associates a single string
