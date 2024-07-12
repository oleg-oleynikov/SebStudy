package main

import (
	"bytes"
	"fmt"
	"net/http"

	pb "SebStudy/proto/resume"

	"google.golang.org/protobuf/proto"
)

func main() {
	url := "http://localhost:8080/"
	testResume := &pb.ResumeSended{
		ResumeId:    "ddddddddddddddddddddddddddddd",
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
		},
		Photo: []byte{},
		Directions: []*pb.Direction{
			{
				Direction: "back-end",
			},
		},
		AboutProjects: "about projects",
		Portfolio:     "https://github.com",
		StudentGroup:  "SA-33",
	}
	data, err := proto.Marshal(testResume)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	req.Header.Set("Ce-Id", "536808d3-88be-4077-5d7a-a3f162705f8")
	req.Header.Set("Ce-Specversion", "1.0")
	req.Header.Set("Ce-Type", "resume.send")
	req.Header.Set("Ce-Source", "example/uri")
	req.Header.Set("Content-Type", "application/json") // application/cloudevents+json

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
