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
	subtree := []Folder{srcFolder}

	// Special case -> where the src Folder is a leaf node
	// just change it's entire srcPath to the dstPath
	if len(f.GetAllChildFolders(srcFolder.OrgId, name)) == 0 {
		subtree[0].Paths = dstFolder.Paths + "." + srcFolder.Name

	} else {
		subtree = append(subtree, f.GetAllChildFolders(srcFolder.OrgId, name)...)

		// Update each folder in the subtree with the new path
		for i := range subtree {
			subtree[i].Paths = findNewPath(subtree[i].Paths, dstFolder.Paths)
		}
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
func findNewPath(srcPath, dstPath string) string {

	// Split the paths into segments by "."
	srcParts := strings.Split(srcPath, ".")
	dstParts := strings.Split(dstPath, ".")

	// Find where the paths diverge
	i := 0
	for i < len(srcParts) && i < len(dstParts) && srcParts[i] != dstParts[i] {
		i++
	}

	// special case where dst folder is a root folder
	if i == len(srcParts) || i == len(dstParts) {
		i = 0
	}

	// Construct the new path with dstPath + remaining srcPath
	result := strings.Join(dstParts, ".")

	// Append any remaining part of srcPath after the point of divergence
	if len(srcParts) > i {
		result += "." + strings.Join(srcParts[i+1:], ".")
	}

	return result
}
