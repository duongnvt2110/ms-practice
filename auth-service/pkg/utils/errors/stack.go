package errors

import (
	"fmt"
	"path"
	"runtime"
)

type Stack []uintptr

func StackTrace(skip int) *Stack {
	s := make(Stack, 64)
	n := runtime.Callers(skip+2, s)
	s = s[:n]
	return &s
}

func (s *Stack) Frames() (*Frames, bool) {
	if s == nil {
		return nil, false
	}
	pc := *s
	if len(pc) == 0 {
		return nil, false
	}
	return CallersFrames(pc), true
}

func (s *Stack) Format(w fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		frames, ok := s.Frames()
		if !ok {
			return
		}

		var format string
		if w.Flag('+') {
			format = "%+" + string(verb) + "\n"
		} else {
			format = "%" + string(verb) + "\n"
		}

		for {
			f, more := frames.Next()
			fmt.Fprintf(w, format, f)

			if !more {
				return
			}
		}
	}
}

func (s *Stack) String() string {
	return fmt.Sprintf("%s", s)
}

type Frames struct {
	Frames *runtime.Frames
}

func CallersFrames(callers []uintptr) *Frames {
	return &Frames{runtime.CallersFrames(callers)}
}

func (ci *Frames) Next() (Frame, bool) {
	f, more := ci.Frames.Next()
	return Frame{f}, more
}

type Frame struct {
	runtime.Frame
}

func (f Frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s\n\t%s:%d", f.Function, f.File, f.Line)
		} else {
			fmt.Fprintf(s, "%s(%s/%s:%d)", path.Base(f.Function), path.Base(path.Dir(f.File)), path.Base(f.File), f.Line)
		}
	}
}

func (f Frame) String() string {
	return fmt.Sprintf("%s", f)
}
