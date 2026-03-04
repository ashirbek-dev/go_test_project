package nats_connection

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"math/rand/v2"
	"time"
)

type Service struct {
	ServiceUrl string
	StreamName string
	ctx        context.Context
	cancel     context.CancelFunc
	stream     jetstream.Stream
	jetStream  jetstream.JetStream
}

var maxConnCount = 16

var connPull []*jetstream.JetStream

func (srv *Service) createJetStream() *jetstream.JetStream {
	connCount := len(connPull)
	if connCount >= maxConnCount {
		return connPull[rand.IntN(connCount-1)]
	}

	nc, err := nats.Connect(srv.ServiceUrl, nats.NoEcho())

	if err != nil {
		log.Panic(err)
		return nil
	}

	jetStream, _ := jetstream.New(nc)

	connPull = append(connPull, &jetStream)

	return &jetStream
}

func (srv *Service) Send(subject string, payload []byte) {
	js := srv.createJetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	lister := (*js).StreamNames(ctx)
	var exists = false
	for name := range lister.Name() {
		exists = exists || name == srv.StreamName
	}

	if !exists {
		_, err := (*js).CreateOrUpdateStream(ctx, jetstream.StreamConfig{
			Name:      srv.StreamName,
			Subjects:  []string{fmt.Sprintf("%s.*", srv.StreamName)},
			MaxAge:    time.Hour * 4,
			Retention: jetstream.WorkQueuePolicy,
		})
		if err != nil {
			log.Panic(err)
			return
		}
	}
	_, err := (*js).Publish(ctx, fmt.Sprintf("%s.%s", srv.StreamName, subject), payload)
	if err != nil {
		log.Panic(err)
		return
	}
}
