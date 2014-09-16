adm - A Dotfile Manager
=======================

## Description

A Dotfile Manager

## Usage

```
NAME:
   adm - A Dotfile Manager

USAGE:
   adm [global options] command [command options] [arguments...]

COMMANDS:
   init
   add
   sync
   discard
   push
   status
   install
   secret
   help, h      Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version
```

## Install

To install, use `go get`:

```bash
$ go get github.com/udzura/adm/cmd/adm
```

## Contribution

1. Fork ([https://github.com/Uchio KONDO/adm/fork](https://github.com/Uchio KONDO/adm/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[Uchio KONDO](https://github.com/udzura)

## License

MIT-LICENSE. See `LICENSE`.
