package mongodbatlas

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGODB_ATLAS_USERNAME", ""),
				Description: "MongoDB Atlas username",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGODB_ATLAS_API_KEY", ""),
				Description: "MongoDB Atlas API Key",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"mongodbatlas_project":   dataSourceProject(),
			"mongodbatlas_container": dataSourceContainer(),
			"mongodbatlas_cluster":   dataSourceCluster(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"mongodbatlas_project":                resourceProject(),
			"mongodbatlas_cluster":                resourceCluster(),
			"mongodbatlas_container":              resourceContainer(),
			"mongodbatlas_vpc_peering_connection": resourceVpcPeeringConnection(),
			"mongodbatlas_ip_whitelist":           resourceIPWhitelist(),
			"mongodbatlas_database_user":          resourceDatabaseUser(),
			"mongodbatlas_alert_configuration":    resourceAlertConfiguration(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AtlasUsername: d.Get("username").(string),
		AtlasAPIKey:   d.Get("api_key").(string),
	}

	client := config.NewClient()

	return client, nil
}
