package main

import (
    "fmt"
    "sync"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
    leftCS, rightCS *Chopstick
}

func (p Philosopher) eat(wg *sync.WaitGroup, id int) {
    defer wg.Done()
    for i := 0; i < 3; i++ { // Each philosopher eats 3 times
        p.leftCS.Lock()
        p.rightCS.Lock()

        fmt.Printf("Philosopher %d is eating\n", id)
        
        p.rightCS.Unlock()
        p.leftCS.Unlock()
    }
}

func main() {
    var wg sync.WaitGroup

    CSticks := make([]*Chopstick, 5)
    for i := 0; i < 5; i++ {
        CSticks[i] = new(Chopstick)
    }

    philosophers := make([]*Philosopher, 5)
    for i := 0; i < 5; i++ {
        philosophers[i] = &Philosopher{CSticks[i], CSticks[(i+1)%5]}
    }

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go philosophers[i].eat(&wg, i)
    }

    wg.Wait()
}
