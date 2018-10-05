package timer

import (
	"github.com/robfig/cron"
	"log"
)
func HealthCheck (){
	log.Println("Running ----------- ")
}

func  SetupTimer () {
	c := cron.New()

	c.AddFunc("0 0/30 * * * *", HealthCheck)

	c.Start()

}
