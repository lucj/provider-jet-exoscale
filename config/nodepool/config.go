package nodepool

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_sks_nodepool", func(r *config.Resource) {
		r.ShortGroup = "nodepool"
		r.ExternalName = config.IdentifierFromProvider
                r.References["cluster_id"] = config.Reference{
                        Type:              "github.com/lucj/provider-jet-exoscale/apis/sks/v1alpha1",
			RefFieldName:      "ClusterIdRef",
                }
                r.References["security_group_ids"] = config.Reference{
			Type:              "github.com/lucj/provider-jet-exoscale/apis/sks/v1alpha1",
			RefFieldName:      "SecurityGroupIdRefs",
			SelectorFieldName: "SecurityGroupIdSelector",
		}
	})
}
