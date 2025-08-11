package structs

type TicketsData struct {
	Total  int `json:"total"`
	Issues []struct {
		Id     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Severity struct {
				Value string `json:"value"`
				Id    string `json:"id"`
			} `json:"customfield_17130"`

			Priority struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"priority"`

			Assignee struct {
				DisplayName string `json:"displayName"`
			} `json:"assignee"`

			Status struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Id          string `json:"id"`

				StatusCategory struct {
					Id   int    `json:"id"`
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`

			Components []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"components"`

			Creator struct {
				Key         string `json:"key"`
				DisplayName string `json:"displayName"`
			} `json:"creator"`

			Reporter struct {
				Key         string `json:"key"`
				DisplayName string `json:"displayName"`
			} `json:"reporter"`

			Issuetype struct {
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"issuetype"`

			Created     string `json:"created"`
			Updated     string `json:"updated"`
			Description string `json:"description"`
			Summary     string `json:"summary"`
		} `json:"fields"`
	} `json:"issues"`
}
