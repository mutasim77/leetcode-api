package model

type QueryLeetCodeResponse struct {
	Data struct {
		MatchedUser struct {
			Username    string `json:"username"`
			GithubUrl   string `json:"githubUrl"`
			TwitterUrl  string `json:"twitterUrl"`
			LinkedinUrl string `json:"linkedinUrl"`
			Profile     struct {
				Ranking       int      `json:"ranking"`
				UserAvatar    string   `json:"userAvatar"`
				RealName      string   `json:"realName"`
				AboutMe       string   `json:"aboutMe"`
				School        string   `json:"school"`
				Websites      []string `json:"websites"`
				CountryName   string   `json:"countryName"`
				Company       string   `json:"company"`
				JobTitle      string   `json:"jobTitle"`
				SkillTags     []string `json:"skillTags"`
				Reputation    int      `json:"reputation"`
				SolutionCount int      `json:"solutionCount"`
			} `json:"profile"`
			LanguageProblemCount []struct {
				LanguageName   string `json:"languageName"`
				ProblemsSolved int    `json:"problemsSolved"`
			} `json:"languageProblemCount"`
			SubmitStatsGlobal struct {
				AcSubmissionNum []struct {
					Difficulty string `json:"difficulty"`
					Count      int    `json:"count"`
				} `json:"acSubmissionNum"`
			} `json:"submitStatsGlobal"`
			TagProblemCounts struct {
				Advanced []struct {
					TagName        string `json:"tagName"`
					TagSlug        string `json:"tagSlug"`
					ProblemsSolved int    `json:"problemsSolved"`
				} `json:"advanced"`
				Intermediate []struct {
					TagName        string `json:"tagName"`
					TagSlug        string `json:"tagSlug"`
					ProblemsSolved int    `json:"problemsSolved"`
				} `json:"intermediate"`
				Fundamental []struct {
					TagName        string `json:"tagName"`
					TagSlug        string `json:"tagSlug"`
					ProblemsSolved int    `json:"problemsSolved"`
				} `json:"fundamental"`
			} `json:"tagProblemCounts"`
			ProblemsSolvedBeatsStats []struct {
				Difficulty string  `json:"difficulty"`
				Percentage float64 `json:"percentage"`
			} `json:"problemsSolvedBeatsStats"`
			Badges []struct {
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				DisplayName string `json:"displayName"`
			} `json:"badges"`
			UserCalendar struct {
				Streak          int `json:"streak"`
				TotalActiveDays int `json:"totalActiveDays"`
			} `json:"userCalendar"`
		} `json:"matchedUser"`
		UserContestRanking struct {
			AttendedContestsCount int     `json:"attendedContestsCount"`
			Rating                float64 `json:"rating"`
			GlobalRanking         int     `json:"globalRanking"`
			TotalParticipants     int     `json:"totalParticipants"`
			TopPercentage         float64 `json:"topPercentage"`
		} `json:"userContestRanking"`
	} `json:"data"`
}
