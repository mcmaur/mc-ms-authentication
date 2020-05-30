package server

import (
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/mcmaur/mc-ms-authentication/server/controllers"
	"github.com/mcmaur/mc-ms-authentication/server/models"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for database support

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/amazon"
	"github.com/markbates/goth/providers/apple"
	"github.com/markbates/goth/providers/auth0"
	"github.com/markbates/goth/providers/azuread"
	"github.com/markbates/goth/providers/battlenet"
	"github.com/markbates/goth/providers/bitbucket"
	"github.com/markbates/goth/providers/box"
	"github.com/markbates/goth/providers/dailymotion"
	"github.com/markbates/goth/providers/deezer"
	"github.com/markbates/goth/providers/digitalocean"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/dropbox"
	"github.com/markbates/goth/providers/eveonline"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/fitbit"
	"github.com/markbates/goth/providers/gitea"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/heroku"
	"github.com/markbates/goth/providers/instagram"
	"github.com/markbates/goth/providers/intercom"
	"github.com/markbates/goth/providers/kakao"
	"github.com/markbates/goth/providers/lastfm"
	"github.com/markbates/goth/providers/line"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/meetup"
	"github.com/markbates/goth/providers/microsoftonline"
	"github.com/markbates/goth/providers/naver"
	"github.com/markbates/goth/providers/nextcloud"
	"github.com/markbates/goth/providers/onedrive"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/markbates/goth/providers/paypal"
	"github.com/markbates/goth/providers/salesforce"
	"github.com/markbates/goth/providers/seatalk"
	"github.com/markbates/goth/providers/shopify"
	"github.com/markbates/goth/providers/slack"
	"github.com/markbates/goth/providers/soundcloud"
	"github.com/markbates/goth/providers/spotify"
	"github.com/markbates/goth/providers/steam"
	"github.com/markbates/goth/providers/strava"
	"github.com/markbates/goth/providers/stripe"
	"github.com/markbates/goth/providers/twitch"
	"github.com/markbates/goth/providers/twitter"
	"github.com/markbates/goth/providers/typetalk"
	"github.com/markbates/goth/providers/uber"
	"github.com/markbates/goth/providers/vk"
	"github.com/markbates/goth/providers/wepay"
	"github.com/markbates/goth/providers/xero"
	"github.com/markbates/goth/providers/yahoo"
	"github.com/markbates/goth/providers/yammer"
	"github.com/markbates/goth/providers/yandex"
)

// DB : db connection
//var DB *gorm.DB
//var providerIndex *ProviderIndex
var err error
var store *sessions.CookieStore

var server = controllers.Server{}

// Start : starts the server
func Start() {

	//var config models.Config
	if _, err = toml.DecodeFile("env.toml", &server.Config); err != nil {
		panic("Failed to read enviroment settings")
	}

	os.Setenv("JWTOKEN_SECRET", server.Config.JwtTokenSecret) // TODO improve it

	databaseConnectionSettings := "host=" + server.Config.DB.Host + " port=" + server.Config.DB.Port + " user=" + server.Config.DB.User + " dbname=" + server.Config.DB.DbName + " password=" + server.Config.DB.Password + " sslmode=disable"
	server.DB, err = gorm.Open("postgres", databaseConnectionSettings)
	if err != nil {
		log.Println("DEBUG: ", databaseConnectionSettings)
		log.Println("ERR: ", err)
		panic("failed to connect database")
	}
	defer server.DB.Close()

	// Migrate the schema
	server.DB.AutoMigrate(&models.User{})

	gothic.Store = sessions.NewCookieStore([]byte(server.Config.CookieStore))
	store = sessions.NewCookieStore([]byte(server.Config.CookieStore))

	goth.UseProviders(
		twitter.New(server.Config.TwitterKey, server.Config.TwitterSecret, server.Config.SocialCallbackBaseURL+"/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(server.Config.TWITTER_KEY, server.Config.TWITTER_SECRET, server.Config.SocialCallbackBaseURL+"/auth/twitter/callback"),

		facebook.New(server.Config.FacebookKey, server.Config.FacebookSecret, server.Config.SocialCallbackBaseURL+"/auth/facebook/callback"),
		fitbit.New(server.Config.FitbitKey, server.Config.FitbitSecret, server.Config.SocialCallbackBaseURL+"/auth/fitbit/callback"),
		google.New(server.Config.GoogleKey, server.Config.GoogleSecret, server.Config.SocialCallbackBaseURL+"/auth/google/callback"),
		gplus.New(server.Config.GplusKey, server.Config.GplusSecret, server.Config.SocialCallbackBaseURL+"/auth/gplus/callback"),
		github.New(server.Config.GithubKey, server.Config.GithubSecret, server.Config.SocialCallbackBaseURL+"/auth/github/callback"),
		spotify.New(server.Config.SpotifyKey, server.Config.SpotifySecret, server.Config.SocialCallbackBaseURL+"/auth/spotify/callback"),
		linkedin.New(server.Config.LinkedinKey, server.Config.LinkedinSecret, server.Config.SocialCallbackBaseURL+"/auth/linkedin/callback"),
		line.New(server.Config.LineKey, server.Config.LineSecret, server.Config.SocialCallbackBaseURL+"/auth/line/callback", "profile", "openid", "email"),
		lastfm.New(server.Config.LastfmKey, server.Config.LastfmSecret, server.Config.SocialCallbackBaseURL+"/auth/lastfm/callback"),
		twitch.New(server.Config.TwitchKey, server.Config.TwitchSecret, server.Config.SocialCallbackBaseURL+"/auth/twitch/callback"),
		dropbox.New(server.Config.DropboxKey, server.Config.DropboxSecret, server.Config.SocialCallbackBaseURL+"/auth/dropbox/callback"),
		digitalocean.New(server.Config.DigitaloceanKey, server.Config.DigitaloceanSecret, server.Config.SocialCallbackBaseURL+"/auth/digitalocean/callback", "read"),
		bitbucket.New(server.Config.BitbucketKey, server.Config.BitbucketSecret, server.Config.SocialCallbackBaseURL+"/auth/bitbucket/callback"),
		instagram.New(server.Config.InstagramKey, server.Config.InstagramSecret, server.Config.SocialCallbackBaseURL+"/auth/instagram/callback"),
		intercom.New(server.Config.IntercomKey, server.Config.IntercomSecret, server.Config.SocialCallbackBaseURL+"/auth/intercom/callback"),
		box.New(server.Config.BoxKey, server.Config.BoxSecret, server.Config.SocialCallbackBaseURL+"/auth/box/callback"),
		salesforce.New(server.Config.SalesforceKey, server.Config.SalesforceSecret, server.Config.SocialCallbackBaseURL+"/auth/salesforce/callback"),
		seatalk.New(server.Config.SeatalkKey, server.Config.SeatalkSecret, server.Config.SocialCallbackBaseURL+"/auth/seatalk/callback"),
		amazon.New(server.Config.AmazonKey, server.Config.AmazonSecret, server.Config.SocialCallbackBaseURL+"/auth/amazon/callback"),
		yammer.New(server.Config.YammerKey, server.Config.YammerSecret, server.Config.SocialCallbackBaseURL+"/auth/yammer/callback"),
		onedrive.New(server.Config.OnedriveKey, server.Config.OnedriveSecret, server.Config.SocialCallbackBaseURL+"/auth/onedrive/callback"),
		azuread.New(server.Config.AzureadKey, server.Config.AzureadSecret, server.Config.SocialCallbackBaseURL+"/auth/azuread/callback", nil),
		microsoftonline.New(server.Config.MicrosoftonlineKey, server.Config.MicrosoftonlineSecret, server.Config.SocialCallbackBaseURL+"/auth/microsoftonline/callback"),
		battlenet.New(server.Config.BattlenetKey, server.Config.BattlenetSecret, server.Config.SocialCallbackBaseURL+"/auth/battlenet/callback"),
		eveonline.New(server.Config.EveonlineKey, server.Config.EveonlineSecret, server.Config.SocialCallbackBaseURL+"/auth/eveonline/callback"),
		kakao.New(server.Config.KakaoKey, server.Config.KakaoSecret, server.Config.SocialCallbackBaseURL+"/auth/kakao/callback"),

		//Pointed localhost.com to http://localhost:3000/auth/yahoo/callback through proxy as yahoo
		// does not allow to put custom ports in redirection uri
		yahoo.New(server.Config.YahooKey, server.Config.YahooSecret, "http://localhost.com"),
		typetalk.New(server.Config.TypetalkKey, server.Config.TypetalkSecret, server.Config.SocialCallbackBaseURL+"/auth/typetalk/callback", "my"),
		slack.New(server.Config.SlackKey, server.Config.SlackSecret, server.Config.SocialCallbackBaseURL+"/auth/slack/callback"),
		stripe.New(server.Config.StripeKey, server.Config.StripeSecret, server.Config.SocialCallbackBaseURL+"/auth/stripe/callback"),
		wepay.New(server.Config.WepayKey, server.Config.WepaySecret, server.Config.SocialCallbackBaseURL+"/auth/wepay/callback", "view_user"),

		//By default paypal production auth urls will be used, please set PAYPAL_ENV=sandbox as environment variable for testing
		//in sandbox environment
		paypal.New(server.Config.PaypalKey, server.Config.PaypalSecret, server.Config.SocialCallbackBaseURL+"/auth/paypal/callback"),
		steam.New(server.Config.SteamKey, server.Config.SocialCallbackBaseURL+"/auth/steam/callback"),
		heroku.New(server.Config.HerokuKey, server.Config.HerokuSecret, server.Config.SocialCallbackBaseURL+"/auth/heroku/callback"),
		uber.New(server.Config.UberKey, server.Config.UberSecret, server.Config.SocialCallbackBaseURL+"/auth/uber/callback"),
		soundcloud.New(server.Config.SoundcloudKey, server.Config.SoundcloudSecret, server.Config.SocialCallbackBaseURL+"/auth/soundcloud/callback"),
		gitlab.New(server.Config.GitlabKey, server.Config.GitlabSecret, server.Config.SocialCallbackBaseURL+"/auth/gitlab/callback"),
		dailymotion.New(server.Config.DailymotionKey, server.Config.DailymotionSecret, server.Config.SocialCallbackBaseURL+"/auth/dailymotion/callback", "email"),
		deezer.New(server.Config.DeezerKey, server.Config.DeezerSecret, server.Config.SocialCallbackBaseURL+"/auth/deezer/callback", "email"),
		discord.New(server.Config.DiscordKey, server.Config.DiscordSecret, server.Config.SocialCallbackBaseURL+"/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
		meetup.New(server.Config.MeetupKey, server.Config.MeetupSecret, server.Config.SocialCallbackBaseURL+"/auth/meetup/callback"),

		//Auth0 allocates domain per customer, a domain must be provided for auth0 to work
		auth0.New(server.Config.Auth0Key, server.Config.Auth0Secret, server.Config.SocialCallbackBaseURL+"/auth/auth0/callback", server.Config.Auth0Domain),
		xero.New(server.Config.XeroKey, server.Config.XeroSecret, server.Config.SocialCallbackBaseURL+"/auth/xero/callback"),
		vk.New(server.Config.VkKey, server.Config.VkSecret, server.Config.SocialCallbackBaseURL+"/auth/vk/callback"),
		naver.New(server.Config.NaverKey, server.Config.NaverSecret, server.Config.SocialCallbackBaseURL+"/auth/naver/callback"),
		yandex.New(server.Config.YandexKey, server.Config.YandexSecret, server.Config.SocialCallbackBaseURL+"/auth/yandex/callback"),
		nextcloud.NewCustomisedDNS(server.Config.NextcloudKey, server.Config.NextcloudSecret, server.Config.SocialCallbackBaseURL+"/auth/nextcloud/callback", server.Config.NextcloudURL),
		gitea.New(server.Config.GiteaKey, server.Config.GiteaSecret, server.Config.SocialCallbackBaseURL+"/auth/gitea/callback"),
		shopify.New(server.Config.ShopifyKey, server.Config.ShopifySecret, server.Config.SocialCallbackBaseURL+"/auth/shopify/callback", shopify.ScopeReadCustomers, shopify.ScopeReadOrders),
		apple.New(server.Config.AppleKey, server.Config.AppleSecret, server.Config.SocialCallbackBaseURL+"/auth/apple/callback", nil, apple.ScopeName, apple.ScopeEmail),
		strava.New(server.Config.StravaKey, server.Config.StravaSecret, server.Config.SocialCallbackBaseURL+"/auth/strava/callback"),
	)

	// OpenID Connect is based on OpenID Connect Auto Discovery URL (https://openid.net/specs/openid-connect-discovery-1_0-17.html)
	// because the OpenID Connect provider initialize it self in the New(), it can return an error which should be handled or ignored
	// ignore the error for now
	openidConnect, _ := openidConnect.New(server.Config.OpenidConnectKey, server.Config.OpenidConnectSecret, server.Config.SocialCallbackBaseURL+"/auth/openid-connect/callback", server.Config.OpenidConnectDiscoveryURL)
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

	m := make(map[string]string)
	m["amazon"] = "Amazon"
	m["bitbucket"] = "Bitbucket"
	m["box"] = "Box"
	m["dailymotion"] = "Dailymotion"
	m["deezer"] = "Deezer"
	m["digitalocean"] = "Digital Ocean"
	m["discord"] = "Discord"
	m["dropbox"] = "Dropbox"
	m["eveonline"] = "Eve Online"
	m["facebook"] = "Facebook"
	m["fitbit"] = "Fitbit"
	m["gitea"] = "Gitea"
	m["github"] = "Github"
	m["gitlab"] = "Gitlab"
	m["google"] = "Google"
	m["gplus"] = "Google Plus"
	m["shopify"] = "Shopify"
	m["soundcloud"] = "SoundCloud"
	m["spotify"] = "Spotify"
	m["steam"] = "Steam"
	m["stripe"] = "Stripe"
	m["twitch"] = "Twitch"
	m["uber"] = "Uber"
	m["wepay"] = "Wepay"
	m["yahoo"] = "Yahoo"
	m["yammer"] = "Yammer"
	m["heroku"] = "Heroku"
	m["instagram"] = "Instagram"
	m["intercom"] = "Intercom"
	m["kakao"] = "Kakao"
	m["lastfm"] = "Last FM"
	m["linkedin"] = "Linkedin"
	m["line"] = "LINE"
	m["onedrive"] = "Onedrive"
	m["azuread"] = "Azure AD"
	m["microsoftonline"] = "Microsoft Online"
	m["battlenet"] = "Battlenet"
	m["paypal"] = "Paypal"
	m["twitter"] = "Twitter"
	m["salesforce"] = "Salesforce"
	m["typetalk"] = "Typetalk"
	m["slack"] = "Slack"
	m["meetup"] = "Meetup.com"
	m["auth0"] = "Auth0"
	m["openid-connect"] = "OpenID Connect"
	m["xero"] = "Xero"
	m["vk"] = "VK"
	m["naver"] = "Naver"
	m["yandex"] = "Yandex"
	m["nextcloud"] = "NextCloud"
	m["seatalk"] = "SeaTalk"
	m["apple"] = "Apple"
	m["strava"] = "Strava"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	server.ProviderIndex = &controllers.ProviderIndex{Providers: keys, ProvidersMap: m}

	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", server.RootHandler)
	server.Router.HandleFunc("/auth/{provider}", server.SocialredirectHandler)
	server.Router.HandleFunc("/auth/{provider}/callback", server.SocialCallbackHandler)
	server.Router.HandleFunc("/logout/{provider}", server.LogoutHandler)

	server.Router.HandleFunc("/user_profile", server.UserProfileHandler)

	server.Router.Use(Middleware)

	log.Println("listening on localhost" + server.Config.ServerPort)
	log.Fatal(http.ListenAndServe(server.Config.ServerPort, server.Router))
}

// Middleware : checking for login tokens
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := models.TokenValid(r); err == nil {
			// Authenticated user
			if strings.HasPrefix(r.URL.Path, "/logout") || strings.HasPrefix(r.URL.Path, "/user_profile") {
				next.ServeHTTP(w, r)
			} else {
				w.Header().Set("Location", "/user_profile")
				w.WriteHeader(http.StatusTemporaryRedirect)
			}
		} else {
			// Unknown user
			if strings.HasPrefix(r.URL.Path, "/logout") || strings.HasPrefix(r.URL.Path, "/user_profile") {
				http.Error(w, "Forbidden", http.StatusForbidden)
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
