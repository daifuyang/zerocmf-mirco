// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	UserId int64 `json:"userId"`
}

type RefreshReq struct {
	RefreshToken string `json:"refreshToken,optional"`
}

type ValidationReq struct {
}
