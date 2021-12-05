package database

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
    p.AddResourceConfigurator("exoscale_database", func(r *config.Resource) {
        // we need to override the default group that terrajet generated for
        r.ShortGroup = "database"
        r.ExternalName = config.IdentifierFromProvider
    })
}
