package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"wahyuade.com/simple-job-api/base"
)

func main() {
	log.Println("Starting & Bootstraping...")
	godotenv.Load()
	context := base.New()
	context.Web().Listen(fmt.Sprintf("127.0.0.1:%s", os.Getenv("PORT")))

}
