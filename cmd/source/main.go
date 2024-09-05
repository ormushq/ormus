package main

import (
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/source/delivery/httpserver"
	"github.com/ormushq/ormus/source/delivery/httpserver/statushandler"
)

//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	JWTToken
//	@in							header
//	@name						Authorization

func main() {
	handlers := []httpserver.Handler{
		statushandler.New(),
	}

	httpServer := httpserver.New(config.C().Source, handlers)

	httpServer.Serve()
}
