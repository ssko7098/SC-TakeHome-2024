package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name    string
		srcName string
		dstName string
		want    interface{}
	}{
		{
			name:    "Move picked-forerunner above one branch",
			srcName: "picked-forerunner",
			dstName: "capable-baroness",
			want: []folder.Folder{
				{Name: "capable-baroness", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness"},
				{Name: "expert-buttercup", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup"},
				{Name: "picked-forerunner", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.picked-forerunner"},
				{Name: "finer-witchblade", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.finer-witchblade"},
				{Name: "magnetic-jungle", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.finer-witchblade.magnetic-jungle"},
				{Name: "quick-cornelius", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.quick-cornelius"},
				{Name: "healthy-mongoose", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.quick-cornelius.healthy-mongoose"},
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
		{
			name:    "Move quick-cornelius to expert-buttercup",
			srcName: "quick-cornelius",
			dstName: "expert-buttercup",
			want: []folder.Folder{
				{Name: "capable-baroness", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness"},
				{Name: "expert-buttercup", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup"},
				{Name: "picked-forerunner", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup.picked-forerunner"},
				{Name: "finer-witchblade", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.finer-witchblade"},
				{Name: "magnetic-jungle", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.finer-witchblade.magnetic-jungle"},
				{Name: "quick-cornelius", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup.quick-cornelius"},
				{Name: "healthy-mongoose", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup.quick-cornelius.healthy-mongoose"},
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
		{
			name:    "Move cheerful-horridus to discrete-manta",
			srcName: "cheerful-horridus",
			dstName: "discrete-manta",
			want: []folder.Folder{
				{Name: "capable-baroness", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness"},
				{Name: "expert-buttercup", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup"},
				{Name: "picked-forerunner", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.expert-buttercup.picked-forerunner"},
				{Name: "finer-witchblade", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.finer-witchblade"},
				{Name: "magnetic-jungle", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.finer-witchblade.magnetic-jungle"},
				{Name: "quick-cornelius", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.quick-cornelius"},
				{Name: "healthy-mongoose", OrgId: uuid.FromStringOrNil("97044613-96ab-442c-9f08-a9b09f04933c"), Paths: "capable-baroness.quick-cornelius.healthy-mongoose"},
				{Name: "discrete-manta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta"},
				{Name: "topical-raphael", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.topical-raphael"},
				{Name: "choice-flash", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.topical-raphael.choice-flash"},
				{Name: "valid-vindicator", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator"},
				{Name: "summary-epoch", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.summary-epoch"},
				{Name: "flying-gambit", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.flying-gambit"},
				{Name: "wanted-talisman", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.valid-vindicator.wanted-talisman"},
				{Name: "cheerful-horridus", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.cheerful-horridus"},
				{Name: "sterling-the-leader", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.cheerful-horridus.sterling-the-leader"},
				{Name: "grown-wallop", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.cheerful-horridus.sterling-the-leader.grown-wallop"},
				{Name: "creative-microbe", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.cheerful-horridus.sterling-the-leader.creative-microbe"},
				{Name: "meet-sauron", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "discrete-manta.cheerful-horridus.sterling-the-leader.meet-sauron"},
			},
		},
		{
			name:    "Move quick-cornelius to a different organization",
			srcName: "quick-cornelius",
			dstName: "meet-sauron",
			want:    errors.New("cannot move a folder to a different organization"),
		},
		{
			name:    "Move capable-baroness to a child of itself",
			srcName: "capable-baroness",
			dstName: "picked-forerunner",
			want:    errors.New("cannot move a folder to a child of itself"),
		},
		{
			name:    "Move capable-baroness to capable-baroness",
			srcName: "capable-baroness",
			dstName: "capable-baroness",
			want:    errors.New("cannot move a folder to itself"),
		},
		{
			name:    "Move invalid folder to capable-baroness",
			srcName: "invalid_folder",
			dstName: "capable-baroness",
			want:    errors.New("source folder does not exist"),
		},
		{
			name:    "Move capable-baroness folder to invalid folder",
			srcName: "capable-baroness",
			dstName: "invalid_folder",
			want:    errors.New("destination folder does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(folder.GetAllFolders())
			got, errorMessage := f.MoveFolder(tt.srcName, tt.dstName)

			if errorMessage == nil {
				assert.ElementsMatch(t, tt.want, got)
			} else {
				assert.Equal(t, tt.want, errorMessage)
			}
		})
	}
}
