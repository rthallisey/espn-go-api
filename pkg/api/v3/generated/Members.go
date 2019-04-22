// This file is generated by running `make generateAPI`
package generated

type Members struct {
	Members []struct {
		DisplayName          string `json:"displayName"`
		FirstName            string `json:"firstName"`
		ID                   string `json:"id"`
		IsLeagueCreator      bool   `json:"isLeagueCreator"`
		IsLeagueManager      bool   `json:"isLeagueManager"`
		LastName             string `json:"lastName"`
		NotificationSettings []struct {
			Enabled bool   `json:"enabled"`
			ID      string `json:"id"`
			Type    string `json:"type"`
		} `json:"notificationSettings"`
	} `json:"members"`
}