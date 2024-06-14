package main

import (
	"SebStudy/app/server"
	"SebStudy/infrastructure"
)

func main() {
	// disp :=
	ceMapper := infrastructure.NewCeMapper()
	ce := server.NewCloudEventsClient(8080)
	s := server.NewCloudServer(nil, ce, ceMapper)

	s.Run()
	// u := uuid.New()
	// fmt.Println(u)
}
