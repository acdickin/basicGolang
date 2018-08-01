package main
import(
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
  "math/big"
)

func main(){
  url:= "http://services.explorecalifornia.org/json/tours.php"
  // resp, err := http.Get(url)
  // if err!=nil{
  //   panic(err)
  // }
  // fmt.Printf("Resp type: %T\n",resp)
  content:= contentFromServer(url)
  tours:= toursFromJson(content)
  // fmt.Println(tours)
  for _,tour := range tours{
    price,_,_:= big.ParseFloat(tour.Price, 10,2,big.ToZero)
    fmt.Printf("%v ($%.2f)\n",tour.Name, price)
  }
}
type Tour struct{
  //Uppercase names from api to export the field for the rest of the application
  Name, Price string
}

func checkError(err error){
  if err != nil{
    panic(err)
  }
}
func contentFromServer(url string) string{
  resp, err := http.Get(url)
  checkError(err)
  defer resp.Body.Close()
  bytes , err := ioutil.ReadAll(resp.Body)
  checkError(err)
  return string(bytes)
}
func toursFromJson(content string) []Tour{
  tours := make([]Tour,0,20)

  decoder:= json.NewDecoder(strings.NewReader(content))
  //removes unwanted bracket around data
  //first object is the token
  _,err:= decoder.Token()
  checkError(err)

  var tour Tour
  //parses the data and gets the values for struct
  //checks if there is more data and moves to that next data
  for decoder.More(){
      err:=decoder.Decode(&tour)
      checkError(err)
      tours=append(tours,tour)
  }
  return tours

}
