package vinderman

import (
	"fmt"
	"github.com/0xDistrust/Vinderman/consts"
	"net/http"
)

func (c Client) QueryLockerItems(credentials UserCredentials, deploymentId string) (resp *http.Response, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err = c.Request("GET", fmt.Sprintf("%s/api/locker/v4/%s/account/%s/items", consts.FORTNITE_SERVICE, deploymentId, credentials.AccountID), headers, "")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("failed to favorite mnemonic: %s", resp.Status)
	}

	return nil, nil
	// https://fngw-svc-gc-livefn.ol.epicgames.com/api/locker/v4/:deploymentId/account/:accountId/items
}

func (c Client) UpdateActiveLockerLoadout(credentials UserCredentials, deploymentId string) (resp *http.Response, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err = c.Request("PUT", fmt.Sprintf("%s/api/locker/v4/%s/account/%s/active-loadout-group", consts.FORTNITE_SERVICE, deploymentId, credentials.AccountID), headers, "{}")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("failed to favorite mnemonic: %s", resp.Status)
	}

	return nil, nil
	// https://fngw-svc-gc-livefn.ol.epicgames.com/api/locker/v4/:deploymentId/account/:accountId/active-loadout-group
}
