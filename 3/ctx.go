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

func service1Start(ctx context.Context) error {
	log.Println("服务1启动中...")
	//return errors.New("服务1启动异常")
	log.Println("服务1启动成功")

	for {
		time.Sleep(time.Second)
		//do service
		log.Println("service1 running...")

		select {
		case <-ctx.Done():
			log.Println("服务1已关闭")
			return ctx.Err()
		default:
		}
	}
}

func service1Stop(cancel context.CancelFunc) error {
	log.Println("服务1关闭中...")
	cancel()
	return nil
}

func service2Start(ctx context.Context) error {
	log.Println("服务2启动中...")
	//return errors.New("服务2启动异常")
	log.Println("服务2启动成功")

	for {
		time.Sleep(time.Second)
		//do service
		log.Println("service2 running...")

		select {
		case <-ctx.Done():
			log.Println("服务2已关闭")
			return ctx.Err()
		default:
		}
	}
}

func service2Stop(cancel context.CancelFunc) error {
	log.Println("服务2关闭中...")
	cancel()
	return nil
}

func main() {
	rootCtx, rootCancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(rootCtx)

	service1Ctx, service1Cancel := context.WithCancel(ctx)
	service2Ctx, service2Cancel := context.WithCancel(ctx)

	g.Go(func() error {
		return service1Start(service1Ctx)
	})
	g.Go(func() error {
		<-ctx.Done()
		return service1Stop(service1Cancel)
	})

	g.Go(func() error {
		return service2Start(service2Ctx)
	})
	g.Go(func() error {
		<-ctx.Done()
		return service2Stop(service2Cancel)
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
