package jwt

import (
	"blog-main/global"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var SECRET_KEY = "1234567890"

// jwtSecret JWT密钥(建议通过配置文件或环境变量设置)
var jwtSecret = []byte(SECRET_KEY)

const (
	// AccessTokenExpire 访问token过期时间
	AccessTokenExpire = 2 * time.Hour
	// RefreshTokenExpire 刷新token过期时间
	RefreshTokenExpire = 7 * 24 * time.Hour
)

// Claims JWT声明结构体
type Claims struct {
	UserId    int    `json:"user_id"`    //用户Id
	Username  string `json:"username"`   //用户名
	TokenType string `json:"token_type"` //token类型：“access”或“refresh”
	jwt.RegisteredClaims
}

// TokenResponse token响应结构体
type TokenResponse struct {
	AccessToken  string `json:"access_token"`  //访问token
	RefreshToken string `json:"refresh_token"` //刷新token
	ExpiresIn    int64  `json:"expires_in"`    //过期时间（秒）
}

// 后台管理生成token
func GenerateToken(userId int, username string) (string, error) {
	tokenResponse, err := GenerateTokenPair(userId, username)
	if err != nil {
		return "", err
	}
	return tokenResponse.AccessToken, nil
}

// GenerateTokenPair 生成访问和刷新token对
func GenerateTokenPair(userId int, username string) (*TokenResponse, error) {
	accessClaims := Claims{
		UserId:    userId,
		Username:  username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	//生成刷新token
	refreshClaims := Claims{
		UserId:    userId,
		Username:  username,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	//将token存储到Redis中
	err = StoreTokenInRedis(userId, accessTokenString, refreshTokenString)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    int64(AccessTokenExpire.Seconds()),
	}, nil
}

func StoreTokenInRedis(userId int, accessToken, refreshToken string) error {
	ctx := context.TODO()
	userKey := fmt.Sprintf("user_token:%d", userId)
	// 使用hash存储用户的token信息
	err := global.RedisClient.HMSet(ctx, userKey, "access_token", accessToken, "refresh_token", refreshToken, "created_at", time.Now().Unix()).Err()
	if err != nil {
		return err
	}
	// 设置过期时间为刷新token的过期时间
	return global.RedisClient.Expire(ctx, userKey, RefreshTokenExpire).Err()
}

func IsTokenValidInRedis(userId int, token, tokenType string) bool {
	ctx := context.TODO()
	userKey := fmt.Sprintf("user_token:%d", userId)
	var redisToken string
	var err error
	if tokenType == "access" {
		redisToken, err = global.RedisClient.HGet(ctx, userKey, "access_token").Result()
	} else {
		redisToken, err = global.RedisClient.HGet(ctx, userKey, "refresh_token").Result()
	}
	if err != nil {
		return false
	}
	return redisToken == token
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if !IsTokenValidInRedis(claims.UserId, tokenString, claims.TokenType) {
		return nil, errors.New("token失效")
	}
	return claims, nil
}

func RevokeToken(userId int) error {
	ctx := context.TODO()
	userKey := fmt.Sprintf("user_token:%d", userId)
	return global.RedisClient.Del(ctx, userKey).Err()
}
