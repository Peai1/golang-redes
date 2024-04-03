package main

import (
	"fmt"
	"time"
	"sync"
)

// go run .\ejemplo.go

func tarea(id int, duracion time.Duration, wg *sync.WaitGroup){
	defer wg.Done() // se ejecuta al finalizar la funcion
	fmt.Printf("Tarea %d comenzando\n", id)
	time.Sleep(duracion)
	fmt.Printf("Tarea %d completada\n", id)
}

func main(){
	var wg sync.WaitGroup
	numeroDeTareas := 5
	tareasEnParalelo := 3
	duraciones := []time.Duration{1*time.Second, 2*time.Second, 3*time.Second, 2*time.Second, 1*time.Second}

	for i := 0; i < numeroDeTareas; i++ {
		wg.Add(1)
		go tarea(i+1, duraciones[i % len(duraciones)], &wg)
		if (i+1) % tareasEnParalelo == 0 {
			wg.Wait()
		}
	}
	wg.Wait()
	fmt.Println("Tareas completadas")

}