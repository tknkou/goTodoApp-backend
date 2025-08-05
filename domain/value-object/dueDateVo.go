package value_object

import(
	"time"
	"errors"
)

type DueDate struct{
	value time.Time
}

func NewDueDate(value string)(*DueDate, error){
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}
	return &DueDate{value : parsedTime}, nil	
}

func ValidateDueDateNotPast(d *DueDate) error {
	if d == nil {
		return nil
	}
	// 現在の日付（時刻は0時にリセット）
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// dueDateの日付部分を取り出す
	due := d.Value()
	dueDateOnly := time.Date(due.Year(), due.Month(), due.Day(), 0, 0, 0, 0, due.Location())

	if dueDateOnly.Before(today) {
		return errors.New("due date cannot be in the past")
	}
	return nil
}

func (d DueDate) Value() time.Time{
	return d.value
}