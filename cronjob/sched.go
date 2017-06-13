package main
// CronJob schedule  https://godoc.org/github.com/robfig/cron
import (
	//"fmt"
	"log"
	"time"
	"github.com/jasonlvhit/gocron"

)

func fireAnAction() {
	log.Println("Fire...")

}

const ONE_SECOND = 1 * time.Second + 10 * time.Millisecond

func main() {
	gocron.Every(1).Second().Do(fireAnAction)
	// function Start start all the pending jobs
	<- gocron.Start()

}

