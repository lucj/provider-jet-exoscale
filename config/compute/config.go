package compute

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_compute", func(r *config.Resource) {
		// we need to override the default group that terrajet generated for
		r.Kind = "Instance"
		r.ShortGroup = "compute"
		r.ExternalName = config.IdentifierFromProvider
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"template_id", "security_group_ids"},
		}
	})
}
