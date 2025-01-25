package graphql

import (
	stdhttp "net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/rustic-beans/spotify-viewer/generated"
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/resolver"
	"github.com/rustic-beans/spotify-viewer/internal/services"
	"github.com/vektah/gqlparser/v2/ast"
)

func NewServer(
	sharedService *services.Shared,
	playerStateWebsocketHandler *models.PlayerStateWebsocketHandler,
) *handler.Server {
	server := handler.New(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver.Resolver{
				SharedService:               sharedService,
				PlayerStateWebsocketHandler: playerStateWebsocketHandler,
			}},
		),
	)

	server.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(_ *stdhttp.Request) bool {
				return true // TODO: Add origin check
			},
		},
		//nolint:mnd // no need to change the ping interval
		KeepAlivePingInterval: 10 * time.Second,
	})
	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})

	//nolint:mnd // no need to change the cache size
	server.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	server.Use(extension.Introspection{})
	server.Use(extension.AutomaticPersistedQuery{
		//nolint:mnd // no need to change the cache size
		Cache: lru.New[string](100),
	})

	return server
}
