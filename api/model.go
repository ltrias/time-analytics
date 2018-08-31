package api

import (
	"database/sql/driver"
	"time"

	"github.com/go-ozzo/ozzo-validation"
)

const TIME_FORMAT string = "2006-01-02"

type MyDate struct {
	time time.Time
}

func (md MyDate) MarshalJSON() ([]byte, error) {
	s := md.time.Format(TIME_FORMAT)

	return []byte("\"" + s + "\""), nil
}

func (md *MyDate) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = s[1 : len(s)-1]

	t, err := time.Parse(TIME_FORMAT, s)
	if err != nil {
		return err
	}

	md.time = t

	return nil
}

func (md MyDate) Value() (driver.Value, error) {
	return md.time, nil
}

func (md *MyDate) Scan(value interface{}) error {
	if value != nil {
		md.time = value.(time.Time)
	}

	return nil
}

func (md MyDate) Validate() error {
	return validation.ValidateStruct(&md,
		validation.Field(&md.time, validation.Required))
}

//TimeEvent represents a time event such as a meeting
type TimeEvent struct {
	ID         int    `json:"id"`
	Day        MyDate `json:"day"`
	Type       string `json:"type"`
	Who        string `json:"who"`
	Duration   int    `json:"duration"`
	Subject    string `json:"subject"`
	Department string `json:"department"`
	Recurrent  bool   `json:"recurrent"`
}

func (t TimeEvent) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Day, validation.Required),
		validation.Field(&t.Type, validation.Required),
		validation.Field(&t.Who, validation.Required),
		validation.Field(&t.Duration, validation.Required),
		validation.Field(&t.Subject, validation.Required),
		validation.Field(&t.Recurrent, validation.NotNil))
}

//Suggest holds information about suggest
type Suggest struct {
	Type       []string `json:"type"`
	Who        []string `json:"who"`
	Duration   []int    `json:"duration"`
	Subject    []string `json:"subject"`
	Department []string `json:"department"`
}
