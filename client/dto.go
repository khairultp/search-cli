package client

type FriendDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserDto struct {
	Id       string       `json:"_id"`
	Index    int          `json:"index"`
	Guid     string       `json:"guid"`
	IsActive bool         `json:"isActive"`
	Balance  string       `json:"balance"`
	Tags     []string     `json:"tags"`
	Friends  []*FriendDto `json:"friends"`
}
