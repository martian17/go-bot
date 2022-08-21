package main

import (
    "time"
)


func main(){
    println("tick")
    ticker := NewCustomTicker(time.Duration(1000)*time.Millisecond,time.Duration(500)*time.Millisecond)
    for{
        select{
        case <- ticker.C:
            println("tick")
        }
    }
}