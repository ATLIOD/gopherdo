package update

import (
	"encoding/json"
	"net/http"
)

type Release struct {
    Version string `json:"tag_name"`
    URL     string `json:"html_url"`
}

func CheckForUpdates(currentVersion string) (*Release, error) {
    resp, err := http.Get("https://api.github.com/repos/yourusername/gopherdo/releases/latest")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var release Release
    if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
        return nil, err
    }

    if release.Version != currentVersion {
        return &release, nil
    }
    return nil, nil
}
