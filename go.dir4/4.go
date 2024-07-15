package main

// const urlInRequest = "https://jsonplaceholder.typicode.com/todos?_limit=2"

// func main() {
// 	res, err := http.Get(urlInRequest)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(body))

// }

// func main() {

// 	dataCh := make(chan string)
// 	errorCh := make(chan string)
// 	go func() {
// 		resp, err := http.Get(urlInRequest)
// 		if err != nil {
// 			errorCh <- err.Error()
// 			return
// 		}
// 		defer resp.Body.Close()

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			errorCh <- err.Error()
// 			return
// 		}

// 		dataCh <- string(body)
// 	}()

// 	select {
// 	case data := <-dataCh:
// 		fmt.Println(data)
// 	case err := <-errorCh:
// 		fmt.Println(err)
// 	}
// }

// type ResponseData struct {
// 	Todos []map[string]interface{}
// 	Posts []map[string]interface{}
// 	Users []map[string]interface{}
// }

// func main() {
// 	urls := []string{
// 		"https://jsonplaceholder.typicode.com/todos?_limit=2",
// 		"https://jsonplaceholder.typicode.com/posts?_limit=2",
// 		"https://jsonplaceholder.typicode.com/users?_limit=2",
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(len(urls))

// 	var response ResponseData

// 	for i, url := range urls {
// 		go func(i int, url string) {
// 			defer wg.Done()
// 			resp, err := http.Get(url)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer resp.Body.Close()

// 			var data []map[string]interface{}
// 			err = json.NewDecoder(resp.Body).Decode(&data)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			switch i {
// 			case 0:
// 				response.Todos = data
// 			case 1:
// 				response.Posts = data
// 			case 2:
// 				response.Users = data
// 			}
// 		}(i, url)
// 	}

// 	wg.Wait()

// 	jsonData, err := json.MarshalIndent(response, "", "  ")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))
// }
