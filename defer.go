package main
import(
  "fmt"
)
func main(){
  defer fmt.Println("Close the File!")
  fmt.Println("Open the File!")
  defer fmt.Println("Statement 1")
  defer fmt.Println("Statement 2")
  myFunc()
  defer fmt.Println("Statement 3")
  defer fmt.Println("Statement 4")
  fmt.Println("undeffered statement")
  fmt.Println("Defer runs frist in last out")
}

func myFunc(){
  defer fmt.Println("Deffered statement")
  fmt.Println("undeffered statement")

}
