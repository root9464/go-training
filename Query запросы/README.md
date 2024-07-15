# Синхронные и асинхронные запросы в golang

## 1
```go
type ResponseData struct {
	Todos []map[string]interface{}
	Posts []map[string]interface{}
	Users []map[string]interface{}
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos?_limit=2",
		"https://jsonplaceholder.typicode.com/posts?_limit=2",
		"https://jsonplaceholder.typicode.com/users?_limit=2",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	var response ResponseData

	for i, url := range urls {
		go func(i int, url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			var data []map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&data)
			if err != nil {
				fmt.Println(err)
				return
			}

			switch i {
			case 0:
				response.Todos = data
			case 1:
				response.Posts = data
			case 2:
				response.Users = data
			}
		}(i, url)
	}

	wg.Wait()

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
```

> Этот пример может быть полезен в ситуациях, когда необходимо выполнить несколько параллельных HTTP-запросов к разным эндпоинтам, объединить полученные данные и обработать их в дальнейшем, например, для агрегации информации из нескольких источников.
> 
> Функционал:
>   * Определяет структуру `ResponseData` , которая содержит три слайса для хранения данных о задачах (Todos), постах (Posts) и пользователях (Users).
>   * Создает список URL-адресов для HTTP-запросов к API сервиса JSONPlaceholder с параметром `_limit=2`.
>   * Использует горутины для выполнения параллельных HTTP-запросов к каждому URL-адресу.
>   * Декодирует полученные данные в формате JSON и сохраняет их в соответствующий слайс структуры `ResponseData`.
>   * После завершения всех запросов маршализует данные структуры ResponseData в формат JSON с отступами и выводит их на экран.


## 2
```go
const urlInRequest = "https://jsonplaceholder.typicode.com/todos?_limit=2"
func main() {

	dataCh := make(chan string)
	errorCh := make(chan string)
	go func() {
		resp, err := http.Get(urlInRequest)
		if err != nil {
			errorCh <- err.Error()
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorCh <- err.Error()
			return
		}

		dataCh <- string(body)
	}()

	select {
	case data := <-dataCh:
		fmt.Println(data)
	case err := <-errorCh:
		fmt.Println(err)
	}
}

```
> Этот пример может быть полезен в ситуациях, когда необходимо выполнить асинхронный HTTP-запрос и обрабатывать результаты запроса или ошибки без блокирования основного потока выполнения.
> 
> Функционал:
>   * Определяет константу `urlInRequest`, содержащую URL-адрес для HTTP-запроса к сервису JSONPlaceholder.
>   * Создает два канала: `dataCh` для передачи данных и `errorCh` для передачи ошибок.
>   * Использует горутину для выполнения HTTP-запроса к URL-адресу из константы `urlInRequest`.
>   * В случае успешного запроса считывает тело ответа, преобразует его в строку и отправляет в канал `dataCh`.
>   * В случае возникновения ошибки отправляет сообщение об ошибке в канал `errorCh`.
>   * Использует оператор `select` для ожидания данных из каналов `dataCh` и `errorCh`, и выводит либо полученные данные, либо сообщение об ошибке на экран.

## 3

```go
const urlInRequest = "https://jsonplaceholder.typicode.com/todos?_limit=2"

func main() {
	res, err := http.Get(urlInRequest)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}

```

> дефолтный запрос и получение данных с последующим преобразованием



