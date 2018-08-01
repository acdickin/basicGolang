package main
import(
  "fmt"
)
type Dog struct{
  Breed string
  Weight int
}
func main(){
  poodle := Dog{"Poodle",35}
  fmt.Println(poodle)
  fmt.Printf("%+v\n", poodle)
  fmt.Printf("Breed: %v\nweight: %v",poodle.Breed,poodle.Weight)
}
