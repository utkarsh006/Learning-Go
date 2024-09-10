package main
import "fmt"

var a = 10  //global variable

func main() {
  fmt.Println("Hello Dude")
  var age = 20   //local variable
  fmt.Println(age)

  var b int      
  b = 1
  fmt.Println(b)

  c := 10 //declaring using short hand property
  fmt.Println(c)  

  var a2, b2, c2 = "10",100.2,1000
  fmt.Printf("Type of a2 %T, b2 %T, c2 %T\n", a2, b2, c2)

  // int, int8, int16, int32, int64
  // unassigned(cant hold negative values) : uint, uint8, uint16, uint32, uint64
  // float32 float 64
  // byte
  // string
  // bool
  // complex64, complex128

}
