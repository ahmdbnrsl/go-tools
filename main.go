package main 

import "go-gabut/pointer"
import "go-gabut/console"
import "go-gabut/channel"


func main () {
    var a any = "Beni"
    pointer.C(&a, 200)
    console.Log(a)
    
    var b int = 3
    channel.Pow(&b, 4)
    console.Log(b)
    
    var d chan int = make(chan int)
    channel.Pow2(&d, 2, 8)
    console.Log(<- d)
}