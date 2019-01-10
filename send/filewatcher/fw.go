package filewatcher

import (
	"io/ioutil"
	"time"

	s ".."
	u "github.com/cdutwhu/util"
	"github.com/fsnotify/fsnotify"
)

// StartFileWatcherAsync :
func StartFileWatcherAsync() {
	defer func() { uPH(recover(), "./log.txt", true) }()

	watcher, e := fsnotify.NewWatcher()
	uPE(e)

	defer watcher.Close()
	uPE(watcher.Add(sifDir))
	uPE(watcher.Add(xapiDir))

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			lPln("event:", event) // CREATE WRITE REMOVE RENAME
			if event.Op&fsnotify.Create == fsnotify.Create {
				lPln("created file:", event.Name)

			AGAIN:
				bytes, e := ioutil.ReadFile(event.Name)
				if e != nil && sC(e.Error(), "The process cannot access the file because it is being used by another process") {
					fPln("read file failed, trying again ...")
					time.Sleep(500 * time.Millisecond)
					goto AGAIN
				}

				str := u.Str(string(bytes))
				if str.IsJSON() {
					s.XAPI(str.V())
				} else if str.IsXMLSegSimple() {
					s.SIF(str.V())
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
