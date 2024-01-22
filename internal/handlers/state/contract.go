package state

type repository interface {
	IsDataReady() bool
}
