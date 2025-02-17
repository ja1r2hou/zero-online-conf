package model

import "time"

type UserToken struct {
	UserName            string    `json:"user_name"`
	Ip                  string    `json:"ip"`
	UserTokenExpireTime time.Time `json:"user_token_expire_time"`
}
