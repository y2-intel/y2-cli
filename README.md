# Y2 CLI

The official CLI for the [Y2 REST API](https://y2.yeetum.com).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/y2-cli/cmd/y2@latest'
```

### Running Locally

```sh
go run cmd/y2/main.go
```

## Usage

The CLI follows a resource-based command structure:

```sh
y2 [resource] [command] [flags]
```

```sh
y2 reports list
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
