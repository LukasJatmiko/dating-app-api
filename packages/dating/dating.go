package dating

import (
	"log"
	"net/http"
	"time"

	"github.com/LukasJatmiko/dating-app-api/driver"
	apiModel "github.com/LukasJatmiko/dating-app-api/packages/api/models"
	datingModel "github.com/LukasJatmiko/dating-app-api/packages/dating/models"
	datingUsecase "github.com/LukasJatmiko/dating-app-api/packages/dating/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type DatingHandler interface {
	SignIn(string, string) (*datingModel.Data, error)
	SignUp(*datingModel.User) error
}

func NewDatingHandler(driver driver.Driver, privateKey, publicKey []byte, validate *validator.Validate) DatingHandler {
	return &datingUsecase.DatingUsecase{
		Driver:        driver,
		Validate:      validate,
		JWTPrivateKey: privateKey,
		JWTPublicKey:  publicKey,
		JWTExpiration: 60 * time.Hour,
		JWTIssuer:     "dating-app-api",
	}
}

func Mount(h DatingHandler, app fiber.Router) {
	app.Post("/signin", SignIn(h))
	app.Post("/signup", SignUp(h))
}

func SignIn(h DatingHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		credential := new(datingModel.Credential)
		if e := c.BodyParser(credential); e != nil {
			return c.Status(http.StatusOK).JSON(apiModel.APIBaseResponsePayload{Code: http.StatusUnauthorized, Message: "unauthorized", Data: fiber.Map{}})
		} else {
			if data, e := h.SignIn(credential.Email, credential.Password); e != nil {
				return c.Status(http.StatusOK).JSON(apiModel.APIBaseResponsePayload{Code: http.StatusUnauthorized, Message: e.Error(), Data: fiber.Map{}})
			} else {
				return c.Status(http.StatusOK).JSON(apiModel.APIBaseResponsePayload{Code: http.StatusOK, Message: "ok", Data: data})
			}
		}
	}
}

func SignUp(h DatingHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		newUser := new(datingModel.User)
		if e := c.BodyParser(newUser); e != nil {
			log.Println("errornya : " + e.Error())
			return c.Status(http.StatusOK).JSON(apiModel.APIBaseResponsePayload{Message: e.Error(), Code: http.StatusBadRequest, Data: fiber.Map{}})
		} else {
			if e := h.SignUp(newUser); e != nil {
				return c.Status(http.StatusOK).JSON(apiModel.APIBaseResponsePayload{Message: e.Error(), Code: http.StatusInternalServerError, Data: fiber.Map{}})
			} else {
				return c.Status(http.StatusOK).JSON(apiModel.APIBaseResponsePayload{Message: "ok", Code: http.StatusOK, Data: newUser})
			}
		}
	}
}
