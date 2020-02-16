package history

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"path"
)

type History struct {
	AppName  string
	HistPath string
	HistFile string
}

func NewHistory(appname, basedir string) *History {
	retv := &History{
		HistPath: basedir,
		AppName: appname,
	}
	appstr := fmt.Sprintf(".%s.hist", appname)
	retv.HistFile = path.Join(basedir, appstr)
	return retv
}

func (h History) Add(instring string) {
	file, err := os.OpenFile(h.HistFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteString(instring + "\n"); err != nil {
		log.Println(err)
	}
}

func (h History) Read() []string {
	var retv []string
	filehandle, err := os.Open(h.HistFile)

	if err != nil {
		log.Fatal(err)
	}

	defer filehandle.Close()

	scanner := bufio.NewScanner(filehandle)
	for scanner.Scan() {
		retv = append(retv, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return retv

}
