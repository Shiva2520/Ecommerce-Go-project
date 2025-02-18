package types

import "go.mongodb.org/mongo-driver/bson/primitive"

//user model `user` table
type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Phone    string             `json:"phone" bson:"phone"`
	Password string             `json:"password" bson:"password"`
}

// address model `addresses` table
type Address struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Address1 string             `json:"address_1" bson:"address_1"`
	UserId   primitive.ObjectID `json:"user_id" bson:"user_id"`
	City     string             `json:"city" bson:"city"`
	Country  string             `json:"country" bson:"country"`
}



// Verification Model `verification` table
type Verification struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Email  string             `json:"email" bson:"email"`
	Otp    int64              `json:"otp" bson:"otp"`
	Status bool               `json:"status" bson:"status"`
}
