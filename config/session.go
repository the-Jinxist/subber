package config

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
)

var scm *scs.SessionManager

func InitSession() {
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
