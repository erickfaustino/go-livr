package livr

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
)

// Eq - make sure that validated value is equal to some value.
func Eq(args ...interface{}) func(...interface{}) (interface{}, interface{}) {
	var allowed interface{}
	if len(args) > 0 {
		allowed = args[0]
	}

	return func(builders ...interface{}) (interface{}, interface{}) {
		var value interface{}
		if len(builders) > 0 {
			value = builders[0]
		}

		if value == nil || value == "" {
			return nil, nil
		}

		switch value.(type) {
		case float64, string, bool:
		default:
			return nil, errors.New("FORMAT_ERROR")
		}
		if fmt.Sprint(value) != fmt.Sprint(allowed) {
			return nil, errors.New("NOT_ALLOWED_VALUE")
		}

		return value, nil
	}
}

// LengthBetween - make sure that validated value's length is between specified length.
func LengthBetween(args ...interface{}) func(...interface{}) (interface{}, interface{}) {
	var minLength, maxLength float64
	if len(args) > 1 {
		if v, ok := args[0].(float64); ok {
			minLength = v
		}
		if v, ok := args[1].(float64); ok {
			maxLength = v
		}
	}
	return func(builders ...interface{}) (interface{}, interface{}) {
		var value interface{}
		if len(builders) > 0 {
			value = builders[0]
		}
		if value == nil || value == "" {
			return value, nil
		}

		switch v := value.(type) {
		case string:
			if utf8.RuneCountInString(v) > int(maxLength) {
				return nil, errors.New("TOO_LONG")
			}
			if utf8.RuneCountInString(v) < int(minLength) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		case bool:
			if utf8.RuneCountInString(strconv.FormatBool(v)) > int(maxLength) {
				return nil, errors.New("TOO_LONG")
			}
			if utf8.RuneCountInString(strconv.FormatBool(v)) < int(minLength) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		case float64:
			if utf8.RuneCountInString(strconv.FormatFloat(v, 'f', -1, 64)) > int(maxLength) {
				return nil, errors.New("TOO_LONG")
			}
			if utf8.RuneCountInString(strconv.FormatFloat(v, 'f', -1, 64)) < int(minLength) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		default:
			return nil, errors.New("FORMAT_ERROR")
		}
	}
}

// LengthEqual - make sure that validated value's length is equal to specified length.
func LengthEqual(args ...interface{}) func(...interface{}) (interface{}, interface{}) {
	var length float64
	if len(args) > 0 {
		if v, ok := args[0].(float64); ok {
			length = v
		}
	}

	return func(builders ...interface{}) (interface{}, interface{}) {
		var value interface{}
		if len(builders) > 0 {
			value = builders[0]
		}
		if value == nil || value == "" {
			return value, nil
		}

		switch v := value.(type) {
		case string:
			if utf8.RuneCountInString(v) > int(length) {
				return nil, errors.New("TOO_LONG")
			}
			if utf8.RuneCountInString(v) < int(length) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		case bool:
			if utf8.RuneCountInString(strconv.FormatBool(v)) > int(length) {
				return nil, errors.New("TOO_LONG")
			}
			if utf8.RuneCountInString(strconv.FormatBool(v)) < int(length) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		case float64:
			if utf8.RuneCountInString(strconv.FormatFloat(v, 'f', -1, 64)) > int(length) {
				return nil, errors.New("TOO_LONG")
			}
			if utf8.RuneCountInString(strconv.FormatFloat(v, 'f', -1, 64)) < int(length) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		default:
			return nil, errors.New("FORMAT_ERROR")
		}
	}
}

// Like - check that validated value is like specified regexp.
func Like(args ...interface{}) func(...interface{}) (interface{}, interface{}) {
	var re *regexp.Regexp
	var flags string
	if len(args) > 0 {
		if len(args) > 1 {
			if v, ok := args[1].(string); ok {
				if v == "i" {
					flags = "(?i)"
				}
			}
		}

		if v, ok := args[0].(string); ok {
			reg, err := regexp.Compile(flags + v)
			if err != nil {
				re = regexp.MustCompile(".*")
			} else {
				re = reg
			}
		}
	}

	return func(builders ...interface{}) (interface{}, interface{}) {
		var value interface{}
		if len(builders) > 0 {
			value = builders[0]
		}
		if value == nil || value == "" {
			return value, nil
		}

		switch v := value.(type) {
		case string:
			if ok := re.MatchString(v); !ok {
				return nil, errors.New("WRONG_FORMAT")
			}
			return v, nil
		case float64:
			if ok := re.MatchString(strconv.FormatFloat(v, 'f', -1, 64)); !ok {
				return nil, errors.New("WRONG_FORMAT")
			}
			return v, nil
		default:
			return nil, errors.New("FORMAT_ERROR")
		}
	}
}

// MaxLength - check that validated value's length is not longer than specified.
func MaxLength(args ...interface{}) func(...interface{}) (interface{}, interface{}) {
	var minLength float64
	if len(args) > 0 {
		if v, ok := args[0].(float64); ok {
			minLength = v
		}
	}

	return func(builders ...interface{}) (interface{}, interface{}) {
		var value interface{}
		if len(builders) > 0 {
			value = builders[0]
		}
		if value == nil || value == "" {
			return value, nil
		}

		switch v := value.(type) {
		case string:
			if utf8.RuneCountInString(v) > int(minLength) {
				return nil, errors.New("TOO_LONG")
			}
			return v, nil
		case bool:
			if utf8.RuneCountInString(strconv.FormatBool(v)) > int(minLength) {
				return nil, errors.New("TOO_LONG")
			}
			return v, nil
		case float64:
			if utf8.RuneCountInString(strconv.FormatFloat(v, 'f', -1, 64)) > int(minLength) {
				return nil, errors.New("TOO_LONG")
			}
			return v, nil
		default:
			return nil, errors.New("FORMAT_ERROR")
		}
	}
}

// MinLength - check that validated value's length is not shorter than specified.
func MinLength(args ...interface{}) func(...interface{}) (interface{}, interface{}) {
	var minLength float64
	if len(args) > 0 {
		if v, ok := args[0].(float64); ok {
			minLength = v
		}
	}

	return func(builders ...interface{}) (interface{}, interface{}) {
		var value interface{}
		if len(builders) > 0 {
			value = builders[0]
		}
		if value == nil || value == "" {
			return value, nil
		}

		switch v := value.(type) {
		case string:
			if utf8.RuneCountInString(v) < int(minLength) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		case bool:
			if utf8.RuneCountInString(strconv.FormatBool(v)) < int(minLength) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		case float64:
			if utf8.RuneCountInString(strconv.FormatFloat(v, 'f', -1, 64)) < int(minLength) {
				return nil, errors.New("TOO_SHORT")
			}
			return v, nil
		default:
			return nil, errors.New("FORMAT_ERROR")
		}
	}
}

// OneOf - check that validated value is one of specified.
func OneOf(args ...interface{}) func(...interface{}) (interface{}, interface{}) {
	var allowed []interface{}
	if len(args) > 0 {
		if v, ok := args[0].([]interface{}); ok {
			allowed = v
		} else {
			allowed = args[0 : len(args)-1]
		}
	}

	return func(builders ...interface{}) (interface{}, interface{}) {
		var value interface{}
		if len(builders) > 0 {
			value = builders[0]
		}
		if value == nil || value == "" {
			return value, nil
		}

		for _, val := range allowed {
			switch v := val.(type) {
			case bool:
				switch vv := value.(type) {
				case bool:
					if vv == v {
						return val, nil
					}
				case string:
					if vs, _ := strconv.ParseBool(vv); vs == v {
						return val, nil
					}
				default:
					return nil, errors.New("FORMAT_ERROR")
				}
			case string:
				switch vv := value.(type) {
				case bool:
					if vb := strconv.FormatBool(vv); vb == v {
						return val, nil
					}
				case string:
					if vv == v {
						return val, nil
					}
				case float64:
					if vf := strconv.FormatFloat(vv, 'f', -1, 64); vf == v {
						return val, nil
					}
				default:
					return nil, errors.New("FORMAT_ERROR")
				}
			case float64:
				switch vv := value.(type) {
				case string:
					if vs := strconv.FormatFloat(v, 'f', -1, 64); vs == vv {
						return val, nil
					}
				case float64:
					if vv == v {
						return val, nil
					}
				default:
					return nil, errors.New("FORMAT_ERROR")
				}
			}
		}
		return nil, errors.New("NOT_ALLOWED_VALUE")
	}
}

// String - check that validated value is string-like.
func String(args ...interface{}) func(...interface{}) (interface{}, interface{}) {
	return func(builders ...interface{}) (interface{}, interface{}) {
		var value interface{}
		if len(builders) > 0 {
			value = builders[0]
		}
		if value == nil || value == "" {
			return value, nil
		}

		switch v := value.(type) {
		case string:
			return v, nil
		case bool:
			return strconv.FormatBool(v), nil
		case float64:
			return strconv.FormatFloat(v, 'f', -1, 64), nil
		default:
			return nil, errors.New("FORMAT_ERROR")
		}
	}
}
