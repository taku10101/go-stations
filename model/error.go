package model

type ErrNotFound struct {
	Err string
}

func (err *ErrNotFound) Error() string {
	return err.Err
}
