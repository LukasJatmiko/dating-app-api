package main

import (
	sqlDriver "database/sql/driver"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/LukasJatmiko/dating-app-api/constants"

	"github.com/LukasJatmiko/dating-app-api/driver"
	"github.com/LukasJatmiko/dating-app-api/packages/dating/models"
	"github.com/LukasJatmiko/dating-app-api/packages/dating/usecases"
	"github.com/LukasJatmiko/dating-app-api/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}
type AnyString struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v sqlDriver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func (a AnyString) Match(v sqlDriver.Value) bool {
	_, ok := v.(string)
	return ok
}

// a successful case
func TestSignIn(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	dialector := postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	var d driver.Driver
	pg := new(driver.PostgresConnectionPool)
	pg.Instance = sqlDB
	pg.GormInstance = gormDB
	d = pg

	RSAPrivateKey, e := os.ReadFile(utils.GetEnvOrDefaultString(string(constants.ENVAuthJWTPrivateKey), "jwtRS256.key"))
	if e != nil {
		t.Errorf("test error while loading key: %s", e)
	}
	RSAPublicKey, e := os.ReadFile(utils.GetEnvOrDefaultString(string(constants.ENVAuthJWTPublicKey), "jwtRS256.key.pub"))
	if e != nil {
		t.Errorf("test error while loading key: %s", e)
	}

	duc := &usecases.DatingUsecase{
		Driver:        d,
		Validate:      validator.New(),
		JWTPrivateKey: RSAPrivateKey,
		JWTPublicKey:  RSAPublicKey,
		JWTExpiration: 60 * time.Hour,
		JWTIssuer:     "dating-app-api",
	}

	birthday, _ := time.Parse("2006-01-02", "1988-01-01")
	createdAt, _ := time.Parse("2006-01-02 15:04:05.999", "2023-12-23 16:08:36.147")
	updatedAt, _ := time.Parse("2006-01-02 15:04:05.999", "2023-12-23 16:08:36.147")

	//test sign in
	{
		rowUser := sqlmock.NewRows(
			[]string{"id", "name", "nickname", "email", "password", "profile_picture_id", "birthday", "gender_id", "created_at", "updated_at", "deleted_at"},
		).AddRow(1, "Mock User", "mockuser", "lucasjatmiko@gmail.com", "$2a$10$vZRrPdBgQnukokYKmA/yf.SorAYTae/7Ev3IXl1xrGfooZ.ejoDty", "", birthday, 1, createdAt, updatedAt, nil)
		rowGender := sqlmock.NewRows(
			[]string{"id", "name"},
		).AddRow(1, "Male")
		mock.ExpectQuery(`SELECT * `).WithArgs("lucasjatmiko@gmail.com").WillReturnRows(rowUser)
		mock.ExpectQuery(`SELECT * `).WithArgs(1).WillReturnRows(rowGender)

		if _, e := duc.SignIn("lucasjatmiko@gmail.com", "lukas123"); e != nil {
			t.Error("should be success (" + e.Error() + ")")
		}
	}

	//test sign up
	{
		user := &models.User{
			Name:           "New User",
			NickName:       "nu",
			Email:          "newuser@email.com",
			Password:       "lukas123",
			Birthday:       datatypes.Date(birthday),
			Gender:         models.Gender{ID: 1, Name: "Male"},
			GenderInterest: []models.Gender{{ID: 2, Name: "Female"}},
		}
		result := sqlmock.NewResult(1, 1)
		rowInsert := sqlmock.NewRows(
			[]string{"id", "id"},
		).AddRow(1, 1)
		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT (.+)").WithArgs(AnyTime{}, AnyTime{}, nil, "Male", 1).WillReturnRows(rowInsert)
		mock.ExpectQuery("^INSERT (.+)").WithArgs(AnyTime{}, AnyTime{}, nil, user.Name, user.Email, AnyString{}, "", user.NickName, AnyTime{}, 1).WillReturnRows(rowInsert)
		mock.ExpectQuery("^INSERT (.+)").WithArgs(AnyTime{}, AnyTime{}, nil, "Female", 2).WillReturnRows(rowInsert)
		mock.ExpectExec("^INSERT (.+)").WithArgs(0, 2).WillReturnResult(result)
		mock.ExpectCommit()

		if e := duc.SignUp(user); e != nil {
			t.Error("should be success")
		}
	}
}
