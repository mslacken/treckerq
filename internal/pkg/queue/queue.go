package queue

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/mslacken/treckerq/internal/pkg/tlog"
)

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
				task := &Task{}
				task.State = Q

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
