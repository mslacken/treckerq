package queue

import "time"

type TreckerQueue struct {
	Format string
	Exec   string
	Cnt    int `comment:"Total count of job tasks" key:"cnt"`
	Twk    int `comment:"Job tasks being worked on (marked as running)" key:"twk"`
	Tdn    int `comment:"Successfully finished job tasks." key:"tdn"`
	Tfl    int `comment:"Failed job tasks" key:"tfl"`
	Tsk    int `comment:"Skipped job task" key:"tsk"`
	Nwk    int `comment:"The number of workers workers for this queue." key:"nwk"`
	Mxt    int `comment:"Maximum recorded task execution time in second" key:"mxt"`
	Mnt    int `comment:"Minimum recorded task execution time in seconds" key:"mnt"`
	Smc    int `comment:"Job task count that went into the execution time sum" key:"smc"`
	Smt    int `comment:"Sum of all execution times in seconds" key:"smt"`
	Nxt    int `comment:"Whole-file byte offset of next entry to work on" key:"nxt"`
	Tasks  []Task
}

type Task struct {
	State JState
	Id    int
	Cmd   int
	Crc   int
	Ret   int16
}

type JState int

const (
	Q       JState = 0
	W       JState = 1
	D       JState = 2
	F       JState = 3
	S       JState = 4
	H       JState = 5
	queued  JState = 0
	running JState = 1
	done    JState = 2
	failed  JState = 3
	skipped JState = 4
	held    JState = 5
)

/*
struct for holding event records
*/
type Log struct {
	Event  LEvent
	Id     int
	Time   time.Time
	Worker int
	Exit   int
}

type LEvent int

const (
	B LEvent = 0
	E LEvent = 1
	R LEvent = 2
	A LEvent = 3
)
