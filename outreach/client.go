package outreach

import (
	"encoding/json"
	"fmt"
	// "go/token"
	"io/ioutil"
	"log"
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
		log.Println("[Token Error]: ",err )
		return nil, err
	}
	client := &http.Client{Timeout: 10 * time.Second}
	req.Header.Add("Authorization", "Bearer "+Token)
	req.Header.Add("content-type", "application/vnd.api+json")
	re, err := client.Do(req)
	if err != nil {
		log.Println("[Token Error]: ",err )
		return nil, err
	}
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
	// redirect_uri := "https://clevertap.com//oauth/outreach"
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
		ClientID:      "UwxFB_r-MdrxJ8Q0XH3qKqn5HikGaa2ObawmKpOY9IY",
		ClientSecrete: "aYI-SoGPPjZwwvQ4ncNna31FTar1iLTHt-iZGnpqAGU",
		RedirectURL:   "https://clevertap.com//oauth/outreach",
		GrantType:     "refresh_token",
		RefreshToken:  "mgrMCv1kMW0HhUoa5k-SmA9Hg0ru6XG8uHmLect3e38",
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
		log.Println("[Token Error]: ",err )
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
		log.Println("[Do req Error]: ",err )
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[Do req Error]: ",err )
		return nil, err
	}

	if res.StatusCode == 400 || res.StatusCode == 401 || res.StatusCode == 429 ||res.StatusCode == 422 ||res.StatusCode == 404{
		return nil, fmt.Errorf("status of doreq: %d, body: %s  ", res.StatusCode, body)
	}
	return body, err
}

func (c *Client) GetDataSourceUser(email string) (*User, error) {
	req, err := http.NewRequest("GET", "https://api.outreach.io/api/v2/users", nil)
    if err != nil {
		log.Println("[GetUser Error]: ",err )
		return nil, err
	}

	r, err := c.doRequest(req)
	if err != nil {
		log.Println("[GetUser Error]: ",err )
		return nil, err
	}
	userlist:= ListUser{}
    err = json.Unmarshal(r, &userlist)
	if err != nil {
		log.Println("[GetUser Error]: ",err )
		return nil, err
	}
	var data *User
	userList:=userlist.List
	for _,cont:= range userList{
          if(cont.Attributes.Email== email){
			  data = &cont
			  break
		  }
	}
	if(data==nil){
		return nil,fmt.Errorf("User with email %s not found",email)
	}
	return data,nil
}


func (c *Client) GetUserData(UserId string) (*Data, error) {
	req, err := http.NewRequest("GET", "https://api.outreach.io/api/v2/users/"+UserId, nil)

	if err != nil {
		log.Println("[GetUser Error]: ",err )
		return nil, err
	}

	r, err := c.doRequest(req)
	if err != nil {
		log.Println("[GetUser Error]: ",err )
		return nil, err
	}
	data := Data{}
	err = json.Unmarshal(r, &data)
	if err != nil {
		log.Println("[GetUser Error]: ",err )
		return nil, err
	}
	return &data, nil
}
func (c *Client) CreateUser(userCreateInfo Data) (*Data, error) {
	reqb, err := json.Marshal(userCreateInfo)
	if err != nil {
		log.Println("[CreateUser Error]: ",err )
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.outreach.io/api/v2/users", strings.NewReader(string(reqb)))
	if err != nil {
		log.Println("[CreateUser Error]: ",err )
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println("[CreateUser Error]: ",err )
		return nil, err
	}
	user := Data{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("[CreateUser Error]: ",err )
		return nil, err
	}
	return &user, nil
}

func (c *Client) UpdateUser(UserId string, userUpdateInfo Data) (*Data, error) {
	reqb, err := json.Marshal(userUpdateInfo)
	if err != nil {
		log.Println("[UpdateUser Error]: ",err )
		return nil, err
	}

	URL := "https://api.outreach.io/api/v2/users/" + UserId

	req, err := http.NewRequest("PATCH", URL, strings.NewReader(string(reqb)))
	if err != nil {
		log.Println("[UpdateUser Error]: ",err )
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println("[UpdateUser Error]: ",err )
		return nil, err
	}

	user := Data{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("[UpdateUser Error]: ",err )
		return nil, err
	}

	return &user, nil
}
