package utility

import (
	"fmt"
	"reflect"
	"strconv"
)

type UnsupportedTypeError struct {
	Type string
	For  string
}

func (a *UnsupportedTypeError) Error() string {
	return fmt.Sprintf("Type %v is not supported for %v", a.Type, a.For)
}

type NotPointerError struct {
	Type string
}

func (a *NotPointerError) Error() string {
	return fmt.Sprintf("Type %v must be a pointer of this type", a.Type)
}

type BadTypeError struct {
	WantedType string
	GotType    string
}

func (a *BadTypeError) Error() string {
	return fmt.Sprintf("Error interface{} must be %v but is %v", a.WantedType, a.GotType)
}

// GetUnderlyingType retourne informations about the type
// behind the interface{} and the ptr ( must be )
func GetUnderlyingType(t interface{}) (string, reflect.Type, reflect.Value, error) {
	typeT := reflect.TypeOf(t)
	typeName := typeT.Name()
	values := reflect.ValueOf(t)
	if typeT.Kind() != reflect.Ptr {
		return typeName, typeT, values, &NotPointerError{
			Type: typeName,
		}
	}
	typeT = typeT.Elem()
	return typeT.Name(), typeT, values.Elem(), nil
}

func getFieldQuery(elem string, sf *reflect.StructField, fv *reflect.Value) error {
	switch fv.Type().Kind() {
	case reflect.String:
		fv.SetString(elem)
	case reflect.Int:
		i, err := strconv.Atoi(elem)
		if err != nil {
			return err
		}
		fv.SetInt(int64(i))
	}

	return nil
}

// QueryToStruct get the value of the struct from the value of the query
func QueryToStruct(q map[string][]string, t interface{}) error {
	_, typeT, values, err := GetUnderlyingType(t)
	if err != nil {
		return err
	}
	for i := 0; i < typeT.NumField(); i++ {
		ft := typeT.Field(i)
		fv := values.Field(i)
		if elems, ok := q[ft.Name]; ok {
			if fv.Type().Kind() == reflect.Array {

			} else {
				if len(elems) == 1 {
					err = getFieldQuery(elems[0], &ft, &fv)
					if err != nil {
						return err
					}
				} else {
					// erreur field unique
				}
			}
		} else {
			// valide si required
		}
	}
	return nil
}
