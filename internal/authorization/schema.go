package authorization

import (
	"net"
	"regexp"
	"strings"

	"github.com/authelia/authelia/internal/configuration/schema"
)

// PolicyToLevel converts a string policy to int authorization level.
func PolicyToLevel(policy string) Level {
	switch policy {
	case "bypass":
		return Bypass
	case "one_factor":
		return OneFactor
	case "two_factor":
		return TwoFactor
	case "deny":
		return Denied
	}
	// By default the deny policy applies.
	return Denied
}

func schemaSubjectToACLSubject(subjectRule string) (subject AccessControlSubject) {
	if strings.HasPrefix(subjectRule, userPrefix) {
		user := strings.Trim(subjectRule[len(userPrefix):], " ")

		return AccessControlUser{Name: user}
	}

	if strings.HasPrefix(subjectRule, groupPrefix) {
		group := strings.Trim(subjectRule[len(groupPrefix):], " ")

		return AccessControlGroup{Name: group}
	}

	return nil
}

func schemaDomainsToACL(domainRules []string) (domains []AccessControlDomain) {
	for _, domainRule := range domainRules {
		domain := AccessControlDomain{}

		if strings.HasPrefix(domainRule, "*.") {
			domain.Wildcard = true
			domain.Name = domainRule[1:]
		} else {
			domain.Name = domainRule
		}

		domains = append(domains, domain)
	}

	return domains
}

func schemaResourcesToACL(resourceRules []string) (resources []AccessControlResource) {
	for _, resourceRule := range resourceRules {
		resources = append(resources, AccessControlResource{regexp.MustCompile(resourceRule)})
	}

	return resources
}

func schemaMethodsToACL(methodRules []string) (methods []string) {
	for _, method := range methodRules {
		methods = append(methods, strings.ToUpper(method))
	}

	return methods
}

func schemaNetworksToACL(networkRules []string, networksMap map[string][]*net.IPNet, networksCacheMap map[string]*net.IPNet) (networks []*net.IPNet) {
	for _, network := range networkRules {
		if _, ok := networksMap[network]; !ok {
			if _, ok := networksCacheMap[network]; ok {
				networks = append(networks, networksCacheMap[network])
			} else {
				cidr, err := parseNetwork(network)
				if err == nil {
					networks = append(networks, cidr)
				}
			}
		} else {
			networks = append(networks, networksMap[network]...)
		}
	}

	return networks
}

func parseSchemaNetworks(schemaNetworks []schema.ACLNetwork) (networksMap map[string][]*net.IPNet, networksCacheMap map[string]*net.IPNet) {
	networksMap = map[string][]*net.IPNet{}    // Used to store ptr's for the networks in the access_control.networks section.
	networksCacheMap = map[string]*net.IPNet{} // Used to store ptr's for the networks in the access_control.rules section, to prevent unnecessary additional addressing.

	for _, aclNetwork := range schemaNetworks {
		var networks []*net.IPNet

		for _, networkRule := range aclNetwork.Networks {
			cidr, err := parseNetwork(networkRule)
			if err == nil {
				networks = append(networks, cidr)
				networksCacheMap[cidr.String()] = cidr
			}
		}

		if _, ok := networksMap[aclNetwork.Name]; !ok {
			networksMap[aclNetwork.Name] = networks
		}
	}

	return networksMap, networksCacheMap
}

func parseNetwork(networkRule string) (cidr *net.IPNet, err error) {
	if !strings.Contains(networkRule, "/") {
		ip := net.ParseIP(networkRule)
		if ip.To4() != nil {
			_, cidr, err = net.ParseCIDR(networkRule + "/32")
		} else {
			_, cidr, err = net.ParseCIDR(networkRule + "/128")
		}
	} else {
		_, cidr, err = net.ParseCIDR(networkRule)
	}

	return cidr, err
}

func schemaSubjectsToACL(subjectRules [][]string) (subjects []AccessControlSubjects) {
	for _, subjectRule := range subjectRules {
		subject := AccessControlSubjects{}

		for _, subjectRuleItem := range subjectRule {
			subject.AddSubject(subjectRuleItem)
		}

		subjects = append(subjects, subject)
	}

	return subjects
}
