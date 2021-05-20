package outreach

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
	"time"
)

const HostURL string = "https://api.outreach.io/api/v2"

type Client struct {
	HostUrl     string
	HTTPClient  *http.Client
	AccessToken string
}

type AuthStruct struct {
	ClientID      string `json:"client_id"`
	ClientSecrete string `json:"client_secrete"`
	RedirectURL   string `json:"redirect_uri"`
	GrantType     string `json:"grant_type"`
	RefreshToken  string `json:"refresh_token"`
}
type AuthResp struct {
	AccToken string `json:"access_token"`
}

func NewClient(client_id string, client_secret string, auth_code string, refresh_token string, acc_token string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostUrl:    HostURL,
	}
	c.AccessToken = acc_token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("authorization", c.AccessToken)
	req.Header.Add("content-type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status12: %d, body: %s  %s", res.StatusCode, body, c.AccessToken)
	}
	return body, err
}

func (c *Client) GetUserData(UserId string) (*Data, error) {
	req, err := http.NewRequest("GET", "https://api.outreach.io/api/v2/users/"+UserId, nil)

	if err != nil {
		return nil, err
	}

	r, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	data := Data{}
	err = json.Unmarshal(r, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
