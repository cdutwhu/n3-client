package filewatcher

import (
	"io/ioutil"

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
				bytes, e := ioutil.ReadFile(event.Name)
				uPE(e)
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
