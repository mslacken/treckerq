package queue

type TreckerQueue struct {
	format string
	exec   string
	cnt    int `comment:"Total count of job tasks" key:"cnt"`
	twk    int
	tfl    int
	tsk    int
	nwk    int
	mxt    int
	mnt    int
	smc    int
	smt    int
	nxt    int
	tasks  []Task
}

type Task struct {
	state State
	cmd   int
	crc   int
	ret   int16
}

type State int

const (
	Q       State = 0
	W             = 1
	D             = 2
	F             = 3
	S             = 4
	H             = 5
	queued        = 0
	running       = 1
	done          = 2
	failed        = 3
	skipped       = 4
	held          = 5
)
