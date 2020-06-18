package app

import (
	"encoding/gob"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	// SQL Dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/iammarkps/triamudom-room-api/handler"
	"github.com/iammarkps/triamudom-room-api/models"
)

// New function create new echo app
func New() (*echo.Echo, *gorm.DB) {
	gob.Register(time.Time{})
	e := echo.New()

	db, err := gorm.Open("postgres", os.Getenv("DB"))
	e.Logger.Info("Connecting to database...")

	if err != nil {
		panic(err)
	}

	e.Logger.Info("Successfully connected to database")
	db.BlockGlobalUpdate(true)
	db.DB().SetMaxIdleConns(100)

	// defer db.Close()
	db.AutoMigrate(&models.User{})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowOrigins:     []string{"http://localhost:3000", "https://room.triamudom.ac.th"},
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())

	h := &handler.Handler{DB: db}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "😼Triam Udom Suksa School's room API is running!")
	})

	e.GET("/student/:id", h.Student)

	return e, db
}
