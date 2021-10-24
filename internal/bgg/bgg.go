package bgg

import (
	"encoding/xml"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func New() *Client {
	r := resty.New()
	r.SetHostURL("https://www.boardgamegeek.com/xmlapi2")
	r.SetRetryCount(3)
	r.SetRetryWaitTime(time.Second)
	r.AddRetryCondition(func(r *resty.Response, e error) bool {
		return r.StatusCode() == http.StatusAccepted || r.StatusCode() == http.StatusTooManyRequests
	})
	// TODO: handle errors from API better

	return &Client{client: r}
}

func (c *Client) GetCollection(username string) (*Collection, error) {
	resp, err := c.client.R().
		SetQueryParam("username", username).
		Get("/collection")
	if err != nil {
		return nil, err
	}

	collection := &Collection{}
	err = xml.Unmarshal(resp.Body(), collection)
	return collection, err
}

func (c *Client) GetThing(id int) (*Thing, error) {
	resp, err := c.client.R().
		SetQueryParam("id", strconv.Itoa(id)).
		SetQueryParam("stats", "1").
		Get("/thing")
	if err != nil {
		return nil, err
	}

	thing := &Thing{}
	if err := xml.Unmarshal(resp.Body(), thing); err != nil {
		return nil, err
	}

	return thing, nil
}
