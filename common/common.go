package common

// Default service addresses and URLS of Argo CD internal services
const (
	// DefaultRepoServerAddr is the gRPC address of the Argo CD repo server
	DefaultRepoServerAddr = "argocd-repo-server:8081"
	// DefaultDexServerAddr is the HTTP address of the Dex OIDC server, which we run a reverse proxy against
	DefaultDexServerAddr = "http://argocd-dex-server:5556"
	// DefaultRedisAddr is the default redis address
	DefaultRedisAddr = "argocd-redis:6379"
)

// Kubernetes ConfigMap and Secret resource names which hold Argo CD settings
const (
	ArgoCDConfigMapName     = "argocd-cm"
	ArgoCDSecretName        = "argocd-secret"
	ArgoCDRBACConfigMapName = "argocd-rbac-cm"
	// Contains SSH known hosts data for connecting repositories. Will get mounted as volume to pods
	ArgoCDKnownHostsConfigMapName = "argocd-ssh-known-hosts-cm"
	// Contains TLS certificate data for connecting repositories. Will get mounted as volume to pods
	ArgoCDTLSCertsConfigMapName = "argocd-tls-certs-cm"
)

// Default system namespace
const (
	DefaultSystemNamespace = "kube-system"
)

// Default listener ports for ArgoCD components
const (
	DefaultPortAPIServer              = 8080
	DefaultPortRepoServer             = 8081
	DefaultPortArgoCDMetrics          = 8082
	DefaultPortArgoCDAPIServerMetrics = 8083
	DefaultPortRepoServerMetrics      = 8084
)

// Argo CD application related constants
const (
	// ArgoCDAdminUsername is the username of the 'admin' user
	ArgoCDAdminUsername = "admin"
	// ArgoCDUserAgentName is the default user-agent name used by the gRPC API client library and grpc-gateway
	ArgoCDUserAgentName = "argocd-client"
	// AuthCookieName is the HTTP cookie name where we store our auth token
	AuthCookieName = "argocd.token"
)

// Dex related constants
const (
	// DexAPIEndpoint is the endpoint where we serve the Dex API server
	DexAPIEndpoint = "/api/dex"
	// LoginEndpoint is Argo CD's shorthand login endpoint which redirects to dex's OAuth 2.0 provider's consent page
	LoginEndpoint = "/auth/login"
	// CallbackEndpoint is Argo CD's final callback endpoint we reach after OAuth 2.0 login flow has been completed
	CallbackEndpoint = "/auth/callback"
	// DexCallbackEndpoint is Argo CD's final callback endpoint when Dex is configured
	DexCallbackEndpoint = "/api/dex/callback"
	// ArgoCDClientAppName is name of the Oauth client app used when registering our web app to dex
	ArgoCDClientAppName = "Argo CD"
	// ArgoCDClientAppID is the Oauth client ID we will use when registering our app to dex
	ArgoCDClientAppID = "argo-cd"
	// ArgoCDCLIClientAppName is name of the Oauth client app used when registering our CLI to dex
	ArgoCDCLIClientAppName = "Argo CD CLI"
	// ArgoCDCLIClientAppID is the Oauth client ID we will use when registering our CLI to dex
	ArgoCDCLIClientAppID = "argo-cd-cli"
)

// Resource metadata labels and annotations (keys and values) used by Argo CD components
const (
	// LabelKeySecretType contains the type of argocd secret (currently: 'cluster')
	LabelKeySecretType = "argocd.argoproj.io/secret-type"
	// LabelValueSecretTypeCluster indicates a secret type of cluster
	LabelValueSecretTypeCluster = "cluster"

	// AnnotationKeyManagedBy is annotation name which indicates that k8s resource is managed by an application.
	AnnotationKeyManagedBy = "managed-by"
	// AnnotationValueManagedByArgoCD is a 'managed-by' annotation value for resources managed by Argo CD
	AnnotationValueManagedByArgoCD = "argocd.argoproj.io"
)

// Environment variables for tuning and debugging Argo CD
const (
	// EnvVarSSODebug is an environment variable to enable additional OAuth debugging in the API server
	EnvVarSSODebug = "ARGOCD_SSO_DEBUG"
	// EnvVarRBACDebug is an environment variable to enable additional RBAC debugging in the API server
	EnvVarRBACDebug = "ARGOCD_RBAC_DEBUG"
)

const (
	// MinClientVersion is the minimum client version that can interface with this API server.
	// When introducing breaking changes to the API or datastructures, this number should be bumped.
	// The value here may be lower than the current value in VERSION
	MinClientVersion = "1.3.0"
	// CacheVersion is a objects version cached using util/cache/cache.go.
	// Number should be bumped in case of backward incompatible change to make sure cache is invalidated after upgrade.
	CacheVersion = "1.0.0"
)
