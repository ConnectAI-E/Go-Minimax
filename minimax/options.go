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

func WithBaseUrl(url string) Option {
	return func(cli *Client) error {
		if len(url) == 0 {
			return errors.New("base url can not empty")
		}
		cli.client.SetBaseURL(url)
		return nil
	}
}

func WithGroupId(groupId string) Option {
	return func(cli *Client) error {
		if len(groupId) == 0 {
			return errors.New("group id can not empty")
		}
		cli.groupId = groupId
		return nil
	}
}
