package filewatcher

import (
	"io/ioutil"
	"time"

	send ".."
	c "../../config"
	u "github.com/cdutwhu/go-util"
	"github.com/fsnotify/fsnotify"
)

// StartFileWatcherAsync :
func StartFileWatcherAsync() {
	defer func() { uPH(recover(), "./log.txt", true) }()

	watcher, e := fsnotify.NewWatcher()
	uPE(e)

	defer watcher.Close()
	uPE(watcher.Add(c.Cfg.Filewatcher.Dirsif))
	uPE(watcher.Add(c.Cfg.Filewatcher.Dirxapi))

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			lPln("event:", event) // CREATE WRITE REMOVE RENAME
			if event.Op&fsnotify.Create == fsnotify.Create {
				time.Sleep(2 * time.Second)
				lPln("created file:", event.Name)

			READ_AGAIN:
				bytes, e := ioutil.ReadFile(event.Name)
				if e != nil && sC(e.Error(), "The process cannot access the file because it is being used by another process") {
					fPln("read file failed, trying again ...")
					time.Sleep(1000 * time.Millisecond)
					goto READ_AGAIN
				}

				str := u.Str(string(bytes))
				if str.IsJSON() {
					send.XAPI(str.V())
				} else if str.IsXMLSegSimple() {
					send.SIF(str.V())
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			lPln("error:", err)
		}
	}
}
