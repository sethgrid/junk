package stringpipe

import (
	"sort"
	"strings"
)

type StringPipe string

func (s StringPipe) Grep(substr string) StringPipe {
	var returnLines []string
	lines := strings.Split(string(s), "\n")
	for _, line := range lines {
		if strings.Contains(line, substr) {
			returnLines = append(returnLines, line)
		}
	}
	return StringPipe(strings.Join(returnLines, "\n"))
}

func (s StringPipe) VGrep(substr string) StringPipe {
	var returnLines []string
	lines := strings.Split(string(s), "\n")
	for _, line := range lines {
		if !strings.Contains(line, substr) {
			returnLines = append(returnLines, line)
		}
	}
	return StringPipe(strings.Join(returnLines, "\n"))
}

func (s StringPipe) Sort() StringPipe {
	lines := strings.Split(string(s), "\n")
	sort.Strings(lines)
	return StringPipe(strings.Join(lines, "\n"))
}

func (s StringPipe) Head(n int) StringPipe {
	var returnLines []string
	lines := strings.Split(string(s), "\n")
	for i, line := range lines {
		if i < n {
			returnLines = append(returnLines, line)
		} else {
			break
		}
	}
	return StringPipe(strings.Join(returnLines, "\n"))
}

func (s StringPipe) Tail(n int) StringPipe {
	var returnLines []string
	lines := strings.Split(string(s), "\n")
	for i := len(lines) - n; i < len(lines); i++ {
		if i < 0 {
			i = 0
			continue
		}
		returnLines = append(returnLines, lines[i])
	}
	return StringPipe(strings.Join(returnLines, "\n"))
}

func (s StringPipe) CutOutColumns(cols ...int) StringPipe {
	lines := strings.Split(string(s), "\n")
	for i, line := range lines {
		var keepParts []string
		parts := strings.Split(line, " ")
		for j, p := range parts {
			if !in(j, cols) {
				keepParts = append(keepParts, p)
			}
		}
		lines[i] = strings.Join(keepParts, " ")
	}
	return StringPipe(strings.Join(lines, "\n"))
}

func (s StringPipe) String() string {
	return string(s)
}

func in(i int, elements []int) bool {
	for _, e := range elements {
		if i == e {
			return true
		}
	}
	return false
}
