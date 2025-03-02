package agent

import (
	"log"
	"os"
	"strconv"
)

func StartAgent() {
	computingPower := 1
	if cpStr := os.Getenv("COMPUTING_POWER"); cpStr != "" {
		if cp, err := strconv.Atoi(cpStr); err == nil && cp > 0 {
			computingPower = cp
		}
	}
	log.Printf("Запуск агента с %d воркерами", computingPower)
	for i := 0; i < computingPower; i++ {
		go worker()
	}
	select {}
}
