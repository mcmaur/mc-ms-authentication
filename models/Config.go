package models

// Config : configurations settings
type Config struct {
	SocialCallbackBaseURL     string
	CookieStore               string
	TwitterKey                string
	TwitterSecret             string
	FacebookKey               string
	FacebookSecret            string
	FitbitKey                 string
	FitbitSecret              string
	GoogleKey                 string
	GoogleSecret              string
	GplusKey                  string
	GplusSecret               string
	GithubKey                 string
	GithubSecret              string
	SpotifyKey                string
	SpotifySecret             string
	LinkedinKey               string
	LinkedinSecret            string
	LineKey                   string
	LineSecret                string
	LastfmKey                 string
	LastfmSecret              string
	TwitchKey                 string
	TwitchSecret              string
	DropboxKey                string
	DropboxSecret             string
	DigitaloceanKey           string
	DigitaloceanSecret        string
	BitbucketKey              string
	BitbucketSecret           string
	InstagramKey              string
	InstagramSecret           string
	IntercomKey               string
	IntercomSecret            string
	BoxKey                    string
	BoxSecret                 string
	SalesforceKey             string
	SalesforceSecret          string
	SeatalkKey                string
	SeatalkSecret             string
	AmazonKey                 string
	AmazonSecret              string
	YammerKey                 string
	YammerSecret              string
	OnedriveKey               string
	OnedriveSecret            string
	AzureadKey                string
	AzureadSecret             string
	MicrosoftonlineKey        string
	MicrosoftonlineSecret     string
	BattlenetKey              string
	BattlenetSecret           string
	EveonlineKey              string
	EveonlineSecret           string
	KakaoKey                  string
	KakaoSecret               string
	YahooKey                  string
	YahooSecret               string
	TypetalkKey               string
	TypetalkSecret            string
	SlackKey                  string
	SlackSecret               string
	StripeKey                 string
	StripeSecret              string
	WepayKey                  string
	WepaySecret               string
	PaypalKey                 string
	PaypalSecret              string
	SteamKey                  string
	HerokuKey                 string
	HerokuSecret              string
	UberKey                   string
	UberSecret                string
	SoundcloudKey             string
	SoundcloudSecret          string
	GitlabKey                 string
	GitlabSecret              string
	DailymotionKey            string
	DailymotionSecret         string
	DeezerKey                 string
	DeezerSecret              string
	DiscordKey                string
	DiscordSecret             string
	MeetupKey                 string
	MeetupSecret              string
	Auth0Key                  string
	Auth0Secret               string
	Auth0Domain               string
	XeroKey                   string
	XeroSecret                string
	VkKey                     string
	VkSecret                  string
	NaverKey                  string
	NaverSecret               string
	YandexKey                 string
	YandexSecret              string
	NextcloudKey              string
	NextcloudSecret           string
	NextcloudURL              string
	GiteaKey                  string
	GiteaSecret               string
	ShopifyKey                string
	ShopifySecret             string
	AppleKey                  string
	AppleSecret               string
	StravaKey                 string
	StravaSecret              string
	OpenidConnectKey          string
	OpenidConnectSecret       string
	OpenidConnectDiscoveryURL string
	DB                        Database `toml:"Database"`
}

// Database : dabatase connection infos
type Database struct {
	User     string
	Password string
	DbName   string
}
