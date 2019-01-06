package testdata

//go:generate dekit -destination=./decoders_dekit.go github.com/go-gad/dekit/looker/testdata CreateOrderReq

type CreateOrderReq struct {
	Name string `query:"name"`
	Body string `json:"body"`
}

type Req1 struct{}

type Req2 struct{}
