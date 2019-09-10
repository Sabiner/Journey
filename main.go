package main

import (
	"Journey/web/routers"
	"fmt"
	"net/http"
	"time"
)

func main()  {

	routersInit := routers.InitRouter()

	s := &http.Server{
		Addr:              ":8084",
		Handler:           routersInit,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("fail to start http server, due to: %v", err)
	}
}
