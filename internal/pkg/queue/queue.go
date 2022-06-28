package queue

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"

	"github.com/mslacken/treckerq/internal/pkg/tlog"
)

func OpenQueueFile(fileName string) (TreckerQueue, error) {
	var queue TreckerQueue
	f, err := os.Open(fileName)
	if err != nil {
		return queue, fmt.Errorf("Could not open queue file", err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.Comma = ' '
	data, err := reader.ReadAll()
	if err != nil {
		return queue, fmt.Errorf("Error in parsing queue file", err)
	}
	for i, line := range data {
		// parse header seperately
		if i == 0 {
			if len(line) != 2 {
				return queue, fmt.Errorf("wrong header format", err)
			}
			queue.format = line[0]
			queue.exec = line[1]
		} else if len(line) >= 2 {
			dType := reflect.TypeOf(queue)
			dhVal := reflect.ValueOf(queue)
			for j := 0; j < dType.Elem().NumField(); j++ {
				field := dType.Elem().Field(j)
				key := field.Tag.Get("mapper")

				kind := field.Type.Kind()
			}
		} else {
			tlog.Warn("Found garbage in queue %s:%u %s\n", fileName, i, line)
		}
	}
	return queue, nil
}
