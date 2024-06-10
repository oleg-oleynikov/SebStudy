package main

import (
	"SebStudy/infrastructure"
	"fmt"
)

type CloudServer struct {
	Dispatcher *infrastructure.Dispatcher
}

func StartServer() {
	fmt.Println("Невозможно! Все пятеро в потоке")
}
