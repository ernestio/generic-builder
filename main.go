/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"os"
	"runtime"

	l "github.com/ernestio/builder-library"
)

var s l.Scheduler

func main() {
	s.Setup(os.Getenv("NATS_URI"))

	// Configure VPCs Hooks
	s.ProcessRequest("vpcs.create", "vpc.create")
	s.ProcessRequest("vpcs.delete", "vpc.delete")

	s.ProcessSuccessResponse("vpc.create.done", "vpc.create", "vpcs.create.done")
	s.ProcessSuccessResponse("vpc.delete.done", "vpc.delete", "vpcs.delete.done")

	s.ProcessFailedResponse("vpc.create.error", "vpcs.create.error")
	s.ProcessFailedResponse("vpc.delete.error", "vpcs.delete.error")

	// Configure Routers Hooks
	s.ProcessRequest("routers.create", "router.create")
	s.ProcessRequest("routers.delete", "router.delete")

	s.ProcessSuccessResponse("router.create.done", "router.create", "routers.create.done")
	s.ProcessSuccessResponse("router.delete.done", "router.delete", "routers.delete.done")

	s.ProcessFailedResponse("router.create.error", "routers.create.error")
	s.ProcessFailedResponse("router.delete.error", "routers.delete.error")

	// Configure Networks Hooks
	s.ProcessRequest("networks.create", "network.create")
	s.ProcessRequest("networks.delete", "network.delete")

	s.ProcessSuccessResponse("network.create.done", "network.create", "networks.create.done")
	s.ProcessSuccessResponse("network.delete.done", "network.delete", "networks.delete.done")

	s.ProcessFailedResponse("network.create.error", "networks.create.error")
	s.ProcessFailedResponse("network.delete.error", "networks.delete.error")

	// Configure Instance Hooks
	s.ProcessRequest("instances.create", "instance.create")
	s.ProcessRequest("instances.update", "instance.update")
	s.ProcessRequest("instances.delete", "instance.delete")

	s.ProcessSuccessResponse("instance.create.done", "instance.create", "instances.create.done")
	s.ProcessSuccessResponse("instance.update.done", "instance.update", "instances.update.done")
	s.ProcessSuccessResponse("instance.delete.done", "instance.delete", "instances.delete.done")

	s.ProcessFailedResponse("instance.create.error", "instances.create.error")
	s.ProcessFailedResponse("instance.update.error", "instances.update.error")
	s.ProcessFailedResponse("instance.delete.error", "instances.delete.error")

	// Configure Firewall Hooks
	s.ProcessRequest("firewalls.create", "firewall.create")
	s.ProcessRequest("firewalls.update", "firewall.update")
	s.ProcessRequest("firewalls.delete", "firewall.delete")

	s.ProcessSuccessResponse("firewall.create.done", "firewall.create", "firewalls.create.done")
	s.ProcessSuccessResponse("firewall.update.done", "firewall.update", "firewalls.update.done")
	s.ProcessSuccessResponse("firewall.delete.done", "firewall.delete", "firewalls.delete.done")

	s.ProcessFailedResponse("firewall.create.error", "firewalls.create.error")
	s.ProcessFailedResponse("firewall.update.error", "firewalls.update.error")
	s.ProcessFailedResponse("firewall.delete.error", "firewalls.delete.error")

	// Configure Nat Hooks
	s.ProcessRequest("nats.create", "nat.create")
	s.ProcessRequest("nats.update", "nat.update")
	s.ProcessRequest("nats.delete", "nat.delete")

	s.ProcessSuccessResponse("nat.create.done", "nat.create", "nats.create.done")
	s.ProcessSuccessResponse("nat.update.done", "nat.update", "nats.update.done")
	s.ProcessSuccessResponse("nat.delete.done", "nat.delete", "nats.delete.done")

	s.ProcessFailedResponse("nat.create.error", "nats.create.error")
	s.ProcessFailedResponse("nat.update.error", "nats.update.error")
	s.ProcessFailedResponse("nat.delete.error", "nats.delete.error")

	// Configure ELBs Hooks
	s.ProcessRequest("elbs.create", "elb.create")
	s.ProcessRequest("elbs.delete", "elb.delete")
	s.ProcessRequest("elbs.update", "elb.update")

	s.ProcessSuccessResponse("elb.create.done", "elb.create", "elbs.create.done")
	s.ProcessSuccessResponse("elb.delete.done", "elb.delete", "elbs.delete.done")
	s.ProcessSuccessResponse("elb.update.done", "elb.update", "elbs.update.done")

	s.ProcessFailedResponse("elb.create.error", "elbs.create.error")
	s.ProcessFailedResponse("elb.delete.error", "elbs.delete.error")
	s.ProcessFailedResponse("elb.update.error", "elbs.update.error")

	// Configure S3 Hooks
	s.ProcessRequest("s3s.create", "s3.create")
	s.ProcessRequest("s3s.delete", "s3.delete")
	s.ProcessRequest("s3s.update", "s3.update")

	s.ProcessSuccessResponse("s3.create.done", "s3.create", "s3s.create.done")
	s.ProcessSuccessResponse("s3.delete.done", "s3.delete", "s3s.delete.done")
	s.ProcessSuccessResponse("s3.update.done", "s3.update", "s3s.update.done")

	s.ProcessFailedResponse("s3.create.error", "s3s.create.error")
	s.ProcessFailedResponse("s3.delete.error", "s3s.delete.error")
	s.ProcessFailedResponse("s3.update.error", "s3s.update.error")

	// Configure EBSs Hooks
	s.ProcessRequest("ebs_volumes.create", "ebs_volume.create")
	s.ProcessRequest("ebs_volumes.delete", "ebs_volume.delete")
	s.ProcessRequest("ebs_volumes.update", "ebs_volume.update")

	s.ProcessSuccessResponse("ebs_volume.create.done", "ebs_volume.create", "ebs_volumes.create.done")
	s.ProcessSuccessResponse("ebs_volume.delete.done", "ebs_volume.delete", "ebs_volumes.delete.done")
	s.ProcessSuccessResponse("ebs_volume.update.done", "ebs_volume.update", "ebs_volumes.update.done")

	s.ProcessFailedResponse("ebs_volume.create.error", "ebs_volumes.create.error")
	s.ProcessFailedResponse("ebs_volume.delete.error", "ebs_volumes.delete.error")
	s.ProcessFailedResponse("ebs_volume.update.error", "ebs_volumes.update.error")

	// Configure RDS cluster Hooks
	s.ProcessRequest("rds_clusters.create", "rds_cluster.create")
	s.ProcessRequest("rds_clusters.delete", "rds_cluster.delete")
	s.ProcessRequest("rds_clusters.update", "rds_cluster.update")

	s.ProcessSuccessResponse("rds_cluster.create.done", "rds_cluster.create", "rds_clusters.create.done")
	s.ProcessSuccessResponse("rds_cluster.delete.done", "rds_cluster.delete", "rds_clusters.delete.done")
	s.ProcessSuccessResponse("rds_cluster.update.done", "rds_cluster.update", "rds_clusters.update.done")

	s.ProcessFailedResponse("rds_cluster.create.error", "rds_clusters.create.error")
	s.ProcessFailedResponse("rds_cluster.delete.error", "rds_clusters.delete.error")
	s.ProcessFailedResponse("rds_cluster.update.error", "rds_clusters.update.error")

	// Configure RDS instance Hooks
	s.ProcessRequest("rds_instances.create", "rds_instance.create")
	s.ProcessRequest("rds_instances.delete", "rds_instance.delete")
	s.ProcessRequest("rds_instances.update", "rds_instance.update")

	s.ProcessSuccessResponse("rds_instance.create.done", "rds_instance.create", "rds_instances.create.done")
	s.ProcessSuccessResponse("rds_instance.delete.done", "rds_instance.delete", "rds_instances.delete.done")
	s.ProcessSuccessResponse("rds_instance.update.done", "rds_instance.update", "rds_instances.update.done")

	s.ProcessFailedResponse("rds_instance.create.error", "rds_instances.create.error")
	s.ProcessFailedResponse("rds_instance.delete.error", "rds_instances.delete.error")
	s.ProcessFailedResponse("rds_instance.update.error", "rds_instances.update.error")

	// Configure Route53 Hooks
	s.ProcessRequest("route53s.create", "route53.create")
	s.ProcessRequest("route53s.delete", "route53.delete")
	s.ProcessRequest("route53s.update", "route53.update")

	s.ProcessSuccessResponse("route53.create.done", "route53.create", "route53s.create.done")
	s.ProcessSuccessResponse("route53.delete.done", "route53.delete", "route53s.delete.done")
	s.ProcessSuccessResponse("route53.update.done", "route53.update", "route53s.update.done")

	s.ProcessFailedResponse("route53.create.error", "route53s.create.error")
	s.ProcessFailedResponse("route53.delete.error", "route53s.delete.error")
	s.ProcessFailedResponse("route53.update.error", "route53s.update.error")

	// Configure ElasticSearches Hooks
	s.ProcessRequest("elasticsearches.create", "elasticsearch.create")
	s.ProcessRequest("elasticsearches.delete", "elasticsearch.delete")
	s.ProcessRequest("elasticsearches.update", "elasticsearch.update")

	s.ProcessSuccessResponse("elasticsearch.create.done", "elasticsearch.create", "elasticsearches.create.done")
	s.ProcessSuccessResponse("elasticsearch.delete.done", "elasticsearch.delete", "elasticsearches.delete.done")
	s.ProcessSuccessResponse("elasticsearch.update.done", "elasticsearch.update", "elasticsearches.update.done")

	s.ProcessFailedResponse("elasticsearch.create.error", "elasticsearches.create.error")
	s.ProcessFailedResponse("elasticsearch.delete.error", "elasticsearches.delete.error")
	s.ProcessFailedResponse("elasticsearch.update.error", "elasticsearches.update.error")

	runtime.Goexit()
}
