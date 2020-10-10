// Package structs consists of all structs in the app
package structs

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
