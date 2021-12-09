package securitygroup

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_security_group", func(r *config.Resource) {
		// we need to override the default group that terrajet generated for
		r.Kind = "SecurityGroup"
		r.ShortGroup = "securitygroup"
		r.ExternalName = config.IdentifierFromProvider
	})
}
