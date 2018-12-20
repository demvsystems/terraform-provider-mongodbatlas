---
layout: "mongodbatlas"
page_title: "MongoDB Atlas: cluster"
sidebar_current: "docs-mongodbatlas-datasource-cluster"
description: |-
    Provides details about a specific Cluster
---

# Data Source: mongodbatlas_cluster

`mongodbatlas_cluster` provides details about a specific Cluster.

This data source can prove useful when looking up the details of a previously created Cluster.

-> **NOTE:** Groups and projects are synonymous terms. `group` arguments on resources are the project ID.

## Example Usage

```hcl
data "mongodbatlas_project" "project" {
  name = "my-project"
}

data "mongodbatlas_cluster" "example" {
  group      = "${data.mongodbatlas_project.project.id}"
  name = "my-cluster-name"
}
```

## Argument Reference

* `group` - (Required) The ID of the project that the desired cluster belongs to.
* `identifier` - (Required) The name of the desired cluster.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The same as `identifier`.
* `mongodb_major_version` - Version of the cluster.
* `backup` - If continuous backups are enabled.
* `provider_backup` - If cloud provider snapshots are enabled.
* `size` - Instance size of all data-bearing servers in the cluster.
* `provider_name` - Name of the cloud provider. One of: `AWS`, `GCP`, `AZURE`, `TENANT`.
* `backing_provider` - The cloud service provider for a shared tier cluster. One of `AWS`, `GCP`, `AZURE`.
* `region` - Atlas-style name of the region in which the cluster was created. e.g. `US_EAST_1`.
* `disk_size_gb` - AWS/GCP only. Size in GB of the server's root volume.
* `disk_gb_enabled` - If disk auto-scaling is enabled.
* `replication_factor` - Number of replica set members. Each shard is a replica set with the specified replication factor if a sharded cluster.
* `state` - Current state of the cluster. One of: `IDLE`, `CREATING`,
    `UPDATING`, `DELETING`, `DELETED`, `REPAIRING`.
* `num_shards` - The number of active shards.
* `paused` - Flag that indicates whether the cluster is paused.
* `mongodb_version` - Version of MongoDB deployed. Major.Minor.Patch.
* `mongo_uri` - Base connection string for the cluster. See `mongo_uri_with_options` for a more usable connection string.
* `mongo_uri_updated` - When the connection string was last updated. Connection string changes, for example, if you change a replica set to a sharded cluster.
* `mongo_uri_with_options` - Connection string for connecting to the Atlas cluster. Includes necessary query parameters with values appropriate for the cluster. Include a username and password for a MongoDB user associated with the project after the `mongodb://` to actually connect.
