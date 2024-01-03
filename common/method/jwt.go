package method

import (
	"github.com/golang-jwt/jwt/v5"
	"lancer/global"
	"time"
)

func GenerateJwt(id int64, roleId int64, userName string) (string, error) {
	claims := LancerClaims{
		UserId:   id,
		RoleId:   roleId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.JwtData.Expire) * time.Minute)), //expire
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                         //sign time
			NotBefore: jwt.NewNumericDate(time.Now()),                                                         //effective time
		},
	}

	//HS256
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := withClaims.SignedString([]byte(global.JwtData.Secret))
	return signedString, err
}

func ParseJwt(tokenString string) (*LancerClaims, error) {

	t, err := jwt.ParseWithClaims(tokenString, &LancerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JwtData.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*LancerClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
