package papermc

import "time"

// BuildResponse represents a build response object
type BuildResponse struct {
	ProjectID   string              `json:"project_id" validate:"regexp=^[a-z]+$"`
	ProjectName string              `json:"project_name"`
	Version     string              `json:"version" validate:"regexp=^[0-9.]+-?(?:pre|SNAPSHOT)?(?:[0-9.]+)?$"`
	Build       int32               `json:"build"`
	Time        time.Time           `json:"time"`
	Channel     ChannelType         `json:"channel"`
	Promoted    bool                `json:"promoted"`
	Changes     []Change            `json:"changes"`
	Downloads   map[string]Download `json:"downloads"`
}

// BuildsResponse represents builds response
type BuildsResponse struct {
	ProjectID   string         `json:"project_id" validate:"regexp=^[a-z]+$"`
	ProjectName string         `json:"project_name"`
	Version     string         `json:"version" validate:"regexp=^[0-9.]+-?(?:pre|SNAPSHOT)?(?:[0-9.]+)?$"`
	Builds      []VersionBuild `json:"builds"`
}

// Change represents a change object
type Change struct {
	Commit  string `json:"commit"`
	Summary string `json:"summary"`
	Message string `json:"message"`
}

// Download represents a download object
type Download struct {
	Name   string `json:"name" validate:"regexp=^[a-z0-9._-]+$"`
	Sha256 string `json:"sha256" validate:"regexp=^[a-f0-9]{64}$"`
}

// ProjectResponse represents a project response
type ProjectResponse struct {
	ProjectID     string   `json:"project_id" validate:"regexp=^[a-z]+$"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

// ProjectsResponse represents list of projects
type ProjectsResponse struct {
	Projects []string `json:"projects"`
}

// VersionBuild represents build info in a version
type VersionBuild struct {
	Build     int32               `json:"build"`
	Time      time.Time           `json:"time"`
	Channel   ChannelType         `json:"channel"`
	Promoted  bool                `json:"promoted"`
	Changes   []Change            `json:"changes"`
	Downloads map[string]Download `json:"downloads"`
}

// VersionFamilyBuild represents build info in a version family
type VersionFamilyBuild struct {
	Version   string              `json:"version" validate:"regexp=^[0-9.]+-?(?:pre|SNAPSHOT)?(?:[0-9.]+)?$"`
	Build     int32               `json:"build"`
	Time      time.Time           `json:"time"`
	Channel   ChannelType         `json:"channel"`
	Promoted  bool                `json:"promoted"`
	Changes   []Change            `json:"changes"`
	Downloads map[string]Download `json:"downloads"`
}

// VersionFamilyBuildsResponse represents builds for a version family
type VersionFamilyBuildsResponse struct {
	ProjectID    string               `json:"project_id" validate:"regexp=^[a-z]+$"`
	ProjectName  string               `json:"project_name"`
	VersionGroup string               `json:"version_group" validate:"regexp=^[0-9.]+-?(?:pre|SNAPSHOT)?(?:[0-9.]+)?$"`
	Versions     []string             `json:"versions"`
	Builds       []VersionFamilyBuild `json:"builds"`
}

// VersionFamilyResponse represents version family response
type VersionFamilyResponse struct {
	ProjectID    string   `json:"project_id" validate:"regexp=^[a-z]+$"`
	ProjectName  string   `json:"project_name"`
	VersionGroup string   `json:"version_group" validate:"regexp=^[0-9.]+-?(?:pre|SNAPSHOT)?(?:[0-9.]+)?$"`
	Versions     []string `json:"versions"`
}

// VersionResponse represents version response
type VersionResponse struct {
	ProjectID   string  `json:"project_id" validate:"regexp=^[a-z]+$"`
	ProjectName string  `json:"project_name"`
	Version     string  `json:"version" validate:"regexp=^[0-9.]+-?(?:pre|SNAPSHOT)?(?:[0-9.]+)?$"`
	Builds      []int32 `json:"builds"`
}

// ChannelType defines enum for channel
type ChannelType string

const (
	ChannelDefault      ChannelType = "default"
	ChannelExperimental ChannelType = "experimental"
)
