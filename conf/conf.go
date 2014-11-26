package conf

import (
	"bitbucket.org/kardianos/osext"
	//"github.com/grooveshark/golib/gslog"
	"github.com/mediocregopher/flagconfig"
	"os"
)

func Parse() (*flagconfig.FlagConfig, error) {
	// Change working dir to that of the executable
	exeFolder, _ := osext.ExecutableFolder()
	os.Chdir(exeFolder)

	f := flagconfig.New("gopic")
	f.StrParam("loglevel", "logging level (DEBUG, INFO, WARN, ERROR, FATAL)", "DEBUG")
	f.StrParam("logfile", "path to log file", "")
	f.StrParam("cachepath", "path to cache folder for thumbnails", ".cache")
	f.RequiredStrParam("imagepath", "path to images")
	f.RequiredStrParam("listen", "a string of the form 'IP:PORT' which program will listen on")
	f.FlagParam("V", "show version/build information", false)

	if err := f.Parse(); err != nil {
		return nil, err
	}
	return f, nil
}
