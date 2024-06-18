package main

import (
	"context"
	"log"
	"time"

	pb "SebStudy/proto/resume"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/protobuf/proto"
)

func main() {
	p, _ := cloudevents.NewHTTP()

	c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	testResume := pb.Resume{
		ResumeId:    1,
		FirstName:   "Алексей",
		MiddleName:  "Валерьевич",
		LastName:    "Кузнецов",
		PhoneNumber: "79295132116",
		Educations: []*pb.Education{
			{
				Education: "IT",
			},
		},
		AboutMe: "I am a student",
		Skills: []*pb.Skill{
			{
				Skill: "Golang",
			},
		},
		Photo: "https://www.google.com",
		Directions: []*pb.Direction{
			{
				Direction: "back-end",
			},
		},
		AboutProjects: "about projects",
		Portfolio:     "https://github.com",
		StudentGroup:  "SA-33",
	}

	t := time.Now()
	current_time := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)

	event := cloudevents.NewEvent()
	event.SetTime(current_time)
	event.SetSource("example/uri")
	event.SetType("resume.send2")
	event.SetTime(current_time)
	b, err := proto.Marshal(&testResume)
	event.SetData("application/protobuf", b)
	ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")

	if result := c.Send(ctx, event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	} else {
		log.Printf("sent: %v", event)
		log.Printf("result: %v", result)
	}

}
