// simple program to demonstrate use of buffered channel

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var goRoutine sync.WaitGroup

func main() {
	rand.Seed(time.Now().Unix())

	// Create a beffered channel to manage the employee vs project load.
	projects := make(chan string, 10)
	for i := 1; i <= 5; i++ {
		go employee(projects, i)
	}

	for j := 1; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}

	// close the channel so the goroutines will quit
	close(projects)
	goRoutine.Wait()
}

func employee(projects chan string, employee int) {
	defer goRoutine.Done()
	for {
		// Wait for project to be assigned.
		project, result := <-projects
		if result == false {
			// this means the channel is empty and closed.
			fmt.Printf("employee : %d : Exit\n")
			return
		}

		fmt.Printf("employee : %d : Started %s\n", employee, project)

		// randomly wait to simulate work time.
		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// display time to wait
		fmt.Println("\nTime to sleep:", sleep, "ms/n")

		// display project completed by employee
		fmt.Printf("Employee : %d : Completed %s\n", employee, project)
	}
}
