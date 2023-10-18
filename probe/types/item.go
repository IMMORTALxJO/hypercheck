package types

type Item interface {
	IsFailed() bool
	GetMessage() string
	Enrich()
}
