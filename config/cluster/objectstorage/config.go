package objectstorage

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cancom_object_storage_bucket", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
	})

	p.AddResourceConfigurator("cancom_object_storage_user", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
	})
}
