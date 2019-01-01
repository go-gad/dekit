package pizza

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeCreateOrderReq(t *testing.T) {
	payload := []byte(`{"body":"lalala"}`)
	req, _ := http.NewRequest("POST", "/orders?name=Ivan", bytes.NewBuffer(payload))
	r, err := decodeCreateOrderReq(req)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("Ivan", r.Name)
	assert.Equal("lalala", r.Body)
}
