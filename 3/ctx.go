package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"sync"
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
	isClose  bool
	err      error
	sync.Mutex
}

func NewService(name string) *service {
	return &service{name: name, stopChan: make(chan struct{}), err: nil}
}

func (s *service) Start() error {
	log.Printf("%s启动中...", s.name)

	for {
		time.Sleep(time.Second)
		log.Printf("%s running...", s.name)
		//s.err = errors.New("运行异常")

		select {
		case <-s.stopChan:
			log.Printf("%s已关闭", s.name)
			return s.err
		default:
			if s.err != nil {
				_ = s.Stop()
			}
		}
	}
}

func (s *service) Stop() error {
	s.Lock()
	defer s.Unlock()
	if s.isClose == false {
		log.Printf("%s关闭中...", s.name)
		close(s.stopChan)
		s.isClose = true
	}
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

	//信号处理
	g.Go(func() error {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT:
					rootCancel()
				default:
					log.Println("undefined signal")
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
