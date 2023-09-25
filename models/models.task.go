package models

type SignIn struct{
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
} 

type SignUp struct{
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirmpassword" bson:"confirmpassword"`
}

