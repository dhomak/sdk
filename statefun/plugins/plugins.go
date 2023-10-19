

// Foliage statefun plugins package.
// Provides unified interfaces for stateful functions plugins
package plugins

import (
	"fmt"
	"sync"

	"github.com/foliagecp/easyjson"

	"github.com/foliagecp/sdk/statefun/cache"
)

type StatefunAddress struct {
	Typename string
	ID       string
}

type SignalProvider int

const (
	JetstreamGlobalSignal SignalProvider = iota
)

type RequestProvider int

const (
	NatsCoreGlobalRequest RequestProvider = iota
	GolangLocalRequest
)

type StatefunContextProcessor struct {
	GlobalCache        *cache.Store
	GetFunctionContext func() *easyjson.JSON
	SetFunctionContext func(*easyjson.JSON)
	GetObjectContext   func() *easyjson.JSON
	SetObjectContext   func(*easyjson.JSON)
	// TODO: DownstreamSignal(<function type>, <links filters>, <payload>, <options>)
	Signal           func(SignalProvider, string, string, *easyjson.JSON, *easyjson.JSON) error
	Request          func(RequestProvider, string, string, *easyjson.JSON, *easyjson.JSON) (*easyjson.JSON, error)
	Self             StatefunAddress
	Caller           StatefunAddress
	Payload          *easyjson.JSON
	Options          *easyjson.JSON
	RequestReplyData *easyjson.JSON // when requested in function: nil - function was signaled, !nil - function was requested
}

type StatefunExecutor interface {
	Run(contextProcessor *StatefunContextProcessor) error
	BuildError() error
}

type StatefunExecutorConstructor func(alias string, source string) StatefunExecutor

type TypenameExecutorPlugin struct {
	alias                      string
	source                     string
	idExecutors                sync.Map
	executorContructorFunction StatefunExecutorConstructor
}

func NewTypenameExecutor(alias string, source string, executorContructorFunction StatefunExecutorConstructor) *TypenameExecutorPlugin {
	tnex := TypenameExecutorPlugin{alias: alias, source: source, executorContructorFunction: executorContructorFunction}
	return &tnex
}

func (tnex *TypenameExecutorPlugin) AddForID(id string) {
	if tnex.executorContructorFunction == nil {
		fmt.Printf("Cannot create new StatefunExecutor for id=%s: missing newExecutor function\n", id)
		tnex.idExecutors.Store(id, nil)
	} else {
		fmt.Printf("______________ Created StatefunExecutor for id=%s\n", id)
		executor := tnex.executorContructorFunction(tnex.alias, tnex.source)
		tnex.idExecutors.Store(id, executor)
	}
}

func (tnex *TypenameExecutorPlugin) RemoveForID(id string) {
	tnex.idExecutors.Delete(id)
}

func (tnex *TypenameExecutorPlugin) GetForID(id string) StatefunExecutor {
	value, _ := tnex.idExecutors.Load(id)
	return value.(StatefunExecutor)
}
