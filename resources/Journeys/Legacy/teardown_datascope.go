package Legacy

import (
	"context"

	ct "github.com/autopilothq/banks/contract/types"
	"github.com/autopilothq/banks/types"
	"github.com/autopilothq/lg"
)

// TeardownDatascope discards components in an underlying database after they
// will no longer be used by the given instance
func TeardownDatascope(
	ctx context.Context, dsID types.DatascopeID, log lg.Log, aux ct.Auxiliary,
) error {
	return nil
}
