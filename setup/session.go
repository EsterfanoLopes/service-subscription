package setup

import (
	"encoding/gob"
	"net/http"
	"subscription-service/data"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
)

func InitSession() *scs.SessionManager {
	gob.Register(data.User{})

	// set up session
	session := scs.New()
	session.Store = redisstore.New(InitRedis())
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}
