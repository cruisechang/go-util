

/*
func (u *User) GetFieldInt(fieldName string) (resValue int, resErr error) {
	defer func() {
		u.mutex.Unlock()
		if r := recover(); r != nil {
			resErr = errors.New("GetFieldInt panic")
		}
	}()

	u.mutex.Lock()
	r := reflect.ValueOf(u)
	f := reflect.Indirect(r).FieldByName(fieldName)
	if f.IsValid() && f.Kind() == reflect.Int {
		resValue = int(f.Int())
	} else {
		resErr = errors.New("GetIntField has no field:" + fieldName)
	}
	return
}
func (u *User) GetFieldFloat64(fieldName string) (resValue float64, resErr error) {
	defer func() {
		u.mutex.Unlock()
		if r := recover(); r != nil {
			resErr = errors.New("GetFieldFloat64 panic")
		}
	}()

	u.mutex.Lock()
	r := reflect.ValueOf(u)
	f := reflect.Indirect(r).FieldByName(fieldName)
	if f.IsValid() && f.Kind() == reflect.Float64 {
		resValue = f.Float()
	} else {
		resErr = errors.New("GetIntField has no field:" + fieldName)
	}
	return
}
func (u *User) GetFieldBool(fieldName string) (resValue bool, resErr error) {
	defer func() {
		u.mutex.Unlock()
		if r := recover(); r != nil {
			resErr = errors.New("GetFieldBool panic")
		}
	}()

	u.mutex.Lock()
	r := reflect.ValueOf(u)
	f := reflect.Indirect(r).FieldByName(fieldName)
	if f.IsValid() && f.Kind() == reflect.Bool {
		resValue = f.Bool()
	} else {
		resErr = errors.New("GetIntField has no field:" + fieldName)
	}
	return
}
func (u *User) GetFieldString(fieldName string) (resValue string, resErr error) {
	defer func() {
		u.mutex.Unlock()
		if r := recover(); r != nil {
			resErr = errors.New("GetFieldString panic")
		}
	}()

	u.mutex.Lock()
	r := reflect.ValueOf(u)
	f := reflect.Indirect(r).FieldByName(fieldName)
	if f.IsValid() && f.Kind() == reflect.String {
		resValue = f.String()
	} else {
		resErr = errors.New("GetIntField has no field:" + fieldName)
	}
	return
}

func (u *User) SetField(fieldName string, value interface{}) (resErr error) {
	defer func() {
		u.mutex.Unlock()
		if r := recover(); r != nil {
			resErr = errors.New("SetField panic")
		}
	}()

	u.mutex.Lock()
	r := reflect.ValueOf(u)
	f := reflect.Indirect(r).FieldByName(fieldName)
	if f.IsValid() {
		return errors.New("filed not exist")
	}

	return nil
}
*/