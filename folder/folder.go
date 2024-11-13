package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type driver struct {
	// define attributes here
	// data structure to store folders
	// or preprocessed data

	// Maps a folder path to its child folders
	parentChildMap map[string][]Folder

	// example: feel free to change the data structure, if slice is not what you want
	folders []Folder
}

func NewDriver(folders []Folder) IDriver {

	d := &driver{
		folders:        folders,
		parentChildMap: make(map[string][]Folder),
	}

	d.buildParentChildMap()
	return d
}

// constructs a map where each path points to its immediate children
func (d *driver) buildParentChildMap() {
	for _, folder := range d.folders {
		parentPath := getParentPath(folder.Paths)
		d.parentChildMap[parentPath] = append(d.parentChildMap[parentPath], folder)
	}
}

// extracts the parent path from a folder path
func getParentPath(path string) string {
	lastDot := strings.LastIndex(path, ".")
	if lastDot == -1 {
		return ""
	}
	return path[:lastDot]
}
