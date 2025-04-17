package messagesys

import (
	"fmt"

	"github.com/nats-io/nats.go"

	"github.com/mch735/education/work5/config"
)

type MessageSys struct {
	*nats.Conn
}

func New(conf *config.NATS) (*MessageSys, error) {
	opts := nats.Options{
		Url:            conf.URL,
		AllowReconnect: true,
		MaxReconnect:   conf.MaxReconnect,
		ReconnectWait:  conf.ReconnectWait,
		Timeout:        conf.Timeout,
	}

	ns, err := opts.Connect()
	if err != nil {
		return nil, fmt.Errorf("nats connect error: %w", err)
	}

	return &MessageSys{ns}, nil
}
