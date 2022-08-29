package model

import (
	"fmt"
	"time"

	// "os"
	"regexp"
	"strconv"
	"strings"
)

type xsdType string

const (
	TYPE_STRING       = "string"
	TYPE_BOOLEAN      = "boolean"
	TYPE_DECIMAL      = "decimal"
	TYPE_FLOAT        = "float"
	TYPE_DOUBLE       = "double"
	TYPE_DURATION     = "duration"
	TYPE_DATETIME     = "dateTime"
	TYPE_TIME         = "time"
	TYPE_DATE         = "date"
	TYPE_G_YEAR_MONTH = "gYearMonth"
	TYPE_G_YEAR       = "gYear"
	TYPE_G_MONTH_DAY  = "gMonthDay"
	TYPE_G_DAY        = "gDay"
	TYPE_G_MONTH      = "gMonth"
	TYPE_HEXBINARY    = "hexBinary"
	TYPE_HEX64BINARY  = "base64Binary"
	TYPE_ANYURI       = "anyURI"
	TYPE_QNAME        = "QName"
	TYPE_NOTATION     = "NOTATION"

	//derivative data type
	TYPE_NORMALIZED_STRING    = "normalizedString"
	TYPE_TOKEN                = "token"
	TYPE_LANGUAGE             = "language"
	TYPE_NMTOKEN              = "NMTOKEN"
	TYPE_NMTOKENS             = "NMTOKENS"
	TYPE_NAME                 = "Name"
	TYPE_NCNAME               = "NCName"
	TYPE_ID                   = "ID"
	TYPE_IDREF                = "IDREF"
	TYPE_IDREFS               = "IDREFS"
	TYPE_ENTITY               = "ENTITY"
	TYPE_ENTITIES             = "ENTITIES"
	TYPE_INTEGER              = "integer"
	TYPE_NON_POSITIVE_INTEGER = "nonPositiveInteger"
	TYPE_NEGATIVE_INTEGER     = "negativeInteger"
	TYPE_LONG                 = "long"
	TYPE_INT                  = "int"
	TYPE_SHORT                = "short"
	TYPE_BYTE                 = "byte"
	TYPE_NON_NEGATIVE_INTEGER = "nonNegativeInteger"
	TYPE_ULONG                = "unsignedLong"
	TYPE_UINT                 = "unsignedInt"
	TYPE_USHORT               = "unsignedShort"
	TYPE_UBYTE                = "unsignedByte"
	TYPE_POSITIVE_INTEGER     = "positiveInteger"
	TYPE_YEAR_MONTH_DURATION  = "yearMonthDuration"
	TYPE_DAY_TIME_DURATION    = "dayTimeDuration"
	TYPE_DATE_TIMESTAMP       = "dateTimeStamp"
)

type Duration struct {
	negative bool
	year     int
	month    int
	day      int
	hour     int
	minute   int
	second   int
}

func (d Duration) String() string {
	var ret string
	if d.negative {
		ret = "-"
	}
	ret += fmt.Sprintf("P%dY%dM%dDT%dH%dM", d.year, d.month, d.day, d.hour, d.minute)
	ret += fmt.Sprintf("%d", d.second/nano)
	
	rem := d.second%nano
	if rem!=0{
		ret+="."
		for mul := nano/10; rem!=0; mul/=10 {
			ret+=fmt.Sprintf("%d", rem/mul)
			rem %= mul 
		}
	}
	ret += "S"
	return ret
}

func (x xsdType) IsValid(s Schema) error {
	switch x {
	case TYPE_STRING, TYPE_BOOLEAN, TYPE_DECIMAL, TYPE_FLOAT, TYPE_DOUBLE, TYPE_DURATION, TYPE_DATETIME, TYPE_TIME, TYPE_DATE,
		TYPE_G_YEAR_MONTH, TYPE_G_YEAR, TYPE_G_MONTH_DAY, TYPE_G_DAY, TYPE_G_MONTH, TYPE_HEXBINARY, TYPE_HEX64BINARY, TYPE_ANYURI,
		TYPE_QNAME, TYPE_NOTATION, TYPE_NORMALIZED_STRING, TYPE_TOKEN, TYPE_LANGUAGE, TYPE_NMTOKEN, TYPE_NMTOKENS, TYPE_NAME,
		TYPE_NCNAME, TYPE_ID, TYPE_IDREF, TYPE_IDREFS, TYPE_ENTITY, TYPE_ENTITIES, TYPE_INTEGER, TYPE_NON_POSITIVE_INTEGER,
		TYPE_NEGATIVE_INTEGER, TYPE_LONG, TYPE_INT, TYPE_SHORT, TYPE_BYTE, TYPE_NON_NEGATIVE_INTEGER, TYPE_ULONG, TYPE_UINT,
		TYPE_USHORT, TYPE_UBYTE, TYPE_POSITIVE_INTEGER, TYPE_YEAR_MONTH_DURATION, TYPE_DAY_TIME_DURATION, TYPE_DATE_TIMESTAMP:
		return nil
	default:
		return fmt.Errorf("%s data type not supported yet", x)
	}
}

func (x xsdType) getValue(s string) (any, error) {
	switch x {
	case TYPE_STRING:
		return s, nil
	case TYPE_BOOLEAN:
		return getBooleanValue(s)
	case TYPE_DECIMAL:
		return getDecimalValue(s)
	case TYPE_FLOAT, TYPE_DOUBLE:
		return getFloatValue(s)
	case TYPE_DURATION:
		return getDurationValue(s)
	case TYPE_DATETIME:
		return getDateTime(s)

		// , , , , TYPE_TIME, TYPE_DATE,
		// TYPE_G_YEAR_MONTH, TYPE_G_YEAR, TYPE_G_MONTH_DAY, TYPE_G_DAY, TYPE_G_MONTH, TYPE_HEXBINARY, TYPE_HEX64BINARY, TYPE_ANYURI,
		// TYPE_QNAME, TYPE_NOTATION, TYPE_NORMALIZED_STRING, TYPE_TOKEN, TYPE_LANGUAGE, TYPE_NMTOKEN, TYPE_NMTOKENS, TYPE_NAME,
		// TYPE_NCNAME, TYPE_ID, TYPE_IDREF, TYPE_IDREFS, TYPE_ENTITY, TYPE_ENTITIES, TYPE_INTEGER, TYPE_NON_POSITIVE_INTEGER,
		// TYPE_NEGATIVE_INTEGER, TYPE_LONG, TYPE_INT, TYPE_SHORT, TYPE_BYTE, TYPE_NON_NEGATIVE_INTEGER, TYPE_ULONG, TYPE_UINT,
		// TYPE_USHORT, TYPE_UBYTE, TYPE_POSITIVE_INTEGER, TYPE_YEAR_MONTH_DURATION, TYPE_DAY_TIME_DURATION, TYPE_DATE_TIMESTAMP:
	}
	return s, nil
}

func getBooleanValue(v string) (bool, error) {
	v = strings.TrimSpace(v)
	
	switch v {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid value constraint value '%s', acceptable value is  'true' | 'false' | '1' | '0'", v)
	}
}

var patternDecimal = regexp.MustCompile(`^(\+|-)?([0-9]+(\.[0-9]*)?|\.[0-9]+)$`)

func getDecimalValue(v string) (int64, error) {
	v = strings.TrimSpace(v)
	tf, err := strconv.ParseFloat(v, 64)
	if err!=nil {
		return 0, fmt.Errorf("invalid value constraint value '%s', acceptable pattern is  '%s", v, patternDecimal.String())
	}
	
	if ok := patternDecimal.MatchString(v); !ok {
		return 0, fmt.Errorf("invalid value constraint value '%s', acceptable pattern is  '%s", v, patternDecimal.String())
	}
	
	return int64(tf), nil
}

func getFloatValue(v string) (float64, error) {
	v = strings.TrimSpace(v)

	if val, err := strconv.ParseFloat(v, 64); err != nil {
		return 0, fmt.Errorf("invalid value constraint value '%s'", v)
	} else {
		return val, err
	}
}

var patternDuration = regexp.MustCompile(`^-?P(((([0-9]+Y)([0-9]+M)?([0-9]+D)?|([0-9]+M)([0-9]+D)?|([0-9]+D))(T(([0-9]+H)([0-9]+M)?([0-9]+(\.[0-9]+)?S)?|([0-9]+M)([0-9]+(\.[0-9]+)?S)?|([0-9]+(\.[0-9]+)?S)))?)|(T(([0-9]+H)([0-9]+M)?([0-9]+(\.[0-9]+)?S)?|([0-9]+M)([0-9]+(\.[0-9]+)?S)?|([0-9]+(\.[0-9]+)?S))))$`)
var patternDurationY = regexp.MustCompile(`^[0-9]+Y$`)
var patternDurationM = regexp.MustCompile(`^[0-9]+M$`)
var patternDurationD = regexp.MustCompile(`^[0-9]+D$`)
var patternDurationTH = regexp.MustCompile(`^[0-9]+H$`)
var patternDurationTM = regexp.MustCompile(`^[0-9]+M$`)
var patternDurationTS = regexp.MustCompile(`^[0-9]+(\.[0-9]+)?S$`)

const nano = 1000000000

func getDurationValue(v string) (Duration, error) {
	v = strings.TrimSpace(v)
	if ok := patternDuration.MatchString(v); !ok {
		return Duration{}, fmt.Errorf("invalid value constraint value '%s', acceptable pattern is  '%s", v, patternDuration.String())
	}
	res := patternDuration.FindAllStringSubmatch(v, -1)
	var y, m, d, th, tm, ts string
	var iy, im, id, ith, itm, its int
	var inTime bool
	
	for _, x := range res[0] {
		if strings.HasPrefix(x, "T") {
			inTime = true
		}
		if !inTime {
			switch {
			case len(y) == 0 && patternDurationY.MatchString(x):
				y = x[:len(x)-1]
				iy, _ = strconv.Atoi(y)
			case len(m) == 0 && patternDurationM.MatchString(x):
				m = x[:len(x)-1]
				im, _ = strconv.Atoi(m)
			case len(d) == 0 && patternDurationD.MatchString(x):
				d = x[:len(x)-1]
				id, _ = strconv.Atoi(d)
			}
		} else {
			switch {
			case len(th) == 0 && patternDurationTH.MatchString(x):
				th = x[:len(x)-1]
				ith, _ = strconv.Atoi(th)
			case len(tm) == 0 && patternDurationTM.MatchString(x):
				tm = x[:len(x)-1]
				itm, _ = strconv.Atoi(tm)
			case len(ts) == 0 && patternDurationTS.MatchString(x):
				ts = x[:len(x)-1]
				fts, _ := strconv.ParseFloat(ts, 64)
				its = int(fts * nano)
			}
		}
	}

	return Duration{
		negative: strings.HasPrefix(v, "-"),
		year:     iy,
		month:    im,
		day:      id,
		hour:     ith,
		minute:   itm,
		second:   its,
	}, nil
}

var patterDateTime = regexp.MustCompile(`^-?([1-9][0-9]{3,}|0[0-9]{3})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T(([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?|(24:00:00(\.0+)?))(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00))?$`)
func getDateTime(v string) (time.Time, error){
	v = strings.TrimSpace(v)
	if ok := patterDateTime.MatchString(v); !ok {
		return time.Time{}, fmt.Errorf("invalid value constraint value '%s'", v)
	}
	
	res := patterDateTime.FindAllStringSubmatch(v, -1)
	if len(res[0][9])==0{
		v+="Z"
	}
	
	isBCE :=false
	if strings.HasPrefix(v, "-"){
		isBCE = true
		v = v[1:]
	}
	
	ret, err:=time.Parse(time.RFC3339Nano, v)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid value constraint value '%s'", v)
	}

	if isBCE{
		ret = ret.AddDate(-2*ret.Year(), 0, 0)
	}
	
	return ret, nil
}