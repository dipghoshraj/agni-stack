package karma

import (
	cosmicmodel "brahma-node-creation-hub/cosmic-model"
)

type CosmosHandler struct {
	CosmosDB *cosmicmodel.CosmosDB
}

func Creation(CosmosDB *cosmicmodel.CosmosDB) *CosmosHandler {
	return &CosmosHandler{CosmosDB: CosmosDB}
}
