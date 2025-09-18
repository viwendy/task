package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 代表一个可执行的任务
type Task struct {
	ID       string
	Function func() error
}

// TaskResult 代表任务执行结果
type TaskResult struct {
	TaskID      string
	Error       error
	ExecuteTime time.Duration
}

// TaskScheduler 任务调度器结构体
type TaskScheduler struct {
	tasks   []Task
	results chan TaskResult
	wg      sync.WaitGroup
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make([]Task, 0),
		results: make(chan TaskResult, 100), // 缓冲通道
	}
}

// AddTask 添加任务到调度器
func (ts *TaskScheduler) AddTask(id string, fn func() error) {
	task := Task{
		ID:       id,
		Function: fn,
	}
	ts.tasks = append(ts.tasks, task)
}

// ExecuteAll 并发执行所有任务
func (ts *TaskScheduler) ExecuteAll() []TaskResult {
	// 启动所有任务
	for _, task := range ts.tasks {
		ts.wg.Add(1)
		go ts.executeTask(task)
	}
	// 启动 goroutine 等待所有任务完成
	go func() {
		ts.wg.Wait()
		close(ts.results)
	}()
	// 收集结果
	var results []TaskResult
	for result := range ts.results {
		results = append(results, result)
	}
	return results
}

// executeTask 执行单个任务并记录执行时间
func (ts *TaskScheduler) executeTask(task Task) {
	defer ts.wg.Done()
	startTime := time.Now()
	// 执行任务函数
	err := task.Function()
	// 计算执行时间
	execTime := time.Since(startTime)
	// 发送结果到通道
	ts.results <- TaskResult{
		TaskID:      task.ID,
		Error:       err,
		ExecuteTime: execTime,
	}
}

// 示例任务函数
func sampleTask(name string, duration time.Duration) func() error {
	return func() error {
		fmt.Printf("Task %s started\n", name)
		time.Sleep(duration)
		fmt.Printf("Task %s completed\n", name)
		return nil
	}
}

func main() {
	// 创建任务调度器
	scheduler := NewTaskScheduler()
	// 添加示例任务
	scheduler.AddTask("task1", sampleTask("Task1", 2*time.Second))
	scheduler.AddTask("task2", sampleTask("Task2", 1*time.Second))
	scheduler.AddTask("task3", sampleTask("Task3", 3*time.Second))
	scheduler.AddTask("task4", func() error {
		time.Sleep(500 * time.Millisecond)
		return fmt.Errorf("task4 encountered an error")
	})
	// 执行所有任务并获取结果
	fmt.Println("Starting task execution...")
	startTime := time.Now()
	results := scheduler.ExecuteAll()
	totalTime := time.Since(startTime)
	// 输出结果
	fmt.Println("\n=== Task Execution Results ===")
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("Task ID: %s | Execution Time: %v | Error: %v\n",
				result.TaskID, result.ExecuteTime, result.Error)
		} else {
			fmt.Printf("Task ID: %s | Execution Time: %v | Status: Success\n",
				result.TaskID, result.ExecuteTime)
		}
	}
	fmt.Printf("\nTotal execution time: %v\n", totalTime)
}
