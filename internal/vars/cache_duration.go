package vars

import (
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type ttlPeriod struct {
	ttlShortPeriod  string
	ttlMediumPeriod string
	ttlLongPeriod   string
	ttlDailyPeriod  string
	ttlH4Period     string
	ttlM15Period    string
	ttlM5Period     string
	ttlM1Period     string
}

var ttl *ttlPeriod

func newTtlPeriod() *ttlPeriod {
	ttl := new(ttlPeriod)

	ttlShortPeriod := viper.GetString("cache.ttl.short")
	if ttlShortPeriod == "" {
		ttlShortPeriod = "3h"
	}
	ttl.ttlShortPeriod = ttlShortPeriod

	ttlMediumPeriod := viper.GetString("cache.ttl.medium")
	if ttlMediumPeriod == "" {
		ttlMediumPeriod = "24h"
	}
	ttl.ttlMediumPeriod = ttlMediumPeriod

	ttlLongPeriod := viper.GetString("cache.ttl.long")
	if ttlLongPeriod == "" {
		ttlLongPeriod = "72h"
	}
	ttl.ttlLongPeriod = ttlLongPeriod

	ttlDailyPeriod := viper.GetString("cache.daily.ttl")
	if ttlDailyPeriod == "" {
		ttlDailyPeriod = "24h"
	}
	ttl.ttlDailyPeriod = ttlDailyPeriod

	ttlH4Period := viper.GetString("cache.h4.ttl")
	if ttlH4Period == "" {
		ttlH4Period = "4h"
	}
	ttl.ttlH4Period = ttlH4Period

	ttlM15Period := viper.GetString("cache.m15.ttl")
	if ttlM15Period == "" {
		ttlM15Period = "15m"
	}
	ttl.ttlM15Period = ttlM15Period

	ttlM5Period := viper.GetString("cache.m5.ttl")
	if ttlM5Period == "" {
		ttlM5Period = "5m"
	}
	ttl.ttlM5Period = ttlM5Period

	ttlM1Period := viper.GetString("cache.m1.ttl")
	if ttlM1Period == "" {
		ttlM1Period = "1m"
	}
	ttl.ttlM1Period = ttlM1Period

	return ttl
}

func init() {
	ttl = newTtlPeriod()
}

func GetTtlShortPeriod() time.Duration {
	return convertStringToDuration(ttl.getTtlShortPeriod())
}

func GetTtlMediumPeriod() time.Duration {
	return convertStringToDuration(ttl.getTtlMediumPeriod())
}

func GetTtlLongPeriod() time.Duration {
	return convertStringToDuration(ttl.getTtlLongPeriod())
}

func GetTtlDailyPeriod() time.Duration {
	return convertStringToDuration(ttl.getDailyPeriod())
}

func GetTtlH4Period() time.Duration {
	return convertStringToDuration(ttl.getH4Period())
}

func GetTtlM15Period() time.Duration {
	return convertStringToDuration(ttl.getM15Period())
}

func GetTtlM5Period() time.Duration {
	return convertStringToDuration(ttl.getM5Period())
}

func GetTtlM1Period() time.Duration {
	return convertStringToDuration(ttl.getM1Period())
}

func (cachePeriod *ttlPeriod) getTtlShortPeriod() string {
	return cachePeriod.ttlShortPeriod
}

func (cachePeriod *ttlPeriod) getTtlMediumPeriod() string {
	return cachePeriod.ttlMediumPeriod
}

func (cachePeriod *ttlPeriod) getTtlLongPeriod() string {
	return cachePeriod.ttlLongPeriod
}

func (cachePeriod *ttlPeriod) getDailyPeriod() string {
	return cachePeriod.ttlDailyPeriod
}

func (cachePeriod *ttlPeriod) getH4Period() string {
	return cachePeriod.ttlH4Period
}

func (cachePeriod *ttlPeriod) getM15Period() string {
	return cachePeriod.ttlM15Period
}

func (cachePeriod *ttlPeriod) getM5Period() string {
	return cachePeriod.ttlM5Period
}

func (cachePeriod *ttlPeriod) getM1Period() string {
	return cachePeriod.ttlM1Period
}

func convertStringToDuration(ttl string) time.Duration {
	duration, err := time.ParseDuration(ttl)
	if err != nil {
		log.Errorf("failed parsing duration", err)
		return time.Hour
	}

	return duration
}
