package user

import (
	"errors"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userStore    Store
	jwtSecret    string
}

func NewHandler(userS Store) (*UserHandler, error) {
	secret, ok := os.LookupEnv("Signing_Key")
	if !ok {
		return nil, errors.New("No secret key ")
	}

	return &UserHandler{
		userStore:    userS,
		jwtSecret: secret,
	}, nil
}


func (h *UserHandler) Register(group *echo.Group) {
	skipper := func (c echo.Context) bool {
		// Skip middleware if path is equal 'login'
		if c.Request().URL.Path == "/auth" || c.Request().URL.Path == "/" {
		  return true
		}
		return false
	 }

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(h.jwtSecret),
		Skipper: skipper,
	})

	guestUsers := group.Group("/")
	guestUsers.POST("/singup", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := group.Group("/user", jwtMiddleware)
	user.GET("", h.GetMe)
	user.GET("/:id", h.GetMe)
	user.PUT("/:id", h.UpdateUser)
	user.PUT("", h.UpdateMe)
	user.DELETE("/:id", h.DeleteUser)
}