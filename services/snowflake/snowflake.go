package snowflake

import (
	"errors"
	"github.com/fjlyx97/short_url/utils/log"
	"sync"
	"time"
)

type Snowflake struct {
	// 开始时间戳 (41bit)
	startTimestamp int64
	// 上次生成Id的时间戳
	lastTimestamp int64
	// 序列号所占的位数 (12bit)
	sequenceBits int64
	// 序列号计数器
	sequence int64
	// 序列号掩码（用于保证序列号不超过指定位数）
	sequenceMask int64
	// 工作机器Id (10bit) = 机房（5bit) + 机子（5bit)
	dataCenterIdBits int64
	workerIdBits     int64
	dataCenterId     int64
	workerId         int64
	// 机器Id偏移量
	dataCenterIdShift int64
	workerIdShift     int64
	// 时间戳偏移量
	timestampShift int64
	// 锁，用于自增序列
	lock sync.Mutex
}

// 生成雪花算法模型，目前位数都是Hard code
func NewSnowflake(startTimestamp int64, dataCenterId int64, workerId int64) *Snowflake {
	var maxWorkerId int64 = (-1 << 5) ^ -1
	var maxDataCenterId int64 = (-1 << 5) ^ -1
	// 41位，最大69年
	var maxTime int64 = (-1 << 41) ^ -1
	// 校验机器Id是否在指定位数内
	if dataCenterId < 0 || dataCenterId > maxDataCenterId ||
		workerId < 0 || workerId > maxWorkerId {
		log.GLogger.Errorf("workerId or dataCenterId error , workerId : %d , dataCenterId : %d", workerId, dataCenterId)
		return nil
	}
	// 校验时间戳
	if startTimestamp > time.Now().UnixNano()/int64(time.Millisecond) {
		log.GLogger.Errorf("start time over now , startTime : %d", startTimestamp)
		return nil
	}
	if time.Now().UnixNano()/int64(time.Millisecond)-startTimestamp > maxTime {
		log.GLogger.Errorf("nowTime subtract startTime over 69 years")
		return nil
	}
	return &Snowflake{
		startTimestamp:    startTimestamp,
		lastTimestamp:     -1,
		sequenceBits:      12,
		sequence:          0,
		sequenceMask:      (-1 << 12) ^ -1,
		dataCenterId:      dataCenterId,
		dataCenterIdBits:  5,
		workerId:          workerId,
		workerIdBits:      5,
		dataCenterIdShift: 17,
		workerIdShift:     12,
		timestampShift:    22,
		lock:              sync.Mutex{},
	}
}

func (s *Snowflake) NextId() (int64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	nowTime := s.genMillisecond()
	if nowTime < s.lastTimestamp {
		log.GLogger.Errorf("Clock moved backwards , lastTimeStamp : %d , nowTime : %d.", s.lastTimestamp, nowTime)
		return -1, errors.New("clock moved backwards")
	}
	// 同一毫秒多个请求
	if nowTime == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & s.sequenceMask
		// 同一毫秒Id被占满
		if s.sequence == 0 {
			nowTime = s.waitNextMillisecond()
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = nowTime
	return (nowTime-s.startTimestamp)<<s.timestampShift | (s.dataCenterId << s.dataCenterIdShift) |
		(s.workerId << s.workerIdShift) | s.sequence, nil
}

func (s *Snowflake) genMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (s *Snowflake) waitNextMillisecond() int64 {
	newTime := s.genMillisecond()
	for newTime <= s.lastTimestamp {
		newTime = s.genMillisecond()
	}
	return newTime
}
