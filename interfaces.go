package main

// GameStart is from game_start.json
type GameStart struct {
	PlayerIndex  int64    `json:"playerIndex"`
	ReplayID     string   `json:"replay_id"`
	ChatRoom     string   `json:"chat_room"`
	TeamChatRoom string   `json:"team_chat_room"`
	Usernames    []string `json:"usernames"`
	Teams        []int64  `json:"teams"`
}

// GameUpdate is from game_update.json
type GameUpdate struct {
	Scores []struct {
		Total int64 `json:"total"`
		Tiles int64 `json:"tiles"`
		I     int64 `json:"i"`
		Dead  bool  `json:"dead"`
	} `json:"scores"`
	Turn        int64   `json:"turn"`
	Stars       []int64 `json:"stars"`
	AttackIndex int64   `json:"attackIndex"`
	Generals    []int64 `json:"generals"`
	MapDiff     []int64 `json:"map_diff"`
	CityDiff    []int64 `json:"city_diff"`
}
