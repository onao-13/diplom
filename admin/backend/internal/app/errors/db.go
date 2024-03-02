package errors

type ErrNotFound struct {
}

func (e ErrNotFound) Error() string {
	return "Данные не найдены"
}
