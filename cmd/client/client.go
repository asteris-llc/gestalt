package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
)

var client *Client

// Client is a client for Gestalt
type Client struct {
	Host      string
	Scheme    string
	UserAgent string

	c *http.Client
}

// New returns a new client
func New() *Client {
	c := &Client{}

	c.UserAgent = "gestalt-cli/1.0"
	c.Scheme = viper.GetString("scheme")
	c.Host = viper.GetString("host")

	c.c = &http.Client{
		Timeout: viper.GetDuration("timeout"),
	}

	return c
}

func setupClient() {
	client = New()
}

// Do makes a request
func (c *Client) Do(method, path string, query map[string]interface{}, body io.Reader) (*http.Response, error) {
	values := url.Values{}
	for k, v := range query {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	url := &url.URL{
		Scheme:   c.Scheme,
		Host:     c.Host,
		Path:     path,
		RawQuery: values.Encode(),
	}

	request, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "gestalt/1.0")
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return c.c.Do(request)
}

// prettify formats JSON nicely if it can. If it can't, it fails silently
// (assuming that the JSON is ill-formed but the calling function still wants
// it.)
func prettify(raw []byte) *bytes.Buffer {
	if !viper.GetBool("pretty") {
		return bytes.NewBuffer(raw)
	}

	buf := &bytes.Buffer{}
	err := json.Indent(buf, raw, "", "  ")

	if err != nil {
		return bytes.NewBuffer(raw)
	}

	return buf
}
