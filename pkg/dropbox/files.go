package dropbox

//
//import (
//	"context"
//	"fmt"
//	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
//	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
//	"golang.org/x/oauth2"
//	"log"
//	"net/http"
//	"os"
//)
//
//const (
//	apiVersion    = 2
//	defaultDomain = ".dropboxapi.com"
//	hostAPI       = "api"
//	hostContent   = "content"
//	hostNotify    = "notify"
//)
//
//type DropboxFiles Context
//
//type Config struct {
//	// OAuth2 access token
//	Token string
//	// No need to set -- for testing only
//	Domain string
//	// No need to set -- for testing only
//	Client *http.Client
//	// No need to set -- for testing only
//	HeaderGenerator func(hostType string, style string, namespace string, route string) map[string]string
//	// No need to set -- for testing only
//	URLGenerator func(hostType string, style string, namespace string, route string) string
//}
//
//func NewClient() *DropboxFiles {
//	config := Config{
//		Token: os.Getenv("DROPBOX_APP_TOKEN"),
//	}
//
//	context := NewContext(config)
//
//	f := DropboxFiles(context)
//	return &f
//}
//
//// Context is the base client context used to implement per-namespace clients.
//type Context struct {
//	Config          Config
//	Client          *http.Client
//	HeaderGenerator func(hostType string, style string, namespace string, route string) map[string]string
//	URLGenerator    func(hostType string, style string, namespace string, route string) string
//}
//
//// OAuthEndpoint constructs an `oauth2.Endpoint` for the given domain
//func OAuthEndpoint(domain string) oauth2.Endpoint {
//	if domain == "" {
//		domain = defaultDomain
//	}
//	authURL := fmt.Sprintf("https://meta%s/1/oauth2/authorize", domain)
//	tokenURL := fmt.Sprintf("https://api%s/1/oauth2/token", domain)
//	if domain == defaultDomain {
//		authURL = "https://www.dropbox.com/1/oauth2/authorize"
//	}
//	return oauth2.Endpoint{AuthURL: authURL, TokenURL: tokenURL}
//}
//
//// NewContext returns a new Context with the given Config.
//func NewContext(c Config) Context {
//	domain := c.Domain
//	if domain == "" {
//		domain = defaultDomain
//	}
//
//	client := c.Client
//	if client == nil {
//		var conf = &oauth2.Config{Endpoint: OAuthEndpoint(domain)}
//		tok := &oauth2.Token{AccessToken: c.Token}
//		client = conf.Client(context.Background(), tok)
//	}
//
//	headerGenerator := c.HeaderGenerator
//	if headerGenerator == nil {
//		headerGenerator = func(hostType string, style string, namespace string, route string) map[string]string {
//			return map[string]string{}
//		}
//	}
//
//	urlGenerator := c.URLGenerator
//	if urlGenerator == nil {
//		hostMap := map[string]string{
//			hostAPI:     hostAPI + domain,
//			hostContent: hostContent + domain,
//			hostNotify:  hostNotify + domain,
//		}
//		urlGenerator = func(hostType string, style string, namespace string, route string) string {
//			fqHost := hostMap[hostType]
//			return fmt.Sprintf("https://%s/%d/%s/%s", fqHost, apiVersion, namespace, route)
//		}
//	}
//
//	return Context{c, client, headerGenerator, urlGenerator}
//}
//
//func (f *DropboxFiles) GetAllInPath(path string) (files.IsMetadata, error) {
//	f.
//	response, err := f.Client.ListFolder(&files.ListFolderArg{
//		Path: path,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	for i := range response.Entries {
//		fmt.Println("entry: ", response.Entries[i])
//	}
//
//	return response.Entries, nil
//}
