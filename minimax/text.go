package minimax

import (
	"context"
	textv1 "github.com/ConnectAI-E/go-minimax/gen/go/minimax/text/v1"
	"google.golang.org/grpc"
)

var _ textv1.MinimaxServiceClient = new(Client)

func (cli *Client) ChatCompletions(ctx context.Context, in *textv1.ChatCompletionsRequest, opts ...grpc.CallOption) (*textv1.ChatCompletionsResponse, error) {
	res := new(struct {
		textv1.ChatCompletionsResponse
	})

	in.Stream = false
	resp, err := cli.client.R().
		SetBody(in).
		SetSuccessResult(res).
		Post("/v1/text/chatcompletion")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}

	return &res.ChatCompletionsResponse, err
}
