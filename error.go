package main
import(
  "fmt"
  "os"
  "errors"
)
func main(){
  f,err := os.Open("filename.ext")

  if err==nil{
    fmt.Println(f)
  }else{
    fmt.Println(err.Error())
  }
  //custom errors
  myError := errors.New("My error string")
  fmt.Println(myError)

  attendance := map[string]bool{
    "Ann":true,
    "Mike":true}

  attended, ok := attendance["mike"]
  if ok{
    fmt.Println(attended)
  }else{
    fmt.Println("no info for mike")
  }
}
