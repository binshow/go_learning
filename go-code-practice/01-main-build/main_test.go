package _1_main_build

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

// -------------------------------------------
// @file          : main_test.go
// @author        : binshow
// @time          : 2022/7/5 10:29 PM
// @description   :	如何优雅的构造服务的main函数：https://www.yuque.com/binshow/gm71hi/dpg52v
// -------------------------------------------

var (
	// ShutdownSignals receives shutdown signals to process
	ShutdownSignalsDrawin = []os.Signal{
		os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
		syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM,
	}

	// ShutdownSignals receives shutdown signals to process
	ShutdownSignalsLinux = []os.Signal{
		os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
		syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM,
	}
)




// 普通的服务中main函数
func TestMain01(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello" , hello)
	http.ListenAndServe(":8000", mux) // main goroutine 阻塞在这
}


// 引入signal
func TestMain02(t *testing.T) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	// http服务改造成异步
	go http.ListenAndServe(":8000", mux)

	// 程序阻塞在这里，除非收到了interrupt或者kill信号
	fmt.Println("receive sig:" , <-sig)
}


// 优雅的退出，希望进行一些原子性的操作，不希望程序运行到一半就突然退出了
func TestMain03(t *testing.T) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
		syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM)

	// 模拟并发进行的处理业务逻辑
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// 我们希望程序能等当前这个周期休眠完，再优雅退出
				time.Sleep(time.Duration(i) * time.Second)
			}
		}(i)
	}
	// 程序阻塞在这里，除非收到了interrupt或者kill信号
	fmt.Println("receive sig:" , <-sig)
}


// 父子goroutine的传递消息：
//1. 父goroutine通知子goroutine准备优雅地关闭，也就是stopCh
//2. 子goroutine通知父goroutine已经关闭完成，也就是finishedCh
func TestMain04(t *testing.T) {
	sig := make(chan os.Signal)
	stopCh := make(chan struct{})
	finishedCh := make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopCh, finishedCh chan struct{}) {
		for {
			select {
			case <-stopCh:
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh, finishedCh)

	<-sig
	stopCh <- struct{}{}
	<-finishedCh
	fmt.Println("finished")
}


// channel 嵌套 channel
func TestMain05(t *testing.T) {
	sig := make(chan os.Signal)
	stopCh := make(chan chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopChh chan chan struct{}) {
		for {
			select {
			case ch := <-stopCh:
				// 结束后，通过ch通知主goroutine
				fmt.Println("stopped")
				ch <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh)

	<-sig
	// ch作为一个channel，传递给子goroutine，待其结束后从中返回
	ch := make(chan struct{})
	stopCh <- ch
	<-ch
	fmt.Println("finished")
}



// 标准方案：引入 context
func TestMain06(t *testing.T) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	ctx, cancel := context.WithCancel(context.Background())
	finishedCh := make(chan struct{})

	go func(ctx context.Context, finishedCh chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				// 结束后，通过ch通知主goroutine
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(ctx, finishedCh)

	<-sig
	cancel()
	<-finishedCh
	fmt.Println("finished")
}

// 一对多的解决方案
func TestMain07(t *testing.T) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	ctx, cancel := context.WithCancel(context.Background())
	num := 10

	// 用wg来控制多个子goroutine的生命周期
	wg := sync.WaitGroup{}
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("stopped")
					return
				default:
					time.Sleep(time.Duration(i) * time.Second)
				}
			}
		}(ctx)
	}

	<-sig
	cancel()
	// 等待所有的子goroutine都优雅退出
	wg.Wait()
	fmt.Println("finished")
}

//curl http://localhost:8000/hello
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "hello\n")
}