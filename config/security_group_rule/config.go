package securitygrouprule

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_security_group_rule", func(r *config.Resource) {
		// we need to override the default group that terrajet generated for
		r.Kind = "SecurityGroupRule"
		r.ShortGroup = "securitygrouprule"
		r.ExternalName = config.IdentifierFromProvider
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"security_group_id"},
		}
	})
}
