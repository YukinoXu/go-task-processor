package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/xuexiangxu/go-task-processor/internal/cache"
	"github.com/xuexiangxu/go-task-processor/internal/config"
	"github.com/xuexiangxu/go-task-processor/internal/db"
	"github.com/xuexiangxu/go-task-processor/internal/model"
	"github.com/xuexiangxu/go-task-processor/internal/mq"
	"github.com/xuexiangxu/go-task-processor/internal/service"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	mq.InitRabbitMQ()
	cache.InitRedis()

	msgs, err := mq.Channel.Consume(
		mq.Queue.Name, // queue
		"",            // consumer tag
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register RabbitMQ consumer: %v", err)
	}

	log.Println("Worker started. Waiting for tasks...")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			go handleTask(d.Body)
		}
	}()

	<-forever
}

func handleTask(body []byte) {
	var task model.Task
	if err := json.Unmarshal(body, &task); err != nil {
		log.Printf("Invalid task format: %v\n", err)
		return
	}

	taskKey := fmt.Sprintf("task:%s", task.ID)
	if !cache.SetIfNotExist(taskKey) {
		log.Printf("Task %s already being processed or done. SKipping...\n", task.ID)
		return
	}

	log.Printf("Received task: %s (%s)\n", task.ID, task.Type)

	if err := service.UpdateTaskStatus(task.ID, model.StatusRunning); err != nil {
		log.Printf("Failed to update task to running: %v\n", err)
		return
	}

	// 模拟处理耗时任务
	time.Sleep(3 * time.Second)

	result := fmt.Sprintf("Task %s completed successfully", task.ID)
	if err := service.FinishTask(task.ID, result, model.StatusSuccess); err != nil {
		log.Printf("Failed to finish task: %v\n", err)
		return
	}

	log.Printf("Task %s done. \n", task.ID)
}
