package util

import (
	"testing"
	"time"
)

func TestMapFetcherRemove(t *testing.T) {
	data := NewMapfetcher(map[string]interface{}{
		"testString": "testing",
		"testInt":    3345678,
		"testBool":   true,
		"testFloat":  300.1,
	})
	expect := 4
	actual := data.Len()
	if expect != actual {
		t.Errorf("length expect:%v, actual:%v", expect, actual)
	}
	data.Remove("testString")

	expect = 3
	actual = data.Len()
	if expect != actual {
		t.Errorf("length expect:%v, actual:%v", expect, actual)
	}
	data.Remove("testInt")

	expect = 2
	actual = data.Len()
	if expect != actual {
		t.Errorf("length expect:%v, actual:%v", expect, actual)
	}
	data.Remove("testBool")

	expect = 1
	actual = data.Len()
	if expect != actual {
		t.Errorf("length expect:%v, actual:%v", expect, actual)
	}
	data.Remove("testFloat")

	expect = 0
	actual = data.Len()
	if expect != actual {
		t.Errorf("length expect:%v, actual:%v", expect, actual)
	}
}

func TestMapFetcherFetchs(t *testing.T) {
	stringKey := "testString"
	boolKey := "testBool"
	intKey := "testInt"
	floatKey := "testFloat"
	timeKey := "testTime"
	stringTimeKey := "testStringTime"

	expectString := "testing"
	expectInt := 3345678
	expectBool := true
	expectFloat := 300.1
	expectTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	stringTimeData := "2009-11-11 00:00:00"
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatal(err)
	}
	shortForm := "2006-01-02 15:04:05"
	expectStringTimeData, err := time.ParseInLocation(shortForm, stringTimeData, loc)
	if err != nil {
		t.Fatal(err)
	}

	mf := NewMapfetcher(map[string]interface{}{
		stringKey:     expectString,
		intKey:        expectInt,
		boolKey:       expectBool,
		floatKey:      expectFloat,
		timeKey:       expectTime,
		stringTimeKey: stringTimeData,
	})
	actualString := mf.FetchString(stringKey)
	if expectString != actualString {
		t.Errorf("expect:%v, actual:%v", expectString, actualString)
	}
	actualInt := mf.FetchInt(intKey)
	if expectInt != actualInt {
		t.Errorf("expect:%v, actual:%v", expectInt, actualInt)
	}
	actualBool := mf.FetchBool(boolKey)
	if expectBool != actualBool {
		t.Errorf("expect:%v, actual:%v", expectBool, actualBool)
	}
	actualFloat := mf.FetchFloat(floatKey)
	if expectFloat != actualFloat {
		t.Errorf("expect:%v, actual:%v", expectFloat, actualFloat)
	}
	actualTime := mf.FetchTime(timeKey)
	if expectTime != actualTime {
		t.Errorf("expect:%v, actual:%v", expectTime, actualTime)
	}
	actualStringTime := mf.FetchTimeFromString(stringTimeKey, "UTC")
	if expectStringTimeData != actualStringTime {
		t.Errorf("expect:%v, actual:%v", expectTime, actualTime)
	}
	if mf.Err() != nil {
		t.Error(mf.Err())
	}
}

func TestMapFetcherFetchAfterError(t *testing.T) {
	stringKey := "testString"
	boolKey := "testBool"
	intKey := "testInt"
	floatKey := "testFloat"
	timeKey := "testTime"
	wrongStringKey := "wrongStringKey"

	mf := NewMapfetcher(map[string]interface{}{
		stringKey: "testing",
		intKey:    3345678,
		boolKey:   true,
		floatKey:  300.1,
		timeKey:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	})
	mf.FetchString(wrongStringKey)

	var expectString string
	actualString := mf.FetchString(stringKey)
	if expectString != actualString {
		t.Errorf("expect:%v, actual:%v", expectString, actualString)
	}

	var expectBool bool
	actualBool := mf.FetchBool(boolKey)
	if expectBool != actualBool {
		t.Errorf("expect:%v, actual:%v", expectBool, actualBool)
	}

	var expectFloat float64
	actualFloat := mf.FetchFloat(floatKey)
	if expectFloat != actualFloat {
		t.Errorf("expect:%v, actual:%v", expectFloat, actualFloat)
	}

	var expectTime time.Time
	actualTime := mf.FetchTime(timeKey)
	if expectTime != actualTime {
		t.Errorf("expect:%v, actual:%v", expectTime, actualTime)
	}
}

func TestMapFetcherFetchTimeFromStringError(t *testing.T) {
	stringTimeKey := "testStringTime"

	mf := NewMapfetcher(map[string]interface{}{
		stringTimeKey: "2009u-10o-102 00:002:00",
	})

	var expectStringTime time.Time

	actualStringTime := mf.FetchTimeFromString(stringTimeKey, "HONGKONG")
	if expectStringTime != actualStringTime {
		t.Errorf("expect:%v, actual:%v", expectStringTime, actualStringTime)
	}
	mf.ClearErr()

	actualStringTime = mf.FetchTimeFromString(stringTimeKey, "UTC")
	if expectStringTime != actualStringTime {
		t.Errorf("expect:%v, actual:%v", expectStringTime, actualStringTime)
	}
	mf.ClearErr()
}

func TestMapFetcherKeyNotExist(t *testing.T) {
	correctStringKey := "correctStringKey"
	correctBoolKey := "correctBoolKey"
	correctIntKey := "correctIntKey"
	correctFloatKey := "correctFloatKey"
	correctTimeKey := "correctTimeKey"
	correctStringTimeKey := "correctStringTimeKey"

	wrongStringKey := "wrongStringKey"
	wrongBoolKey := "wrongBoolKey"
	wrongIntKey := "wrongIntKey"
	wrongFloatKey := "wrongFloatKey"
	wrongTimeKey := "wrongTimeKey"
	wrongStringTimeKey := "wrongStringTimeKey"

	mf := NewMapfetcher(map[string]interface{}{
		correctStringKey:     "testing",
		correctIntKey:        3345678,
		correctBoolKey:       true,
		correctFloatKey:      300.1,
		correctTimeKey:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		correctStringTimeKey: "2019-11-11 00:00:00",
	})
	mf.FetchString(wrongStringKey)
	expect := ErrMapFetcherKeyNotExist
	actual := mf.Err()
	if expect != actual {
		t.Errorf("%v expect:%v, actual:%v", wrongStringKey, expect, actual)
	}
	mf.ClearErr()

	mf.FetchBool(wrongBoolKey)
	actual = mf.Err()
	if expect != actual {
		t.Errorf("%v expect:%v, actual:%v", wrongBoolKey, expect, actual)
	}
	mf.ClearErr()

	mf.FetchInt(wrongIntKey)
	actual = mf.Err()
	if expect != actual {
		t.Errorf("%v expect:%v, actual:%v", wrongIntKey, expect, actual)
	}
	mf.ClearErr()

	mf.FetchFloat(wrongFloatKey)
	actual = mf.Err()
	if expect != actual {
		t.Errorf("%v expect:%v, actual:%v", wrongFloatKey, expect, actual)
	}
	mf.ClearErr()

	mf.FetchTime(wrongTimeKey)
	actual = mf.Err()
	if expect != actual {
		t.Errorf("%v expect:%v, actual:%v", wrongTimeKey, expect, actual)
	}
	mf.ClearErr()

	mf.FetchTimeFromString(wrongTimeKey, "UTC")
	actual = mf.Err()
	if expect != actual {
		t.Errorf("%v expect:%v, actual:%v", wrongStringTimeKey, expect, actual)
	}
	mf.ClearErr()
}

func TestMapFetcherAssertionFaild(t *testing.T) {
	stringKey := "testString"
	boolKey := "testBool"
	intKey := "testInt"
	floatKey := "testFloat"
	timeKey := "testTime"
	stringTimeKey := "testStringTime"

	mf := NewMapfetcher(map[string]interface{}{
		stringKey:     struct{}{},
		intKey:        struct{}{},
		boolKey:       struct{}{},
		floatKey:      struct{}{},
		timeKey:       struct{}{},
		stringTimeKey: struct{}{},
	})
	mf.FetchString(stringKey)
	expect := ErrMapFetcherFaildAssertionStringType
	actual := mf.Err()
	if expect != actual {
		t.Errorf("expect:%v, actual:%v", expect, actual)
	}
	mf.ClearErr()

	mf.FetchBool(boolKey)
	expect = ErrMapFetcherFaildAssertionBoolType
	actual = mf.Err()
	if expect != actual {
		t.Errorf("expect:%v, actual:%v", expect, actual)
	}
	mf.ClearErr()

	mf.FetchInt(intKey)
	expect = ErrMapFetcherFaildAssertionIntType
	actual = mf.Err()
	if expect != actual {
		t.Errorf("expect:%v, actual:%v", expect, actual)
	}
	mf.ClearErr()

	mf.FetchFloat(floatKey)
	expect = util.ErrMapFetcherFaildAssertionFloatType
	actual = mf.Err()
	if expect != actual {
		t.Errorf("expect:%v, actual:%v", expect, actual)
	}
	mf.ClearErr()

	mf.FetchTime(timeKey)
	expect = ErrMapFetcherFaildAssertionTimeType
	actual = mf.Err()
	if expect != actual {
		t.Errorf("expect:%v, actual:%v", expect, actual)
	}
	mf.ClearErr()

	mf.FetchTimeFromString(stringTimeKey, "UTC")
	expect = ErrMapFetcherFaildAssertionStringType
	actual = mf.Err()
	if expect != actual {
		t.Errorf("expect:%v, actual:%v", expect, actual)
	}
	mf.ClearErr()
}

func TestMapFetherActions(t *testing.T) {
	expectData := map[string]interface{}{
		"testString": "testing",
		"testInt":    3345678,
		"testBool":   true,
		"testFloat":  300.1,
	}
	actualData := map[string]interface{}{}
	expect := NewMapfetcher(expectData)
	actual := NewMapfetcher(actualData)
	actual.CopyMap(expect)
	if actual.Len() != expect.Len() {
		t.Errorf("CopyMap length expect:%v, actual:%v", expect.Len(), actual.Len())
	}
	expect.Add("testHaHa", false)
	actual.CopyMap(expect)
	if actual.Len() != expect.Len() {
		t.Errorf("Add length expect:%v, actual:%v", expect.Len(), actual.Len())
	}
	if actual.Err() != nil {
		t.Errorf("actual MapFetcher error:%v", actual.Err())
	}
	if expect.Err() != nil {
		t.Errorf("expect MapFetcher error:%v", expect.Err())
	}
}
