package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "SebStudy/proto/resume"

	_ "github.com/cloudevents/sdk-go/binding/format/protobuf/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {
	p, _ := cloudevents.NewHTTP()

	c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	imageBytes, err := os.ReadFile("C:/Users/Олег/Desktop/images.jpg")
	if err != nil {
		fmt.Println("Картинка хуйня")
		return
	}

	testResume := pb.ResumeSended{
		ResumeId:    "dddddddddddddddddd",
		FirstName:   "Алексей",
		MiddleName:  "Валерьевич",
		LastName:    "Кузнецов",
		PhoneNumber: "79295132116",
		Education:   "РКСИ",
		AboutMe:     "I am a student",
		Skills: []*pb.Skill{
			{
				Skill: "Golang",
			},
			{
				Skill: "Java",
			},
		},
		// Photo: []byte{255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255},
		Photo: imageBytes,
		Directions: []*pb.Direction{
			{
				Direction: "back-end",
			},
			{
				Direction: "fddffdfddf",
			},
		},
		AboutProjects: "about projects",
		Portfolio:     "https://github.com",
		StudentGroup:  "SA-33",
	}

	t := time.Now()
	current_time := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	startTime := time.Now()

	event := cloudevents.NewEvent()
	event.SetID("1234-1234-1234-1234")
	event.SetSource("example/uri")
	event.SetType("resume.send")
	event.SetTime(current_time)
	event.SetSpecVersion("1.0")

	// b, _ := proto.Marshal(&testResume)
	// fmt.Println(b)
	// var protoBytes []byte = make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	// base64.StdEncoding.Encode(protoBytes, b)

	// event.SetData(pbcloudevents.ContentTypeProtobuf, &testResume)
	event.SetData("application/protobuf", &testResume)
	// event.SetData("application/json", &testResume)
	// fmt.Println(event.DataEncoded)
	log.Println(event.Data())
	ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")
	// log.Println(event.Data())

	if result := c.Send(ctx, event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	} else {
		// log.Printf("sent cloudevent, %s\n", event)
		log.Printf("status code: %v\n", result)
		log.Println(time.Since(startTime))
	}
}
