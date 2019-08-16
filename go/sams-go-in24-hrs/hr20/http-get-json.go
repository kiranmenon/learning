package main
import (
 "fmt"
 "log"
 "net/http"
 "encoding/json"
)


func main(){
 type githubJsonApi struct{
  Name string `json:"login"`
	Avatar string `json:"avatar_url"`
 }

 var user githubJsonApi
 res,err := http.Get("https://api.github.com/users/kiranmenon")
 if err !=nil{
  log.Fatal(err)
 }
 defer res.Body.Close()

 err = json.NewDecoder(res.Body).Decode(&user)
 if err !=nil{
  log.Fatal(err)
 }
 fmt.Printf("%+v\n",user)
}

