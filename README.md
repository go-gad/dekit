# dekit
Toolchain to generate decoder based on request's struct

## Install

```
go get -u github.com/go-gad/dekit
```

## Usage

```sh
❯ dekit -h
Usage:
    dekit [options...] <import_path> <parameter_names>

Example:
        dekit -destination=./decoders_dekit.go github.com/go-gad/dekit/examples/pizza CreateOrderReq

  <import_path>
        describes the complete package path where the parameter is located.
  <parameter_names>
        indicates the parameter names that are separated by comma.

Options:
  -build_flags string
        Additional flags for go build.
  -destination string
        Output file; defaults to stdout.
```

## Purpose

```go
type UpdateAuthorReq struct {
    ID int64 `path:"id"`
    Name string `query_string:"name"`
    Body string `json:"body"`
    Slug string `header:"X-Slug"`
}
```

Request:
```sh
curl -X PATCH -H "X-Slug: ivanov" http://site.com/authors/123?name=Alex -d '{"body":"a good person"}'
```

dekitgen will generate a decoder which convert `*http.Request` to struct with filled struct:
```go
...
return UpdateAuthorReq {
    ID: 123,
    Name: "Alex",
    Body: "a good person",
    Slug: "ivanov",
}
```
