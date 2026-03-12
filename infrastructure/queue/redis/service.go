package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"incard.uz/humo/core/app"
	"time"
)

type Service struct {
	Host   string
	Port   string
	ctx    context.Context
	cancel context.CancelFunc
	rc     *redis.Client
}

func (srv *Service) initRedis() error {
	srv.rc = redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%s", srv.Host, srv.Port),
		Password:        "",  // no password set
		DB:              1,   // use default DB
		PoolSize:        512, // max-active
		MaxIdleConns:    1024,
		MinIdleConns:    128,
		ConnMaxIdleTime: time.Duration(60) * time.Minute,
	})
	srv.ctx = context.Background()
	return srv.rc.Ping(srv.ctx).Err()
}

type ListenerConfig struct {
	AppService app.ApplicationService
	Routes     []ListenerRoute
}

type ListenerRoute struct {
	QueueName    string
	ThreadsCount int
	Action       ControllerAction
}

func (srv *Service) Listen(config ListenerConfig) {

	controller := Controller{AppService: config.AppService}

	err := srv.initRedis()

	if err != nil {
		controller.HandleError("redis:conn", err)
		panic(err)
	}

	forever := make(chan bool)
	for _, route := range config.Routes {
		_threadsCount := route.ThreadsCount
		if _threadsCount == 0 {
			_threadsCount = 8
		}
		for i := 0; i < route.ThreadsCount; i++ {
			go listen(srv.rc, route.QueueName, controller, forever)
		}
	}

	<-forever

	controller.HandleError("redis:conn", errors.New("loop is closed"))

}

func listen(rc *redis.Client, queueName string, handler ControllerContract, forever chan bool) {
	ctx := context.Background()
	limit := 1
	for {
		pipe := (*rc).Pipeline()
		results := pipe.LRange(ctx, queueName, 0, int64(limit-1))
		pipe.LTrim(ctx, queueName, int64(limit), -1)
		_, err := pipe.Exec(ctx)
		if err != nil {
			handler.HandleError(queueName, err)
		}
		messages, _ := results.Result()
		for _, message := range messages {
			handler.Handle(queueName, message)
		}
		time.Sleep(100 * time.Millisecond)
	}

}
