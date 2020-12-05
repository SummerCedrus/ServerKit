package timer

import "time"

func CurMidNight() int64{
	t := time.Now()
	y,m,d := t.Date()
	loc := t.Location()
	return Time2Ms(time.Date(y,m,d,0,0,0,0,loc))
}

func Time2Ms(t time.Time) int64{
	return t.UnixNano()/int64(time.Millisecond)
}

func Now() int64{
	return time.Now().UnixNano()/int64(time.Millisecond)
}

func CurHour() int64{
	t := time.Now()
	y,m,d := t.Date()
	h := t.Hour()
	loc := t.Location()
	return Time2Ms(time.Date(y,m,d,h,0,0,0,loc))
}
