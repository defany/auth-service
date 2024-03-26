package config

import "time"

type JWT struct {
	RefreshTokenSecret     string        `json:"refresh_secret_key" env:"JWT_REFRESH_SECRET_KEY" env-required:"true"`
	AccessTokenSecret      string        `json:"access_secret_key" env:"JWT_ACCESS_SECRET_KEY" env-required:"true"`
	RefreshTokenExpiration time.Duration `json:"refresh_token_expiration" env:"JWT_REFRESH_TOKEN_EXPIRATION" env-required:"true"`
	AccessTokenExpiration  time.Duration `json:"access_token_expiration" env:"JWT_ACCESS_TOKEN_EXPIRATION" env-required:"true"`
}

func (j *JWT) RefreshSecretKey() []byte {
	return []byte(j.RefreshTokenSecret)
}

func (j *JWT) AccessSecretKey() []byte {
	return []byte(j.AccessTokenSecret)
}

func (j *JWT) RefreshTokenDuration() time.Duration {
	return j.RefreshTokenExpiration * time.Second
}

func (j *JWT) AccessTokenDuration() time.Duration {
	return j.AccessTokenExpiration * time.Second
}
