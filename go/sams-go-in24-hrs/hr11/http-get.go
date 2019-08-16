package main
import (
 "fmt"
 "log"
 "net/http"
 "time"
)

func getSites(url string) {
 start := time.Now()
 res, err := http.Get(url)
 if err != nil {
  log.Fatal(err)
 }

 defer res.Body.Close()
 elapsed := time.Since(start).Seconds()
 fmt.Printf("Elaspsed time : %v sec for url %s\n",elapsed,url)
}


func main () {
 urls := make([]string, 3)
 urls[0] = "https://www.usa.gov/"
 urls[1] = "https://www.gov.uk/"
 urls[2] = "https://www.india.gov.in/"

 for _,url := range urls {
  getSites(url)
 }
}

