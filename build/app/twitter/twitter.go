package twitter

import (
	"context"
	"fmt"

	"app/models"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/direktiv/apps/go/pkg/apps"
)

func DoTwitter(ctx context.Context, body *models.PostParamsBody, ri *apps.RequestInfo) (map[string]interface{}, error) {

	ret := make(map[string]interface{})

	config := oauth1.NewConfig(*body.Apikey, *body.Apisecret)
	token := oauth1.NewToken(*body.Accesstoken, *body.Accesstokensecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	action := ""
	if body.Action == nil {
		action = "search"
	} else {
		action = *body.Action
	}

	if action == "search" {

		// search
		if body.Query == nil {
			return ret, fmt.Errorf("no query specified")
		}

		if body.Query.Limit == 0 {
			body.Query.Limit = 10
		}

		result, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
			Query: body.Query.Term,
			Count: int(body.Query.Limit),
		})

		if err != nil {
			return ret, err
		}

		if resp.Status != "200 OK" {
			return ret, fmt.Errorf("request failed: %v", resp.Status)
		}

		ret["search"] = result

	} else {

		ri.Logger().Infof("posting tweet")

		// no media supported atm
		tweet, resp, err := client.Statuses.Update(body.Content, nil)
		if err != nil {
			return ret, err
		}

		if resp.Status != "200 OK" {
			return ret, fmt.Errorf("request failed: %v", resp.Status)
		}

		fmt.Printf("TWEET %+v\n", resp)

		ret["tweet"] = tweet

	}

	return ret, nil
}
