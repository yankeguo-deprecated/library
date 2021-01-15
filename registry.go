package library

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type registryListTagResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type dockerHubTagItem struct {
	Layer string `json:"layer"`
	Name  string `json:"name"`
}

func RegistryListTags(ctx context.Context, repo string) (tags []string, err error) {
	comps := strings.Split(repo, "/")
	if len(comps) < 2 {
		err = fmt.Errorf("bad format for docker repository: %s", repo)
		return
	}
	isHub := !strings.Contains(comps[0], ".")
	var urlList string
	if isHub {
		urlList = "https://registry.hub.docker.com/v1/repositories/" + strings.Join(comps, "/") + "/tags"
	} else {
		urlList = "https://" + comps[0] + "/v2/" + strings.Join(comps[1:], "/") + "/tags/list"
	}
	var req *http.Request
	var res *http.Response
	if req, err = http.NewRequest(http.MethodGet, urlList, nil); err != nil {
		return
	}
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad response: %s", res.Status)
		return
	}
	if isHub {
		var body []dockerHubTagItem
		if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
			return
		}
		for _, item := range body {
			tags = append(tags, item.Name)
		}
	} else {
		var body registryListTagResponse
		if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
			return
		}
		tags = body.Tags
	}
	return
}
