package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Check func(v interface{}) error

func parseTag(tag string) (key string, value string) {
	splited := strings.Split(tag, ":")
	if len(splited) != 2 {
		return
	}
	key = splited[0]
	value = splited[1]
	return
}

func Unpack(req *http.Request, ptr interface{}, checks map[string]Check) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		splited := strings.Split(string(tag), ",")
		key, value := parseTag(splited[0])
		if key != "http" {
			break
		}
		name := value
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		key, value = parseTag(splited[1])
		if key != "check" {
			break
		}
		checkName := value
		if checkName != "" {
			if check, ok := checks[checkName]; ok {
				fmt.Println(v.Field(i).Interface())
				err := check(v.Field(i).Interface())
				if err != nil {
					return err
				}
			}
		}
		fields[name] = v.Field(i)
	}

	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}