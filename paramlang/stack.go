package paramlang

import (
	"reflect"

	"golang.org/x/xerrors"
)

type stackItemType string

const (
	stUnknown stackItemType = "unknown"
	// stDefer defers the type until it is used'
	stDefer  stackItemType = "defer"
	stString stackItemType = "string"
	stBool   stackItemType = "bool"
	stInt    stackItemType = "int"
	stFloat  stackItemType = "float"

	// stNoTemplate indicates the string is not using param lang.
	// Treat these as literals
	stNoTemplate stackItemType = "no-template"
)

// stackItem is an item on the stack.
type stackItem struct {
	Value interface{}
	Type  stackItemType
}

func (s stackItem) AsString() (string, error) {
	v, err := s.As(stString)
	if err != nil {
		return "", err
	}

	vi, ok := v.(string)
	if !ok {
		return "", xerrors.Errorf("can't cast to bool")
	}
	return vi, nil
}

func (s stackItem) AsBool() (bool, error) {
	v, err := s.As(stBool)
	if err != nil {
		return false, err
	}

	vi, ok := v.(bool)
	if !ok {
		return false, xerrors.Errorf("can't cast to bool")
	}
	return vi, nil
}

func (s stackItem) AsFloat() (float64, error) {
	v, err := s.As(stFloat)
	if err != nil {
		return -1, err
	}

	vi, ok := v.(float64)
	if !ok {
		return -1, xerrors.Errorf("can't cast to int")
	}
	return vi, nil
}

func (s stackItem) AsInt() (int, error) {
	v, err := s.As(stInt)
	if err != nil {
		return -1, err
	}

	vi, ok := v.(int)
	if !ok {
		return -1, xerrors.Errorf("can't cast to int")
	}
	return vi, nil
}

func (s stackItem) As(t stackItemType) (interface{}, error) {
	if t != s.Type && !(s.Type == stDefer || s.Type == stNoTemplate) {
		return nil, xerrors.Errorf("type expected %s, got %s", s.Type, t)
	}

	if t == s.Type {
		// No conversion needed
		return s.Value, nil
	}

	if s.Type == stDefer || s.Type == stNoTemplate {
		str, ok := s.Value.(string)
		if !ok {
			return nil, xerrors.Errorf("deferred type not of type string, got %s", reflect.TypeOf(s.Value).String())
		}
		switch t {
		case stBool:
			return toBool(str)
		case stInt:
			return toInteger(str)
		case stFloat:
			return toFloat(str)
		case stString:
			if s.Type == stDefer {
				return toString(str)
			}
			// Strings from 'stNoTemplate' are already escaped
			return str, nil
		}
	}

	return nil, xerrors.Errorf("unable to get value as %s", t)
}

func (l *wacListener) pushBool(b bool) {
	l.push(b, stBool)
}

func (l *wacListener) pushString(s string) {
	l.push(s, stString)
}

func (l *wacListener) pushInt(i int) {
	l.push(i, stInt)
}

func (l *wacListener) pushFloat(f float64) {
	l.push(f, stFloat)
}

func (l *wacListener) pushDefer(s string) {
	l.push(s, stDefer)
}

func (l *wacListener) pushStackItem(s *stackItem) {
	l.stack = append(l.stack, s)
}

func (l *wacListener) push(i interface{}, ty stackItemType) {
	l.stack = append(l.stack, &stackItem{
		Value: i,
		Type:  ty,
	})
}

func (l *wacListener) popBool() bool {
	item := l.pop()
	v, err := item.As(stBool)
	if err != nil {
		l.stopParsing(xerrors.Errorf("exp bool"))
	}
	return v.(bool)
}

//nolint:unused
func (l *wacListener) popInt() int {
	item := l.pop()
	v, err := item.As(stInt)
	if err != nil {
		l.stopParsing(xerrors.Errorf("exp int"))
	}
	return v.(int)
}

//nolint:unused
func (l *wacListener) popString() string {
	item := l.pop()
	v, err := item.As(stString)
	if err != nil {
		l.stopParsing(xerrors.Errorf("exp string"))
	}
	return v.(string)
}

func (l *wacListener) pop() *stackItem {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}
