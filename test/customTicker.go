package main

import(
    "time"
)


type CustomTicker struct{
    C chan time.Time
}

func NewCustomTicker(pause time.Duration, d time.Duration) *CustomTicker {
    C := make(chan time.Time)
    
    go func(){
        time.Sleep(pause)
        C <- time.Now()
        ticker := time.NewTicker(d)
        for{
            select{
            case <- ticker.C:
                C <- time.Now()
            case <- C:
                return
            }
        }
    }()
    
    return &CustomTicker{
        C: C,
    }
}

func (this *CustomTicker) Stop(){
    this.C <- time.Now()
}