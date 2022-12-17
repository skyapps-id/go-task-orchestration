package utils

import (
	"go-task-orchestration/tasks"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	log "github.com/sirupsen/logrus"
)

func GetMachineryServer() *machinery.Server {
	log.Info("initing task server")

	taskserver, err := machinery.NewServer(&config.Config{
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
		Broker:          "redis://localhost:6379",
		ResultBackend:   "redis://localhost:6379",
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	taskserver.RegisterTasks(map[string]interface{}{
		"send_webhook": tasks.SendWebhook,
	})

	return taskserver
}
