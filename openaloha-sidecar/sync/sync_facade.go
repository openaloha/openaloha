package sync

import (
	"openaloha.io/openaloha/openaloha-sidecar/config"
	synchandler "openaloha.io/openaloha/openaloha-sidecar/sync/sync_handler"
	"openaloha.io/openaloha/openaloha-sidecar/runfunc"
)

// SyncFacade is the facade for the sync service
type SyncFacade struct {
	Config       config.Config
	syncHandlers []synchandler.SyncHandler
}

// Sync is the method to init and refresh code
func (f *SyncFacade) Sync(initFunc runfunc.InitFunc, refreshFunc runfunc.RefreshFunc) error {
	// init sync handler
	f.syncHandlers = initSyncHandler()

	// get sync handler by sync type
	syncHandler, err := getSyncHandler(f.syncHandlers, f.Config.Sync.Type)
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

// init sync handler
func initSyncHandler() []synchandler.SyncHandler {
	var syncHandlers []synchandler.SyncHandler
	// init sync handler
	syncHandlers = append(syncHandlers, &synchandler.GitSyncHandler{})

	return syncHandlers
}

// get sync handler by sync type
func getSyncHandler(syncHandlers []synchandler.SyncHandler, syncType string) (synchandler.SyncHandler, error) {
	for _, syncHandler := range syncHandlers {
		if syncHandler.IsSupport(syncType) {
			return syncHandler, nil
		}
	}
	return nil, nil
}
