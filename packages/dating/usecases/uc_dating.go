package usecases

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/LukasJatmiko/dating-app-api/driver"
	datingModel "github.com/LukasJatmiko/dating-app-api/packages/dating/models"
	"github.com/LukasJatmiko/dating-app-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DatingUsecase struct {
	Driver        driver.Driver
	Validate      *validator.Validate
	JWTPrivateKey []byte
	JWTPublicKey  []byte
	JWTExpiration time.Duration
	JWTIssuer     string
}

func (h *DatingUsecase) SignIn(email, password string) (*datingModel.Data, error) {
	claims := new(datingModel.UserClaims)
	user := new(datingModel.User)
	sql := h.Driver.GetWrapperInstance().(*gorm.DB)
	if result := sql.Preload("Gender").Where("email = ?", email).First(user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("unknown user")
		} else {
			return nil, fmt.Errorf("internal error")
		}
	} else {
		if e := utils.BcryptCheckPassword(password, user.Password); e != nil {
			return nil, fmt.Errorf("incorrect password")
		}
	}
	if id, e := uuid.NewRandom(); e != nil {
		return nil, fmt.Errorf("error while creating uuid (%v)", e)
	} else {
		claims.ID = id.String()
	}
	claims.UserID = uint(user.ID)
	claims.UserEmail = user.Email
	claims.UserName = user.Name
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(h.JWTExpiration))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.NotBefore = jwt.NewNumericDate(time.Now())
	claims.Issuer = h.JWTIssuer
	claims.Subject = claims.UserName
	claims.Audience = []string{h.JWTIssuer}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	block, _ := pem.Decode(h.JWTPrivateKey)
	privateKey, e := x509.ParsePKCS1PrivateKey(block.Bytes)
	if e != nil {
		return nil, fmt.Errorf("error while parsing pivate key %v", e)
	}
	if tokenString, err := token.SignedString(privateKey); err != nil {
		return nil, fmt.Errorf("error while signing token %v", err)
	} else {
		return &datingModel.Data{Token: tokenString}, nil
	}
}

func (h *DatingUsecase) SignUp(user *datingModel.User) error {
	if e := h.Validate.Struct(user); e != nil {
		return e
	} else {
		//hash user's password
		if hashedPassword, e := utils.BcryptHashPassword(user.Password); e != nil {
			return fmt.Errorf("internal error")
		} else {
			user.Password = hashedPassword
			db := h.Driver.GetWrapperInstance().(*gorm.DB)
			result := db.Create(user)
			if result.Error != nil {
				if result.Error == gorm.ErrDuplicatedKey {
					return fmt.Errorf("user with the same email already exists")
				} else {
					log.Println(result.Error)
					return fmt.Errorf("internal error")
				}
			} else {
				return nil
			}
		}
	}
}
