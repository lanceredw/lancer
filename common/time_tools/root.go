package time_tools

import (
	"database/sql/driver"
	"time"
)

type SerializerTime time.Time

var (
	cst *time.Location = time.Local
)

func init() {
	/*var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}*/
}

// CSTLayout China Standard SerializerTime Layout
const CSTLayout = "2006-01-02 15:04:05.999"

// ConvertSerializerTime 类型转换
func ConvertSerializerTime(t time.Time) SerializerTime {
	return SerializerTime(t)
}

func ConvertTime(t SerializerTime) time.Time {
	return time.Time(t)
}

// SerializerTimeNow 获得一个当前时间
func SerializerTimeNow() SerializerTime {
	return SerializerTime(time.Now())
}

func (t *SerializerTime) UnmarshalJSON(data []byte) (err error) {
	var now time.Time
	if string(data) != "null" || string(data) != "" {
		r := []rune(string(data))
		if string(r[0]) == `"` && string(r[len(r)-1]) == `"` {
			r = r[1 : len(r)-1]
		}
		if string(r) != "" {
			now, err = ParseCSTInLocation(string(r))
			if err != nil {
				var ns string
				ns, err = RFC3339ToCSTLayout(string(r))
				if err != nil {
					return err
				}
				now, err = ParseCSTInLocation(ns)
				if err != nil {
					return err
				}
			}
		}

	}

	*t = SerializerTime(now)
	return
}

func (t SerializerTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(CSTLayout)+2)
	b = append(b, []byte(`"`)...)
	// 时间有值才处理
	if time.Time(t).Unix() > 0 {
		b = time.Time(t).AppendFormat(b, CSTLayout)
	}
	b = append(b, []byte(`"`)...)
	return b, nil
}

func (t *SerializerTime) Scan(value interface{}) error {
	if value == nil {
		*t = SerializerTime(time.Time{})
		return nil
	}
	switch value.(type) {
	case time.Time:
		*t = SerializerTime(value.(time.Time))
	case string:
		r := value.(string)
		tm, err := ParseCSTInLocation(r)
		if err != nil {
			var ns string
			ns, err = RFC3339ToCSTLayout(r)
			if err != nil {
				return err
			}
			tm, err = ParseCSTInLocation(ns)
			if err != nil {
				return err
			}
		}
		*t = SerializerTime(tm)
	}

	return nil
}

func (t SerializerTime) Value() (driver.Value, error) {
	if time.Time(t).Unix() == 0 {
		return nil, nil
	}
	return time.Time(t), nil
}

func (t SerializerTime) String() string {
	if time.Time(t).Unix() == 0 {
		return ""
	}
	return time.Time(t).Format(CSTLayout)
}

// MarshalBinary 实现 BinaryMarshaler 接口
func (t SerializerTime) MarshalBinary() ([]byte, error) {
	b := make([]byte, 0, len(CSTLayout)+2)
	b = append(b, []byte(`"`)...)
	// 时间有值才处理
	if time.Time(t).Unix() > 0 {
		b = time.Time(t).AppendFormat(b, CSTLayout)
	}
	b = append(b, []byte(`"`)...)
	return b, nil
}

// UnmarshalBinary 实现 BinaryUnmarshaler 接口
func (t *SerializerTime) UnmarshalBinary(data []byte) (err error) {
	var now time.Time
	if string(data) != "null" || string(data) != "" {
		r := []rune(string(data))
		if string(r[0]) == `"` && string(r[len(r)-1]) == `"` {
			r = r[1 : len(r)-1]
		}
		if string(r) != "" {
			now, err = ParseCSTInLocation(string(r))
			if err != nil {
				var ns string
				ns, err = RFC3339ToCSTLayout(string(r))
				if err != nil {
					return err
				}
				now, err = ParseCSTInLocation(ns)
				if err != nil {
					return err
				}
			}
		}

	}

	*t = SerializerTime(now)
	return
}

func (t *SerializerTime) UnmarshalText(data []byte) (err error) {
	var now time.Time
	if string(data) != "null" || string(data) != "" {
		r := []rune(string(data))
		if string(r[0]) == `"` && string(r[len(r)-1]) == `"` {
			r = r[1 : len(r)-1]
		}
		if string(r) != "" {
			now, err = ParseCSTInLocation(string(r))
			if err != nil {
				var ns string
				ns, err = RFC3339ToCSTLayout(string(r))
				if err != nil {
					return err
				}
				now, err = ParseCSTInLocation(ns)
				if err != nil {
					return err
				}
			}
		}

	}

	*t = SerializerTime(now)
	return nil
}
