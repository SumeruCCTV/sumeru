package errors

import (
	"github.com/SumeruCCTV/sumeru/pkg/json"
	"github.com/bytedance/sonic"
)

type WebError []byte

func (err WebError) Error() string {
	return string(err)
}

func New(err error) error {
	bytes, err := sonic.Marshal(json.Error(err.Error()))
	if err != nil {
		return err
	}
	return WebError(bytes)
}
