package tpanic

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func Err1() {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error: ", err)
		default:
			fmt.Println("error: ", err)
		}
	}()

	panic("panic error")
}

func Err2() {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error: ", err)
		case nil:
		default:
			fmt.Println("error: ", err)
		}
	}()

	go func() {
		defer func() {
			err := recover()
			switch err.(type) {
			case runtime.Error:
				fmt.Println("runtime error: ", err)
			default:
				fmt.Println("error: ", err)
			}
		}()
		panic("panic error")
	}()

	time.Sleep(time.Second)
}

func Err3() {
	wg := sync.WaitGroup{}
	urls := []string{
		"https://www.baidu.com",
		"https://www.tencent.com",
		"https://www.cramin.com",
		"https://www.alibaba.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("get [%s] success\n", url)
				err := resp.Body.Close()
				if err != nil {
					return
				}
			}
			return
		}(url)

		wg.Wait()
	}
}

func Err4() {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	urls := []string{
		"https://www.baidu.com",
		"https://www.tencent.com",
		"https://www.cramin.com",
		"https://www.alibaba.com",
	}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("Get \"%s\": success\n", url)
				err := resp.Body.Close()
				if err != nil {
					return err
				}
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("all successes")
}
