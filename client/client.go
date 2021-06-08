package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client struct {
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

func NewClient(client_id string, client_secret string, refresh_token string, acc_token string, redirect_url string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
	Token := os.Getenv("outreach_acc_token")
	req, err := http.NewRequest("GET", "https://api.outreach.io/api/v2", nil)
	if err != nil {
		log.Println("[Token Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	client := &http.Client{Timeout: 10 * time.Second}
	req.Header.Add("Authorization", "Bearer "+Token)
	req.Header.Add("content-type", "application/vnd.api+json")
	re, err := client.Do(req)
	if err != nil {
		log.Println("[Token Generation Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	defer re.Body.Close()
	if re.StatusCode != 200 {
			req_json := AuthStruct{
				ClientID:      client_id,
				ClientSecrete: client_secret,
				RedirectURL:   redirect_url,
				GrantType:     "refresh_token",
				RefreshToken:  refresh_token,
			}
			rb,err := json.Marshal(req_json)
			if err != nil {
				log.Println("[Token Generation Error]: ", err)
				return nil, fmt.Errorf("%v", err)
			}
			req, err := http.NewRequest("POST", "https://api.outreach.io/oauth/token", strings.NewReader(string(rb)))
			if err != nil {
				log.Println("[Token Generation Error]: ", err)
				return nil, fmt.Errorf("%v", err)
			}
			req.Header.Add("content-type", "application/json")
			tok, _ := c.doRequest(req)
			ar := AuthResp{}
			json.Unmarshal(tok, &ar)
			os.Setenv("outreach_acc_token", ar.AccToken)
			os.Setenv("outreach_refresh_token", ar.RefreshToken)
			Token = ar.AccToken
	}
	os.Setenv("outreach_acc_token", Token)
	c.AccessToken = "Bearer " + Token
	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("authorization", c.AccessToken)
	req.Header.Add("content-type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Println("[Do req Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[Do req Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	if res.StatusCode == 400 || res.StatusCode == 401 || res.StatusCode == 429 || res.StatusCode == 422 || res.StatusCode == 404 {
		log.Println("[Do req Error]: ", Errors[res.StatusCode])
		return nil, fmt.Errorf("status of doreq: %s ", Errors[res.StatusCode])
	}
	return body, err
}

/*
   User ID is requires to fetch the Outreach user. GetDataSourceUser() function defined to fetch user
   and import the user using email ID.The function fetch the all outreach users of organization and
   extract the user with given email id. But email ID can be changed or update using UI. So can't use
   email ID as primary key.
*/
// func (c *Client) GetDataSourceUser(email string) (*User, error) {
// 	req, err := http.NewRequest("GET", "https://api.outreach.io/api/v2/users", nil)
//     if err != nil {
// 		log.Println("[GetUser Error]: ",err )
// 		return nil, err
// 	}
// 	body, err := c.doRequest(req)
// 	if err != nil {
// 		log.Println("[GetUser Error]: ",err )
// 		return nil, err
// 	}
// 	userlist:=ListUser{}
//     err = json.Unmarshal(body, &userlist)
// 	if err != nil {
// 		log.Println("[GetUser Error]: ",err )
// 		return nil, err
// 	}
// 	var data *User
// 	userList:=userlist.List
// 	for _,cont:= range userList{
//           if(cont.Attributes.Email== email){
// 			  data = &cont
// 			  break
// 		  }
// 	}
// 	file, _ := json.MarshalIndent(userList, "", " ")
// 	_ = ioutil.WriteFile("C:/Users/Kunal/Desktop/Outreach_Provider/client/test.json", file, 0644)
// 	if(data==nil){
// 		return nil,fmt.Errorf("user with email %s not found",email)
// 	}

// 	return data,nil
// }

func (c *Client) GetUserData(UserId string) (*Data, error) {
	req, err := http.NewRequest("GET", "https://api.outreach.io/api/v2/users/"+UserId, nil)
	if err != nil {
		log.Println("[GetUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println("[GetUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	data := Data{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("[GetUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	return &data, nil
}

func (c *Client) CreateUser(userCreateInfo Data) (*Data, error) {
	reqb, err := json.Marshal(userCreateInfo)
	if err != nil {
		log.Println("[CreateUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	req, err := http.NewRequest("POST", "https://api.outreach.io/api/v2/users", strings.NewReader(string(reqb)))
	if err != nil {
		log.Println("[CreateUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println("[CreateUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	user := Data{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("[CreateUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	return &user, nil
}

func (c *Client) UpdateUser(UserId string, userUpdateInfo Data) (*Data, error) {
	reqb, err := json.Marshal(userUpdateInfo)
	if err != nil {
		log.Println("[UpdateUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	req, err := http.NewRequest("PATCH", "https://api.outreach.io/api/v2/users/"+UserId, strings.NewReader(string(reqb)))
	if err != nil {
		log.Println("[UpdateUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println("[UpdateUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	user := Data{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("[UpdateUser Error]: ", err)
		return nil, fmt.Errorf("%v", err)
	}
	return &user, nil
}
