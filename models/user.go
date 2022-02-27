package models

import (
	"context"
	"golang-microsvc/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/golang-jwt/jwt"
)

type User struct {
	Username string `json:"username"`
	Password string `josn:"password"`
	Role     string `json:"role"`
}

type MyToken struct {
	User
	jwt.StandardClaims
}

func (u *User) Login(username, password string) map[string]interface{} {
	filter := bson.M{
		"username": username,
		"password": password,
	}

	coll := mDB.Collection(USER_COLLECTION)
	res := coll.FindOne(
		context.TODO(),
		filter,
		options.FindOne().SetProjection(bson.M{"password": 0}),
	)

	resp := bson.M{}
	err := res.Decode(&resp)
	if err != nil {
		return utils.ErrorResponse("Unable to decode")
	}

	resp["token"] = GenerateJwtToken(&User{		
		Username: username,
		Role: resp["role"].(string),
	})
	
	return utils.SuccessResponse(resp)
}

func GenerateJwtToken(u *User) string {
	expTime := time.Now().Add(time.Minute * 60)
	tk := MyToken {
		*u,
		jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &tk)
	tokenString, _ := token.SignedString([]byte(PASS_KEY))
	return tokenString
}
