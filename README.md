# Y2 CLI

The official CLI for the [Y2 REST API](https://y2.yeetum.com).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/y2-intel/y2-cli/cmd/y2@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/y2/main.go
```

<!-- x-release-please-end -->

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
