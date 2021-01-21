package outputredis

import (
	"context"
	"engine/models"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type OutputRedis struct {
	Ctx   context.Context
	RDB   *redis.Client
	URL   string
	Queue string
}

const ModuleName = "redis"

func InitHandler(ctx context.Context, values map[string]interface{}) (output *OutputRedis, err error) {
	output = new(OutputRedis)
	output.Ctx = ctx
	output.URL = values["url"].(string)
	output.Queue = values["queue"].(string)

	opt, _ := redis.ParseURL(output.URL)
	output.RDB = redis.NewClient(opt)
	return output, nil
}

func (t *OutputRedis) Start(output chan models.Event) {
	for {
		select {
		case <-t.Ctx.Done():
			return
		case event := <-output:
			err := t.RDB.Set(context.Background(), "output", event.Body, 0)
			fmt.Println(err)
		}
	}
}
