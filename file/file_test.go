package file

import (
	"io/fs"
	"testing"
)

func TestFileModel(t *testing.T) {
	fileMode := fs.FileMode(0600)
	t.Logf("fileMode %+v", fileMode)
}
