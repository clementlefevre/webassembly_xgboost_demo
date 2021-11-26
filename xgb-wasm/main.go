// Copyright 2019 Vanessa Sochat. All rights reserved.
// Use of this source code is governed by the Polyform Strict license
// that can be found in the LICENSE file and available at
// https://polyformproject.org/licenses/noncommercial/1.0.0

package main

import (
	"syscall/js"
)

func main() {

	c := make(chan struct{}, 0)
	runnerXGB := XGBRunner{}
	runnerXGB.readModel()
	runnerXGB.readXData()

	js.Global().Set("runXGB", js.FuncOf(runXGB))

	<-c
}

func runXGB(this js.Value, val []js.Value) interface{} {

	return nil
}
