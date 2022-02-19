package user

type registrationRequest struct {
	Username string `json:"username" validate:"required,gt=3,lte=30,alphanum"`
	Password string `json:"password" validate:"required,gt=8,lte=128,password"`
}

type registrationResponse struct {
	ErrorMessage string `json:"message,omitempty"`
}

func (r *registrationRequest) toUser() (user User) {
	user.Username = r.Username
	user.Password = r.Password
	return user
}

func (r *registrationResponse) setErrorMessage(e error) {
	r.ErrorMessage = e.Error()
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	ErrorMessage string `json:"message,omitempty"`
}

func (r *loginRequest) toUser() (user User) {
	user.Username = r.Username
	user.Password = r.Password
	return user
}

func (r *loginResponse) setErrorMessage(e error) {
	r.ErrorMessage = e.Error()
}
