package nitestats

import (
	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type FLToken struct {
	Version string `json:"version"`
	Token   string `json:"fltoken"`
}

func (c Client) FetchFLToken() (FLToken, error) {
	res, err := request.Getf[FLToken]("%s/builds/fltoken", consts.NITESTATS_API)
	return res.Body, err
}
