package update

import "selltech/internal/selltech"

type repository interface {
	UpdateIndividuals(entries []selltech.SDNEntry) error
}
