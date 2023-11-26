package event

import (
	"net"
	"net/url"
)

// This is an implementation of Context class similar to the `CoreExtraContext` in segment js SDK.
type Context struct {
	Active        bool
	IP            net.IP    // Current user's IP address.
	Locale        string    // Locale string for the current user, for example en-US.
	Location      *Location // Dictionary of information about the user’s current location.
	Page          *Page     // Dictionary of information about the current web page.
	UserAgent     string
	UserAgentData *UserAgentData // User agent data returned by the Client Hints API
	Library       *Library       // the name of the library(sdk) and the version of it
	Traits        *Traits
	Campaign      *Campaign // Dictionary of information about the campaign that resulted in the API call, containing name, source, medium, term, content, and any other custom UTM parameter.
	Referrer      *Referrer // Dictionary of information about the way the user was referred to the website or app.
	CustomData    *CustomData
}

type Referrer struct {
	ReferrerType string
	Name         string
	Url          url.URL
	Link         string

	// id           string these three properties were undocumented in segment so is commented them
	// btid         string
	// urid         string
}

type Campaign struct {
	Name       string
	Source     string
	Medium     string
	CustomData CustomData
}

type Library struct {
	Name    string // analytics-node-next/latest
	Version string // 1.43.1
}

type AgentBrandVersion struct {
	Brand   string // i think this is for storing the browsers brand, for example, firefox or chromium
	Version string // also this is the property for storing the version of the browser
}

type UserAgentData struct {
	Mobile          bool
	Platform        string
	Architecture    string
	Bitness         string
	Model           string
	PlatformVersion string
	WOW64           bool
	Brands          []AgentBrandVersion

	//{
	//	"brand": "Chromium",
	//	"version": "119"
	// },
	//{
	//	"brand": "Not?A_Brand",
	//	"version": "24"
	//}
	// This is an example which I've collected from js SDK of segment my browser's name is ARC which is not identified by segment so the brand field is Not?A_brand by also the engine of my browser which is chromioum based is stored

	fullVersionList []AgentBrandVersion // TODO: i don't know why this exists??
	// uaFullVersion   string // also this field is logged as deprecated in segment so i left it as a comment
}

type Page struct {
	Path     string  // academy/
	Referrer url.URL // https://www.foo.com/
	Search   string  // projectId=123
	Title    string  // Analytics Academy
	Url      url.URL // https://segment.com/academy/
}

type Location struct {
	City      string
	Country   string
	Latitude  string
	Longitude string
	Region    string
	// TODO: i'm not sure what speed is in this context?
	Speed int
}
