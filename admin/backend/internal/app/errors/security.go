package errors

type ErrUnauth struct {
}

func (e *ErrUnauth) Error() string {
	return "Вы не авторизованы"
}
