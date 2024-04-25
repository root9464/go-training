### 1/
```go
func Atomic() {
	var (
		wg      sync.WaitGroup
		counter int64
	)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
}
```
> Атомарная асинхронная функция увеличивающа внутренний counter на еденицу. `(Время выполнения 2.0809ms)` p.s: хуй знает почему так ведь анатомарность должна быть быстрее mu и медленнее шардирования 


### 2/
```go
func Mutex() {
	var (
		counter int64
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
}
```
> Мьютексная функция  функция увеличивающа внутренний counter на еденицу. `(Время выполнения 512.7µs)`

### 3/
```go

func ChanMutex() {
	var (
		counter int64
		wg      sync.WaitGroup
		mu      sync.Mutex
		done    = make(chan struct{})
	)

	wg.Add(1000)
	go func() {
		wg.Wait()
		close(done)
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	<-done
}


```

> Канальная мьютексная функция  увеличивающа внутренний counter на еденицу. `(Время выполнения 518.9µs)`


> [Дополнительная статься](https://dev.to/karanpratapsingh/advanced-concurrency-patterns-in-go-2je1)
