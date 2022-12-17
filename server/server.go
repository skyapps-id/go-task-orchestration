package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	task "go-task-orchestration/tasks"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gofiber/fiber"
	"github.com/shemul/go-machinery/utils"
)

func StartServer(taskserver *machinery.Server) {

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Helo")
	})

	app.Post("/send-notif", func(ctx *fiber.Ctx) {
		p := new(task.Payload)
		if err := ctx.BodyParser(p); err != nil {
			utils.Logger.Fatal(err)
		}

		reqJSON, err := json.Marshal(p)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
		fmt.Println("Sampai di rest", p)

		b64EncodedReq := base64.StdEncoding.EncodeToString([]byte(reqJSON))
		eta := time.Now().UTC().Add(time.Second * 5)
		task := tasks.Signature{
			Name: "send_webhook",
			Args: []tasks.Arg{
				{
					Type:  "string",
					Value: b64EncodedReq,
				},
			},
			Headers: map[string]interface{}{
				"test": "test",
			},
			ETA:        &eta,
			RetryCount: 3,
		}

		res, err := taskserver.SendTask(&task)
		if err != nil {
			utils.Logger.Error(err.Error())
		}

		ctx.JSON(&fiber.Map{
			"task_uuid": res.GetState().TaskUUID,
		})

	})

	app.Listen(3000)

}
