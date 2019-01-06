package pizza

//go:generate dekit -destination=./decoders_dekit.go github.com/go-gad/dekit/examples/pizza CreateOrderReq

type CreateOrderReq struct {
	Name string `query:"name"`
	Body string `json:"body"`
}
