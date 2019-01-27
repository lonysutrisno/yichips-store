package exception

// ------------------------------------------------------
// ------------------------------------------------------
// INTERFACE
// ------------------------------------------------------

type Exception interface {
	GetCode() string
	GetMessage() string
}
