package paramlang

import (
	"strconv"

	"golang.org/x/xerrors"
)

// Convert functions are just helper functions to convert a string to
// it's proper typed value.

func toString(s string) (string, error) {
	val, err := strconv.Unquote(s)
	if err != nil {
		return "", xerrors.Errorf("to string: %w", err)
	}
	return val, nil
}

func toFloat(s string) (float64, error) {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1, xerrors.Errorf("to float: %w", err)
	}
	return val, nil
}

func toInteger(s string) (int, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1, xerrors.Errorf("to int: %w", err)
	}
	return int(val), nil
}

func toBool(s string) (bool, error) {
	val, err := strconv.ParseBool(s)
	if err != nil {
		return false, xerrors.Errorf("to bool: %w", err)
	}
	return val, nil
}
