package types

type Item interface {
	IsFailed() bool
	GetMessage() string
	TableName() string
	Enrich()
}
