package config

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/the-Jinxist/subber/data"
)

var scm *scs.SessionManager

func InitSession() {
	gob.Register(data.User{})

	session := scs.New()

	rdb := GetRedis()
	if rdb == nil {
		log.Panic("redis has not been initialized yet")
	}
	session.Store = redisstore.New(rdb)
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	scm = session
}

func GetSession() *scs.SessionManager {
	return scm
}
