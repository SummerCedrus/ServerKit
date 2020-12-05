package timer

import (
	//"github.com/SummerCedrus/ServerKit/misc"
	"github.com/SummerCedrus/ServerKit/misc"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	MINUTE	= 0
	HOUR	= 1
	//DAY		= 2
	WEEK	= 3
	//MONTH	= 4
	ONE_DAY_SEC = 86400
	ONE_DAY_MS = 86400000
	ONE_HOUR_SEC = 3600
	ONE_HOUR_MS = 3600000
	ONE_MINUTE_SEC = 60
	ONE_MINUTE_MS = 60000

	CRON_TIME_ERROR = -2
	CRON_EVERY_TIME = -1
)

type Event struct {
	ID	int32
	TriggerMs int64	//触发时间戳(秒)
	DeltaMs	int64
	IsRepeat bool
	Conf	map[int16]int16
	Func	func(interface{})interface{}
	Args	interface{}
}
type Timer struct {
	Events map[int32]*Event
	FRAME_TIME int64
	MaxID	int32
	lock 	sync.Mutex
}

func NewTimer() *Timer{
	tr := Timer{
		Events: make(map[int32]*Event),
		FRAME_TIME:time.Now().UnixNano()/int64(time.Millisecond),
		MaxID:0,
	}

	return &tr
}

func (self *Timer) Run(){
	go func() {
		for{
			//一秒检查一次
			time.Sleep(time.Second)
			self.FRAME_TIME = time.Now().UnixNano()/int64(time.Millisecond)
			for _, event := range self.Events{
				if self.FRAME_TIME < event.TriggerMs{
					//misc.Debugf("no event %v %v", self.FRAME_TIME, event.TriggerMs)
					continue
				}
				//misc.Debugf("exec event %v %v", self.FRAME_TIME, event.TriggerMs)
				event.Func(event.Args)
				if !event.IsRepeat || event.calNextMs() <= 0{
					self.delEvent(event.ID)//一次性事件
				}
			}
		}
	}()
}
//cronStr: "m,h,d"
func (self *Timer)CreateEvent(deltaMs int64, isRepeat bool, cronStr string, f func(interface{})interface{}, args interface{}){
	if deltaMs <= 0 && "" == cronStr{
		misc.Errorf("params error!!!")
		return
	}
	if deltaMs > 0 && "" != cronStr{
		misc.Warnf("cronStr is shadow by deltaMs!!!")
	}
	event := Event{
		DeltaMs:deltaMs,
		IsRepeat:isRepeat,
		Func:f,
		Args:args,
		Conf:make(map[int16]int16),
	}

	if deltaMs <= 0{
		minute, hour, week := parseCronStr(cronStr)
		if -2 == minute || -2 == hour || -2 == week{
			return
		}
		if -1 !=  minute{
			event.Conf[MINUTE] = int16(minute)
		}

		if -1 != hour{
			event.Conf[HOUR] = int16(hour)
		}

		if -1 != week{
			event.Conf[WEEK] = int16(week)
		}

	}
	event.calNextMs()
	misc.Debug(event)
	self.addEvent(&event)
}

func parseCronStr(str string) (int16, int16 ,int16){
	vStr := strings.Split(str, ",")
	if 3 != len(vStr){
		misc.Errorf("Input Time Error!!!")
		return -2, -2, -2
	}
	week := -1
	hour := -1
	minute := -1
	var err error
	if "*" != vStr[2]{
		week,err = strconv.Atoi(vStr[2])
		if nil != err || week > 6 || week < 0{
			misc.Errorf("Input Week Error %v %v!!!",err, week)
			return -2, -2, -2
		}
	}

	if "*" != vStr[1]{
		hour, err = strconv.Atoi(vStr[1])
		if nil != err || hour > 23 || hour < 0{
			misc.Errorf("Input Hour Error %v %v!!!", err, week)
			return -2, -2, -2
		}
	}

	if "*" != vStr[0]{
		minute, err = strconv.Atoi(vStr[0])
		if nil != err || minute > 59 || minute < 0{
			misc.Errorf("Input Minute Error %v %v!!!", err, week)
			return -2, -2, -2
		}
	}

	return int16(minute), int16(hour), int16(week)
}
func (self *Timer) addEvent(event *Event){
	defer self.lock.Unlock()
	self.lock.Lock()
	self.MaxID++
	event.ID = self.MaxID
	self.Events[self.MaxID] = event
}

func (self *Timer) delEvent(id int32){
	defer self.lock.Unlock()
	self.lock.Lock()
	delete(self.Events, id)
}

func (event *Event) calNextMs() int64{


	nextMs := CurMidNight()
	nextWeek, existWeek := event.Conf[WEEK]
	nextHour, existHour := event.Conf[HOUR]
	nextMinute, existMinute := event.Conf[MINUTE]
	//misc.Debugf("%v %v %v", existWeek,existHour,existMinute)
	if event.DeltaMs > 0{
		nextMs = Now() + event.DeltaMs
	}else if existWeek {//每个周几的几点几分触发
		nowWeek := int16(time.Now().Weekday())
		deltaDay := int16(0)
		if nextWeek < nowWeek {
			deltaDay = 7 - (nowWeek - nextWeek)
		} else {
			deltaDay = nextWeek - nowWeek
		}
		nextMs += int64(deltaDay)*ONE_DAY_MS
		if existHour {
			nextMs += int64(nextHour) * ONE_HOUR_MS
		}
		if existMinute {
			nextMs += int64(nextMinute) * ONE_MINUTE_MS
		}
	} else if existHour {//每天的几点几分触发
		nextMs += int64(nextHour) * ONE_HOUR_MS
		if existMinute {
			nextMs += int64(nextMinute) * ONE_MINUTE_MS
		}
		if Now() > nextMs {
			nextMs += ONE_DAY_MS
		}
	} else if existMinute {//每小时的几分触发
		nextMs = CurHour()
		nextMs += int64(nextMinute) * ONE_MINUTE_MS
		if Now() > nextMs {
			nextMs += ONE_HOUR_MS
		}
	}
	if nextMs < Now(){
		misc.Errorf("nextMs %v < Now %v Error!!!", nextMs, Now())
		return 0
	}
	event.TriggerMs = nextMs
	misc.Debugf("calNextTs %d\n", nextMs/1000)
	return nextMs
}


