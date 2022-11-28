package valueTypes

import (
	"encoding/json"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"time"
)


var inputDateLayout = []string{
	DateTimeFullLayout,
	DateTimeLayout,
	"2006/01/02 15:04:05",
	DateLayout,
	"2006/01/02",
	DateTimeLayoutSecond,
	DateTimeLayoutMinute,
	DateTimeLayoutHour,
	DateTimeLayoutDay,
	DateTimeLayoutMonth,
	DateTimeLayoutYear,
	DateTimeLayoutYear,
}

const (
	DateTimeFullLayout        = time.RFC3339
	DateTimeLayout            = DateLayout + " " + TimeLayout
	DateTimeAltLayout         = DateLayoutDay + "-" + TimeLayoutSecond
	DateTimeLayoutZeroSeconds = DateLayout + " " + TimeLayoutZeroSeconds
	DateTimeLayoutSecond      = DateLayoutDay + TimeLayoutSecond
	DateTimeLayoutMinute      = DateLayoutDay + TimeLayoutMinute
	DateTimeLayoutHour        = DateLayoutDay + TimeLayoutHour
	DateTimeLayoutDay         = DateLayoutDay
	DateTimeLayoutMonth       = DateLayoutMonth
	DateTimeLayoutYear        = DateLayoutYear

	DateLayout            = "2006-01-02"
	DateLayoutDay         = "20060102"
	DateLayoutMonth       = "200601"
	DateLayoutYear        = "2006"

	TimeLayout            = "15:04:05"
	TimeLayoutZeroSeconds = "15:04:00"
	TimeLayoutHourColon   = "15:04"
	TimeLayoutSecond      = "150405"
	TimeLayoutMinute      = "1504"
	TimeLayoutHour        = "15"
)


type DateTime struct {
	string    `json:"string,omitempty"`
	time.Time `json:"time,omitempty"`
	DateType  string
	format    string
	Error     error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		if len(data) == 0 {
			break
		}

		// Store result from string
		dt.Error = json.Unmarshal(data, &dt.string)
		if dt.Error == nil {
			dt.SetString(dt.string)
			break
		}

		// Store result from time
		dt.Error = json.Unmarshal(data, &dt.Time)
		if dt.Error == nil {
			dt.SetValue(dt.Time)
			break
		}

		for _, f := range inputDateLayout {
			dt.Time, dt.Error = time.Parse(f, string(data))
			if dt.Error == nil {
				dt.format = f
				dt.string = dt.Time.Format(DateTimeLayout)
				dt.SetDateType(string(data))
				break
			}
		}

		if dt.Error != nil {
			fmt.Printf("Error:UnmarshalJSON DateTime(%s) - %s\n", string(data), dt.Error)
		}
	}

	return dt.Error
}

// MarshalJSON - Convert value to JSON
func (dt DateTime) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		// data = []byte("\"" + dt.Time.Format(DateTimeLayout) + "\"")
		// data = []byte("\"" + dt.string + "\"")
		// data = []byte("\"" + dt.Original() + "\"")
		data = []byte("\"" + dt.Time.Format(dt.format) + "\"")
	}

	return data, dt.Error
}

func (dt DateTime) Value() time.Time {
	return dt.Time
}

func (dt DateTime) String() string {
	// return dt.Original()
	if dt.IsZero() {
		return "--"
	}
	return dt.Time.Format(DateTimeLayout)
}

func (dt DateTime) Match(comp time.Time) bool {
	if dt.Time == comp {
		return true
	}
	return false
}

func (dt *DateTime) SetString(value string) DateTime {
	for range Only.Once {
		dt.string = value
		dt.Time = time.Time{}

		if value == "" {
			break
		}

		if value == "--" {
			value = ""
			break
		}

		for _, f := range inputDateLayout {
			dt.Time, dt.Error = time.Parse(f, value)
			if dt.Error == nil {
				dt.format = f
				dt.string = dt.Time.Format(f)
				dt.SetDateType(value)
				break
			}
		}

		if dt.Error != nil {
			fmt.Printf("Error:SetString DateTime(%s) - %s\n", value, dt.Error)
		}
	}

	return *dt
}

func (dt *DateTime) SetValue(value time.Time) DateTime {
	for range Only.Once {
		dt.string = ""
		dt.Time = value

		if value.IsZero() {
			break
		}

		dt.string = value.Format(DateTimeLayout)
		dt.format = DateTimeLayout
		dt.DateType = "3"
	}

	return *dt
}

func (dt *DateTime) SetDateType(value string) {
	switch  {
		case len(value) == len(DateTimeLayout):
			dt.DateType = "1"
		case len(value) == len(DateTimeLayoutYear):
			dt.DateType = "3"
		case len(value) == len(DateTimeLayoutMonth):
			dt.DateType = "2"
		case len(value) == len(DateTimeLayoutDay):
			dt.DateType = "1"
		case len(value) == len(DateTimeLayoutHour):
			dt.DateType = "1"
		case len(value) == len(DateTimeLayoutMinute):
			dt.DateType = "1"
		case len(value) == len(DateTimeLayoutSecond):
			dt.DateType = "1"
		case value == "total":
			dt.DateType = "4"
	}
}

func SetDateTimeString(value string) DateTime {
	var t DateTime
	return t.SetString(value)
}

func SetDateTimeValue(value time.Time) DateTime {
	var t DateTime
	return t.SetValue(value)
}

func (dt *DateTime) GetDayStartTimestamp() string {
	var ret string
	f1 := dt.Time.Round(time.Hour * 24)
	ret = f1.Format(DateTimeLayoutSecond)
	return ret
	// return fmt.Sprintf("%s000000", dt.Time.Format(DtLayoutDay))
}

func (dt *DateTime) GetDayEndTimestamp() string {
	var ret string
	f1 := dt.Time.Round(time.Hour * 24).Add(time.Hour * 24)
	ret = f1.Format(DateTimeLayoutSecond)
	return ret
	// return fmt.Sprintf("%s235900", dt.Time.Format(DtLayoutDay))
}

func (dt DateTime) PrintFull() string {
	return dt.Time.Format(DateTimeLayout)
}

func (dt DateTime) Original() string {
	var ret string
	switch dt.DateType {
		case "3":
			ret = dt.Time.Format(DateTimeLayoutYear)
		case "2":
			ret = dt.Time.Format(DateTimeLayoutMonth)
		case "1":
			ret = dt.Time.Format(DateTimeLayoutDay)
		default:
			ret = dt.Time.Format(DateTimeLayout)
	}
	return ret
}

const Now = "now"

func NewDateTime(value string) DateTime {
	var ret DateTime
	for range Only.Once {
		if (value == Now) || (value == "") {
			// value = time.Now().Format(DateTimeLayout)
			ret.SetValue(time.Now())
			ret.SetDateType(ret.string)
			break
		}
		for _, f := range inputDateLayout {
			ret.Time, ret.Error = time.Parse(f, value)
			if ret.Error == nil {
				ret.format = f
				ret.SetValue(ret.Time)
				ret.SetDateType(value)
				break
			}
		}

		if ret.Error != nil {
			fmt.Printf("Error:NewDateTime DateTime(%s) - %s\n", value, ret.Error)
		}
	}
	return ret
}

func TimeNowString() string {
	return time.Now().Format(DateTimeLayout)
}
