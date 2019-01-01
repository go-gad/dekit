package pizza

type CreateOrderReq struct {
	Name string `query:"name"`
	Body string `json:"body"`
}
