package main

import (
	"SebStudy/internal/infrastructure"
	"SebStudy/pb"
	"encoding/base64"
	"log"

	"google.golang.org/protobuf/types/known/anypb"
)

func main() {

	client := infrastructure.NewCloudeventsServiceClient("localhost:50051")

	resumeEvent := &pb.ResumeChanged{
		// ResumeId:  "123456",
		// Education: "PTY",
		// AboutMe:   "м",
		Skills: []*pb.Skill{
			{Skill: "SUCK DICK VERY NICE"},
		},
		// BirthDate:     time.Date(2005, 1, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02"),
		// AboutProjects: "а",
		// Portfolio:     "л",
		// Direction:     "ь",
	}

	protoEvent, err := anypb.New(resumeEvent)
	if err != nil {
		log.Println("Failed to cast to anypb")
		return
	}
	log.Println(base64.StdEncoding.EncodeToString(protoEvent.Value))

	cloudevent := &pb.CloudEvent{
		Id:          "1231312345",
		Source:      "example.com",
		SpecVersion: "1.0",
		Type:        "resume.change",
		Data: &pb.CloudEvent_ProtoData{
			ProtoData: protoEvent,
		},
	}

	log.Println(base64.StdEncoding.EncodeToString(protoEvent.Value))

	client.Publish(cloudevent)
}
