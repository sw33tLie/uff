package ffuf

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

var (
	//VERSION holds the current version number
	VERSION = "2.1.0 sw33tlie unleashed"
	//VERSION_APPENDIX holds additional version definition
	VERSION_APPENDIX = "-dev"
	CONFIGDIR        = filepath.Join(xdg.ConfigHome, "ffuf")
	HISTORYDIR       = filepath.Join(CONFIGDIR, "history")
	SCRAPERDIR       = filepath.Join(CONFIGDIR, "scraper")
	AUTOCALIBDIR     = filepath.Join(CONFIGDIR, "autocalibration")
)
