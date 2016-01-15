package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type activity struct {
	name string
	r    *rand.Rand
	tmin int
	tmax int
}

func (a *activity) run(person string) {
	fmt.Println(person, "started", a.name)
	var s = a.r.Intn(a.tmax-a.tmin) + a.tmin
	time.Sleep(time.Duration(s) * time.Millisecond)
	fmt.Println(person, "spent", s, "seconds", a.name)
}

func (a *activity) runMany(persons ...string) {
	var wg = sync.WaitGroup{}
	wg.Add(len(persons))
	for _, person := range persons {
		go func(person string) {
			defer wg.Done()
			a.run(person)
		}(person)
	}
	wg.Wait()
}

func main() {
	var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var ready = activity{
		name: "getting ready",
		r:    r,
		tmin: 60,
		tmax: 90,
	}
	fmt.Println("Let's go for a walk!")
	ready.runMany("Alice", "Bob")
	fmt.Println("Arming alarm.")
}
