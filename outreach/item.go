package outreach

type Data struct{
	Data User `json:"data"`
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
	// CreateAt string `json:"createdAt,omitempty"`
	Locked bool `json:"locked"`
	UserName string `json:"username,omitempty"`
	// Title string `json:"title,omitempty"`
	// UpdatedAt string `json:"updatedAt,omitempty"`
}