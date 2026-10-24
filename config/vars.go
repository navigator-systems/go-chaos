package config

import "os"

// Vars groups the credentials and auth-related settings used by cloud providers
// and notifications. Never log this value because it contains secrets.
type Vars struct {
	AWS          AWSCredentials
	Kubernetes   KubernetesCredentials
	GCP          GCPCredentials
	DigitalOcean DigitalOceanCredentials
	Slack        SlackCredentials
	Gmail        GmailCredentials
}

type AWSCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Region          string
	Profile         string
}

type KubernetesCredentials struct {
	KubeconfigPath string
}

type GCPCredentials struct {
	CredentialsFile string
}

type DigitalOceanCredentials struct {
	Token string
}

type SlackCredentials struct {
	AuthToken string
}

type GmailCredentials struct {
	AppToken string
}

// ReadVars reads auth-related environment variables once. Call it near the
// application entry point, then pass the returned Vars value (or *Vars) to
// provider and notification functions instead of having each function read
// environment variables itself. Those functions and their callers must accept
// Vars for this to take effect; their signatures are intentionally unchanged
// here.
//
// Empty values are allowed: AWS, GCP, and Kubernetes can use their existing
// SDK/client default credential discovery when explicit settings are absent.
func ReadVars() Vars {
	return Vars{
		AWS: AWSCredentials{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			SessionToken:    os.Getenv("AWS_SESSION_TOKEN"),
			Region:          os.Getenv("AWS_REGION"),
			Profile:         os.Getenv("AWS_PROFILE"),
		},
		Kubernetes: KubernetesCredentials{
			KubeconfigPath: os.Getenv("KUBECONFIG"),
		},
		GCP: GCPCredentials{
			CredentialsFile: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		},
		DigitalOcean: DigitalOceanCredentials{
			Token: os.Getenv("DIGITALOCEAN_TOKEN"),
		},
		Slack: SlackCredentials{
			AuthToken: os.Getenv("SLACK_AUTH_TOKEN"),
		},
		Gmail: GmailCredentials{
			AppToken: os.Getenv("GMAIL_APP_TOKEN"),
		},
	}
}
