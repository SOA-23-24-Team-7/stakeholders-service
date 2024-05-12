package service

import (
	"os"
	"strconv"
	"time"

	"stakeholders-service/dto"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)


func getEnv(key, fallbackValue string) string {
    value, ok := os.LookupEnv(key)
    if !ok {
        return fallbackValue
    }
    return value
}

type JWTService struct {
}

func (controller *JWTService) GenerateAccsessToken(userId int64, username string,personId int64,role string ) (*dto.TokenResponse,error) {

	key := getEnv("JWT_KEY", "explorer_secret_key")
    issuer := getEnv("JWT_ISSUER", "explorer")
    audience := getEnv("JWT_AUDIENCE", "explorer-front.com")

    claims := jwt.MapClaims{
        "jti":       uuid.New().String(),
        "id":        strconv.FormatInt(userId, 10),
        "username":  username,
        "personId":  strconv.FormatInt(personId, 10),
        "role":      role,
        "exp":       time.Now().Add(time.Minute * 60 * 24 * 100).Unix(),
		"iss":       issuer,     
        "aud":       audience,  
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   

    signedToken, err := token.SignedString([]byte(key))
    if err != nil {
        return nil, err
    }

	var returnToken = dto.TokenResponse{
		Id: userId,
		AccessToken: signedToken,
	}

    return &returnToken, nil
}