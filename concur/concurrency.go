package concur

import (
	"fmt"
	"image"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func say(x string) {
	for i := 1; i < 5; i++ {
		fmt.Printf("%s say hello %d\n", x, i)
	}
}

func sing() {
	for i := 1; i < 5; i++ {
		fmt.Printf("sing a song %d\n", i)
	}
}

func dance() {
	for i := 1; i < 5; i++ {
		fmt.Printf("dance dance dance %d\n", i)
	}
}

func ConcurrencyDemo() {
	fmt.Printf("sync invoke\n")
	go say("sync")

	fmt.Printf("async invoke\n")
	say("async")

	time.Sleep(time.Second)
}

var wg sync.WaitGroup

func hello() {
	fmt.Println("hello")
	wg.Done()
}

func WaitDemo() {
	wg.Add(1)
	go hello()
	fmt.Println("你好啊")
	wg.Wait()
}

func Demo() {
	store := 10
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			store--
			fmt.Printf("%d-%d\n", i, store)
			wg.Done()
		}()
	}
	fmt.Println(store)
	wg.Wait()
}

func ChannelDemo() {
	var ch1 chan int
	var ch2 chan bool
	var ch3 chan []int

	fmt.Println(ch1, ch2, ch3, reflect.ValueOf(ch1).IsNil())

	ch4 := make(chan int, 2)

	ch4 <- 10
	ch4 <- 15

	wg.Add(1)
	go func() {
		for v := range ch4 {
			fmt.Println("x", v)
		}
		wg.Done()
	}()

	ch4 <- 20
	y := <-ch4
	fmt.Println("y", y)

	close(ch4)

	wg.Wait()
}

func LimitChannel() {
	prodChan := producer()
	fmt.Println(consumer(prodChan))
}

func producer() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				time.Sleep(200 * time.Millisecond)
				fmt.Printf("push %d\n", i)
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// consumer 参数为接收通道
func consumer(ch <-chan int) int {
	sum := 0
	for v := range ch {
		fmt.Printf("pop %d\n", v)
		sum += v
	}
	return sum
}

func SelectDemo() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)

	ch := []chan int{ch1, ch2, ch3}

	var x int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		select {
		case x = <-ch1:
			fmt.Printf("%d. ch1 pop: %d\n", i, x)
		case x = <-ch2:
			fmt.Printf("%d. ch2 pop: %d\n", i, x)
		case x = <-ch3:
			fmt.Printf("%d. ch3 pop: %d\n", i, x)
		default:
			idx := rand.Intn(3)
			val := rand.Intn(10)
			ch[idx] <- val
			fmt.Printf("%d. ch%d push %d\n", i, idx+1, val)
		}
	}
}

func BugDemo() {
	demo2()
}

// demo1 通道误用导致的bug
func demo1() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				task, ok := <-ch
				if !ok {
					break
				}
				// 这里假设对接收的数据执行某些操作
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// demo2 通道误用导致的bug
func demo2() {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(4 * time.Second): // 较小的超时时间
		return
	}
}

var x int64
var m sync.Mutex
var rwm sync.RWMutex

func CurCompetition() {
	do(writeWithLock, readWithLock, 10, 100)
	do(writeWithRWLock, readWithRWLock, 10, 100)
}

func add() {
	for i := 0; i < 5000; i++ {
		m.Lock()
		x++
		m.Unlock()
	}
	wg.Done()
}

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	m.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	m.Unlock()                        // 解互斥锁
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	m.Lock()                     // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	m.Unlock()                   // 释放互斥锁
	wg.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwm.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwm.Unlock()                      // 释放写锁
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwm.RLock()                  // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwm.RUnlock()                // 释放读锁
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)
}

var icons map[string]image.Image
var iconOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) image.Image {
	iconOnce.Do(loadIcons)
	return icons[name]
}

func loadIcon(path string) image.Image {
	var x image.Image
	return x
}

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

var aMap = make(map[string]int)

func setMap(k string, v int) {
	aMap[k] = v
}
func getMap(k string) int {
	return aMap[k]
}

var sMap = sync.Map{}

func CurMapDemo() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			// int covert string
			key := strconv.Itoa(n)
			sMap.Store(key, n)
			v, _ := sMap.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func AtomicDemo() {
	c1 := CommonCounter{} // 非并发安全
	test(&c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	test(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	test(&c3)
}

type Counter interface {
	Inc()
	Load() int64
}

type CommonCounter struct {
	counter int64
}

func (c *CommonCounter) Inc() {
	c.counter++
}

func (c *CommonCounter) Load() int64 {
	return c.counter
}

type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

var cnt int64 = 0

func CountRand() {
	fmt.Printf("start: %d\n", cnt)
	for res := range calc(genRand()) {
		fmt.Printf("rand sum: %d\n", res)
	}
	fmt.Printf("end: %d\n", cnt)
}

func genRand() <-chan int64 {
	jobChan := make(chan int64, 10)
	rand.Seed(time.Now().Unix())
	go func() {
		for i := 0; i < 100; i++ {
			jobChan <- rand.Int63()
		}
		close(jobChan)
	}()
	return jobChan
}

func calc(jobChan <-chan int64) <-chan int {
	resultChan := make(chan int, 10)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go func() {
			for num := range jobChan {
				fmt.Printf("get rand: %d\n", num)
				sum := 0
				for num > 0 {
					sum = sum + int(num%10)
					num /= 10
				}
				atomic.AddInt64(&cnt, 1)
				resultChan <- sum
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}
