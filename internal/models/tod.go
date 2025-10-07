package models

import (
	"fmt"
	"regexp"
	"strconv"
)

type TOD struct {
	Hours   int
	Minutes int
}

func (tod TOD) String() string {
	return fmt.Sprintf("%s:%s", toTimeStr(tod.Hours), toTimeStr(tod.Minutes))
}

func TimeOfDay(s string) (TOD, error) {
	tod := TOD{}
	return tod, nil
}

func toTimeStr(i int) string {
	s := fmt.Sprintf("%d", i)
	for len(s) < 2 {
		s = "0" + s
	}
	return s
}

func parseTOD(s string) (TOD, error) {
	tod := TOD{}
	re := regexp.MustCompile(`(\d+):(\d+)`)
	found := re.FindStringSubmatch(s)
	if len(found) != 3 {
		return tod, fmt.Errorf("invalid string provided: '%v'", s)
	}
	h, _ := strconv.Atoi(found[1])
	m, _ := strconv.Atoi(found[2])
	tod.Hours, tod.Minutes = h, m
	return tod, validateTOD(tod)
}

func validateTOD(tod TOD) error {
	if tod.Hours < 0 {
		return fmt.Errorf("invalid tod: hours are negative")
	}
	if tod.Minutes < 0 {
		return fmt.Errorf("invalid tod: minutes are negative")
	}
	if tod.Hours > 23 {
		return fmt.Errorf("invalid tod: hours are more than 23")
	}
	if tod.Minutes > 59 {
		return fmt.Errorf("invalid tod: minutes are more than 59")
	}
	return nil
}
