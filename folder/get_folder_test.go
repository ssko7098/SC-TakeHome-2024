package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name  string
		orgID uuid.UUID
		want  []folder.Folder
	}{
		{
			name:  "Retrieve all folders in org1",
			orgID: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"),
			want: []folder.Folder{
				{Name: "capable-baroness", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness"},
				{Name: "expert-buttercup", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup"},
				{Name: "picked-forerunner", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup.picked-forerunner"},
				{Name: "finer-witchblade", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.finer-witchblade"},
				{Name: "magnetic-jungle", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.finer-witchblade.magnetic-jungle"},
				{Name: "quick-cornelius", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.quick-cornelius"},
				{Name: "healthy-mongoose", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.quick-cornelius.healthy-mongoose"},
			},
		},
		{
			name:  "Retrieve all folders in org2",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			want: []folder.Folder{
				{Name: "discrete-manta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta"},
				{Name: "topical-raphael", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.topical-raphael"},
				{Name: "choice-flash", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.topical-raphael.choice-flash"},
				{Name: "valid-vindicator", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator"},
				{Name: "summary-epoch", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.summary-epoch"},
				{Name: "flying-gambit", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.flying-gambit"},
				{Name: "wanted-talisman", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.wanted-talisman"},
				{Name: "cheerful-horridus", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "cheerful-horridus"},
				{Name: "sterling-the-leader", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "cheerful-horridus.sterling-the-leader"},
				{Name: "grown-wallop", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "cheerful-horridus.sterling-the-leader.grown-wallop"},
				{Name: "creative-microbe", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "cheerful-horridus.sterling-the-leader.creative-microbe"},
				{Name: "meet-sauron", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "cheerful-horridus.sterling-the-leader.meet-sauron"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(folder.GetAllFolders())
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
		want       []folder.Folder
	}{
		{
			name:       "Retrieve children of expert-buttercup in org1",
			orgID:      uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"),
			folderName: "expert-buttercup",
			want: []folder.Folder{
				{Name: "picked-forerunner", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup.picked-forerunner"},
			},
		},
		{
			name:       "Retrieve children of valid-vindicator in org2",
			orgID:      uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folderName: "valid-vindicator",
			want: []folder.Folder{
				{Name: "summary-epoch", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.summary-epoch"},
				{Name: "flying-gambit", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.flying-gambit"},
				{Name: "wanted-talisman", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.wanted-talisman"},
			},
		},
		{
			name:       "Retrieve children of leaf folder",
			orgID:      uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folderName: "summary-epoch",
			want:       []folder.Folder{},
		},
		{
			name:       "Folder does not exist",
			orgID:      uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"),
			folderName: "invalidFolder",
			want:       nil,
		},
		{
			name:       "Folder does not exist in specified organisation",
			orgID:      uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folderName: "magnetic-jungle",
			want:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a driver instance with the test folder data
			f := folder.NewDriver(folder.GetAllFolders())

			got := f.GetAllChildFolders(tt.orgID, tt.folderName)

			assert.Equal(t, tt.want, got)
		})
	}
}
