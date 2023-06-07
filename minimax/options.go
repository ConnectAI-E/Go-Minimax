package minimax

import "errors"

type Option func(*Client) error

func WithApiToken(token string) Option {
	return func(cli *Client) error {
		if len(token) == 0 {
			return errors.New("api token can not empty")
		}
		cli.apiToken = token
		return nil
	}
}
