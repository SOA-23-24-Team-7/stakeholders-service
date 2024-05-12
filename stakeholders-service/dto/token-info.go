package dto

type TokenInfo struct {
	UserId   int64  `json:"userId"`
	PresonId int64  `json:"personId"`
	Username string `json:"username"`
	Role     string `json:"role"`
}