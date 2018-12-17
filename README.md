# dekit
Toolchain to generate decoder based on request's struct

### Usage

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
