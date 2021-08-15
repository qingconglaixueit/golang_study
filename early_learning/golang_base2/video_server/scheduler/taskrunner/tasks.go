package taskrunner

import (
	"errors"
	"github.com/wonderivan/logger"
	"log"
	"os"
	"path/filepath"
	"sync"
	"video_server/scheduler/dbops"
)

func deleteVideo(vid string) error {
	path, _ := filepath.Abs(VIDEO_PATH + vid)
	logger.Info(path)
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
	}
	if len(res) == 0 {
		return errors.New("All tasks finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case id := <-dc:
			logger.Info("vid := <-dc")
			//go func(id interface{}) {
				logger.Info("deleteVideo")
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
				}
			//}(vid)
		default:
			logger.Info("break forloop")
			break forloop
		}
	}
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
