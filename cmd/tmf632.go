package main

import (
	//	"tmfEcho/internal/database"
	"tmfEcho/internal/log"
)

func main() {
	//ctx := context.Background()
	log.AppTraceLog.Info(log.AppTraceLogInfo("START SERVER", "ADMIN", "", "", "", ""))
	/*
		db, err := database.NewDB()
		if err != nil {
			log.AppTraceLog.Error(log.AppTraceLogInfo("START SERVER", "ADMIN", "", "", "", err.Error()))
			return
		}
	*/
	// Create service

	//log.AppTraceLog.Info(log.AppTraceLogInfo("Listening on port :"+apiPort+"...", "ADMIN", "", "", "", ""))

	//log.AppTraceLog.Error(log.AppTraceLogInfo("Port :8080...", "ADMIN", "", "", "", http.ListenAndServe(":8080", httpHandler).Error()))

	//log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
