package papermc

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	API_BASE_URL              = "https://api.papermc.io/v2"
	httpClient   *http.Client = &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   time.Second,
			ResponseHeaderTimeout: time.Second * 5,
			IdleConnTimeout:       time.Second,
		},
	}
)

func httpGetAndUnmarshal(path string, v any) error {
	res, err := httpClient.Get(API_BASE_URL + path)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func joinPath(path ...string) string {
	return strings.Join(path, "/")
}

func GetAllAvailableProjects() (*ProjectsResponse, error) {
	var data ProjectsResponse
	return &data, httpGetAndUnmarshal("/projects", &data)
}

func GetProjectInfo(project string) (*ProjectResponse, error) {
	var data ProjectResponse
	return &data, httpGetAndUnmarshal(
		joinPath("/projects", project),
		&data,
	)
}

func GetPorjectVersionInfo(project, version string) (*VersionResponse, error) {
	var data VersionResponse
	return &data, httpGetAndUnmarshal(
		joinPath("/projects", project, "versions", version),
		&data,
	)
}

func GetProjectBuilds(project, version string) (*BuildsResponse, error) {
	var data BuildsResponse
	return &data, httpGetAndUnmarshal(
		joinPath("/projects", project, "versions", version, "builds"),
		&data,
	)
}

func GetProjectBuildInfo(project, version, build string) (*BuildResponse, error) {
	var data BuildResponse
	return &data, httpGetAndUnmarshal(
		joinPath("/projects", project, "versions", version, "builds", build),
		&data,
	)
}

func DownloadJar(project, version, build, download string, filepath string) error {
	res, err := httpClient.Get(API_BASE_URL + joinPath("/projects", project, "versions", version, "builds", build, "downloads", download))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}

	// if we require some debug messages, add it here.
	return nil
}
