package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var reLine = regexp.MustCompile(`^(\S+) (\d+)$`)

func AggregateLogsAI(f *os.File) map[string]int64 {
	totals := make(map[string]int64)
	sc := bufio.NewScanner(f) // looks for the file
	for sc.Scan() {
		m := reLine.FindStringSubmatch(sc.Text()) // for each line run a compile regular exp and splits the string
		if len(m) == 3 {
			n, _ := strconv.Atoi(m[2])       // parses integers with 'Atoi'
			totals[string(m[1])] += int64(n) // hashes into a map. The GC runs after to clean up
		}
	}
}
