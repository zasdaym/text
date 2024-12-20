// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"github.com/zasdaym/text/language"
	"github.com/zasdaym/text/message"
)

// The printer defined here will be picked up by the first print statement
// in main.go.
var printer = message.NewPrinter(language.English)
