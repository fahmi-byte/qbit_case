package config

import "time"

type AuthConfig struct {
	jwtTokenExpiresTTL int
	jwtTokenSecret     string
}

func (c *AuthConfig) AccessTokenSecret() string {
	return c.jwtTokenSecret
}

func (c *AuthConfig) AccessTokenExpiresDate() time.Time {
	duration := time.Duration(c.jwtTokenExpiresTTL)
	return time.Now().UTC().Add(time.Minute * duration)
}

func (c *AuthConfig) RefreshTokenExpiresDate() time.Time {
	duration := time.Duration(c.jwtTokenExpiresTTL)
	return time.Now().UTC().Add(time.Minute * duration * 8)
}
