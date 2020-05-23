package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"v2/models"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

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

// ProviderIndex : mind your own business
type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

// DB : db connection
var DB *gorm.DB
var providerIndex *ProviderIndex
var err error
var store *sessions.CookieStore

func main() {

	os.Setenv("API_SECRET", "SDJKAFHHAJSGFAJSHFAFSAFHHSAJSKFASJJF") // TODO REMOVE IT

	var config models.Config
	if _, err = toml.DecodeFile("env.toml", &config); err != nil {
		panic("Failed to read enviroment settings")
	}

	databaseConnectionSettings := "host=" + config.DB.Host + " port=" + config.DB.Port + " user=" + config.DB.User + " dbname=" + config.DB.DbName + " password=" + config.DB.Password + " sslmode=disable"
	DB, err = gorm.Open("postgres", databaseConnectionSettings)
	if err != nil {
		log.Println("DEBUG: ", databaseConnectionSettings)
		log.Println("ERR: ", err)
		panic("failed to connect database")
	}
	defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&models.User{})

	gothic.Store = sessions.NewCookieStore([]byte(config.CookieStore))
	store = sessions.NewCookieStore([]byte(config.CookieStore))

	goth.UseProviders(
		twitter.New(config.TwitterKey, config.TwitterSecret, config.SocialCallbackBaseURL+"/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(config.TWITTER_KEY, config.TWITTER_SECRET, config.SocialCallbackBaseURL+"/auth/twitter/callback"),

		facebook.New(config.FacebookKey, config.FacebookSecret, config.SocialCallbackBaseURL+"/auth/facebook/callback"),
		fitbit.New(config.FitbitKey, config.FitbitSecret, config.SocialCallbackBaseURL+"/auth/fitbit/callback"),
		google.New(config.GoogleKey, config.GoogleSecret, config.SocialCallbackBaseURL+"/auth/google/callback"),
		gplus.New(config.GplusKey, config.GplusSecret, config.SocialCallbackBaseURL+"/auth/gplus/callback"),
		github.New(config.GithubKey, config.GithubSecret, config.SocialCallbackBaseURL+"/auth/github/callback"),
		spotify.New(config.SpotifyKey, config.SpotifySecret, config.SocialCallbackBaseURL+"/auth/spotify/callback"),
		linkedin.New(config.LinkedinKey, config.LinkedinSecret, config.SocialCallbackBaseURL+"/auth/linkedin/callback"),
		line.New(config.LineKey, config.LineSecret, config.SocialCallbackBaseURL+"/auth/line/callback", "profile", "openid", "email"),
		lastfm.New(config.LastfmKey, config.LastfmSecret, config.SocialCallbackBaseURL+"/auth/lastfm/callback"),
		twitch.New(config.TwitchKey, config.TwitchSecret, config.SocialCallbackBaseURL+"/auth/twitch/callback"),
		dropbox.New(config.DropboxKey, config.DropboxSecret, config.SocialCallbackBaseURL+"/auth/dropbox/callback"),
		digitalocean.New(config.DigitaloceanKey, config.DigitaloceanSecret, config.SocialCallbackBaseURL+"/auth/digitalocean/callback", "read"),
		bitbucket.New(config.BitbucketKey, config.BitbucketSecret, config.SocialCallbackBaseURL+"/auth/bitbucket/callback"),
		instagram.New(config.InstagramKey, config.InstagramSecret, config.SocialCallbackBaseURL+"/auth/instagram/callback"),
		intercom.New(config.IntercomKey, config.IntercomSecret, config.SocialCallbackBaseURL+"/auth/intercom/callback"),
		box.New(config.BoxKey, config.BoxSecret, config.SocialCallbackBaseURL+"/auth/box/callback"),
		salesforce.New(config.SalesforceKey, config.SalesforceSecret, config.SocialCallbackBaseURL+"/auth/salesforce/callback"),
		seatalk.New(config.SeatalkKey, config.SeatalkSecret, config.SocialCallbackBaseURL+"/auth/seatalk/callback"),
		amazon.New(config.AmazonKey, config.AmazonSecret, config.SocialCallbackBaseURL+"/auth/amazon/callback"),
		yammer.New(config.YammerKey, config.YammerSecret, config.SocialCallbackBaseURL+"/auth/yammer/callback"),
		onedrive.New(config.OnedriveKey, config.OnedriveSecret, config.SocialCallbackBaseURL+"/auth/onedrive/callback"),
		azuread.New(config.AzureadKey, config.AzureadSecret, config.SocialCallbackBaseURL+"/auth/azuread/callback", nil),
		microsoftonline.New(config.MicrosoftonlineKey, config.MicrosoftonlineSecret, config.SocialCallbackBaseURL+"/auth/microsoftonline/callback"),
		battlenet.New(config.BattlenetKey, config.BattlenetSecret, config.SocialCallbackBaseURL+"/auth/battlenet/callback"),
		eveonline.New(config.EveonlineKey, config.EveonlineSecret, config.SocialCallbackBaseURL+"/auth/eveonline/callback"),
		kakao.New(config.KakaoKey, config.KakaoSecret, config.SocialCallbackBaseURL+"/auth/kakao/callback"),

		//Pointed localhost.com to http://localhost:3000/auth/yahoo/callback through proxy as yahoo
		// does not allow to put custom ports in redirection uri
		yahoo.New(config.YahooKey, config.YahooSecret, "http://localhost.com"),
		typetalk.New(config.TypetalkKey, config.TypetalkSecret, config.SocialCallbackBaseURL+"/auth/typetalk/callback", "my"),
		slack.New(config.SlackKey, config.SlackSecret, config.SocialCallbackBaseURL+"/auth/slack/callback"),
		stripe.New(config.StripeKey, config.StripeSecret, config.SocialCallbackBaseURL+"/auth/stripe/callback"),
		wepay.New(config.WepayKey, config.WepaySecret, config.SocialCallbackBaseURL+"/auth/wepay/callback", "view_user"),

		//By default paypal production auth urls will be used, please set PAYPAL_ENV=sandbox as environment variable for testing
		//in sandbox environment
		paypal.New(config.PaypalKey, config.PaypalSecret, config.SocialCallbackBaseURL+"/auth/paypal/callback"),
		steam.New(config.SteamKey, config.SocialCallbackBaseURL+"/auth/steam/callback"),
		heroku.New(config.HerokuKey, config.HerokuSecret, config.SocialCallbackBaseURL+"/auth/heroku/callback"),
		uber.New(config.UberKey, config.UberSecret, config.SocialCallbackBaseURL+"/auth/uber/callback"),
		soundcloud.New(config.SoundcloudKey, config.SoundcloudSecret, config.SocialCallbackBaseURL+"/auth/soundcloud/callback"),
		gitlab.New(config.GitlabKey, config.GitlabSecret, config.SocialCallbackBaseURL+"/auth/gitlab/callback"),
		dailymotion.New(config.DailymotionKey, config.DailymotionSecret, config.SocialCallbackBaseURL+"/auth/dailymotion/callback", "email"),
		deezer.New(config.DeezerKey, config.DeezerSecret, config.SocialCallbackBaseURL+"/auth/deezer/callback", "email"),
		discord.New(config.DiscordKey, config.DiscordSecret, config.SocialCallbackBaseURL+"/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
		meetup.New(config.MeetupKey, config.MeetupSecret, config.SocialCallbackBaseURL+"/auth/meetup/callback"),

		//Auth0 allocates domain per customer, a domain must be provided for auth0 to work
		auth0.New(config.Auth0Key, config.Auth0Secret, config.SocialCallbackBaseURL+"/auth/auth0/callback", config.Auth0Domain),
		xero.New(config.XeroKey, config.XeroSecret, config.SocialCallbackBaseURL+"/auth/xero/callback"),
		vk.New(config.VkKey, config.VkSecret, config.SocialCallbackBaseURL+"/auth/vk/callback"),
		naver.New(config.NaverKey, config.NaverSecret, config.SocialCallbackBaseURL+"/auth/naver/callback"),
		yandex.New(config.YandexKey, config.YandexSecret, config.SocialCallbackBaseURL+"/auth/yandex/callback"),
		nextcloud.NewCustomisedDNS(config.NextcloudKey, config.NextcloudSecret, config.SocialCallbackBaseURL+"/auth/nextcloud/callback", config.NextcloudURL),
		gitea.New(config.GiteaKey, config.GiteaSecret, config.SocialCallbackBaseURL+"/auth/gitea/callback"),
		shopify.New(config.ShopifyKey, config.ShopifySecret, config.SocialCallbackBaseURL+"/auth/shopify/callback", shopify.ScopeReadCustomers, shopify.ScopeReadOrders),
		apple.New(config.AppleKey, config.AppleSecret, config.SocialCallbackBaseURL+"/auth/apple/callback", nil, apple.ScopeName, apple.ScopeEmail),
		strava.New(config.StravaKey, config.StravaSecret, config.SocialCallbackBaseURL+"/auth/strava/callback"),
	)

	// OpenID Connect is based on OpenID Connect Auto Discovery URL (https://openid.net/specs/openid-connect-discovery-1_0-17.html)
	// because the OpenID Connect provider initialize it self in the New(), it can return an error which should be handled or ignored
	// ignore the error for now
	openidConnect, _ := openidConnect.New(config.OpenidConnectKey, config.OpenidConnectSecret, config.SocialCallbackBaseURL+"/auth/openid-connect/callback", config.OpenidConnectDiscoveryURL)
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

	providerIndex = &ProviderIndex{Providers: keys, ProvidersMap: m}

	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/auth/{provider}", socialredirect)
	r.HandleFunc("/auth/{provider}/callback", callback)
	r.HandleFunc("/logout/{provider}", logout)

	r.HandleFunc("/user_profile", userProfile)

	r.Use(Middleware)

	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
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

// root : showing login page
func root(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("fe/layout.html", "fe/login.html")
	tmpl.ExecuteTemplate(res, "layout", providerIndex)
}

// socialredirect : redirect to social login page of the provider chosen
func socialredirect(res http.ResponseWriter, req *http.Request) {
	if _, err := gothic.CompleteUserAuth(res, req); err == nil {
		res.Header().Set("Location", "/user_profile")
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		log.Println("ERR: ", err)
		gothic.BeginAuthHandler(res, req)
	}
}

// callback : function executed after return from social network
func callback(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		log.Println(res, err)
		return
	}

	var currentUser models.User
	currentUser.FromGothUser(user)
	DB.FirstOrCreate(&currentUser)

	err = models.CreateToken(res, req, currentUser.ID)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Location", "/user_profile")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

// logout : logut function
func logout(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	models.DeleteToken(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

// userProfile : showing page with user infos
func userProfile(res http.ResponseWriter, req *http.Request) {

	userid, err := models.ExtractTokenID(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{}
	foundUser, err := user.FindUserByID(DB, userid)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, _ := template.ParseFiles("fe/layout.html", "fe/user_info.html")
	tmpl.ExecuteTemplate(res, "layout", foundUser)
}
