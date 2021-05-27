package client

import (
	"encoding/json"
	"fmt"
	
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

var (
    Errors = make(map[int]string)
)

func init() {
	Errors[400] = "Bad Request, StatusCode = 400"
	Errors[404] = "User Does Not Exist , StatusCode = 404"
	Errors[409] = "User Already Exist, StatusCode = 409"
	Errors[401] = "Unautharized Access, StatusCode = 401"
	Errors[429] = "User Has Sent Too Many Request, StatusCode = 429"
	Errors[422] = "Email has already been taken, StatusCode = 422"
}


func NewClient(client_id string, client_secret string,  refresh_token string, acc_token string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostUrl:    HostURL,
	}
	Token := os.Getenv("acc_token")
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
		tok, _ := TokenGen(client_id, client_secret,  refresh_token, acc_token)
		ar := AuthResp{}
		json.Unmarshal(tok, &ar)
		os.Setenv("acc_token", ar.AccToken)
		Token = ar.AccToken
	}

	c.AccessToken = "Bearer " + Token
	return &c, nil
}

func TokenGen(client_id string, client_secret string,  refresh_token string, acc_token string) ([]byte, error) {
	req_json:=AuthStruct{
		ClientID:      client_id,
		ClientSecrete: client_secret,
		RedirectURL:   "https://clevertap.com//oauth/outreach",
		GrantType:     "refresh_token",
		RefreshToken:  refresh_token,
	}

	
	rb, _ := json.Marshal(req_json)

	req, err := http.NewRequest("POST", "https://api.outreach.io/oauth/token", strings.NewReader(string(rb)))
	if err != nil {
		return nil, fmt.Errorf("%v",err)
	}
	req.Header.Add("content-type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		log.Println("[Token Error]: ",err )
		return nil, fmt.Errorf("%v", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[Token Error]: ",err)
		return nil, fmt.Errorf("%v",err)
	}
	temp := os.Getenv("acc_token")
	if res.StatusCode != 200 {
		log.Println("[Token Error]: ",Errors[res.StatusCode])
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
		log.Println("[Do req Error]: ",Errors[res.StatusCode])
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
	userlist:=ListUser{}
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
		return nil,fmt.Errorf("user with email %s not found",email)
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

	

	req, err := http.NewRequest("PATCH","https://api.outreach.io/api/v2/users/" + UserId, strings.NewReader(string(reqb)))
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
