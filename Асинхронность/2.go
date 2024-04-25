//Обработка ошибокб + Промисы, async/await
//========================================

package main

// func printNumbers() {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Printf("%d ", i)
// 	}
// }

// func printLetters() {
// 	for i := 'a'; i < 'e'; i++ {
// 		time.Sleep(300 * time.Millisecond)
// 		fmt.Printf("%c ", i)
// 	}
// }
// func main() {
// 	go printNumbers()
// 	go printLetters()

// 	time.Sleep(3 * time.Second)
// }

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for num := range ch {
// 		fmt.Println(num)
// 	}
// }

// func main() {

// 	runtime.GOMAXPROCS(2)

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	ch := make(chan int)

// 	for i := 0; i < 2; i++ {
// 		go func() {
// 			defer wg.Done()
// 			for i := 0; i < 10; i++ {
// 				ch <- i
// 			}
// 		}()
// 	}

// 	for i := 0; i < 2; i++ {
// 		go func() {
// 			defer wg.Done()
// 			for i := 0; i < 10; i++ {
// 				fmt.Println(<-ch)
// 			}
// 		}()
// 	}

// 	wg.Wait()
// }

// type Counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *Counter) Increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *Counter) Decrement() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count--
// }

// func (c *Counter) GetValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {
// 	counter := Counter{}

// 	var wg sync.WaitGroup
// 	wg.Add(100)

// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			defer wg.Done()
// 			counter.Increment()
// 		}()
// 	}

// 	wg.Wait()

//		fmt.Println("итог:", counter.GetValue())
//	}
// func fetchData(url string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("Ошибка при выполнении запроса к %s: %v\n", url, err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("Ошибка при чтении ответа от %s: %v\n", url, err)
// 		return
// 	}

// 	fmt.Printf("Данные для %s: %s\n", url, string(body))
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i <= 10; i++ {
// 		wg.Add(2)
// 		go fetchData(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i), &wg)
// 		go fetchData(fmt.Sprintf("https://jsonplaceholder.typicode.com/comments/%d", i), &wg)
// 	}

// 	wg.Wait()
// }

// var (
// 	counter1 int
// 	counter2 int
// 	mutex    sync.RWMutex
// 	wg       sync.WaitGroup
// )

// func Inc1() {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		mutex.Lock()
// 		counter1++
// 		fmt.Printf("Inc1 counter: %d\n", counter1)
// 		mutex.Unlock()
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func Inc2() {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		mutex.Lock()
// 		counter2++
// 		mutex.Unlock()
// 		time.Sleep(100 * time.Millisecond)
// 	}
// 	mutex.RLock()
// 	fmt.Printf("Inc2 counter: %d\n", counter2)
// 	mutex.RUnlock()
// }

// func main() {
// 	wg.Add(2)
// 	go Inc1()
// 	go Inc2()

// 	wg.Wait()
// }
