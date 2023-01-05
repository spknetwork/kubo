package main

import (
	"net/http"

	"github.com/spknetwork/kubo/profile"
)

func init() {
	http.HandleFunc("/debug/stack",
		func(w http.ResponseWriter, _ *http.Request) {
			_ = profile.WriteAllGoroutineStacks(w)
		},
	)
}
