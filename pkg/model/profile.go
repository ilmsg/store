package model

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Location  string `json:"location"`
}

type ProfileResponse struct {
}
