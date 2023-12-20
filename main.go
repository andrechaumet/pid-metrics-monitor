package main

import (
	"net/http"
	"pid-metrics-monitor/router"
 )
 
 func main() {
	r := router.SetupRouter()
	http.ListenAndServe(":8081", r)
 }