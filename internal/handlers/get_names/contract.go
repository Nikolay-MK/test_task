package get_names

import (
	"selltech/internal/selltech"
)

type repository interface {
	GetNames(query string, args ...interface{}) []selltech.NameResult
	GetStrongNames(name string) []selltech.NameResult
	GetWeakNames(name string) []selltech.NameResult
	GetAllNames(name string) []selltech.NameResult
}
