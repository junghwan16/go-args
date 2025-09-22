package args

import "strings"

type Spec struct {
	opts map[string]Option
}

type Option struct {
	Name string
	Kind kind
}

type kind int

const (
	kindFlag kind = iota
)

type Result struct {
	values map[string]any
}

func New() *Spec {
	return &Spec{opts: make(map[string]Option)}
}

func (s *Spec) Boolean(name string) *Spec {
	s.opts[name] = Option{Name: name, Kind: kindFlag}
	return s
}

func (s *Spec) Parse(argv []string) (Result, error) {
	// 초기값 채우기
	values := make(map[string]any, len(s.opts))
	for name, opt := range s.opts {
		switch opt.Kind {
		case kindFlag:
			values[name] = false
		}
	}

	// 초기 단순한 파싱 - "-<name>" 만나면 true
	for i := 0; i < len(argv); i++ {
		tok := argv[i]
		if isFlag(tok) {
			name := strings.TrimLeft(tok, "-")
			if opt, ok := s.opts[name]; ok && opt.Kind == kindFlag {
				values[name] = true
			}
			// 우선 무시
		}
	}

	return Result{values: values}, nil
}

func isFlag(s string) bool {
	return strings.HasPrefix(s, "-") && len(s) > 1
}

func (r Result) Bool(name string) bool {
	v, _ := r.values[name].(bool)
	return v
}
