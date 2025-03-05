package main

import (
	"log"
	"net/http"
	"time"

	auth "gateway/grpc/client/auth"
	video "gateway/grpc/client/video"
	authorize "gateway/internal/auth"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
)

type Application struct {
	Port           string
	AuthService    auth.AuthServiceClient
	VideoService   video.VideoUploadServiceClient
	SessionManager *scs.SessionManager
	Authorize      *authorize.Authorize
}

func main() {
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "redis_db:6379")
		},
	}

	sessionManager := scs.New()
	sessionManager.Store = redisstore.New(pool)
	sessionManager.Cookie.Name = "kk_session"
	sessionManager.Lifetime = 24 * time.Hour

	app := &Application{
		Port:           ":8088",
		AuthService:    auth.ConnectToAuth(),
		VideoService:   video.ConnectToVideo(),
		SessionManager: sessionManager,
		Authorize:      &authorize.Authorize{Session: sessionManager},
	}

	mux := app.handleRoutes()

	err := http.ListenAndServe(app.Port, mux)
	if err != nil {
		log.Println(err)
	}
}
