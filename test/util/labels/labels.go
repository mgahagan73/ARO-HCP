package labels

import (
<<<<<<< HEAD
	. "github.com/onsi/ginkgo/v2"
)

var (
	Negative = Label("Negative")
	Positive = Label("Positive")
=======
	ginkgo "github.com/onsi/ginkgo/v2"
)

var (
	Negative = ginkgo.Label("Negative")
>>>>>>> upstream/main
)

// Test cases importance
var (
<<<<<<< HEAD
	Low      = Label("Low")
	Medium   = Label("Medium")
	High     = Label("High")
	Critical = Label("Critical")
=======
	Low      = ginkgo.Label("Low")
	Medium   = ginkgo.Label("Medium")
	High     = ginkgo.Label("High")
	Critical = ginkgo.Label("Critical")
>>>>>>> upstream/main
)
