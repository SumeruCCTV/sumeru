package pkg

import (
	"errors"
	"github.com/SumeruCCTV/sumeru/pkg/utils"
	"github.com/SumeruCCTV/sumeru/service"
	"reflect"
	"unsafe"
)

// TODO: This needs to be optimized, it's currently very inefficient.
// TODO: Instead of using fieldName, just find the field that has the type of the service.

func (app *Application) injectFields(name string, svc service.Service) error {
	svcRef := reflect.ValueOf(svc)
	if !svcRef.IsValid() || svcRef.Kind() != reflect.Ptr {
		return errors.New("invalid service")
	}
	for _, _svc := range app.services {
		if err := _injectField(name, svcRef, _svc); err != nil {
			return err
		}
	}
	for fieldName, fieldRef := range app.injectables {
		if err := _injectField(fieldName, reflect.ValueOf(fieldRef), svc); err != nil {
			return err
		}
	}
	return nil
}

func _injectField(fieldName string, fieldRef reflect.Value, svc service.Service) error {
	ref := reflect.ValueOf(svc)
	if !ref.IsValid() || ref.Kind() != reflect.Ptr {
		return errors.New("invalid service: invalid or not a pointer")
	}
	ref = reflect.Indirect(ref)
	if !ref.IsValid() || ref.Kind() != reflect.Struct {
		return errors.New("invalid service: invalid or not a struct")
	}
	v := ref.FieldByName(fieldName)
	if !v.IsValid() {
		return nil
	}
	if v.Kind() != reflect.Ptr {
		return errors.New("invalid service: not a pointer")
	}
	if !v.CanSet() {
		v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
	}
	v.Set(fieldRef)
	return nil
}

func (app *Application) createLogger(svc service.Service) *utils.Logger {
	return app.log.Named(svc.Name())
}
