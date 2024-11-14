package folder

import (
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	// Helper function to find folder by name
	findFolder := func(name string) (Folder, bool) {
		for _, folder := range f.folders {
			if folder.Name == name {
				return folder, true
			}
		}
		return Folder{}, false
	}

	// Find source and destination folders
	srcFolder, srcExists := findFolder(name)
	dstFolder, dstExists := findFolder(dst)

	// Different cases for error handling
	switch {
	case !srcExists:
		return nil, fmt.Errorf("source folder does not exist")
	case !dstExists:
		return nil, fmt.Errorf("destination folder does not exist")
	case srcFolder.Paths == dstFolder.Paths:
		return nil, fmt.Errorf("cannot move a folder to itself")
	case strings.HasPrefix(dstFolder.Paths, srcFolder.Paths):
		return nil, fmt.Errorf("cannot move a folder to a child of itself")
	case srcFolder.OrgId != dstFolder.OrgId:
		return nil, fmt.Errorf("cannot move a folder to a different organization")
	}

	// Move Logic
	// Find the subtree (including the src as the root folder)
	subtree := append([]Folder{srcFolder}, f.GetAllChildFolders(srcFolder.OrgId, name)...)

	for i := range subtree {
		subtree[i].Paths = findNewPath(subtree[i].Paths, dstFolder.Paths, name)
	}

	// integrates the subtree's updated paths with the rest of the folders
	subtreeMap := make(map[string]string) // Map folder name to new path
	for _, folder := range subtree {
		subtreeMap[folder.Name] = folder.Paths
	}

	res := []Folder{}
	for i := range f.folders {
		res = append(res, f.folders[i])

		if newPath, exists := subtreeMap[f.folders[i].Name]; exists {
			res[i].Paths = newPath
		}
	}

	return res, nil
}

// constructs a new path by replacing the root of the subtree with dstPath.
func findNewPath(srcPath, dstPath, name string) string {

	// Split the paths into segments by "."
	srcParts := strings.Split(srcPath, ".")
	dstParts := strings.Split(dstPath, ".")

	// Find the index of the root of the subtree in the source path
	// (i.e., find the index of name)
	var index int
	for i, part := range srcParts {
		if part == name {
			index = i - 1
			break
		}
	}

	// Replace the source path up to the name with the destination path
	// Concatenate dstParts with the remaining parts of srcParts after "name"
	newPathParts := append(dstParts, srcParts[index+1:]...)
	result := strings.Join(newPathParts, ".")

	return result
}
