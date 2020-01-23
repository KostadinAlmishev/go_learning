package main

import "fmt"

var var1 = 1 // OK
// var2 := 2 // ERROR

func main() {
    // MAIN TYPES
    
    // string
    // bool
    // int int8 int16 int32 int64
    // uint uint8 uint16 uint32 uint64
    // byte - alias for uint8
    // rune - alias for uint32
    // float32 float64
    // complex64 complex128


    // Using var
    var name = "Brad"
    var age = 32
    const isCool = true
    // isCool = false // ERROR
    var3 := 15.1

    // Shorthand
    name1 := "Brad"
    email := "brad@mail.ru"

    name2, email1 := "Brad", "brad@mail.ru"


    fmt.Println(name1, email)
    fmt.Println(name2, email1)

    fmt.Println(name, age, isCool, var3)
    fmt.Printf("%T\n", isCool)
}
