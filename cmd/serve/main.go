package main

import (
	log "github.com/blockchain-abstraction-middleware/rest-api/pkg/logger"
	"github.com/blockchain-abstraction-middleware/rest-api/pkg/routes"
	"github.com/blockchain-abstraction-middleware/rest-api/pkg/server"

	config "github.com/IPFS-Media-Platform/vault-abstraction/pkg/config"
	vm "github.com/IPFS-Media-Platform/vault-abstraction/pkg/routes"
)

func main() {
	abstractionConfig := config.NewConfig()

	serverConfig := server.Config{
		BasePath:       "/ipfs-media-platform",
		Name:           abstractionConfig.Name,
		Port:           abstractionConfig.Port,
		MetricsEnabled: abstractionConfig.MetricsEnabled,
	}
	srv := server.New(&serverConfig)

	healthResource := routes.HealthResource{}
	swaggerResource := routes.SwaggerResource{}
	vaultManagerResource := vm.ManagerResource{}

	srv.RegisterResource(healthResource.NewResource("/health"))
	srv.RegisterResource(swaggerResource.NewResource("/swagger", "/ipfs-media-platform/swagger"))

	srv.RegisterResource(vaultManagerResource.NewResource("/manager", abstractionConfig))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Serving failed")
	}
}
