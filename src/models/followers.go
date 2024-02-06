package models

type Followers struct {
	FollowerCount int    `json:"followerCount"`
	Followers     []User `json:"followers"`
}
