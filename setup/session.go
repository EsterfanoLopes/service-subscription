package setup

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
)

func InitSession() *scs.SessionManager {
	// set up session
	session := scs.New()
	session.Store = redisstore.New(setup.InitRedis())
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}
