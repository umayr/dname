# `dname`
>generate names like docker does when it creates a new container

## Install
```
λ go get -u github.com/umayr/dname
```

Or, you can download pre-built binaries from [here](https://github.com/umayr/dname/releases).

## Usage

#### CLI:
```
λ dname -h
  NAME:
     dname - generate names like docker does when it creates a new container

  USAGE:
     dname [command] [arguments...]

  VERSION:
     0.0.1

  COMMANDS:
       help, h  Shows a list of commands or help for one command

  GLOBAL OPTIONS:
     --number value, -n value  total number of generated names (default: 1)
     --help, -h                show help
     --version, -v             print the version

λ dname
  elated_jennings

λ dname -n 10
  sleepy_morse
  thirsty_shaw
  jovial_dubinsky
  hungry_joliot
  big_ardinghelli
  nostalgic_borg
  drunk_carson
  serene_franklin
  evil_noether
  evil_goodall
```

#### API:
```go
import "github.com/umayr/dname"

func main() {
    fmt.Println(dname.Generate())
}
```

