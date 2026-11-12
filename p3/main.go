package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Name     string
	Duration time.Duration
	Status   string
}

type taskResult struct {
	TaskID int
	Status string
	Error  error
}
type taskManager struct {
	Tasks []Task
	mu sync.Mutex
}

func (tm *taskManager) AddTask (task Task){
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.Tasks = append(tm.Tasks, task)
}

const (
	Pending   = "PENDING"
	Running   = "RUNNING"
	Completed = "COMPLETED"
	Cancelled = "CANCELLED"
	Failed    = "FAILED"
)

func runTask(ctx context.Context, task *Task, wg *sync.WaitGroup, results chan taskResult, 	mu sync.Mutex) {
	mu.Lock()
	defer wg.Done()
	defer mu.Unlock()

	task.Status = Running
	fmt.Println("Task", task.Name, "is Running")

	select {
	case <-time.After(task.Duration):
		task.Status = Completed
		results <- taskResult{TaskID: task.ID, Status: task.Status}
	case <-ctx.Done():
		task.Status = Cancelled
		results <- taskResult{TaskID: task.ID, Status: task.Status, Error: ctx.Err()}
	}
}

func (tm *taskManager) RunTasks(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(len(tm.Tasks))

	results := make(chan taskResult, len(tm.Tasks))

	for i := range tm.Tasks{
		go runTask(ctx, &tm.Tasks[i], &wg, results,tm.mu)
	}
	go func(){
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(
			"Task ID:", result.TaskID,
			"| Status:", result.Status,
		)
	}



}

func main() {
	manager := taskManager{}

	manager.AddTask(Task{
	ID:       1,
	Name:     "Task 1",
	Duration: 1 * time.Second,
	Status:   "PENDING",
})

manager.AddTask(Task{
	ID:       2,
	Name:     "Task 2",
	Duration: 3 * time.Second,
	Status:   "PENDING",
})

manager.AddTask(Task{
	ID:       3,
	Name:     "Task 3",
	Duration: 2 * time.Second,
	Status:   "PENDING",
})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	manager.RunTasks(ctx)

	
}