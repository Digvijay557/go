package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
type Task struct {
    Name     string
    Duration time.Duration
	Status   string
}



func runTask(ctx context.Context,task *Task, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()
	task.Status = "RUNNING"
	fmt.Println("Task " + task.Name + " is Running")
	
	select {
	case <- time.After(task.Duration):
		results <- "Task " + task.Name + " is Completed"
		task.Status = "COMPLETED"	
	case <-ctx.Done():
		// fmt.Println("Task " + name + " is Cancelled")
		task.Status = "CANCELLED"
		results <- "Task " + task.Name+ " is Cancelled"
		
	}

}

func runError(ctx context.Context,name string, duration time.Duration, wg *sync.WaitGroup, errors chan string) {
	defer wg.Done()

	fmt.Println("Task " + name + " is Running")
	select{
		case <- time.After(duration):
			errors <- "Task " + name + " has error"
			
		case <- ctx.Done():
			errors <- "Task " + name + " has delayed error"
		}
}

func main() {
	tasks := []Task{
    {"Task 1", 1 * time.Second,"PENDING"},
    {"Task 2", 3 * time.Second,"PENDING"},
    {"Task 3", 2 * time.Second,"PENDING"},
}

	var wg sync.WaitGroup

	wg.Add(len(tasks))
	fmt.Println(tasks[1].Status)
	ch := make(chan string, len(tasks))
	er := make(chan string, len(tasks))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()
	for i := range tasks {
		go runTask(ctx, &tasks[i], &wg, ch)
	}
	fmt.Println(tasks[1].Status)
	
	go func() {
		wg.Wait()
		close(ch)
		close(er)
		}()
		// timeout := time.After(2 * time.Second)
		
		for ch != nil || er != nil {
			select {
			case message, ok := <-ch:
				if !ok {
					ch = nil
					continue
				}
				fmt.Println(message)
				
			case message, ok := <-er:
				if !ok {
					er = nil
					continue
				}
				fmt.Println(message)
			}
		}
		fmt.Println(tasks[1].Status)

	
}