package model

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getValueString(t *testing.T) {
	typ := xsdType("string")
	val, err := typ.getValue("This is just same Random String")
	assert.Nil(t, err)
	assert.Equal(t, "This is just same Random String", val)
}

func Test_getValueBoolean(t *testing.T) {
	typ := xsdType("boolean")
	val, err := typ.getValue("true ")
	assert.Nil(t, err)
	assert.True(t, val.(bool))

	val, err = typ.getValue("1")
	assert.Nil(t, err)
	assert.True(t, val.(bool))

	val, err = typ.getValue(" false")
	assert.Nil(t, err)
	assert.False(t, val.(bool))

	val, err = typ.getValue(" 0")
	assert.Nil(t, err)
	assert.False(t, val.(bool))

	val, err = typ.getValue("True")
	assert.NotNil(t, err)
	assert.False(t, val.(bool))

	val, err = typ.getValue("False")
	assert.NotNil(t, err)
	assert.False(t, val.(bool))

	val, err = typ.getValue("2")
	assert.NotNil(t, err)
	assert.False(t, val.(bool))
}

func Test_getValueDecimal(t *testing.T) {
	typ := xsdType("decimal")
	val, err := typ.getValue(" +5757676.99 ")
	assert.Nil(t, err)
	assert.Equal(t, int64(5757676), val.(int64))

	val, err = typ.getValue(" +5757676 ")
	assert.Nil(t, err)
	assert.Equal(t, int64(5757676), val.(int64))

	val, err = typ.getValue(" 5757676 ")
	assert.Nil(t, err)
	assert.Equal(t, int64(5757676), val.(int64))

	val, err = typ.getValue(" -092.8391 ")
	assert.Nil(t, err)
	assert.Equal(t, int64(-92), val.(int64))

	val, err = typ.getValue(" -92.8391 ")
	assert.Nil(t, err)
	assert.Equal(t, int64(-92), val.(int64))

	val, err = typ.getValue(" NaN ")
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), val.(int64))

	val, err = typ.getValue(" 92.8391. ")
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), val.(int64))

	val, err = typ.getValue(" +-92.8391 ")
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), val.(int64))
}

func Test_getValueFloat(t *testing.T) {
	typ := xsdType("float")
	val, err := typ.getValue("INF")
	assert.Nil(t, err)
	assert.True(t, math.IsInf(val.(float64), 1))

	val, err = typ.getValue("+INF")
	assert.Nil(t, err)
	assert.True(t, math.IsInf(val.(float64), 1))

	val, err = typ.getValue("-INF")
	assert.Nil(t, err)
	assert.True(t, math.IsInf(val.(float64), -1))

	val, err = typ.getValue("NaN")
	assert.Nil(t, err)
	assert.True(t, math.IsNaN(val.(float64)))

	_, err = typ.getValue("+NaN")
	assert.NotNil(t, err)

	_, err = typ.getValue("-NaN")
	assert.NotNil(t, err)

	val, err = typ.getValue(" +923e-1")
	assert.Nil(t, err)
	assert.Equal(t, float64(92.3), val)

	val, err = typ.getValue(" -9.23e1")
	assert.Nil(t, err)
	assert.Equal(t, float64(-92.3), val)

	val, err = typ.getValue(" -9.23E1")
	assert.Nil(t, err)
	assert.Equal(t, float64(-92.3), val)

	val, err = typ.getValue(" .3")
	assert.Nil(t, err)
	assert.Equal(t, float64(0.3), val)
	
	val, err = typ.getValue(" -+9.23e1")
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), val)
}

func Test_getValueDouble(t *testing.T) {
	typ := xsdType("float")
	val, err := typ.getValue("INF")
	assert.Nil(t, err)
	assert.True(t, math.IsInf(val.(float64), 1))

	val, err = typ.getValue("+INF")
	assert.Nil(t, err)
	assert.True(t, math.IsInf(val.(float64), 1))

	val, err = typ.getValue("-INF")
	assert.Nil(t, err)
	assert.True(t, math.IsInf(val.(float64), -1))

	val, err = typ.getValue("NaN")
	assert.Nil(t, err)
	assert.True(t, math.IsNaN(val.(float64)))

	_, err = typ.getValue("+NaN")
	assert.NotNil(t, err)

	_, err = typ.getValue("-NaN")
	assert.NotNil(t, err)

	val, err = typ.getValue(" +923e-1")
	assert.Nil(t, err)
	assert.Equal(t, float64(92.3), val)

	val, err = typ.getValue(" -9.23e1")
	assert.Nil(t, err)
	assert.Equal(t, float64(-92.3), val)

	val, err = typ.getValue(" -9.23E1")
	assert.Nil(t, err)
	assert.Equal(t, float64(-92.3), val)

	val, err = typ.getValue(" -+9.23e1")
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), val)
}

func Test_getValueDuration(t *testing.T) {
	typ := xsdType("duration")
	val, err := typ.getValue("P2Y6M5DT12H35M30S") //2 years, 6 months, 5 days, 12 hours, 35 minutes, 30 seconds
	assert.Nil(t, err)
	d:= val.(Duration)
	assert.False(t, d.negative)
	assert.Equal(t, 2, d.year)
	assert.Equal(t, 6, d.month)
	assert.Equal(t, 5, d.day)
	assert.Equal(t, 12, d.hour)
	assert.Equal(t, 35, d.minute)
	assert.Equal(t, 30*nano, d.second)
	assert.Equal(t, "P2Y6M5DT12H35M30S", d.String())

	val, err = typ.getValue("P1DT2H") //1 day, 2 hours
	assert.Nil(t, err)
	d= val.(Duration)
	assert.False(t, d.negative)
	assert.Equal(t, 0, d.year)
	assert.Equal(t, 0, d.month)
	assert.Equal(t, 1, d.day)
	assert.Equal(t, 2, d.hour)
	assert.Equal(t, 0, d.minute)
	assert.Equal(t, 0, d.second)
	assert.Equal(t, "P0Y0M1DT2H0M0S", d.String())

	val, err = typ.getValue("P20M") //20 months (the number of months can be more than 12)
	assert.Nil(t, err)
	d= val.(Duration)
	assert.False(t, d.negative)
	assert.Equal(t, 0, d.year)
	assert.Equal(t, 20, d.month)
	assert.Equal(t, 0, d.day)
	assert.Equal(t, 0, d.hour)
	assert.Equal(t, 0, d.minute)
	assert.Equal(t, 0, d.second)
	assert.Equal(t, "P0Y20M0DT0H0M0S", d.String())

	val, err = typ.getValue("PT20M") //20 minutes
	assert.Nil(t, err)
	d= val.(Duration)
	assert.False(t, d.negative)
	assert.Equal(t, 0, d.year)
	assert.Equal(t, 0, d.month)
	assert.Equal(t, 0, d.day)
	assert.Equal(t, 0, d.hour)
	assert.Equal(t, 20, d.minute)
	assert.Equal(t, 0, d.second)
	assert.Equal(t, "P0Y0M0DT0H20M0S", d.String())

	val, err = typ.getValue(" P0Y20M0D ") //20 months (0 is permitted as a number, but is not required)
	assert.Nil(t, err)
	d= val.(Duration)
	assert.False(t, d.negative)
	assert.Equal(t, 0, d.year)
	assert.Equal(t, 20, d.month)
	assert.Equal(t, 0, d.day)
	assert.Equal(t, 0, d.hour)
	assert.Equal(t, 0, d.minute)
	assert.Equal(t, 0, d.second)
	assert.Equal(t, "P0Y20M0DT0H0M0S", d.String())

	val, err = typ.getValue("P0Y") //0 years
	assert.Nil(t, err)
	d= val.(Duration)
	assert.False(t, d.negative)
	assert.Equal(t, 0, d.year)
	assert.Equal(t, 0, d.month)
	assert.Equal(t, 0, d.day)
	assert.Equal(t, 0, d.hour)
	assert.Equal(t, 0, d.minute)
	assert.Equal(t, 0, d.second)
	assert.Equal(t, "P0Y0M0DT0H0M0S", d.String())
	
	val, err = typ.getValue("-P60D") //minus 60 days
	assert.Nil(t, err)
	d= val.(Duration)
	assert.True(t, d.negative)
	assert.Equal(t, 0, d.year)
	assert.Equal(t, 0, d.month)
	assert.Equal(t, 60, d.day)
	assert.Equal(t, 0, d.hour)
	assert.Equal(t, 0, d.minute)
	assert.Equal(t, 0, d.second)
	assert.Equal(t, "-P0Y0M60DT0H0M0S", d.String())


	val, err = typ.getValue("PT1M30.050S") //1 minute, 30.5 seconds
	assert.Nil(t, err)
	d= val.(Duration)
	assert.False(t, d.negative)
	assert.Equal(t, 0, d.year)
	assert.Equal(t, 0, d.month)
	assert.Equal(t, 0, d.day)
	assert.Equal(t, 0, d.hour)
	assert.Equal(t, 1, d.minute)
	assert.Equal(t, int(30.05*nano), d.second)
	assert.Equal(t, "P0Y0M0DT0H1M30.05S", d.String())

	_, err = typ.getValue("P-20M")
	assert.NotNil(t, err)

	_, err = typ.getValue("P20MT")
	assert.NotNil(t, err)

	_, err = typ.getValue("P1YM5D")
	assert.NotNil(t, err)

	_, err = typ.getValue("P15.5Y")
	assert.NotNil(t, err)

	_, err = typ.getValue("P1D2H")
	assert.NotNil(t, err)

	_, err = typ.getValue("1Y2M")
	assert.NotNil(t, err)

	_, err = typ.getValue("P2M1Y")
	assert.NotNil(t, err)

	_, err = typ.getValue("P")
	assert.NotNil(t, err)

	_, err = typ.getValue("PT15.S")
	assert.NotNil(t, err)
}

func Test_getValueDateTime(t *testing.T) {
	typ := xsdType("dateTime")
	val, err := typ.getValue("2004-04-12T13:20:00") //1:20 pm on April 12, 2004
	assert.Nil(t, err)
	v := val.(time.Time)
	assert.Equal(t, 2004, v.Year())
	assert.Equal(t, time.Month(4), v.Month())
	assert.Equal(t, 12, v.Day())
	assert.Equal(t, 13, v.Hour())
	assert.Equal(t, 20, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())

	val, err = typ.getValue("2004-04-12T13:20:15.5") //1:20 pm and 15.5 seconds on April 12, 2004
	assert.Nil(t, err)
	v = val.(time.Time)
	assert.Equal(t, 15, v.Second())
	assert.Equal(t, int(0.5*nano), v.Nanosecond())
	
	val, err = typ.getValue("-2004-04-12T13:20:15.5") //1:20 pm and 15.5 seconds on April 12, 2004
	assert.Nil(t, err)
	v = val.(time.Time)
	assert.Equal(t, -2004, v.Year())
	
	_, err = typ.getValue("2004-04-12T13:20:00-05:00") //1:20 pm on April 12, 2004, US Eastern Standard Time
	assert.Nil(t, err)

	_, err = typ.getValue("2004-04-12T13:20:00Z") //1:20 pm on April 12, 2004, Coordinated Universal Time (UTC)
	assert.Nil(t, err)
	
	_, err = typ.getValue("2004-04-12T13:20:00+01:00") //1:20 pm on April 12, 2004, Coordinated Universal Time (UTC)
	assert.Nil(t, err)

	_, err = typ.getValue("2004-04-12T13:00") //seconds must be specified
	assert.NotNil(t, err)

	_, err = typ.getValue("2004-04-1213:20:00") //the letter T is required
	assert.NotNil(t, err)

	_, err = typ.getValue("99-04-12T13:00") //the century must not be left truncated
	assert.NotNil(t, err)

	_, err = typ.getValue("2004-04-12") //the time is required
	assert.NotNil(t, err)
}