package api

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	validate = validator.New()
)

func ParseBody(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return err
	}
	return validateStruct(obj)
}

func validateStruct(obj interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = indirect(v)
	}

	switch v.Type().Kind() {
	case reflect.Struct:
		return validate.Struct(obj)
	}

	// Simply ignore validation for other cases.
	return nil
}

func indirect(v reflect.Value) reflect.Value {
	// Recursively digs into v if the type of v is a pointer.
	for {
		if (v.Kind() == reflect.Interface && !v.IsNil()) || v.Kind() == reflect.Ptr {
			v = v.Elem()
			continue
		}
		return v
	}
}
