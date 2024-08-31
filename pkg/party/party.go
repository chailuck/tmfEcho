package party

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type PartyHandler struct {
	DB *sqlx.DB
}

func (h *PartyHandler) Initialize(db *sqlx.DB) {

	//db.AutoMigrate(&Customer{})

	h.DB = db
}

func (h *PartyHandler) GetIndividual(c echo.Context) error {
	return nil
}
