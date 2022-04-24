package main
import (
   "io/ioutil"
   "log"
   "net/http"
)

func main() {

	resp, err := http.Get("https://api.github.com/search/users?q=alex+schapelle")
	if err != nil {
   		log.Fatalln(err)
	}
body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   log.Printf(sb)
}
