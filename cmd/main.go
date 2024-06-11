package main

import "SebStudy/app/server"

func main() {
	// disp :=
	ce := server.NewCloudEventsClient(8080)
	s := server.NewCloudServer(nil, ce)

	s.Run()
}
