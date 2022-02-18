package user

import (
	"log"
	"net/http"

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
}

type registrationRequest struct {
	Username string `json:"username" validate:"required,gt=3,lte=30,alphanum"`
	Password string `json:"password" validate:"required,gt=8,lte=128,password"`
}

type registrationResponse struct {
	ErrorMessage string `json:"message,omitempty"`
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

func (r *registrationRequest) toUser() (user User) {
	user.Username = r.Username
	user.Password = r.Password
	return user
}

func (r *registrationResponse) setErrorMessage(e error) {
	r.ErrorMessage = e.Error()
}
