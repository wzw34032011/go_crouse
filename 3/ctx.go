package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var signals []os.Signal
var signalChan = make(chan os.Signal)

func registerSignal(s ...os.Signal) {
	signals = append(signals, s...)
}

func init() {
	//注册信号
	registerSignal(os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	//监听信号
	signal.Notify(signalChan, signals...)
}

type service struct {
	name     string
	stopChan chan struct{}
	err      error
}

func NewService(name string) *service {
	return &service{name: name, stopChan: make(chan struct{}), err: nil}
}

func (s service) Start() error {
	//return errors.New("启动异常")
	log.Printf("%s启动中...", s.name)
	log.Printf("%s启动成功", s.name)

	for {
		time.Sleep(time.Second)
		//do service
		log.Printf("%s running...", s.name)

		select {
		case <-s.stopChan:
			log.Printf("%s已关闭", s.name)
			return s.err
		default:
		}
	}
}

func (s service) Stop() error {
	log.Printf("%s关闭中...", s.name)
	close(s.stopChan)
	return nil
}

func main() {
	rootCtx, rootCancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(rootCtx)

	s1 := NewService("service A")
	s2 := NewService("service B")

	//服务A
	g.Go(func() error {
		return s1.Start()
	})
	g.Go(func() error {
		<-ctx.Done()
		return s1.Stop()
	})

	//服务B
	g.Go(func() error {
		return s2.Start()
	})
	g.Go(func() error {
		<-ctx.Done()
		return s2.Stop()
	})

	g.Go(func() error {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT:
					rootCancel()
				default:
					log.Println("undefined s")
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
	log.Println("系统已下线")
}
