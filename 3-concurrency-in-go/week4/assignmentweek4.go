package main

import (
    "fmt"
    "sync"
    "time"
)

const numberOfPhilosophers = 5
const eatingTimes = 3

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
    id                int
    leftChopstick, rightChopstick *Chopstick
    eatingCount       int
}

func (p *Philosopher) eat(wg *sync.WaitGroup, host chan bool) {
    defer wg.Done()
    for p.eatingCount < eatingTimes {
        <-host
        p.leftChopstick.Lock()
        p.rightChopstick.Lock()

        fmt.Printf("starting to eat %d\n", p.id)
        time.Sleep(time.Second) // Simulate eating
        fmt.Printf("finishing eating %d\n", p.id)

        p.rightChopstick.Unlock()
        p.leftChopstick.Unlock()

        p.eatingCount++
        host <- true
    }
}

func main() {
    var wg sync.WaitGroup
    chopsticks := make([]*Chopstick, numberOfPhilosophers)
    for i := 0; i < numberOfPhilosophers; i++ {
        chopsticks[i] = new(Chopstick)
    }

    philosophers := make([]*Philosopher, numberOfPhilosophers)
    for i := 0; i < numberOfPhilosophers; i++ {
        philosophers[i] = &Philosopher{
            id:              i + 1,
            leftChopstick:   chopsticks[i],
            rightChopstick:  chopsticks[(i+1)%numberOfPhilosophers],
            eatingCount:     0,
        }
    }

    host := make(chan bool, 2) // Host allows up to 2 philosophers to eat concurrently
    for i := 0; i < 2; i++ {
        host <- true
    }

    wg.Add(numberOfPhilosophers)
    for _, p := range philosophers {
        go p.eat(&wg, host)
    }

    wg.Wait()
}
