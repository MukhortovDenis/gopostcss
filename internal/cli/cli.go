package cli

import (
	"errors"
	"flag"
)

type Flags struct {
	strings []*string
}

func getFlags() []*string {
	nf := flag.String("nf", "", "\"new-filename\" полное имя нового файла, которое хотите получить, стандартное имя `new_YOURFILENAME`")
	return []*string{nf}
}

// InitFlags is func where creates flags for cli
func InitFlags() (string, string, error) {
	flags := new(Flags)
	flags.strings = getFlags()
	flag.Parse()
	var filename string
	var newFilename string
	if len(flag.Args()) == 0 {
		return "", "", errors.New("undefined filename")
	}
	filename = flag.Args()[0]

	if *flags.strings[0] == "" {
		newFilename = filename
	} else {
		newFilename = *flags.strings[0]
	}

	return filename, newFilename, nil
}
