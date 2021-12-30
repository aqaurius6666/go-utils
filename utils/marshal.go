package utils

import (
	"reflect"
	"strconv"

	"golang.org/x/xerrors"
)

// MarshalOffsetLimit take `Offset`, `Limit` fields of src that have string types, and apply them to `Offset`, `Limit` fields of dest that have int types.
// if `Offset` ,`Limit` has not been defined, it will be recieved 0, 10, respectively.
func MarshalOffsetLimit(dest interface{}, src interface{}) error {
	var err error
	var destValue, srcValue, offsetValue, limitValue reflect.Value
	var destType, srcType, offsetType, limitType reflect.Type
	var offset, limit string
	var offsetI, limitI int
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return xerrors.New("dest must be pointer")
	}
	if reflect.TypeOf(dest).Elem().Kind() != reflect.Struct {
		return xerrors.New("dest value must be struct")
	}
	destType = reflect.TypeOf(dest).Elem()
	destValue = reflect.ValueOf(dest).Elem()
	if reflect.TypeOf(src).Kind() == reflect.Ptr {
		if reflect.TypeOf(src).Elem().Kind() != reflect.Struct {
			return xerrors.New("src value must be struct")
		}
		srcValue = reflect.ValueOf(src).Elem()
		srcType = reflect.TypeOf(src).Elem()
	} else {
		if reflect.TypeOf(src).Kind() != reflect.Struct {
			return xerrors.New("src value must be struct")
		}
		srcValue = reflect.ValueOf(src)
		srcType = reflect.TypeOf(src)
	}
	if reflect.TypeOf(dest).Kind() == reflect.Ptr {
		if reflect.TypeOf(dest).Elem().Kind() != reflect.Struct {
			return xerrors.New("dest value must be struct")
		}
	} else {
		if reflect.TypeOf(dest).Kind() != reflect.Struct {
			return xerrors.New("dest value must be struct")
		}
	}

	// Hanle Offset
	offsetValue = srcValue.FieldByName("Offset")
	offsetStructField, ok := srcType.FieldByName("Offset")
	if !ok {
		return xerrors.New("missing offset field")
	}
	if offsetStructField.Type.Kind() == reflect.Ptr {
		offsetType = offsetStructField.Type.Elem()
	} else {
		offsetType = offsetStructField.Type
	}
	if offsetType.Kind() != reflect.String {
		return xerrors.New("offset is not string type")
	}
	offset = offsetValue.String()
	if offsetI, err = strconv.Atoi(offset); err != nil {
		return xerrors.New("offset invalid")
	}
	// Hanle Limit
	limitValue = srcValue.FieldByName("Limit")
	limitStructField, ok := srcType.FieldByName("Limit")
	if !ok {
		return xerrors.New("missing limit field")
	}
	if limitStructField.Type.Kind() == reflect.Ptr {
		limitType = limitStructField.Type.Elem()
	} else {
		limitType = limitStructField.Type
	}
	if limitType.Kind() != reflect.String {
		return xerrors.New("limit is not string type")
	}
	limit = limitValue.String()
	if limitI, err = strconv.Atoi(limit); err != nil {
		return xerrors.New("limit invalid")
	}
	// Handle dest

	if structField, ok := destType.FieldByName("Offset"); !ok {
		return xerrors.New("dest missing offset field")
	} else {
		offsetType = structField.Type
		offsetValue = destValue.FieldByName("Offset")
	}
	if offsetType.Kind() != reflect.Int {
		return xerrors.New("offset limit not int type")
	}
	if structField, ok := destType.FieldByName("Limit"); !ok {
		return xerrors.New("dest missing limit field")
	} else {
		limitType = structField.Type
		limitValue = destValue.FieldByName("Limit")
	}
	if limitType.Kind() != reflect.Int {
		return xerrors.New("dest limit not int type")
	}
	if !limitValue.CanSet() {
		return xerrors.New("limit field can not be set")
	}
	limitValue.Set(reflect.ValueOf(limitI))

	if !offsetValue.CanSet() {
		return xerrors.New("offset field can not be set")
	}
	offsetValue.Set(reflect.ValueOf(offsetI))
	return nil
}
