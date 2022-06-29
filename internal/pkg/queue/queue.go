package queue

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/mslacken/treckerq/internal/pkg/tlog"
)

/*
Read in a queue file and return a queue object.
*/
func OpenQueueFile(fileName string) (TreckerQueue, error) {
	var queue TreckerQueue
	f, err := os.Open(fileName)
	if err != nil {
		return queue, fmt.Errorf("Could not open queue file", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		// parse header separately
		if i == 0 {
			if len(line) != 2 {
				return queue, fmt.Errorf("wrong header format", err)
			}
			queue.Format = line[0]
			queue.Exec = line[1]
		} else if len(line) >= 2 {
			switch line[0] {
			case "Q":
				task, err := readTask(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				task.State = Q
				queue.Tasks = append(queue.Tasks, task)
			case "W":
				task, err := readTask(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				task.State = W
				queue.Tasks = append(queue.Tasks, task)
			case "D":
				task, err := readTask(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				task.State = D
				queue.Tasks = append(queue.Tasks, task)
			case "F":
				task, err := readTask(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				task.State = F
				queue.Tasks = append(queue.Tasks, task)
			case "S":
				task, err := readTask(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				task.State = S
				queue.Tasks = append(queue.Tasks, task)
			case "H":
				task, err := readTask(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				task.State = H
				queue.Tasks = append(queue.Tasks, task)
			case "B":
				event, err := readEvent(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				event.Event = B
				queue.Events = append(queue.Events, event)
			case "E":
				event, err := readEvent(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				event.Event = E
				queue.Events = append(queue.Events, event)
			case "R":
				event, err := readEvent(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				event.Event = R
				queue.Events = append(queue.Events, event)
			case "A":
				event, err := readEvent(line)
				if err != nil {
					return queue, fmt.Errorf("failed to parse line %d: %w", i, err)
				}
				event.Event = A
				queue.Events = append(queue.Events, event)

			default:
				dType := reflect.TypeOf(&queue)
				dhVal := reflect.ValueOf(&queue)
				for j := 0; j < dType.Elem().NumField(); j++ {
					field := dType.Elem().Field(j)
					key := field.Tag.Get("key")
					if line[0] == key {
						kind := field.Type.Kind()
						result := dhVal.Elem().Field(j)
						if kind == reflect.String {
							result.SetString(line[1])
						} else if kind == reflect.Int {
							val, err := strconv.Atoi(line[1])
							if err != nil {
								return queue, fmt.Errorf("conversion error", err)
							}
							result.SetInt(int64(val))
						} else {
							return queue, errors.New("only supports string and int")
						}
					}
				}
			}
		} else {
			tlog.Warn("Found garbage in queue file %s:%u %s\n", fileName, i, line)
		}
		i++
	}
	return queue, nil
}

func readTask(line []string) (task Task, err error) {
	if len(line) != 5 {
		return task, fmt.Errorf("error parsing line %v\n", line)
	}
	task.Id, err = strconv.Atoi(line[1])
	task.Cmd, err = strconv.Atoi(line[2])
	task.Crc, err = strconv.Atoi(line[3])
	ret, err := strconv.Atoi(line[4])
	task.Ret = int16(ret)
	return task, err
}

func readEvent(line []string) (event Log, err error) {
	if len(line) < 4 {
		return event, fmt.Errorf("error parsing line %v", line)
	}
	event.Id, err = strconv.Atoi(line[1])
	iTime, err := strconv.ParseInt(line[2], 10, 64)
	if err != nil {
		return event, fmt.Errorf("error pasring unix time %w", err)
	}
	event.Time = time.Unix(iTime, 0)
	event.Worker, err = strconv.Atoi(line[3])
	if len(line) > 4 {
		event.Exit, err = strconv.Atoi(line[4])
	}
	return event, err
}
