package minimax

import (
	"context"
	"errors"
	textv1 "github.com/ConnectAI-E/go-minimax/gen/go/minimax/text/v1"
	"github.com/ConnectAI-E/go-minimax/internal"
	"google.golang.org/grpc"
	"io"
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

func (cli *Client) ChatCompletionStream(ctx context.Context, in *textv1.ChatCompletionsRequest, opts ...grpc.CallOption) (textv1.MinimaxService_ChatCompletionStreamClient, error) {

	in.Stream = true
	in.UseStandardSse = true
	resp, err := cli.client.R().
		DisableAutoReadResponse().
		SetBody(in).
		Post("/v1/text/chatcompletion")

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	return internal.NewStreamReader[*textv1.ChatCompletionsResponse](resp.Body), err
}
