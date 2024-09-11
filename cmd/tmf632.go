package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"tmfEcho/internal/database"
	"tmfEcho/internal/log"
	"tmfEcho/pkg/party"

	"github.com/labstack/echo"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
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

	//	e.Logger.Fatal(e.Start(":8082"))

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	wg := sync.WaitGroup{}

	httpServer := http.Server{
		Addr:    ":8082",
		Handler: e,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			lg := log.GenErrLog(fmt.Sprintf("error when starting HTTP server: %v", err), log.LogTracing{}, log.E000000, err)
			log.AppTraceLog.Error(lg)

		} else {
			lg := log.GenAppLog("HTTP server stopped serving requests", log.LogTracing{})
			log.AppTraceLog.Info(lg)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done() // wait for ctrl+c

		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(timeout); err != http.ErrServerClosed {
			lg := log.GenErrLog(fmt.Sprintf("error when shutting HTTP server: %v", err), log.LogTracing{}, log.E000000, err)
			log.AppTraceLog.Error(lg)
		} else {
			lg := log.GenAppLog("HTTP server shut down", log.LogTracing{})
			log.AppTraceLog.Info(lg)

		}
	}()

	autoTLSManager := autocert.Manager{
		//HostPolicy: autocert.HostWhitelist(conf.Host),
		HostPolicy: autocert.HostWhitelist("mydomainname.com"),
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	tlsServer := http.Server{
		Addr:    ":443",
		Handler: e,
		TLSConfig: &tls.Config{
			GetCertificate: autoTLSManager.GetCertificate,
			NextProtos:     []string{acme.ALPNProto},
		},
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := tlsServer.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
			lg := log.GenErrLog(fmt.Sprintf("error when starting HTTPS server: %v", err), log.LogTracing{}, log.E000000, err)
			log.AppTraceLog.Error(lg)
		} else {
			lg := log.GenAppLog("HTTPS server stopped serving requests", log.LogTracing{})
			log.AppTraceLog.Info(lg)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done() // wait for ctrl+c

		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := tlsServer.Shutdown(timeout); err != http.ErrServerClosed {
			lg := log.GenErrLog(fmt.Sprintf("error when shutting HTTPS server: %v", err), log.LogTracing{}, log.E000000, err)
			log.AppTraceLog.Error(lg)
		} else {
			lg := log.GenAppLog("HTTPS server shut down", log.LogTracing{})
			log.AppTraceLog.Info(lg)
		}
	}()

	wg.Wait() // wait for all goroutines to end - server listeners and shutdown routines

	// Create service

	//log.AppTraceLog.Info(log.AppTraceLogInfo("Listening on port :"+apiPort+"...", "ADMIN", "", "", "", ""))

	//log.AppTraceLog.Error(log.AppTraceLogInfo("Port :8080...", "ADMIN", "", "", "", http.ListenAndServe(":8080", httpHandler).Error()))

	//log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
