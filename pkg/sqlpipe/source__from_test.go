package sqlpipe

import (
	"testing"

	"github.com/octohelm/storage/pkg/sqlfrag"
	"github.com/octohelm/storage/pkg/sqlfrag/testutil"
	"github.com/octohelm/storage/testdata/model"
	testingx "github.com/octohelm/x/testing"
)

func TestSourceFrom(t *testing.T) {
	src := FromAll[model.User]()

	testingx.Expect[sqlfrag.Fragment](t, src, testutil.BeFragment(`
SELECT *
FROM t_user
`))
}
