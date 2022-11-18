package types

type User struct {
	Id    string `bson:"_id" json:"id" `
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}
