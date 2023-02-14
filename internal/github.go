package internal

import (
	"context"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func GetContributions(username, token string) (Calendar, error) {
	tokenSrc := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client := githubv4.NewClient(oauth2.NewClient(context.Background(), tokenSrc))
	var q struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					//TotalContributions int
					Weeks []struct {
						ContributionDays []struct {
							//ContributionCount int
							//Date              string
							//Color             string
							ContributionLevel string
						}
					}
				}
			}
		} `graphql:"user(login: $username)"`
	}
	variables := map[string]interface{}{
		"username": githubv4.String(username),
	}
	if err := client.Query(context.Background(), &q, variables); err != nil {
		return nil, err
	}

	weeks := q.User.ContributionsCollection.ContributionCalendar.Weeks
	cal := make([][]int8, len(weeks))
	for i := range weeks {
		cal[i] = make([]int8, len(weeks[i].ContributionDays))
		for j := range cal[i] {
			switch weeks[i].ContributionDays[j].ContributionLevel {
			case "NONE":
				cal[i][j] = 0
			case "FIRST_QUARTILE":
				cal[i][j] = 1
			case "SECOND_QUARTILE":
				cal[i][j] = 2
			case "THIRD_QUARTILE":
				cal[i][j] = 3
			case "FOURTH_QUARTILE":
				cal[i][j] = 4
			}
		}
	}

	return cal, nil
}
