package user

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
	log     *log.Logger
}

func RegisterHandlers(e *echo.Echo, s *Service, l *log.Logger) {
	handler := Handler{
		service: s,
		log:     l,
	}

	e.POST("/register", handler.createUser)
	e.POST("/login", handler.loginUser)
}

func (h *Handler) createUser(ctx echo.Context) error {
	request := registrationRequest{}
	response := registrationResponse{}

	err := ctx.Bind(&request)
	if err != nil {
		h.log.Println("could not bind registration request", err)
		return err
	}
	err = request.Validate()
	if err != nil {
		h.log.Println("invalid request data", err)
		response.setErrorMessage(err)
		ctx.JSON(http.StatusBadRequest, response)
	}

	user := request.toUser()
	user, err = h.service.CreateUser(user)
	if err != nil {
		h.log.Println("could not create user", err)
		response.setErrorMessage(err)
		ctx.JSON(http.StatusBadRequest, response)
	}

	ctx.NoContent(http.StatusOK)
	return nil
}

func (h *Handler) loginUser(ctx echo.Context) error {
	request := loginRequest{}
	response := loginResponse{}

	err := ctx.Bind(&request)
	if err != nil {
		h.log.Println("could not bind login request", err)
		return err
	}
	err = request.Validate()
	if err != nil {
		h.log.Println("invalid request data", err)
		response.setErrorMessage(err)
		ctx.JSON(http.StatusBadRequest, response)
	}

	user := request.toUser()
	user, err = h.service.LoginUser(user)
	if err != nil {
		h.log.Println("could not login user", err)
		response.setErrorMessage(err)
		ctx.JSON(http.StatusBadRequest, response)
	}

	h.issueAuthCookie(ctx)
	return ctx.NoContent(http.StatusOK)
}

func (*Handler) issueAuthCookie(ctx echo.Context) {
	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["authenticated"] = true
	sess.Save(ctx.Request(), ctx.Response())
}
