package querierpatcher

import (
	"github.com/octohelm/storage/deprecated/pkg/dal"
	"github.com/octohelm/storage/pkg/sqlbuilder"
)

func ApplyToQuerier[M sqlbuilder.Model](q dal.Querier, patchers ...Typed[M]) dal.Querier {
	return q.Apply(CastSliceAsAnySlice(patchers...)...)
}
