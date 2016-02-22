package main

import (
  "fmt"
  // "io"
  "log"
  "net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is your new api : )")
}
func main() {

  http.HandleFunc("/", testHandler)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
    return
  }
}
