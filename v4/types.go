package authelia

// UserProvider is the interface for checking user password and gathering user details.
type UserProvider interface {
	CheckUserPassword(username string, password string) (success bool, err error)
	GetDetails(username string) (details *UserDetails, err error)
	UpdatePassword(username string, newPassword string) (err error)
}

// NotificationProvider interface for sending the user messages. This includes information that is security sensitive.
type NotificationProvider interface {
	Send(recipient, subject, body, htmlBody string) (err error)
	StartupCheck() (success bool, err error)
}

// UserDetails represent the details retrieved for a given user.
type UserDetails struct {
	Username    string
	DisplayName string
	Emails      []string
	Groups      []string
}

// PluginInformation provides information about plugins.
type PluginInformation struct {
	Type    PluginType
	Name    string
	Version string
	Author  string
}
