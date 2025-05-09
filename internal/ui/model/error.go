package model

func NewErrorModel() *ErrorModel {
	return &ErrorModel{}
}

type ErrorModel struct {
	err error
}

func (m *ErrorModel) Error() error {
	return m.err
}

func (m *ErrorModel) SetError(err error) {
	m.err = err
}
