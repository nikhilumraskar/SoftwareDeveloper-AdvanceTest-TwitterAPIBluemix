package main

import (
	//"flag"
	"fmt"
	//"log"
	//"os"
	//"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	 "encoding/json"
	 "bytes"
	"io/ioutil"
	//"reflect"
	//"golang.org/x/oauth2"
	"net/http"
)

type Payload struct {
	Text string `json:"text"`
}

type twitter_struct struct {
	Statuses []struct {
		Contributors interface{} `json:"contributors"`
		Coordinates interface{} `json:"coordinates"`
		CreatedAt string `json:"created_at"`
		CurrentUserRetweet interface{} `json:"current_user_retweet"`
		Entities struct {
			Hashtags []interface{} `json:"hashtags"`
			Media []struct {
				Indices []int `json:"indices"`
				DisplayURL string `json:"display_url"`
				ExpandedURL string `json:"expanded_url"`
				URL string `json:"url"`
				ID int64 `json:"id"`
				IDStr string `json:"id_str"`
				MediaURL string `json:"media_url"`
				MediaURLHTTPS string `json:"media_url_https"`
				SourceStatusID int64 `json:"source_status_id"`
				SourceStatusIDStr string `json:"source_status_id_str"`
				Type string `json:"type"`
				Sizes struct {
					Thumb struct {
						W int `json:"w"`
						H int `json:"h"`
						Resize string `json:"resize"`
					} `json:"thumb"`
					Large struct {
						W int `json:"w"`
						H int `json:"h"`
						Resize string `json:"resize"`
					} `json:"large"`
					Medium struct {
						W int `json:"w"`
						H int `json:"h"`
						Resize string `json:"resize"`
					} `json:"medium"`
					Small struct {
						W int `json:"w"`
						H int `json:"h"`
						Resize string `json:"resize"`
					} `json:"small"`
				} `json:"sizes"`
				VideoInfo struct {
					AspectRatio []int `json:"aspect_ratio"`
					DurationMillis int `json:"duration_millis"`
					Variants interface{} `json:"variants"`
				} `json:"video_info"`
			} `json:"media"`
			Urls []interface{} `json:"urls"`
			UserMentions []interface{} `json:"user_mentions"`
		} `json:"entities"`
		FavoriteCount int `json:"favorite_count"`
		Favorited bool `json:"favorited"`
		FilterLevel string `json:"filter_level"`
		ID int64 `json:"id"`
		IDStr string `json:"id_str"`
		InReplyToScreenName string `json:"in_reply_to_screen_name"`
		InReplyToStatusID int `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
		InReplyToUserID int `json:"in_reply_to_user_id"`
		InReplyToUserIDStr string `json:"in_reply_to_user_id_str"`
		Lang string `json:"lang"`
		PossiblySensitive bool `json:"possibly_sensitive"`
		RetweetCount int `json:"retweet_count"`
		Retweeted bool `json:"retweeted"`
		RetweetedStatus interface{} `json:"retweeted_status"`
		Source string `json:"source"`
		Scopes interface{} `json:"scopes"`
		Text string `json:"text"`
		Place interface{} `json:"place"`
		Truncated bool `json:"truncated"`
		User struct {
			ContributorsEnabled bool `json:"contributors_enabled"`
			CreatedAt string `json:"created_at"`
			DefaultProfile bool `json:"default_profile"`
			DefaultProfileImage bool `json:"default_profile_image"`
			Description string `json:"description"`
			Email string `json:"email"`
			Entities struct {
				URL struct {
					Hashtags interface{} `json:"hashtags"`
					Media interface{} `json:"media"`
					Urls interface{} `json:"urls"`
					UserMentions interface{} `json:"user_mentions"`
				} `json:"url"`
				Description struct {
					Hashtags interface{} `json:"hashtags"`
					Media interface{} `json:"media"`
					Urls []interface{} `json:"urls"`
					UserMentions interface{} `json:"user_mentions"`
				} `json:"description"`
			} `json:"entities"`
			FavouritesCount int `json:"favourites_count"`
			FollowRequestSent bool `json:"follow_request_sent"`
			Following bool `json:"following"`
			FollowersCount int `json:"followers_count"`
			FriendsCount int `json:"friends_count"`
			GeoEnabled bool `json:"geo_enabled"`
			ID int64 `json:"id"`
			IDStr string `json:"id_str"`
			IDTranslator bool `json:"id_translator"`
			Lang string `json:"lang"`
			ListedCount int `json:"listed_count"`
			Location string `json:"location"`
			Name string `json:"name"`
			Notifications bool `json:"notifications"`
			ProfileBackgroundColor string `json:"profile_background_color"`
			ProfileBackgroundImageURL string `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
			ProfileBackgroundTile bool `json:"profile_background_tile"`
			ProfileBannerURL string `json:"profile_banner_url"`
			ProfileImageURL string `json:"profile_image_url"`
			ProfileImageURLHTTPS string `json:"profile_image_url_https"`
			ProfileLinkColor string `json:"profile_link_color"`
			ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor string `json:"profile_sidebar_fill_color"`
			ProfileTextColor string `json:"profile_text_color"`
			ProfileUseBackgroundImage bool `json:"profile_use_background_image"`
			Protected bool `json:"protected"`
			ScreenName string `json:"screen_name"`
			ShowAllInlineMedia bool `json:"show_all_inline_media"`
			Status interface{} `json:"status"`
			StatusesCount int `json:"statuses_count"`
			TimeZone string `json:"time_zone"`
			URL string `json:"url"`
			UtcOffset int `json:"utc_offset"`
			Verified bool `json:"verified"`
			WithheldInCountries string `json:"withheld_in_countries"`
			WithheldScope string `json:"withheld_scope"`
		} `json:"user"`
		WithheldCopyright bool `json:"withheld_copyright"`
		WithheldInCountries interface{} `json:"withheld_in_countries"`
		WithheldScope string `json:"withheld_scope"`
		ExtendedEntities struct {
			Media []struct {
				Indices []int `json:"indices"`
				DisplayURL string `json:"display_url"`
				ExpandedURL string `json:"expanded_url"`
				URL string `json:"url"`
				ID int64 `json:"id"`
				IDStr string `json:"id_str"`
				MediaURL string `json:"media_url"`
				MediaURLHTTPS string `json:"media_url_https"`
				SourceStatusID int64 `json:"source_status_id"`
				SourceStatusIDStr string `json:"source_status_id_str"`
				Type string `json:"type"`
				Sizes struct {
					Thumb struct {
						W int `json:"w"`
						H int `json:"h"`
						Resize string `json:"resize"`
					} `json:"thumb"`
					Large struct {
						W int `json:"w"`
						H int `json:"h"`
						Resize string `json:"resize"`
					} `json:"large"`
					Medium struct {
						W int `json:"w"`
						H int `json:"h"`
						Resize string `json:"resize"`
					} `json:"medium"`
					Small struct {
						W int `json:"w"`
						H int `json:"h"`
						Resize string `json:"resize"`
					} `json:"small"`
				} `json:"sizes"`
				VideoInfo struct {
					AspectRatio []int `json:"aspect_ratio"`
					DurationMillis int `json:"duration_millis"`
					Variants interface{} `json:"variants"`
				} `json:"video_info"`
			} `json:"media"`
		} `json:"extended_entities"`
		QuotedStatusID int `json:"quoted_status_id"`
		QuotedStatusIDStr string `json:"quoted_status_id_str"`
		QuotedStatus interface{} `json:"quoted_status"`
	} `json:"statuses"`
	SearchMetadata struct {
		Count int `json:"count"`
		SinceID int `json:"since_id"`
		SinceIDStr string `json:"since_id_str"`
		MaxID int64 `json:"max_id"`
		MaxIDStr string `json:"max_id_str"`
		RefreshURL string `json:"refresh_url"`
		NextResults string `json:"next_results"`
		CompletedIn float64 `json:"completed_in"`
		Query string `json:"query"`
	} `json:"search_metadata"`
}

type watson_struct struct {
	DocumentTone struct {
		ToneCategories []struct {
			Tones []struct {
				Score float64 `json:"score"`
				ToneID string `json:"tone_id"`
				ToneName string `json:"tone_name"`
			} `json:"tones"`
			CategoryID string `json:"category_id"`
			CategoryName string `json:"category_name"`
		} `json:"tone_categories"`
	} `json:"document_tone"`
	SentencesTone []struct {
		SentenceID int `json:"sentence_id"`
		Text string `json:"text"`
		InputFrom int `json:"input_from"`
		InputTo int `json:"input_to"`
		ToneCategories []struct {
			Tones []struct {
				Score float64 `json:"score"`
				ToneID string `json:"tone_id"`
				ToneName string `json:"tone_name"`
			} `json:"tones"`
			CategoryID string `json:"category_id"`
			CategoryName string `json:"category_name"`
		} `json:"tone_categories"`
	} `json:"sentences_tone"`
}

type final_display struct {
	twitter_struct
	watson_struct
}

type final_display_arr []final_display

func main() {
	config := oauth1.NewConfig("", "")
	token := oauth1.NewToken("", "")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)



	// Search Tweets
	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
	    Query: "#Goa",
	    Count: 1,
	})
	tweet_res,_ :=json.Marshal(search)
	
	tweet_byte := bytes.NewReader(tweet_res)
	tweet_str, _ := ioutil.ReadAll(tweet_byte)
	byte1 := []byte(tweet_str)
	var oggy twitter_struct
    json.Unmarshal(byte1, &oggy)
    //fmt.Println(oggy)
    display_arr := final_display{
    	
    }
    for l := range oggy.Statuses {
    	fmt.Println("========================")
		fmt.Println()
		fmt.Printf(oggy.Statuses[l].Text)

		data :=  Payload{
	        Text: oggy.Statuses[l].Text,
	    }

	    payloadBytes, err := json.Marshal(data)
		if err != nil {
			// handle err
		}
		body := bytes.NewReader(payloadBytes)

		req, err := http.NewRequest("POST", "https://gateway.watsonplatform.net/tone-analyzer/api/v3/tone?version=2016-05-19", body)
		if err != nil {
			// handle err
		}
		req.SetBasicAuth("", "")
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// handle err
		}
		//defer resp.Body.Close
		body1, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body1))
		fmt.Println()
		fmt.Println("========================endhere")
		fmt.Println()

    }


}