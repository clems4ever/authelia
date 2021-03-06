package schema

// AccessControlConfiguration represents the configuration related to ACLs.
type AccessControlConfiguration struct {
	DefaultPolicy string       `mapstructure:"default_policy"`
	Networks      []ACLNetwork `mapstructure:"networks"`
	Rules         []ACLRule    `mapstructure:"rules"`
}

// ACLNetwork represents one ACL network group entry; "weak" coerces a single value into slice.
type ACLNetwork struct {
	Name     string   `mapstructure:"name"`
	Networks []string `mapstructure:"networks"`
}

// ACLRule represents one ACL rule entry; "weak" coerces a single value into slice.
type ACLRule struct {
	Domains   []string   `mapstructure:"domain,weak"`
	Policy    string     `mapstructure:"policy"`
	Subjects  [][]string `mapstructure:"subject,weak"`
	Networks  []string   `mapstructure:"networks"`
	Resources []string   `mapstructure:"resources"`
	Methods   []string   `mapstructure:"methods"`
}

// DefaultACLNetwork represents the default configuration related to access control network group configuration.
var DefaultACLNetwork = []ACLNetwork{
	{
		Name:     "localhost",
		Networks: []string{"127.0.0.1"},
	},
	{
		Name:     "internal",
		Networks: []string{"10.0.0.0/8"},
	},
}

// DefaultACLRule represents the default configuration related to access control rule configuration.
var DefaultACLRule = []ACLRule{
	{
		Domains: []string{"public.example.com"},
		Policy:  "bypass",
	},
	{
		Domains: []string{"singlefactor.example.com"},
		Policy:  "one_factor",
	},
	{
		Domains: []string{"secure.example.com"},
		Policy:  "two_factor",
	},
}
