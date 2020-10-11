// Package structs consists of all structs in the app
package structs

import "time"

// UserIDStruct Struct for new user ID
type UserIDStruct struct {
	UID  string `json:"uid"`
	Used bool   `json:"used"`
}

// SessIDStruct Struct for new session ID
type SessIDStruct struct {
	SessID string `json:"sessid"`
	Used   bool   `json:"used"`
}

// User Struct for a single user
type User struct {
	UID        string    `json:"uid"`
	Uname      string    `json:"uname"`
	Pass       string    `json:"pass"`
	DateJoined time.Time `json:"dateJoined"`
}
