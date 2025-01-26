package main

import {
	"fmt"
	"time"
	"github.com/yourusername/nova/internal/task"
}

func main() {
	newTask := task.NewTask("Do the dishes", 0, 30)
	fmt.Println(newTask)
}
