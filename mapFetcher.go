package util

import (
	"errors"
	"time"
)

var (
	// ErrMapFetcherKeyNotExist means the input key is not exist.
	ErrMapFetcherKeyNotExist = errors.New("map fetcher: key not exist")
	// ErrMapFetcherFaildAssertionTimeType means the time.Time type assertion faild.
	ErrMapFetcherFaildAssertionTimeType = errors.New("time.Time type assertion failed")
	// ErrMapFetcherFaildAssertionFloatType means the float64 type assertion faild.
	ErrMapFetcherFaildAssertionFloatType = errors.New("float64 type assertion failed")
	// ErrMapFetcherFaildAssertionStringType means the string type assertion faild.
	ErrMapFetcherFaildAssertionStringType = errors.New("string type assertion failed")
	// ErrMapFetcherFaildAssertionIntType means the int type assertion faild.
	ErrMapFetcherFaildAssertionIntType = errors.New("int type assertion failed")
	// ErrMapFetcherFaildAssertionBoolType means the bool type assertion faild.
	ErrMapFetcherFaildAssertionBoolType = errors.New("bool type assertion failed")
)

// MapFetcher is the wrapper for map[string]interface{},
// and provides actions for easy usage of map actions.
type MapFetcher struct {
	err  error
	data map[string]interface{}
}

// NewMapfetcher returns MapFetcher instance.
func NewMapfetcher(data map[string]interface{}) *MapFetcher {
	return &MapFetcher{
		data: data,
	}
}

// Err returns MapFetcher holds err, it is using this article's thoughts:
// https://blog.golang.org/errors-are-values
func (mf *MapFetcher) Err() error {
	return mf.err
}

// ClearErr clears the error.
func (mf *MapFetcher) ClearErr() {
	mf.err = nil
}

// Fetch returns interface{} in the holding map by input key.
func (mf *MapFetcher) Fetch(key string) interface{} {
	if mf.Err() != nil {
		return nil
	}
	v, ok := mf.data[key]
	if !ok {
		mf.err = ErrMapFetcherKeyNotExist
		return nil
	}
	return v
}

// CopyMap copies a MapFetcher's data map into self data map.
func (mf *MapFetcher) CopyMap(from *MapFetcher) {
	for k, v := range from.GetMap() {
		mf.data[k] = v
	}
}

// Len returns the self data map length.
func (mf *MapFetcher) Len() int {
	return len(mf.data)
}

// GetMap returns whole map.
func (mf *MapFetcher) GetMap() map[string]interface{} {
	return mf.data
}

// Add adds key value into map.
func (mf *MapFetcher) Add(key string, value interface{}) {
	mf.data[key] = value
}

// Remove removes a key from map.
func (mf *MapFetcher) Remove(key string) {
	delete(mf.data, key)
}

// FetchTimeFromString converts specific key value from string time into time.Time then returns.
func (mf *MapFetcher) FetchTimeFromString(key, location string) time.Time {
	value := mf.Fetch(key)
	if value == nil {
		return time.Time{}
	}
	sv, ok := value.(string)
	if !ok {
		mf.err = ErrMapFetcherFaildAssertionStringType
		return time.Time{}
	}
	loc, err := time.LoadLocation(location)
	if err != nil {
		mf.err = err
		return time.Time{}
	}
	shortForm := "2006-01-02 15:04:05"
	v, err := time.ParseInLocation(shortForm, sv, loc)
	if err != nil {
		mf.err = err
		return time.Time{}
	}
	return v
}

// FetchTime converts specific key value into time.Time then returns.
func (mf *MapFetcher) FetchTime(key string) time.Time {
	value := mf.Fetch(key)
	if value == nil {
		return time.Time{}
	}
	v, ok := value.(time.Time)
	if !ok {
		mf.err = ErrMapFetcherFaildAssertionTimeType
	}
	return v
}

// FetchFloat converts specific key value into float64 then returns.
func (mf *MapFetcher) FetchFloat(key string) float64 {
	value := mf.Fetch(key)
	if value == nil {
		return 0
	}
	v, ok := value.(float64)
	if !ok {
		mf.err = ErrMapFetcherFaildAssertionFloatType
		return 0
	}
	return v
}

// FetchString converts specific key value into string then returns.
func (mf *MapFetcher) FetchString(key string) string {
	value := mf.Fetch(key)
	if value == nil {
		return ""
	}
	v, ok := value.(string)
	if !ok {
		mf.err = ErrMapFetcherFaildAssertionStringType
		return ""
	}
	return v
}

// FetchInt converts specific key value into int then returns.
func (mf *MapFetcher) FetchInt(key string) int {
	value := mf.Fetch(key)
	if value == nil {
		return 0
	}
	v, ok := value.(int)
	if !ok {
		mf.err = ErrMapFetcherFaildAssertionIntType
		return 0
	}
	return v
}

// FetchBool converts specific key value into boolean then returns.
func (mf *MapFetcher) FetchBool(key string) bool {
	value := mf.Fetch(key)
	if value == nil {
		return false
	}
	v, ok := value.(bool)
	if !ok {
		mf.err = ErrMapFetcherFaildAssertionBoolType
		return false
	}
	return v
}
