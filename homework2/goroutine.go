package main

import (
	"context"
	"fmt"
	"time"
)

func print1() {
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
	}
}

func print2() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
	}
}

// GoroutineHomeWork1 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func GoroutineHomeWork1() {
	go print1()
	go print2()
}

// GoroutineHomeWork2 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func GoroutineHomeWork2() {
	tasks := []Task{
		func(ctx context.Context) error {
			time.Sleep(100 * time.Millisecond)
			return nil
		},
		func(ctx context.Context) error {
			time.Sleep(200 * time.Millisecond)
			return fmt.Errorf("任务1失败")
		},
		func(ctx context.Context) error {
			select {
			case <-time.After(500 * time.Millisecond):
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
		func(ctx context.Context) error {
			time.Sleep(50 * time.Millisecond)
			return nil
		},
	}

	scheduler := NewScheduler(2, 300*time.Millisecond)
	results := scheduler.Run(tasks)

	fmt.Println("任务结果：")
	for _, r := range results {
		status := "成功"
		if r.Error != nil {
			status = "失败 " + r.Error.Error()
		}
		fmt.Printf("任务[%d]: 耗时=%v, 状态=%s\n", r.ID, r.Duration, status)
	}

}
