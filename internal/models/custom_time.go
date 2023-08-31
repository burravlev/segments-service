package models

import (
	"fmt"
	"strings"
	"time"
)

const layout = "2006-01-02"

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(layout, s)
	if err != nil {
		return err
	}
	*ct = CustomTime(nt)
	return
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *CustomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(layout))
}
