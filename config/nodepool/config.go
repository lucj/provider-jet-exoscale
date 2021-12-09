package nodepool

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_sks_nodepool", func(r *config.Resource) {
		// we need to override the default group that terrajet generated for
		r.ShortGroup = "nodepool"
		r.ExternalName = config.IdentifierFromProvider
	})
}
