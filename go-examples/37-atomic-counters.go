package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var opsAtomic, ops uint64
    var wg sync.WaitGroup
    fmt.Println(`Lancement de 50 goroutines qui parcourt
une boucle de 1 à 1000 incrementant
la variable ops et opsAtomic`)
    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {
                atomic.AddUint64(&opsAtomic, 1)
                ops++
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("Compteur atomic ops:", opsAtomic)
    fmt.Println("Compteur non atomic ops:", ops)
}
