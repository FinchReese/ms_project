package time_format

import "time"

func ConvertTimeToString(t time.Time) string {
	return t.Format("2006-01-02_15:04:05")
}

func ConvertTimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func ConvertTimeToYMD(t time.Time) string {
	return t.Format("2006-01-02")
}
func ConvertMsecToString(t int64) string {
	return time.UnixMilli(t).Format("2006-01-02 15:04:05")
}

func ParseTimeStr(timeStr string) (int64, error) {
	t, error := time.Parse("2006-01-02 15:04", timeStr)
	if error != nil {
		return 0, error
	}
	return t.UnixMilli(), nil
}
