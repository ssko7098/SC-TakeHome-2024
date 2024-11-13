package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// Custom UUID mappings for simplicity
var orgIDMap = map[string]uuid.UUID{
	"org1": uuid.Must(uuid.NewV4()),
	"org2": uuid.Must(uuid.NewV4()),
}

// Sample folder data shared across tests
var folders = []folder.Folder{
	{Name: "alpha", Paths: "alpha", OrgId: orgIDMap["org1"]},
	{Name: "bravo", Paths: "alpha.bravo", OrgId: orgIDMap["org1"]},
	{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgIDMap["org1"]},
	{Name: "delta", Paths: "alpha.delta", OrgId: orgIDMap["org1"]},
	{Name: "echo", Paths: "echo", OrgId: orgIDMap["org1"]},
	{Name: "foxtrot", Paths: "foxtrot", OrgId: orgIDMap["org2"]},
}

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:    "Retrieve all folders in org1",
			orgID:   orgIDMap["org1"],
			folders: folders,
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgIDMap["org1"]},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgIDMap["org1"]},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgIDMap["org1"]},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgIDMap["org1"]},
				{Name: "echo", Paths: "echo", OrgId: orgIDMap["org1"]},
			},
		},
		{
			name:    "Retrieve all folders in org2",
			orgID:   orgIDMap["org2"],
			folders: folders,
			want: []folder.Folder{
				{Name: "foxtrot", Paths: "foxtrot", OrgId: orgIDMap["org2"]},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got := f.GetFoldersByOrgID(tt.orgID)

			assert.Equal(t, tt.want, got)
		})
	}
}

// Test GetAllChildFolders function
func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		orgID      uuid.UUID
		folderName string
		folders    []folder.Folder
		want       []folder.Folder
	}{
		{
			name:       "Get children of 'alpha' in org1",
			orgID:      orgIDMap["org1"],
			folderName: "alpha",
			folders:    folders,
			want: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgIDMap["org1"]},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgIDMap["org1"]},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgIDMap["org1"]},
			},
		},
		{
			name:       "Get children of 'bravo' in org1",
			orgID:      orgIDMap["org1"],
			folderName: "bravo",
			folders:    folders,
			want: []folder.Folder{
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgIDMap["org1"]},
			},
		},
		{
			name:       "Get children of 'charlie' in org1",
			orgID:      orgIDMap["org1"],
			folderName: "charlie",
			folders:    folders,
			want:       []folder.Folder{},
		},
		{
			name:       "Get children of 'echo' in org1",
			orgID:      orgIDMap["org1"],
			folderName: "echo",
			folders:    folders,
			want:       []folder.Folder{},
		},
		{
			name:       "Folder does not exist",
			orgID:      orgIDMap["org1"],
			folderName: "invalidFolder",
			folders:    folders,
			want:       nil,
		},
		{
			name:       "Folder does not exist in specified organisation",
			orgID:      orgIDMap["org1"],
			folderName: "foxtrot",
			folders:    folders,
			want:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a driver instance with the test folder data
			f := folder.NewDriver(tt.folders)

			got := f.GetAllChildFolders(tt.orgID, tt.folderName)

			assert.Equal(t, tt.want, got)
		})
	}
}
