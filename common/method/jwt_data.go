package method

import "github.com/golang-jwt/jwt/v5"

type LancerClaims struct {
	UserId   int64  `json:"user_id"`
	RoleId   int64  `json:"role_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}
