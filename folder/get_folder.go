package folder

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {

	// Find the folder with the given name and orgID to get its path
	var rootPath string
	folderExists := false

	for _, f := range f.folders {
		if f.OrgId == orgID && f.Name == name {
			folderExists = true
			rootPath = f.Paths
			break
		}
	}

	// If the folder wasn't found, return a null value
	if !folderExists {
		return nil
	}

	// If the parent folder was not found, return a null value
	if rootPath == "" {
		return nil
	}

	// Collect all descendants starting from the root path
	return f.collectDescendants(rootPath)
}

// recursively collects all folders starting from a given path
func (f *driver) collectDescendants(path string) []Folder {
	descendants := []Folder{}
	children := f.parentChildMap[path]

	for _, child := range children {
		descendants = append(descendants, child)
		descendants = append(descendants, f.collectDescendants(child.Paths)...) // Recursively collect children
	}

	return descendants
}
