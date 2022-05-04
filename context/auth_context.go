package context

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/session"
	"lapi/config"
	"lapi/middleware"
	"lapi/model"
	"lapi/utils"
)

// OAuth2.0 github code
// 1. in front-end, user clicks the link(https://github.com/login/oauth/authorize?client_id=&redirect_uri=)
// 2. if user has not login the github in the same browser redirect to github login page
// 3. if user has login the github show the page that user permit the backend-app access the user's public data
// 4. if user permit redirect to redirect_uri and give the uri parameter code
// 5. in backend, OAUTHGithubRedirect use httpClient send a request to github use client_id&client_secret&code to get access_token
// 6. if the backend network is ok, backend get scope&token_type&access_token store in the session, the session can store in redis,sql,or default in the memory
// 7. in the front-end, get the access_token from back_end and send request(https://api.github.com/user) in header (Authorization:bearer access_token)
// 8. user can see the github username and github avatar in the front-end page!

func OAUTHGithubFetchClientId(ctx *fiber.Ctx) error {
	return utils.OkResponse(ctx,config.Config("GITHUB_CLIENT_ID"))
}
func OAUTHGithubFetchAccessToken(ctx *fiber.Ctx) error {
	sessData,err := middleware.MySessionStore.Get(ctx)
	tokenType,tokenValue := sessData.Get("token_type"),sessData.Get("token_value")
	if err == nil && tokenType != nil && tokenValue != nil {
		if tokenType == "" || tokenValue == "" {
			return utils.RequestFailureResponse(ctx,"Token is nil")
		} else {
			return utils.OkResponse(ctx,tokenType,tokenValue)
		}
	} else {
		return utils.ErrorResponse(ctx,err)
	}
}
func OAUTHGithubRedirect(ctx *fiber.Ctx) error {
	code := ctx.Query("code","")
	if len(code) < 1 {
		return utils.RequestFailureResponse(ctx,"Bad Request")
	}

	httpClientAgent := fiber.AcquireAgent()
	req := httpClientAgent.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.Set("accept","application/json")
	req.SetRequestURI(fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",model.ClientID,model.ClientSecret,code))
	if err := httpClientAgent.Parse() ; err != nil {
		return utils.ErrorResponse(ctx,"HTTP Request fails",err)
	}
	var retCode int
	var retBody []byte
	var errs []error
	var tResponse *model.OAuthAccessResponse
	if retCode,retBody,errs = httpClientAgent.Struct(&tResponse) ; len(errs) > 0 {
		return utils.ErrorResponse(ctx,retCode,retBody,errs)
	}
	httpClientAgent.ConnectionClose()
	var sess *session.Session
	var err error
	if sess, err = middleware.MySessionStore.Get(ctx); err == nil {
		sess.Set("token_type",tResponse.TokenType)
		sess.Set("token_value",tResponse.AccessToken)
		sess.Save()
		if err != nil {
			return utils.ErrorResponse(ctx,err)
		}
		return utils.RedirectResponse(ctx,"/welcome.html")
	}
	return utils.RedirectResponse(ctx,"/index.html")
}
func OAUTHGithubProtected(ctx *fiber.Ctx) error {
	sessData, err := middleware.MySessionStore.Get(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx,err)
	}

	tkValue := sessData.Get("token_value")

	if tkValue == nil {
		sessData.Destroy()
		return utils.RedirectResponse(ctx,"/index.html")
	}
	return utils.OkResponse(ctx)
}
func OAUTHGithubHandler(ctx *fiber.Ctx) error {
	return utils.NonAuthResponse(ctx)
}
func OAUTHGithubDisconnect(ctx *fiber.Ctx) error {
	sessData, err := middleware.MySessionStore.Get(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx,err)
	}
	sessData.Destroy()
	return utils.RedirectResponse(ctx,"/index.html")
}

func OAUTHQQRedirect(ctx *fiber.Ctx) error {
	return nil
}
func OAUTHQQDisconnect(ctx *fiber.Ctx) error {
	return nil
}

func OAUTHGoogleRedirect(ctx *fiber.Ctx) error {

	return nil
}
func OAUTHGoogleDisconnect(ctx *fiber.Ctx) error {
	return nil
}

// BasicAuth
// 1.encode `postman:your-password` https://www.base64encode.org -> encodedString
// 2.Set Authorization Basic Auth Username & Password
// 3.Set HTTP Header Authorization Basic encodedString
// 4.Send your http request
func BasicAuth(ctx *fiber.Ctx) error {
	basicauth.New(basicauth.Config{
		Users: map[string]string{
			"Chisato@2084team.com":"2084team-lapi-admin",
		},
		Realm: "Admin",
		Authorizer: func(user,pass string) bool {
			if user == "Chisato@2084team.com" && pass == "2084team-lapi-admin" {
				return true
			} else {
				return false
			}
		},
		Unauthorized: func(ctx *fiber.Ctx) error {
			return utils.NonAuthResponse(ctx)
		},
	})
	return ctx.Next()
}

