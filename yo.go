package main

import (
	"errors"
	"net/http"
	"net/url"
)

var YO_API = "http://api.justyo.co"

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) YoAll() error {
	data := url.Values{}
	data.Set("api_token", c.Token)
	res, err := http.PostForm(YO_API+"/yoall/", data)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return errors.New("Received response with non 201 status code.")
	}

	return nil
}

func (c *Client) YoUser(username string) error {
	data := url.Values{}
	data.Set("api_token", c.Token)
	data.Set("username", username)
	res, err := http.PostForm(YO_API+"/yoall/", data)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return errors.New("Received response with non 201 status code.")
	}

	return nil
}
