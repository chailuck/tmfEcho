package main

import (
	"tmfEcho/internal/database"
	"tmfEcho/internal/log"
	"tmfEcho/pkg/party"

	"github.com/labstack/echo"
)

func main() {
	//ctx := context.Background()
	log.AppTraceLog.Info(log.AppTraceLogInfo("START SERVER", "ADMIN", "", "", "", ""))

	db, err := database.NewDB()
	if err != nil {
		log.AppTraceLog.Error(log.AppTraceLogInfo("START SERVER", "ADMIN", "", "", "", err.Error()))
		return
	}

	e := echo.New()

	h := party.PartyHandler{}

	h.Initialize(db)

	e.GET("/individual", h.GetIndividual)
	/*	e.POST("/customers", h.SaveCustomer)
		e.GET("/customers/:id", h.GetCustomer)
		e.PUT("/customers/:id", h.UpdateCustomer)
		e.DELETE("/customers/:id", h.DeleteCustomer)
	*/
	e.Logger.Fatal(e.Start(":8080"))

	// Create service

	//log.AppTraceLog.Info(log.AppTraceLogInfo("Listening on port :"+apiPort+"...", "ADMIN", "", "", "", ""))

	//log.AppTraceLog.Error(log.AppTraceLogInfo("Port :8080...", "ADMIN", "", "", "", http.ListenAndServe(":8080", httpHandler).Error()))

	//log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
