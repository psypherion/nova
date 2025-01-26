package task

import "time"

type Task struct {
	ID        		int				`json:"id"`
	Name      		string			`json:"name"`
	CreatedAt 		time.Time		`json:"created_at"`
	StartTime 		time.Time		`json:"start_time"`
	Duration  		time.Duration 	`json:"duration"`
	Enjoyability 	int				`json:"enjoyability"`
	Completed 		bool			`json:"completed"`
	Exhaustion 		int				`json:"exhaustion"`
}

func NewTask(name string, starttime int, duration int) Task {
	return Task{
		Name: name,
		StartTime: time.Unix(int64(starttime), 0),
		Duration: time.Duration(duration),
		CreatedAt: time.Now(),
		Enjoyability: 5,
		Completed: false,
		Exhaustion: 5,
	}
}
