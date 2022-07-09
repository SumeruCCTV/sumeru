package constants

import (
	"regexp"
	"time"
)

const (
	SumeruName    = "sumeru"
	SumeruVersion = "0.0.1"
	IsDev         = true
)

const (
	CookieTokenKey   = "sumeru_token"
	HeaderCaptchaKey = "X-Captcha-Key"

	TokenLength         = 128
	UuidLength          = 32
	UsernameMinLength   = 3
	UsernameMaxLength   = 12
	CameraNameMinLength = 3
	CameraNameMaxLength = 18

	TokenExpiration = (14 * 24 * time.Hour) / time.Second // 2 weeks in seconds

	RedisTokenUuidKey = "__sumeru_redis_token__"

	UuidTokenAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	TokenRegex      = regexp.MustCompile(`^[A-Za-z\d]*$`)
	UsernameRegex   = TokenRegex
	CameraNameRegex = TokenRegex
)
