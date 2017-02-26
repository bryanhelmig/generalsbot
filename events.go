package main

// GameStart is from game_start.json
type GameStart struct {
	PlayerIndex  int      `json:"playerIndex"`
	ReplayID     string   `json:"replay_id"`
	ChatRoom     string   `json:"chat_room"`
	TeamChatRoom string   `json:"team_chat_room"`
	Usernames    []string `json:"usernames"`
	Teams        []int    `json:"teams"`
}

// GameUpdate is from game_update.json
type GameUpdate struct {
	Scores []struct {
		Total int  `json:"total"`
		Tiles int  `json:"tiles"`
		I     int  `json:"i"`
		Dead  bool `json:"dead"`
	} `json:"scores"`
	Turn        int   `json:"turn"`
	Stars       []int `json:"stars"`
	AttackIndex int   `json:"attackIndex"`
	Generals    []int `json:"generals"`
	MapDiff     []int `json:"map_diff"`
	CityDiff    []int `json:"city_diff"`
}
