package outreach

import (
	"encoding/json"
	"fmt"
	// "go/token"
	"io/ioutil"
	// "log"
	"net/http"
	"os"
	// "strconv"
	"strings"

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
	ClientSecrete string `json:"client_secret"`
	RedirectURL   string `json:"redirect_uri"`
	GrantType     string `json:"grant_type"`
	RefreshToken  string `json:"refresh_token"`
}
type AuthResp struct {
	AccToken     string `json:"access_token"`
	Token_Type   string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

// var acc string

func NewClient(client_id string, client_secret string, auth_code string, refresh_token string, acc_token string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostUrl:    HostURL,
	}
	Token := os.Getenv("acc_token")
    // fmt.Println(Token)
	// fmt.Println("enterd init")
	req, err := http.NewRequest("GET", "https://api.outreach.io/api/v2", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: 10 * time.Second}
	req.Header.Add("Authorization", "Bearer "+Token)
	req.Header.Add("content-type", "application/vnd.api+json")
	re, _ := client.Do(req)
	if re.StatusCode != 200 {
		fmt.Println("enterd for new tok")
		tok, _ := TokenGen(client_id, client_secret, auth_code, refresh_token, acc_token)
		ar := AuthResp{}
		json.Unmarshal(tok, &ar)
		os.Setenv("acc_token", ar.AccToken)
		Token = ar.AccToken
	}

	c.AccessToken = "Bearer " + Token

	// c.AccessToken = acc_token
	return &c, nil
}

func TokenGen(client_id string, client_secret string, auth_code string, refresh_token string, acc_token string) ([]byte, error) {
	
	// if (client_id != "") && (client_secret != "") && (redirect_uri != "") {
	// }
	// grant_type := "refresh_token"
	// req_json:=AuthStruct{
	// 	ClientID:      client_id,
	// 	ClientSecrete: client_secret,
	// 	RedirectURL:   redirect_uri,
	// 	GrantType:     grant_type,
	// 	RefreshToken:  refresh_token,
	// }

	req_json := AuthStruct{
		ClientID:      "",
		ClientSecrete: "",
		RedirectURL:   "",
		GrantType:     "refresh_token",
		RefreshToken:  "",
	}
	rb, _ := json.Marshal(req_json)
	// if err != nil {
	// 	return "",fmt.Errorf("%v",err)
	// }

	req, _ := http.NewRequest("POST", "https://api.outreach.io/oauth/token", strings.NewReader(string(rb)))
	// time.Sleep(10 * time.Second)
	// if err != nil {
	// 	return "", fmt.Errorf("%v",err)
	// }
	req.Header.Add("content-type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	// res,err:=c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer res.Body.Close()
	// log.Println(res.StatusCode)
	body, _ := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return "", fmt.Errorf("%v",err)
	// }

	// ar := AuthResp{}
	//  json.Unmarshal(body, &ar)
	// if err != nil {
	// 	return "",fmt.Errorf("%v",err)
	// }
	// if res.StatusCode != 200 {
	// 	return nil, fmt.Errorf("status: %d, body: %s  %s token", res.StatusCode, body, ar.AccToken)
	// }
	// if res.StatusCode == 200 {
	// 	return nil, fmt.Errorf("status: %d, body: %s  %s token", res.StatusCode, body, ar.AccToken)
	// }
	// log.Println(ar.AccToken)
	// token := ar.AccToken
	temp := os.Getenv("acc_token")
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status: %d, body: %s %s tok", res.StatusCode, body, temp)
	}

	return body, nil
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

	if res.StatusCode == 400 || res.StatusCode == 401 || res.StatusCode == 429 ||res.StatusCode == 422 ||res.StatusCode == 404{
		return nil, fmt.Errorf("status of doreq: %d, body: %s  ", res.StatusCode, body)
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
func (c *Client) CreateUser(userCreateInfo Data) (*Data, error) {
	reqb, err := json.Marshal(userCreateInfo)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.outreach.io/api/v2/users", strings.NewReader(string(reqb)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	user := Data{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) UpdateUser(UserId string, userUpdateInfo Data) (*Data, error) {
	reqb, err := json.Marshal(userUpdateInfo)
	if err != nil {
		return nil, err
	}

	URL := "https://api.outreach.io/api/v2/users/" + UserId

	req, err := http.NewRequest("PATCH", URL, strings.NewReader(string(reqb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	user := Data{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
