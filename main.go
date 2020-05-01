package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"sort"

	"log"

	"github.com/BurntSushi/toml"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
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

type config struct {
	SERVER_BASE_URL              string
	COOKIEPSW                    string
	TWITTER_KEY                  string
	TWITTER_SECRET               string
	FACEBOOK_KEY                 string
	FACEBOOK_SECRET              string
	FITBIT_KEY                   string
	FITBIT_SECRET                string
	GOOGLE_KEY                   string
	GOOGLE_SECRET                string
	GPLUS_KEY                    string
	GPLUS_SECRET                 string
	GITHUB_KEY                   string
	GITHUB_SECRET                string
	SPOTIFY_KEY                  string
	SPOTIFY_SECRET               string
	LINKEDIN_KEY                 string
	LINKEDIN_SECRET              string
	LINE_KEY                     string
	LINE_SECRET                  string
	LASTFM_KEY                   string
	LASTFM_SECRET                string
	TWITCH_KEY                   string
	TWITCH_SECRET                string
	DROPBOX_KEY                  string
	DROPBOX_SECRET               string
	DIGITALOCEAN_KEY             string
	DIGITALOCEAN_SECRET          string
	BITBUCKET_KEY                string
	BITBUCKET_SECRET             string
	INSTAGRAM_KEY                string
	INSTAGRAM_SECRET             string
	INTERCOM_KEY                 string
	INTERCOM_SECRET              string
	BOX_KEY                      string
	BOX_SECRET                   string
	SALESFORCE_KEY               string
	SALESFORCE_SECRET            string
	SEATALK_KEY                  string
	SEATALK_SECRET               string
	AMAZON_KEY                   string
	AMAZON_SECRET                string
	YAMMER_KEY                   string
	YAMMER_SECRET                string
	ONEDRIVE_KEY                 string
	ONEDRIVE_SECRET              string
	AZUREAD_KEY                  string
	AZUREAD_SECRET               string
	MICROSOFTONLINE_KEY          string
	MICROSOFTONLINE_SECRET       string
	BATTLENET_KEY                string
	BATTLENET_SECRET             string
	EVEONLINE_KEY                string
	EVEONLINE_SECRET             string
	KAKAO_KEY                    string
	KAKAO_SECRET                 string
	YAHOO_KEY                    string
	YAHOO_SECRET                 string
	TYPETALK_KEY                 string
	TYPETALK_SECRET              string
	SLACK_KEY                    string
	SLACK_SECRET                 string
	STRIPE_KEY                   string
	STRIPE_SECRET                string
	WEPAY_KEY                    string
	WEPAY_SECRET                 string
	PAYPAL_KEY                   string
	PAYPAL_SECRET                string
	STEAM_KEY                    string
	HEROKU_KEY                   string
	HEROKU_SECRET                string
	UBER_KEY                     string
	UBER_SECRET                  string
	SOUNDCLOUD_KEY               string
	SOUNDCLOUD_SECRET            string
	GITLAB_KEY                   string
	GITLAB_SECRET                string
	DAILYMOTION_KEY              string
	DAILYMOTION_SECRET           string
	DEEZER_KEY                   string
	DEEZER_SECRET                string
	DISCORD_KEY                  string
	DISCORD_SECRET               string
	MEETUP_KEY                   string
	MEETUP_SECRET                string
	AUTH0_KEY                    string
	AUTH0_SECRET                 string
	AUTH0_DOMAIN                 string
	XERO_KEY                     string
	XERO_SECRET                  string
	VK_KEY                       string
	VK_SECRET                    string
	NAVER_KEY                    string
	NAVER_SECRET                 string
	YANDEX_KEY                   string
	YANDEX_SECRET                string
	NEXTCLOUD_KEY                string
	NEXTCLOUD_SECRET             string
	NEXTCLOUD_URL                string
	GITEA_KEY                    string
	GITEA_SECRET                 string
	SHOPIFY_KEY                  string
	SHOPIFY_SECRET               string
	APPLE_KEY                    string
	APPLE_SECRET                 string
	STRAVA_KEY                   string
	STRAVA_SECRET                string
	OPENID_CONNECT_KEY           string
	OPENID_CONNECT_SECRET        string
	OPENID_CONNECT_DISCOVERY_URL string
}

func main() {

	var config config
	if _, err := toml.DecodeFile("env.toml", &config); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	gothic.Store = sessions.NewCookieStore([]byte(config.COOKIEPSW))

	goth.UseProviders(
		twitter.New(config.TWITTER_KEY, config.TWITTER_SECRET, config.SERVER_BASE_URL+"/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(config.TWITTER_KEY, config.TWITTER_SECRET, config.SERVER_BASE_URL+"/auth/twitter/callback"),

		facebook.New(config.FACEBOOK_KEY, config.FACEBOOK_SECRET, config.SERVER_BASE_URL+"/auth/facebook/callback"),
		fitbit.New(config.FITBIT_KEY, config.FITBIT_SECRET, config.SERVER_BASE_URL+"/auth/fitbit/callback"),
		google.New(config.GOOGLE_KEY, config.GOOGLE_SECRET, config.SERVER_BASE_URL+"/auth/google/callback"),
		gplus.New(config.GPLUS_KEY, config.GPLUS_SECRET, config.SERVER_BASE_URL+"/auth/gplus/callback"),
		github.New(config.GITHUB_KEY, config.GITHUB_SECRET, config.SERVER_BASE_URL+"/auth/github/callback"),
		spotify.New(config.SPOTIFY_KEY, config.SPOTIFY_SECRET, config.SERVER_BASE_URL+"/auth/spotify/callback"),
		linkedin.New(config.LINKEDIN_KEY, config.LINKEDIN_SECRET, config.SERVER_BASE_URL+"/auth/linkedin/callback"),
		line.New(config.LINE_KEY, config.LINE_SECRET, config.SERVER_BASE_URL+"/auth/line/callback", "profile", "openid", "email"),
		lastfm.New(config.LASTFM_KEY, config.LASTFM_SECRET, config.SERVER_BASE_URL+"/auth/lastfm/callback"),
		twitch.New(config.TWITCH_KEY, config.TWITCH_SECRET, config.SERVER_BASE_URL+"/auth/twitch/callback"),
		dropbox.New(config.DROPBOX_KEY, config.DROPBOX_SECRET, config.SERVER_BASE_URL+"/auth/dropbox/callback"),
		digitalocean.New(config.DIGITALOCEAN_KEY, config.DIGITALOCEAN_SECRET, config.SERVER_BASE_URL+"/auth/digitalocean/callback", "read"),
		bitbucket.New(config.BITBUCKET_KEY, config.BITBUCKET_SECRET, config.SERVER_BASE_URL+"/auth/bitbucket/callback"),
		instagram.New(config.INSTAGRAM_KEY, config.INSTAGRAM_SECRET, config.SERVER_BASE_URL+"/auth/instagram/callback"),
		intercom.New(config.INTERCOM_KEY, config.INTERCOM_SECRET, config.SERVER_BASE_URL+"/auth/intercom/callback"),
		box.New(config.BOX_KEY, config.BOX_SECRET, config.SERVER_BASE_URL+"/auth/box/callback"),
		salesforce.New(config.SALESFORCE_KEY, config.SALESFORCE_SECRET, config.SERVER_BASE_URL+"/auth/salesforce/callback"),
		seatalk.New(config.SEATALK_KEY, config.SEATALK_SECRET, config.SERVER_BASE_URL+"/auth/seatalk/callback"),
		amazon.New(config.AMAZON_KEY, config.AMAZON_SECRET, config.SERVER_BASE_URL+"/auth/amazon/callback"),
		yammer.New(config.YAMMER_KEY, config.YAMMER_SECRET, config.SERVER_BASE_URL+"/auth/yammer/callback"),
		onedrive.New(config.ONEDRIVE_KEY, config.ONEDRIVE_SECRET, config.SERVER_BASE_URL+"/auth/onedrive/callback"),
		azuread.New(config.AZUREAD_KEY, config.AZUREAD_SECRET, config.SERVER_BASE_URL+"/auth/azuread/callback", nil),
		microsoftonline.New(config.MICROSOFTONLINE_KEY, config.MICROSOFTONLINE_SECRET, config.SERVER_BASE_URL+"/auth/microsoftonline/callback"),
		battlenet.New(config.BATTLENET_KEY, config.BATTLENET_SECRET, config.SERVER_BASE_URL+"/auth/battlenet/callback"),
		eveonline.New(config.EVEONLINE_KEY, config.EVEONLINE_SECRET, config.SERVER_BASE_URL+"/auth/eveonline/callback"),
		kakao.New(config.KAKAO_KEY, config.KAKAO_SECRET, config.SERVER_BASE_URL+"/auth/kakao/callback"),

		//Pointed localhost.com to http://localhost:3000/auth/yahoo/callback through proxy as yahoo
		// does not allow to put custom ports in redirection uri
		yahoo.New(config.YAHOO_KEY, config.YAHOO_SECRET, "http://localhost.com"),
		typetalk.New(config.TYPETALK_KEY, config.TYPETALK_SECRET, config.SERVER_BASE_URL+"/auth/typetalk/callback", "my"),
		slack.New(config.SLACK_KEY, config.SLACK_SECRET, config.SERVER_BASE_URL+"/auth/slack/callback"),
		stripe.New(config.STRIPE_KEY, config.STRIPE_SECRET, config.SERVER_BASE_URL+"/auth/stripe/callback"),
		wepay.New(config.WEPAY_KEY, config.WEPAY_SECRET, config.SERVER_BASE_URL+"/auth/wepay/callback", "view_user"),

		//By default paypal production auth urls will be used, please set PAYPAL_ENV=sandbox as environment variable for testing
		//in sandbox environment
		paypal.New(config.PAYPAL_KEY, config.PAYPAL_SECRET, config.SERVER_BASE_URL+"/auth/paypal/callback"),
		steam.New(config.STEAM_KEY, config.SERVER_BASE_URL+"/auth/steam/callback"),
		heroku.New(config.HEROKU_KEY, config.HEROKU_SECRET, config.SERVER_BASE_URL+"/auth/heroku/callback"),
		uber.New(config.UBER_KEY, config.UBER_SECRET, config.SERVER_BASE_URL+"/auth/uber/callback"),
		soundcloud.New(config.SOUNDCLOUD_KEY, config.SOUNDCLOUD_SECRET, config.SERVER_BASE_URL+"/auth/soundcloud/callback"),
		gitlab.New(config.GITLAB_KEY, config.GITLAB_SECRET, config.SERVER_BASE_URL+"/auth/gitlab/callback"),
		dailymotion.New(config.DAILYMOTION_KEY, config.DAILYMOTION_SECRET, config.SERVER_BASE_URL+"/auth/dailymotion/callback", "email"),
		deezer.New(config.DEEZER_KEY, config.DEEZER_SECRET, config.SERVER_BASE_URL+"/auth/deezer/callback", "email"),
		discord.New(config.DISCORD_KEY, config.DISCORD_SECRET, config.SERVER_BASE_URL+"/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
		meetup.New(config.MEETUP_KEY, config.MEETUP_SECRET, config.SERVER_BASE_URL+"/auth/meetup/callback"),

		//Auth0 allocates domain per customer, a domain must be provided for auth0 to work
		auth0.New(config.AUTH0_KEY, config.AUTH0_SECRET, config.SERVER_BASE_URL+"/auth/auth0/callback", config.AUTH0_DOMAIN),
		xero.New(config.XERO_KEY, config.XERO_SECRET, config.SERVER_BASE_URL+"/auth/xero/callback"),
		vk.New(config.VK_KEY, config.VK_SECRET, config.SERVER_BASE_URL+"/auth/vk/callback"),
		naver.New(config.NAVER_KEY, config.NAVER_SECRET, config.SERVER_BASE_URL+"/auth/naver/callback"),
		yandex.New(config.YANDEX_KEY, config.YANDEX_SECRET, config.SERVER_BASE_URL+"/auth/yandex/callback"),
		nextcloud.NewCustomisedDNS(config.NEXTCLOUD_KEY, config.NEXTCLOUD_SECRET, config.SERVER_BASE_URL+"/auth/nextcloud/callback", config.NEXTCLOUD_URL),
		gitea.New(config.GITEA_KEY, config.GITEA_SECRET, config.SERVER_BASE_URL+"/auth/gitea/callback"),
		shopify.New(config.SHOPIFY_KEY, config.SHOPIFY_SECRET, config.SERVER_BASE_URL+"/auth/shopify/callback", shopify.ScopeReadCustomers, shopify.ScopeReadOrders),
		apple.New(config.APPLE_KEY, config.APPLE_SECRET, config.SERVER_BASE_URL+"/auth/apple/callback", nil, apple.ScopeName, apple.ScopeEmail),
		strava.New(config.STRAVA_KEY, config.STRAVA_SECRET, config.SERVER_BASE_URL+"/auth/strava/callback"),
	)

	// OpenID Connect is based on OpenID Connect Auto Discovery URL (https://openid.net/specs/openid-connect-discovery-1_0-17.html)
	// because the OpenID Connect provider initialize it self in the New(), it can return an error which should be handled or ignored
	// ignore the error for now
	openidConnect, _ := openidConnect.New(config.OPENID_CONNECT_KEY, config.OPENID_CONNECT_SECRET, config.SERVER_BASE_URL+"/auth/openid-connect/callback", config.OPENID_CONNECT_DISCOVERY_URL)
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

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, user)
	})

	p.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.New("foo").Parse(userTemplate)
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("login_page.html")
		t.Execute(res, providerIndex)
	})
	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", p))
}

// ProviderIndex : mind your own business
type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

var userTemplate = `
		<p><a href="/logout/{{.Provider}}">logout</a></p>
		<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
		<p>Email: {{.Email}}</p>
		<p>NickName: {{.NickName}}</p>
		<p>Location: {{.Location}}</p>
		<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
		<p>Description: {{.Description}}</p>
		<p>UserID: {{.UserID}}</p>
		<p>AccessToken: {{.AccessToken}}</p>
		<p>ExpiresAt: {{.ExpiresAt}}</p>
		<p>RefreshToken: {{.RefreshToken}}</p>
		`
