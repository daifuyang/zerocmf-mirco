// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshReq struct {
	RefreshToken string `json:"refreshToken,optional"`
}

type ValidationReq struct {
}
