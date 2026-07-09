// Package contentmcp is the Content MCP togo plugin: A JSON-RPC Model Context Protocol server exposing site content to agents (read tools public; bearer-gated write tools create real posts in registered tables).
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package contentmcp

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("content-mcp", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/content-mcp/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"content-mcp","status":"ok"}`))
		})
		return nil
	})
}
