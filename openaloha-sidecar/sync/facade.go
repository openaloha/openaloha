package sync

import (
	"openaloha.io/openaloha/openaloha-sidecar/config"
	_ "openaloha.io/openaloha/openaloha-sidecar/internal/sync/handler"
	"openaloha.io/openaloha/openaloha-sidecar/runfunc"
	"openaloha.io/openaloha/openaloha-sidecar/sync/factory"
	"openaloha.io/openaloha/openaloha-sidecar/sync/handler"
)

// SyncFacade is the facade for the sync service
type SyncFacade struct {
	Config       config.Config
}

// Sync is the method to init and refresh code
func (f *SyncFacade) Sync(initFunc runfunc.InitFunc, refreshFunc runfunc.RefreshFunc) error {
	// get sync handler by sync type
	syncHandler, err := getSyncHandler( f.Config.Sync.Type)
	if err != nil {
		return err
	}

	// init code by sync handler
	if err := syncHandler.Init(f.Config.Workspace, f.Config.Sync, initFunc); err != nil {
		return err
	}

	// refresh code by sync handler
	if err := syncHandler.Refresh(f.Config.Workspace, f.Config.Sync, refreshFunc); err != nil {
		return err
	}

	return nil
}


// get sync handler by sync type
func getSyncHandler(syncType string) (handler.SyncHandler, error) {
	syncHandler, err := factory.New(syncType)
	if err != nil {
		return nil, err
	}
	return syncHandler, nil
}
