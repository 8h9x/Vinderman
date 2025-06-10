package vinderman

import (
	"gitlab.com/8h9x/vinderman/consts"
	"gitlab.com/8h9x/vinderman/request"
	"net/http"
	"time"
)

type FortniteMCPVersion struct {
	App                       string    `json:"app"`
	ServerDate                time.Time `json:"serverDate"`
	OverridePropertiesVersion string    `json:"overridePropertiesVersion"`
	Cln                       string    `json:"cln"`
	Build                     string    `json:"build"`
	ModuleName                string    `json:"moduleName"`
	BuildDate                 time.Time `json:"buildDate"`
	Version                   string    `json:"version"`
	Branch                    string    `json:"branch"`
	Modules                   struct {
		EpicLightSwitchAccessControlCore struct {
			Cln       string    `json:"cln"`
			Build     string    `json:"build"`
			BuildDate time.Time `json:"buildDate"`
			Version   string    `json:"version"`
			Branch    string    `json:"branch"`
		} `json:"Epic-LightSwitch-AccessControlCore"`
		EpicCommonCore struct {
			Cln       string    `json:"cln"`
			Build     string    `json:"build"`
			BuildDate time.Time `json:"buildDate"`
			Version   string    `json:"version"`
			Branch    string    `json:"branch"`
		} `json:"epic-common-core"`
	} `json:"modules"`
}

func (c *Client) GetMCPVersion() (FortniteMCPVersion, error) {
	req, err := http.NewRequest("GET", consts.FortniteMCPService+"/fortnite/api/version", nil)
	if err != nil {
		return FortniteMCPVersion{}, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return FortniteMCPVersion{}, err
	}

	res, err := request.ResponseParser[FortniteMCPVersion](resp)
	if err != nil {
		return FortniteMCPVersion{}, err
	}

	return res.Body, err
}
