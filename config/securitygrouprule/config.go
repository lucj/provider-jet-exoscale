package securitygrouprule

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_security_group_rule", func(r *config.Resource) {
		// we need to override the default group that terrajet generated for
		r.Kind = "SecurityGroupRule"
		r.ShortGroup = "securitygrouprule"
		r.ExternalName = config.IdentifierFromProvider
		r.References["security_group_id"] = config.Reference{
			Type:              "github.com/lucj/provider-jet-exoscale/apis/securitygroup/v1alpha1.SecurityGroup",
			RefFieldName:      "SecurityGroupIdRef",
			SelectorFieldName: "SecurityGroupIdSelector",
		}
		r.References["user_security_group_id"] = config.Reference{
			Type:              "github.com/lucj/provider-jet-exoscale/apis/securitygroup/v1alpha1.SecurityGroup",
			RefFieldName:      "UserSecurityGroupIdRef",
			SelectorFieldName: "UserSecurityGroupIdSelector",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"security_group", "security_group_id", "user_security_group", "user_security_group_id"},
		}
	})
}
