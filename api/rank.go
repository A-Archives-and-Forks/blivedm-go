package api

import (
	"fmt"
	"net/http"
)

type ContributionRankResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Msg     string               `json:"msg"`
	Data    ContributionRankData `json:"data"`
}

type ContributionRankData struct {
	Count int                    `json:"count"`
	Item  []ContributionRankUser `json:"item"`
}

type ContributionRankUser struct {
	Uid        int64                  `json:"uid"`
	Name       string                 `json:"name"`
	Face       string                 `json:"face"`
	Rank       int                    `json:"rank"`
	Score      int64                  `json:"score"`
	GuardLevel int                    `json:"guard_level"`
	MedalInfo  *ContributionMedalInfo `json:"medal_info"`
}

type ContributionMedalInfo struct {
	MedalName       string `json:"medal_name"`
	Level           int    `json:"level"`
	MedalColorStart int    `json:"medal_color_start"`
}

func (m *ContributionMedalInfo) MedalColor() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("#%06x", m.MedalColorStart)
}

func GetContributionRank(roomID, ruid int, cookie string, page, pageSize int) (*ContributionRankResponse, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 100
	}

	rawURL := fmt.Sprintf(
		"https://api.live.bilibili.com/xlive/general-interface/v1/rank/queryContributionRank?ruid=%d&room_id=%d&page=%d&page_size=%d&type=online_rank&switch=contribution_rank&platform=web&web_location=0.0",
		ruid, roomID, page, pageSize,
	)
	signedURL, err := WbiKeysSignString(rawURL)
	if err != nil {
		return nil, err
	}

	result := &ContributionRankResponse{}
	headers := &http.Header{}
	headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:137.0) Gecko/20100101 Firefox/137.0")
	headers.Set("Referer", "https://live.bilibili.com/")
	if cookie != "" {
		headers.Set("Cookie", cookie)
	}
	if err := GetJsonWithHeader(signedURL, headers, result); err != nil {
		return nil, err
	}
	return result, nil
}
