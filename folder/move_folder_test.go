package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

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
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgIDMap["org1"]},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: orgIDMap["org2"]},
		{Name: "golf", Paths: "golf", OrgId: orgIDMap["org1"]},
	}

	tests := [...]struct {
		name    string
		srcName string
		dstName string
		folders []folder.Folder
		want    interface{}
	}{
		{
			name:    "Move bravo to delta",
			srcName: "bravo",
			dstName: "delta",
			folders: folders,
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgIDMap["org1"]},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: orgIDMap["org1"]},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: orgIDMap["org1"]},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgIDMap["org1"]},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgIDMap["org1"]},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: orgIDMap["org2"]},
				{Name: "golf", Paths: "golf", OrgId: orgIDMap["org1"]},
			},
		},
		{
			name:    "Move bravo to golf",
			srcName: "bravo",
			dstName: "golf",
			folders: folders,
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgIDMap["org1"]},
				{Name: "bravo", Paths: "golf.bravo", OrgId: orgIDMap["org1"]},
				{Name: "charlie", Paths: "golf.bravo.charlie", OrgId: orgIDMap["org1"]},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgIDMap["org1"]},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgIDMap["org1"]},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: orgIDMap["org2"]},
				{Name: "golf", Paths: "golf", OrgId: orgIDMap["org1"]},
			},
		},
		// {
		// 	name:    "Move bravo to charlie",
		// 	srcName: "bravo",
		// 	dstName: "charlie",
		// 	folders: folders,
		// 	want:    errors.New("cannot move a folder to a child of itself"),
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, errorMessage := f.MoveFolder(tt.srcName, tt.dstName)

			if errorMessage == nil {
				assert.ElementsMatch(t, tt.want, got)
			} else {
				assert.Equal(t, tt.want, errorMessage)
			}
		})
	}
}
