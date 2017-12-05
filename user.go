package confluence

type User struct {
	Key            string             `json:"userKey"`
	Username       string             `json:"username"`
	DisplayName    string             `json:"displayName"`
	ProfilePicture UserProfilePicture `json:"profilePicture"`
	Links          Links              `json:"_links"`
}

type UserProfilePicture struct {
	Path      string `json:"path"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsDefault bool   `json:"is_default"`
}
