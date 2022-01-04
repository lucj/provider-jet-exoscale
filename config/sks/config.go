package sks

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_sks_cluster", func(r *config.Resource) {
                r.Kind = "Cluster"
		r.ShortGroup = "sks"
		r.ExternalName = config.IdentifierFromProvider
	})

        p.AddResourceConfigurator("exoscale_sks_nodepool", func(r *config.Resource) {
                r.Kind = "Nodepool"
                r.ShortGroup = "sks"
                r.ExternalName = config.IdentifierFromProvider
                r.References["cluster_id"] = config.Reference{
                        Type:              "Cluster",
                        RefFieldName:      "ClusterIdRef",
                }
                /*r.References["security_group_ids"] = config.Reference{
                        Type:              "SecurityGroup",
                        RefFieldName:      "SecurityGroupIdRefs",
                        SelectorFieldName: "SecurityGroupIdSelector",
                }*/
        })
}
