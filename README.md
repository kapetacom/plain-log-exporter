# plain-log-exporter

## Initial Setup

This section is intended to help developers and contributors get a working copy of
`plain-log-exporter` on their end

Clone this repository

```sh
git clone https://github.com/kapetacom/plain-log-exporter
cd plain-log-exporter
```


## Local Development

This section will guide you to setup a fully-functional local copy of `plain-log-exporter`.


### Installing dependencies

To install all dependencies associated with `plain-log-exporter`, run the
command

```sh
go mod tidy
```

### Running Tests

```sh
go test ./...
```

### Running `plain-log-exporter`

To run plain-log-exporter, use the command

```sh
go run main.go
```
or
```sh
go build && ./plain-log-exporter
```


