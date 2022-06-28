package queue

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

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
			switch line[0] {
			case "cnt":
				queue.cnt, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("cnt has wrong format %s\n", line[1])
				}
			case "twk":
				queue.twk, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("twk has wrong format %s\n", line[1])
				}
			case "tfl":
				queue.tfl, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("tfl has wrong format %s\n", line[1])
				}
			case "tsk":
				queue.tsk, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("tsk has wrong format %s\n", line[1])
				}
			case "nwk":
				queue.nwk, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("nwk has wrong format %s\n", line[1])
				}
			case "mxt":
				queue.mxt, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("mxt has wrong format %s\n", line[1])
				}
			case "mnt":
				queue.mnt, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("mnt has wrong format %s\n", line[1])
				}
			case "smc":
				queue.smc, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("smc has wrong format %s\n", line[1])
				}
			case "nxt":
				queue.nxt, err = strconv.Atoi(line[1])
				if err != nil {
					tlog.Warn("nxt has wrong format %s\n", line[1])
				}
			case "Q":
			}

		} else {
			tlog.Warn("Found garbage in queue %s:%u %s\n", fileName, i, line)
		}
	}
	return queue, nil
}
