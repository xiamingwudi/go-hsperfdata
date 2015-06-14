package hsperfdata

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type HsperfdataRepository struct {
	dir string
}

func New() (*HsperfdataRepository, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	return NewUser(user.Username)
}

func NewUser(userName string) (*HsperfdataRepository, error) {
	dir := filepath.Join(os.TempDir(), "hsperfdata_"+userName)
	return &HsperfdataRepository{dir}, nil
}

func (repository *HsperfdataRepository) GetFile(pid string) HsperfdataFile {
	return HsperfdataFile{filepath.Join(repository.dir, pid)}
}

func (repository *HsperfdataRepository) GetFiles() ([]HsperfdataFile, error) {
	files, err := ioutil.ReadDir(repository.dir)
	if err != nil {
		return nil, err
	}
	retval := make([]HsperfdataFile, len(files))
	for i, f := range files {
		retval[i] = HsperfdataFile{filepath.Join(repository.dir, f.Name())}
	}

	return retval, nil
}