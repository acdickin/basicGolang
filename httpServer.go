package main
import(
  "fmt"
  "net/http"
)
type Hello struct{}

func(h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w,"<h1>Hello from go web server!</h1>")

}

func main(){
  var h Hello
  fmt.Println("Running on localhost:4000")
  err:= http.ListenAndServe("localhost:4000", h)
  checkError(err)
}
func checkError(err error){
  if err != nil{
    panic(err)
  }
}
