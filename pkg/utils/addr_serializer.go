package utils

import (
	"context"
	"gorm.io/gorm/schema"
	"net"
	"reflect"
)

type AddrSerializer struct{}

func (AddrSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, val interface{}) error {
	parsed := net.ParseIP(val.(string))
	field.ReflectValueOf(ctx, dst).Set(reflect.ValueOf(parsed))
	return nil
}

func (AddrSerializer) Value(_ context.Context, _ *schema.Field, _ reflect.Value, val interface{}) (interface{}, error) {
	return val.(net.Addr).String(), nil
}
