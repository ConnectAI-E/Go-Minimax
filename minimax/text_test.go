package minimax

import (
	"context"
	"fmt"
	textv1 "github.com/ConnectAI-E/go-minimax/gen/go/minimax/text/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_ChatCompletions(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
	res, err := cli.ChatCompletions(context.Background(),
		&textv1.ChatCompletionsRequest{
			Messages: []*textv1.Message{
				{
					SenderType: "USER",
					Text:       "hello",
				},
			},
			Model:       "abab5-chat",
			Temperature: 0.7,
		})
	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Printf("%+v\n", res)
}

func TestClient_ChatCompletionStream(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
	stream, err := cli.ChatCompletionStream(context.Background(),
		&textv1.ChatCompletionsRequest{
			Messages: []*textv1.Message{
				{
					SenderType: "USER",
					Text:       "你好",
				},
			},
			Model:       "abab5-chat",
			Temperature: 0.7,
		})
	assert.Nil(t, err)
	assert.NotNil(t, stream)
	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("res: %+v\n", res)
	}
}
