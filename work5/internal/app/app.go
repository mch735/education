package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"

	"github.com/mch735/education/work5/config"
	"github.com/mch735/education/work5/internal/controller/web"
	"github.com/mch735/education/work5/internal/repo/messagesys"
	"github.com/mch735/education/work5/internal/repo/usercache"
	"github.com/mch735/education/work5/internal/repo/userrepo"
	"github.com/mch735/education/work5/internal/usecase"
	"github.com/mch735/education/work5/pkg/logger"
)

//nolint:nlreturn
func Run(conf *config.Config) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	_, err := logger.New(&conf.Log)
	if err != nil {
		panic(err)
	}

	repo, err := userrepo.New(&conf.PG)
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	cache, err := usercache.New(&conf.Redis)
	if err != nil {
		panic(err)
	}
	defer cache.Close()

	ms, err := messagesys.New(&conf.NATS)
	if err != nil {
		panic(err)
	}
	defer ms.Close()

	sub, err := ms.Subscribe("methods", func(m *nats.Msg) {
		fmt.Printf("Method call: %s\n", string(m.Data))
	})
	if err != nil {
		panic(err)
	}

	uc := usecase.New(repo, cache, ms)

	server := web.NewServer(&conf.HTTP)
	server.Handler = web.NewRouter(uc)

	g, c := errgroup.WithContext(ctx)

	g.Go(func() error {
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-c.Done()
		return server.Shutdown(context.Background()) //nolint:contextcheck
	})
	g.Go(func() error {
		<-c.Done()
		return sub.Unsubscribe()
	})

	err = g.Wait()
	if err != nil {
		fmt.Printf("%s \n", err)
	}
}
