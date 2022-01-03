package main

import "hand_by_hand_framework/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
