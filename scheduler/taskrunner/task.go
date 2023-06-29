package taskrunner

import (
	"errors"
	"log"
	"os"
	"stream-media/scheduler/dbops"
	"sync"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)

	if err != nil && !os.IsNotExist(err){
		log.Printf("Deleting video error: %v",err)
	}
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasked finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil  
}

func VideoClearExecutor(da dataChan) error {
	errMap := &sync.Map{}
	var err error

	forLoop:
		for {
			select {
			case vid :=<- dc:
				go func(id interface{})  {
					if err := deleteVideo(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
				}(vid)
				default:
					break forLoop
			}
		}
	errMap.Range(func(k, v, interface{})bool{
		err = v.(error)
		if err != nil {
			return false
		}
	})
}