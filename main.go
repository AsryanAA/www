package main

import ("fmt"
        "net/http")

type Data struct {
  data string
}

func test1(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Go easy start!")
}

// обработка url адресов
func handleRequest()  {
  http.HandleFunc("/", test1)
  http.ListenAndServe(":8080", nil)
}

// точка входа в приложение
func main()  {
  fmt.Println("Start Go Project!")
  handleRequest()
}
