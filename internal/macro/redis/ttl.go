package redis

import "time"

const (
	// TenMinute 10分钟
	TenMinute     = time.Second * 60 * 10
	UserKeyTTL    = time.Minute * 30
	ArticleKeyTTL = time.Hour * 6
)
