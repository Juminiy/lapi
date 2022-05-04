package router

import (
	"github.com/gofiber/fiber/v2"
	"zhaoxin-api/context"
)

func AuthBaseApi(authBase fiber.Router) {

	authGithub := authBase.Group("/github")
	authGithub.Get("/client_id",context.OAUTHGithubFetchClientId)
	authGithub.Get("/access_token",context.OAUTHGithubFetchAccessToken)
	authGithub.Get("/logout",context.OAUTHGithubDisconnect)
	authGithub.Get("/protected",context.OAUTHGithubProtected,context.OAUTHGithubHandler)
	authGithub.Get("/redirect",context.OAUTHGithubRedirect)

	authQQ := authBase.Group("/qq")
	authQQ.Get("/redirect",context.OAUTHQQRedirect)
	authQQ.Get("/logout",context.OAUTHQQDisconnect)

	authGoogle := authBase.Group("/google")
	authGoogle.Get("/redirect",context.OAUTHGoogleRedirect)
	authGoogle.Get("/logout",context.OAUTHGoogleDisconnect)
}