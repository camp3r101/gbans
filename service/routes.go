package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/leighmacdonald/gbans/config"
	"net/http"
)

type Route string

const (
	routeDist             Route = "dist"
	routeHome             Route = "home"
	routeServers          Route = "servers"
	routeLogin            Route = "login"
	routeLogout           Route = "logout"
	routeLoginCallback    Route = "login_callback"
	routeAPIBans          Route = "api_bans"
	routeAPIFilteredWords Route = "filtered_words"
	routeServerAPIAuth    Route = "sapi_auth"
	routeServerAPIBan     Route = "sapi_ban"
	routeServerAPICheck   Route = "sapi_check"
	routeServerAPIMessage Route = "sapi_message"
)

func initRouter() {
	ses := sessions.Sessions("gbans", cookie.NewStore([]byte(config.HTTP.CookieKey)))
	router.Use(gin.Logger())
	// Dont use session for static assets
	router.StaticFS("/dist", http.Dir("frontend/dist"))

	session := router.Group("")
	session.Use(ses, authMiddleWare())

	session.GET(routeRaw(string(routeHome)), onIndex())
	session.GET(routeRaw(string(routeServers)), onServers())
	session.GET(routeRaw(string(routeLoginCallback)), onOpenIDCallback())
	session.GET(routeRaw(string(routeLogin)), onGetLogin())
	session.GET(routeRaw(string(routeLogout)), onGetLogout())
	session.GET(routeRaw(string(routeAPIBans)), onGetBans())
	session.GET(routeRaw(string(routeAPIFilteredWords)), onGetFilteredWords())
	session.POST(routeRaw(string(routeServerAPIAuth)), onPostServerAuth())

	// Game server specific API
	authed := router.Group("/", checkServerAuth)
	authed.GET(string(routeServerAPIBan), onGetServerBan())
	authed.POST(string(routeServerAPICheck), onPostServerCheck())
	authed.POST(routeRaw(string(routeServerAPIMessage)), onPostLogMessage())
}

func init() {
	routes = map[Route]string{
		routeHome:             "/",
		routeDist:             "/dist",
		routeServers:          "/servers",
		routeLogin:            "/auth/login",
		routeLoginCallback:    "/auth/callback",
		routeLogout:           "/auth/logout",
		routeAPIFilteredWords: "/api/v1/filtered_words",
		routeAPIBans:          "/api/v1/bans",
		routeServerAPIAuth:    "/sapi/v1/auth",
		routeServerAPIBan:     "/sapi/v1/ban",
		routeServerAPICheck:   "/sapi/v1/check",
		routeServerAPIMessage: "/sapi/v1/message",
	}
}
