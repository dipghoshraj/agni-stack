package karma

import (
	cosmicmodel "brahma-node-creation-hub/cosmic-model"
)

type CosmosHandler struct {
	DbManager *cosmicmodel.CosmosDB
}

func Creation(DbManager *cosmicmodel.CosmosDB) *CosmosHandler {
	return &CosmosHandler{DbManager: DbManager}
}
