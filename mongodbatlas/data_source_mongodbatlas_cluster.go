package mongodbatlas

import (
	"fmt"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceClusterRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mongodb_major_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"provider_backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backing_provider": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disk_size_gb": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"replication_factor": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"num_shards": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"paused": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disk_gb_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mongodb_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mongo_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mongo_uri_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mongo_uri_with_options": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_spec": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:     schema.TypeString,
							Required: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"electable_nodes": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"read_only_nodes": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},
		},
	}
}

func dataSourceClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ma.Client)
	group := d.Get("group").(string)
	name := d.Get("name").(string)
	c, _, err := client.Clusters.Get(group, name)
	if err != nil {
		return fmt.Errorf("Error reading MongoDB Cluster with name %s: %s", name, err)
	}
	replicationSpecs := []interface{}{}
	for region, replicationSpec := range c.ReplicationSpec {
		spec := map[string]interface{}{
			"region":          region,
			"priority":        replicationSpec.Priority,
			"electable_nodes": replicationSpec.ElectableNodes,
			"read_only_nodes": replicationSpec.ReadOnlyNodes,
		}
		replicationSpecs = append(replicationSpecs, spec)
	}

	d.SetId(c.ID)
	d.Set("name", c.Name)
	d.Set("group", c.GroupID)
	d.Set("mongodb_major_version", c.MongoDBMajorVersion)
	d.Set("backup", c.BackupEnabled)
	d.Set("provider_backup", c.ProviderBackupEnabled)
	d.Set("size", c.ProviderSettings.InstanceSizeName)
	d.Set("provider_name", c.ProviderSettings.ProviderName)
	d.Set("backing_provider", c.ProviderSettings.BackingProviderName)
	d.Set("region", c.ProviderSettings.RegionName)
	d.Set("disk_size_gb", c.DiskSizeGB)
	d.Set("disk_gb_enabled", c.AutoScaling.DiskGBEnabled)
	d.Set("replication_factor", c.ReplicationFactor)
	d.Set("state", c.StateName)
	d.Set("num_shards", c.NumShards)
	d.Set("paused", c.Paused)
	d.Set("mongodb_version", c.MongoDBVersion)
	d.Set("mongo_uri", c.MongoURI)
	d.Set("mongo_uri_updated", c.MongoURIUpdated)
	d.Set("mongo_uri_with_options", c.MongoURIWithOptions)
	d.Set("replication_spec", replicationSpecs)

	return nil
}
