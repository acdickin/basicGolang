package main
import(
  "fmt"
  "io/ioutil"
)
func main(){
 fileName := "./fromBytes.txt"
 content,err := ioutil.ReadFile(fileName)
 checkError(err)

 result := string(content)

 fmt.Println("Read from file:", content)
 fmt.Println("Read Text:", result)
}
func checkError(err error){
  if err!=nil{
    panic(err)
  }
}
