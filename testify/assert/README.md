# assert


## func New
```go
func New(t TestingT) *Assertions
```
 Assertions provides assertion methods around the TestingT interface.


## func Contains
``` go
func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) bool
```
 Contains asserts that the specified string, list(array, slice...) or map contains the
 specified substring or element.

    a.Contains("Hello World", "World", "But 'Hello World' does contain 'World'")
    a.Contains(["Hello", "World"], "World", "But ["Hello", "World"] does contain 'World'")
    a.Contains({"Hello": "World"}, "Hello", "But {'Hello': 'World'} does contain 'Hello'")

 Returns whether the assertion was successful (true) or not (false).


## func Empty
 ``` go
func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) bool
 ```
 Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
 a slice or a channel with len == 0.

  a.Empty(obj)

 Returns whether the assertion was successful (true) or not (false).


## func Equal
``` go
func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```
 Equal asserts that two objects are equal.

    a.Equal(123, 123, "123 and 123 should be equal")

 Returns whether the assertion was successful (true) or not (false).

 Pointer variable equality is determined based on the equality of the
 referenced values (as opposed to the memory addresses).


## func EqualError
``` go
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) bool
```
 EqualError asserts that a function returned an error (i.e. not `nil`)
 and that it is equal to the provided error.

   actualObj, err := SomeFunction()
   a.EqualError(err,  expectedErrorString, "An error was expected")

 Returns whether the assertion was successful (true) or not (false).


## func EqualValues
``` go
func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```
 EqualValues asserts that two objects are equal or convertable to the same types
 and equal.

    a.EqualValues(uint32(123), int32(123), "123 and 123 should be equal")

 Returns whether the assertion was successful (true) or not (false).


## func Error
``` go
func (a *Assertions) Error(err error, msgAndArgs ...interface{}) bool
```
 Error asserts that a function returned an error (i.e. not `nil`).

   actualObj, err := SomeFunction()
   if a.Error(err, "An error was expected") {
 	   assert.Equal(t, err, expectedError)
   }

 Returns whether the assertion was successful (true) or not (false).


## func Exactly
``` go
func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```
 Exactly asserts that two objects are equal is value and type.

    a.Exactly(int32(123), int64(123), "123 and 123 should NOT be equal")

 Returns whether the assertion was successful (true) or not (false).


## func FailNow
``` go
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{}) bool
```
 FailNow fails test


## func Fail
 ``` go
 func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) bool
 ```  
 Fail reports a failure through


## func False
``` go
func (a *Assertions) False(value bool, msgAndArgs ...interface{}) bool
```
 False asserts that the specified value is false.

    a.False(myBool, "myBool should be false")

 Returns whether the assertion was successful (true) or not (false).


## func HTTPBodyContains
``` go
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}) bool
```
 HTTPBodyContains asserts that a specified handler returns a
 body that contains a string.

  a.HTTPBodyContains(myHandler, "www.google.com", nil, "I'm Feeling Lucky")

 Returns whether the assertion was successful (true) or not (false).


## func HTTPBodyNotContains
``` go
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}) bool
```
 HTTPBodyNotContains asserts that a specified handler returns a
 body that does not contain a string.

  a.HTTPBodyNotContains(myHandler, "www.google.com", nil, "I'm Feeling Lucky")

 Returns whether the assertion was successful (true) or not (false).


## func HTTPError
``` go
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values) bool
```
 HTTPError asserts that a specified handler returns an error status code.

  a.HTTPError(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}

 Returns whether the assertion was successful (true) or not (false).


## func HTTPRedirect
``` go
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values) bool
```
 HTTPRedirect asserts that a specified handler returns a redirect status code.

  a.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}

 Returns whether the assertion was successful (true) or not (false).


## func HTTPSuccess
``` go
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values) bool
```
 HTTPSuccess asserts that a specified handler returns a success status code.

  a.HTTPSuccess(myHandler, "POST", "http:www.google.com", nil)

 Returns whether the assertion was successful (true) or not (false).


## func Implements
``` go
func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool
```
 Implements asserts that an object is implemented by the specified interface.


## func InDelta
``` go
func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool
```
 InDelta asserts that the two numerals are within delta of each other.

 	 a.InDelta(math.Pi, (22 / 7.0), 0.01)

 Returns whether the assertion was successful (true) or not (false).


## func InDeltaSlice
``` go
func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool
```
 InDeltaSlice is the same as InDelta, except it compares two slices.


## func InEpsilon
``` go
func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool
```
 InEpsilon asserts that expected and actual have a relative error less than epsilon

 Returns whether the assertion was successful (true) or not (false).


## func InEpsilonSlice
``` go
func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool
```
 InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.


## func IsType
``` go
func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool
```
 IsType asserts that the specified objects are of the same type.


## func JSONEq
``` go
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) bool
```
 JSONEq asserts that two JSON strings are equivalent.

  a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)

 Returns whether the assertion was successful (true) or not (false).


## func Len
``` go
func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) bool
```
 Len asserts that the specified object has specific length.
 Len also fails if the object has a type that len() not accept.

    a.Len(mySlice, 3, "The size of slice is not 3")

 Returns whether the assertion was successful (true) or not (false).


## func Nil
``` go
func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) bool
```
 Nil asserts that the specified object is nil.

    a.Nil(err, "err should be nothing")

 Returns whether the assertion was successful (true) or not (false).


## func NoError
``` go
func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) bool
```

 NoError asserts that a function returned no error (i.e. `nil`).

   actualObj, err := SomeFunction()
   if a.NoError(err) {
 	   assert.Equal(t, actualObj, expectedObj)
   }

 Returns whether the assertion was successful (true) or not (false).


## func NotContains
``` go
func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) bool
```
 NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
 specified substring or element.

    a.NotContains("Hello World", "Earth", "But 'Hello World' does NOT contain 'Earth'")
    a.NotContains(["Hello", "World"], "Earth", "But ['Hello', 'World'] does NOT contain 'Earth'")
    a.NotContains({"Hello": "World"}, "Earth", "But {'Hello': 'World'} does NOT contain 'Earth'")

 Returns whether the assertion was successful (true) or not (false).


## func NotEmpty
``` go
func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) bool
```
 NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
 a slice or a channel with len == 0.

  if a.NotEmpty(obj) {
    assert.Equal(t, "two", obj[1])
  }

 Returns whether the assertion was successful (true) or not (false).


## func NotEqual
``` go
func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```
 NotEqual asserts that the specified values are NOT equal.

    a.NotEqual(obj1, obj2, "two objects shouldn't be equal")

 Returns whether the assertion was successful (true) or not (false).

 Pointer variable equality is determined based on the equality of the
 referenced values (as opposed to the memory addresses).


## func NotNil
``` go
func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) bool
```
 NotNil asserts that the specified object is not nil.

    a.NotNil(err, "err should be something")

 Returns whether the assertion was successful (true) or not (false).


## func NotPanics
``` go
func (a *Assertions) NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) bool
```
 NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.

   a.NotPanics(func(){
     RemainCalm()
   }, "Calling RemainCalm() should NOT panic")

 Returns whether the assertion was successful (true) or not (false).


## func NotRegexp
``` go
func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool
```
 NotRegexp asserts that a specified regexp does not match a string.

  a.NotRegexp(regexp.MustCompile("starts"), "it's starting")
  a.NotRegexp("^start", "it's not starting")

 Returns whether the assertion was successful (true) or not (false).


## func NotZero
``` go
func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) bool
```
 NotZero asserts that i is not the zero value for its type and returns the truth.


## func Panics
``` go
func (a *Assertions) Panics(f PanicTestFunc, msgAndArgs ...interface{}) bool
```
 Panics asserts that the code inside the specified PanicTestFunc panics.

   a.Panics(func(){
     GoCrazy()
   }, "Calling GoCrazy() should panic")

 Returns whether the assertion was successful (true) or not (false).


## func Regexp
``` go
func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool
```
 Regexp asserts that a specified regexp matches a string.

  a.Regexp(regexp.MustCompile("start"), "it's starting")
  a.Regexp("start...$", "it's not starting")

 Returns whether the assertion was successful (true) or not (false).


## func True
``` go
func (a *Assertions) True(value bool, msgAndArgs ...interface{}) bool
```
 True asserts that the specified value is true.

    a.True(myBool, "myBool should be true")

 Returns whether the assertion was successful (true) or not (false).


## func WithinDuration
``` go
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool
```
 WithinDuration asserts that the two times are within duration delta of each other.

   a.WithinDuration(time.Now(), time.Now(), 10*time.Second, "The difference should not be more than 10s")

 Returns whether the assertion was successful (true) or not (false).


## func Zero
``` go
func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) bool
```
 Zero asserts that i is the zero value for its type and returns the truth.
