package errors

type ScreamerError string

const (
	ErrGameStop ScreamerError = "game stopped"
)

func (e ScreamerError) Error() string {
	return string(e)
}
