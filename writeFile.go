package main
import(
  "fmt"
  "io"
  "io/ioutil"
  "os"
)
func main(){
  content := "This text was created by writeFile.go!!"

  file, err := os.Create("./fromString.txt")
  checkError(err)
  defer file.Close()
  length, err := io.WriteString(file,content)
  checkError(err)
  fmt.Printf("All done with file of %v characters", length)

  bytes := []byte(content)
  ioutil.WriteFile("./fromBytes.txt",bytes, 0644)
}
func checkError(err error){
  if err!=nil{
    panic(err)
  }
}
