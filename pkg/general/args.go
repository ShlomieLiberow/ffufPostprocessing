package general

import (
	"flag"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
)

func GetArguments() _struct.Configuration {

	OriginalFfufResultFile := flag.String("result-file", "", "Path to the original ffuf result file")
	NewFfufResultFile := flag.String("new-result-file", "", "Path to the new ffuf result file")

	FfufBodiesFolder := flag.String("bodies-folder", "", "Path to the ffuf bodies folder")
	DeleteUnnecessaryBodyFiles := flag.Bool("delete-bodies", false, "Delete unnecessary body files")

	flag.Parse()

	return _struct.Configuration{
		OriginalFfufResultFile:     *OriginalFfufResultFile,
		NewFfufResultFile:          *NewFfufResultFile,
		FfufBodiesFolder:           *FfufBodiesFolder,
		DeleteUnnecessaryBodyFiles: *DeleteUnnecessaryBodyFiles,
	}
}
