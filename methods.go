package main
import(
  "fmt"
)
type Dog struct{
  Breed string
  Weight int
  Sound string
}
func(d Dog) Speak(){
  fmt.Println(d.Sound)
}
//* makes d a pointer
func (d *Dog)SpeakThreeTimes(){
    d.Sound =fmt.Sprintf("%v! %v! %v!",d.Sound,d.Sound,d.Sound)
    fmt.Println(d.Sound)
}
func main(){
  poodle := Dog{"Poodle",35, "woof"}
  poodle.Speak()

  poodle.Sound= "Arf"
  poodle.Speak()

  poodle.SpeakThreeTimes()
  poodle.SpeakThreeTimes()
}
//instance of dog
