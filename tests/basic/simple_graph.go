package main

import (
	"github.com/foliagecp/easyjson"
	lg "github.com/foliagecp/sdk/statefun/logger"

	"github.com/foliagecp/sdk/statefun"
	sfPlugins "github.com/foliagecp/sdk/statefun/plugins"
	"github.com/foliagecp/sdk/statefun/system"
)

func CreateTestGraph(runtime *statefun.Runtime) {
	lg.Logln(lg.DebugLevel, ">>> Test started: simple graph creation")

	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "rt", easyjson.NewJSONObject().GetPtr(), nil))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "a", easyjson.NewJSONObject().GetPtr(), nil))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "b", easyjson.NewJSONObject().GetPtr(), nil))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "c", easyjson.NewJSONObject().GetPtr(), nil))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "d", easyjson.NewJSONObject().GetPtr(), nil))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "e", easyjson.NewJSONObject().GetPtr(), nil))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "f", easyjson.NewJSONObject().GetPtr(), nil))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "g", easyjson.NewJSONObject().GetPtr(), nil))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.vertex.create", "h", easyjson.NewJSONObject().GetPtr(), nil))

	var v easyjson.JSON

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("a"))
	v.SetByPath("type", easyjson.NewJSON("type1"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t1", "t2"}))
	v.SetByPath("name", easyjson.NewJSON("2a"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "rt", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("a"))
	v.SetByPath("type", easyjson.NewJSON("type2"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t2", "t4"}))
	v.SetByPath("name", easyjson.NewJSON("2a"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "rt", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("b"))
	v.SetByPath("type", easyjson.NewJSON("type2"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t2"}))
	v.SetByPath("name", easyjson.NewJSON("2b"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "rt", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("c"))
	v.SetByPath("type", easyjson.NewJSON("type1"))
	v.SetByPath("tags", easyjson.NewJSONObject())
	v.SetByPath("name", easyjson.NewJSON("2c"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "rt", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("e"))
	v.SetByPath("type", easyjson.NewJSON("type3"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t3"}))
	v.SetByPath("name", easyjson.NewJSON("2e"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "a", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("e"))
	v.SetByPath("type", easyjson.NewJSON("type4"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t1", "t2", "t3"}))
	v.SetByPath("name", easyjson.NewJSON("2e"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "b", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("d"))
	v.SetByPath("type", easyjson.NewJSON("type3"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t1"}))
	v.SetByPath("name", easyjson.NewJSON("2d"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "c", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("b"))
	v.SetByPath("type", easyjson.NewJSON("type1"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t1", "t3"}))
	v.SetByPath("name", easyjson.NewJSON("2b"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "d", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("b"))
	v.SetByPath("type", easyjson.NewJSON("type2"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t4"}))
	v.SetByPath("name", easyjson.NewJSON("2b"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "e", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("f"))
	v.SetByPath("type", easyjson.NewJSON("type1"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t1", "t4"}))
	v.SetByPath("name", easyjson.NewJSON("2f"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "e", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("g"))
	v.SetByPath("type", easyjson.NewJSON("type5"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t1", "t2", "t3", "t4"}))
	v.SetByPath("name", easyjson.NewJSON("2g"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "f", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("d"))
	v.SetByPath("type", easyjson.NewJSON("type2"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{"t5"}))
	v.SetByPath("name", easyjson.NewJSON("2d"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "g", &v, nil))

	v = easyjson.NewJSONObject()
	v.SetByPath("to", easyjson.NewJSON("h"))
	v.SetByPath("type", easyjson.NewJSON("type2"))
	v.SetByPath("tags", easyjson.JSONFromArray([]string{}))
	v.SetByPath("name", easyjson.NewJSON("2h"))
	system.MsgOnErrorReturn(runtime.Request(sfPlugins.AutoRequestSelect, "functions.graph.api.link.create", "g", &v, nil))

	lg.Logln(lg.DebugLevel, "<<< Test ended: simple graph creation")
}
