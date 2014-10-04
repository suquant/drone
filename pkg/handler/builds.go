package handler

import (
	"net/http"

	"github.com/suquant/drone/pkg/database"
	. "github.com/suquant/drone/pkg/model"
)

type BuildResult struct {
	Status string
}

// Returns the combined stdout / stderr for an individual Build.
func BuildOut(w http.ResponseWriter, r *http.Request, u *User, repo *Repo) error {
	branch := r.FormValue("branch")
	if branch == "" {
		branch = "master"
	}

	hash := r.FormValue(":commit")
	labl := r.FormValue(":label")

	// get the commit from the database
	commit, err := database.GetCommitBranchHash(branch, hash, repo.ID)
	if err != nil {
		return err
	}

	// get the build from the database
	build, err := database.GetBuildSlug(labl, commit.ID)
	if err != nil {
		return err
	}

	return RenderText(w, build.Stdout, http.StatusOK)
}

// Returns the combined stdout / stderr for an individual Build.
func BuildStatus(w http.ResponseWriter, r *http.Request, repo *Repo) error {
	branch := r.FormValue("branch")
	if branch == "" {
		branch = "master"
	}

	hash := r.FormValue(":commit")
	labl := r.FormValue(":label")

	// get the commit from the database
	commit, err := database.GetCommitBranchHash(branch, hash, repo.ID)
	if err != nil {
		return err
	}

	// get the build from the database
	build, err := database.GetBuildSlug(labl, commit.ID)
	if err != nil {
		return err
	}

	build_result := BuildResult{build.Status}

	return RenderJson(w, build_result)
}

// Returns the gzipped stdout / stderr for an individual Build
func BuildOutGzip(w http.ResponseWriter, r *http.Request, u *User) error {
	// TODO
	return nil
}
