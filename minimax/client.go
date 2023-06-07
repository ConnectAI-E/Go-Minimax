package minimax

import (
	"github.com/imroc/req/v3"
	"github.com/moul/http2curl"
	"log"
	"os"
)

const minimaxBaseUrl = "https://api.minimax.chat"

type Client struct {
	client   *req.Client
	apiToken string
}

func New(opts ...Option) (*Client, error) {
	cli := &Client{
		client: req.C().SetBaseURL(minimaxBaseUrl),
	}

	cli.client.OnBeforeRequest(func(client *req.Client,
		req *req.Request) error {
		if len(cli.apiToken) > 0 {
			req.SetHeader("Authorization", cli.apiToken)
		}
		return nil
	}).OnAfterResponse(func(client *req.Client, resp *req.Response) error {
		if os.Getenv("APP_ENV") == "debug" {
			curl, err := http2curl.GetCurlCommand(resp.Request.RawRequest)
			if err != nil {
				return err
			}
			log.Println(curl)
		}
		return nil
	})
	for _, opt := range opts {
		err := opt(cli)
		if err != nil {
			return nil, err
		}
	}
	return cli, nil
}
