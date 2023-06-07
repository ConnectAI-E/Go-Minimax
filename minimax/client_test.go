package minimax

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getTestClient() (*Client, error) {
	return New(WithApiToken(os.Getenv("TEST_MINIMAX_API_TOKEN")))
}

func TestNew(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
}
