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

	e.GET("/partyManagement/v5/individual", h.GetIndividual)
	e.GET("/partyManagement/v5/individual/:id", h.GetIndividualById)
	e.POST("/partyManagement/v5/individual", h.SaveIndividual)
	e.PATCH("/partyManagement/v5/individual/:id", h.UpdateIndividual)
	e.DELETE("/partyManagement/v5/individual/:id", h.DeleteIndividual)

	e.GET("/partyManagement/v5/organization", h.GetOrganization)
	e.GET("/partyManagement/v5/organization/:id", h.GetOrganizationById)
	e.POST("/partyManagement/v5/organization", h.SaveOrganization)
	e.PATCH("/partyManagement/v5/organization/:id", h.UpdateOrganization)
	e.DELETE("/partyManagement/v5/organization/:id", h.DeleteOrganization)

	e.Logger.Fatal(e.Start(":8082"))

	// Create service

	//log.AppTraceLog.Info(log.AppTraceLogInfo("Listening on port :"+apiPort+"...", "ADMIN", "", "", "", ""))

	//log.AppTraceLog.Error(log.AppTraceLogInfo("Port :8080...", "ADMIN", "", "", "", http.ListenAndServe(":8080", httpHandler).Error()))

	//log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
