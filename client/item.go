package client

type Data struct{
	Data User `json:"data"`
}

type ListUser struct{
	List []User `json:"data"`
}

type User struct {
	Type string `json:"type"`
	ID int `json:"id,omitempty"`
    Attributes Attributes `json:"attributes"`
}

type Attributes struct{
    Email string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Locked bool `json:"locked"`
	UserName string `json:"username,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Title string `json:"title,omitempty"`
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