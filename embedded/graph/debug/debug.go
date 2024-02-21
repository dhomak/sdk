

// Foliage graph store debug package.
// Provides debug stateful functions for the graph store
package debug

import (
	"fmt"

	"github.com/foliagecp/sdk/embedded/graph/crud"
	"github.com/foliagecp/sdk/statefun"
	lg "github.com/foliagecp/sdk/statefun/logger"
	sfPlugins "github.com/foliagecp/sdk/statefun/plugins"
)

func RegisterAllFunctionTypes(runtime *statefun.Runtime) {
	statefun.NewFunctionType(runtime, "functions.graph.api.object.debug.print", LLAPIObjectDebugPrint, *statefun.NewFunctionTypeConfig())
	statefun.NewFunctionType(runtime, "functions.graph.api.object.debug.print.graph", LLAPIPrintGraph, *statefun.NewFunctionTypeConfig())
}

/*
Prints to caonsole the content of an object the function being called on along with all its input and output links.
*/
func LLAPIObjectDebugPrint(executor sfPlugins.StatefunExecutor, ctx *sfPlugins.StatefunContextProcessor) {
	self := ctx.Self

	objectContext := ctx.GetObjectContext()
	lg.Logf(lg.DebugLevel, "************************* Object's body (id=%s):\n", self.ID)
	lg.Logln(lg.DebugLevel, objectContext.ToString())
	lg.Logf(lg.DebugLevel, "************************* In links:\n")
	for _, key := range ctx.Domain.Cache().GetKeysByPattern(fmt.Sprintf(crud.InLinkKeyPrefPattern+crud.LinkKeySuff1Pattern, self.ID, ">")) {
		lg.Logln(lg.DebugLevel, key)
	}
	lg.Logf(lg.DebugLevel, "************************* Out links:\n")
	for _, key := range ctx.Domain.Cache().GetKeysByPattern(fmt.Sprintf(crud.OutLinkBodyKeyPrefPattern+crud.LinkKeySuff1Pattern, self.ID, ">")) {
		lg.Logln(lg.DebugLevel, key)
		if j, err := ctx.Domain.Cache().GetValueAsJSON(key); err == nil {
			lg.Logln(lg.DebugLevel, j.ToString())
		}
	}
	lg.Logln(lg.DebugLevel)
}
