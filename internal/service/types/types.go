package types

import "user-api/internal/storage/types"

type CreateUser struct {
	Id       string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email" `
	Password string `bson:"password"`
}

type GetAllUsersResponse struct {
	Data       []types.User `json:"data"`
	TotalCount int64        `json:"totalCount"`
}
