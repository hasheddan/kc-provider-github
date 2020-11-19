module github.com/hasheddan/kc-provider-github

go 1.13

require (
	github.com/crossplane/crossplane-runtime v0.10.0
	github.com/crossplane/crossplane-tools v0.0.0-20201007233256-88b291e145bb
	github.com/google/go-github/v32 v32.1.0
	github.com/pkg/errors v0.9.1
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/controller-tools v0.2.4
)
