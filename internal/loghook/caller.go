package loghook

import (
	"runtime"
	"strings"
)

func findCaller(skip int) (string, string, int) {
	var (
		pc       uintptr
		file     string
		function string
		line     int
	)
	for i := 0; i < 10; i++ {
		pc, file, line = getCaller(skip + i)
		if !(strings.HasPrefix(file, "logrus") ||
			strings.HasPrefix(file, "config") ||
			strings.HasPrefix(file, "glog")) {
			break
		}
	}
	if pc != 0 {
		frames := runtime.CallersFrames([]uintptr{pc})
		frame, _ := frames.Next()
		function = frame.Function
	}

	return file, function, line
}

func getCaller(skip int) (uintptr, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return 0, "", 0
	}

	// n := 0

	dirs := strings.Split(file, "/")
	dirs = dirs[len(dirs)-2:]
	file = strings.Join(dirs, "/")
	// file = fmt.Sprintf("%v/%v", dirs[len(dirs)-1], dirs[len(dirs)-2])

	/*
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				n++
				if n >= 2 {
					file = file[i+1:]
					break
				}
			}
		}
	*/

	return pc, file, line
}
