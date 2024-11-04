package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUserReq struct {
	Name string `json:"name" validate:"required"`
}

type UserRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserListRes struct {
	Users []*UserRes `json:"users"`
}
