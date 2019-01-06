package testdata

type CreateOrderReq struct {
	Name string `query:"name"`
	Body string `json:"body"`
}

type Req1 struct{}

type Req2 struct{}
