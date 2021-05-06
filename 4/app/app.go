package app

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//默认退出信号
var exitSignals = []os.Signal{os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT}

type App struct {
	name      string
	version   string
	ctx       context.Context
	ctxCancel func()
	option
}

type Server interface {
	Start() error
	Stop() error
}

type option struct {
	meta    map[string]string
	signals []os.Signal
	server  []Server
}

type OptFun func(*option)

func AddMeta(key, value string) OptFun {
	return func(o *option) {
		o.meta[key] = value
	}
}

func AddSignal(sig ...os.Signal) OptFun {
	return func(o *option) {
		o.signals = append(o.signals, sig...)
	}
}

func AddServer(s ...Server) OptFun {
	return func(o *option) {
		o.server = append(o.server, s...)
	}
}

func NewApp(name, version string, optionFuns ...OptFun) *App {
	app := &App{
		name:    name,
		version: version,
		option: option{
			signals: exitSignals,
		},
	}

	for _, optionFun := range optionFuns {
		optionFun(&app.option)
	}

	app.ctx, app.ctxCancel = context.WithCancel(context.Background())
	return app
}

func (app *App) Run() {
	g, ctx := errgroup.WithContext(app.ctx)

	for _, s := range app.server {
		s := s
		g.Go(func() error {
			return s.Start()
		})
		g.Go(func() error {
			<-ctx.Done()
			return s.Stop()
		})
	}

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, app.signals...)
	//信号处理
	g.Go(func() error {
		for {
			select {
			case s := <-signalChan:
				for _, exitSig := range exitSignals {
					if s == exitSig {
						close(signalChan)
						app.Stop()
					}
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	err := g.Wait()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("系统已关闭")
}

func (app *App) Stop() {
	log.Println("系统开始关闭")
	app.ctxCancel()
}
