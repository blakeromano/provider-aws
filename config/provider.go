// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"github.com/upbound/provider-aws/config/acm"
	"github.com/upbound/provider-aws/config/acmpca"
	"github.com/upbound/provider-aws/config/apigateway"
	"github.com/upbound/provider-aws/config/apigatewayv2"
	"github.com/upbound/provider-aws/config/apprunner"
	"github.com/upbound/provider-aws/config/appstream"
	"github.com/upbound/provider-aws/config/athena"
	"github.com/upbound/provider-aws/config/autoscaling"
	"github.com/upbound/provider-aws/config/backup"
	"github.com/upbound/provider-aws/config/bedrock"
	"github.com/upbound/provider-aws/config/budgets"
	"github.com/upbound/provider-aws/config/cloudformation"
	"github.com/upbound/provider-aws/config/cloudfront"
	"github.com/upbound/provider-aws/config/cloudsearch"
	"github.com/upbound/provider-aws/config/cloudwatch"
	"github.com/upbound/provider-aws/config/cloudwatchevents"
	"github.com/upbound/provider-aws/config/cloudwatchlogs"
	"github.com/upbound/provider-aws/config/cognitoidentity"
	"github.com/upbound/provider-aws/config/cognitoidp"
	"github.com/upbound/provider-aws/config/connect"
	"github.com/upbound/provider-aws/config/cur"
	"github.com/upbound/provider-aws/config/datasync"
	"github.com/upbound/provider-aws/config/dax"
	"github.com/upbound/provider-aws/config/devicefarm"
	"github.com/upbound/provider-aws/config/directconnect"
	"github.com/upbound/provider-aws/config/dms"
	"github.com/upbound/provider-aws/config/docdb"
	"github.com/upbound/provider-aws/config/ds"
	"github.com/upbound/provider-aws/config/dynamodb"
	"github.com/upbound/provider-aws/config/ebs"
	"github.com/upbound/provider-aws/config/ec2"
	"github.com/upbound/provider-aws/config/ecr"
	"github.com/upbound/provider-aws/config/ecrpublic"
	"github.com/upbound/provider-aws/config/ecs"
	"github.com/upbound/provider-aws/config/efs"
	"github.com/upbound/provider-aws/config/eks"
	"github.com/upbound/provider-aws/config/elasticache"
	"github.com/upbound/provider-aws/config/elb"
	"github.com/upbound/provider-aws/config/elbv2"
	"github.com/upbound/provider-aws/config/firehose"
	"github.com/upbound/provider-aws/config/fsx"
	"github.com/upbound/provider-aws/config/gamelift"
	"github.com/upbound/provider-aws/config/globalaccelerator"
	"github.com/upbound/provider-aws/config/glue"
	"github.com/upbound/provider-aws/config/grafana"
	"github.com/upbound/provider-aws/config/iam"
	"github.com/upbound/provider-aws/config/identitystore"
	"github.com/upbound/provider-aws/config/iot"
	"github.com/upbound/provider-aws/config/kafka"
	"github.com/upbound/provider-aws/config/kafkaconnect"
	"github.com/upbound/provider-aws/config/kendra"
	"github.com/upbound/provider-aws/config/kinesis"
	"github.com/upbound/provider-aws/config/kinesisanalytics"
	kinesisanalytics2 "github.com/upbound/provider-aws/config/kinesisanalyticsv2"
	"github.com/upbound/provider-aws/config/kms"
	"github.com/upbound/provider-aws/config/lakeformation"
	"github.com/upbound/provider-aws/config/lambda"
	"github.com/upbound/provider-aws/config/licensemanager"
	"github.com/upbound/provider-aws/config/medialive"
	"github.com/upbound/provider-aws/config/memorydb"
	"github.com/upbound/provider-aws/config/mq"
	"github.com/upbound/provider-aws/config/neptune"
	"github.com/upbound/provider-aws/config/networkfirewall"
	"github.com/upbound/provider-aws/config/networkmanager"
	"github.com/upbound/provider-aws/config/opensearch"
	"github.com/upbound/provider-aws/config/opensearchserverless"
	"github.com/upbound/provider-aws/config/opsworks"
	"github.com/upbound/provider-aws/config/organization"
	"github.com/upbound/provider-aws/config/qldb"
	"github.com/upbound/provider-aws/config/ram"
	"github.com/upbound/provider-aws/config/rds"
	"github.com/upbound/provider-aws/config/redshift"
	"github.com/upbound/provider-aws/config/redshiftserverless"
	"github.com/upbound/provider-aws/config/rolesanywhere"
	"github.com/upbound/provider-aws/config/route53"
	"github.com/upbound/provider-aws/config/route53recoverycontrolconfig"
	"github.com/upbound/provider-aws/config/route53resolver"
	"github.com/upbound/provider-aws/config/s3"
	"github.com/upbound/provider-aws/config/sagemaker"
	"github.com/upbound/provider-aws/config/secretsmanager"
	"github.com/upbound/provider-aws/config/servicecatalog"
	"github.com/upbound/provider-aws/config/servicediscovery"
	"github.com/upbound/provider-aws/config/sfn"
	"github.com/upbound/provider-aws/config/sns"
	"github.com/upbound/provider-aws/config/sqs"
	"github.com/upbound/provider-aws/config/ssoadmin"
	"github.com/upbound/provider-aws/config/transfer"
)

func init() {
	ProviderConfiguration.AddConfig(acm.Configure)
	ProviderConfiguration.AddConfig(acmpca.Configure)
	ProviderConfiguration.AddConfig(apigateway.Configure)
	ProviderConfiguration.AddConfig(apigatewayv2.Configure)
	ProviderConfiguration.AddConfig(apprunner.Configure)
	ProviderConfiguration.AddConfig(appstream.Configure)
	ProviderConfiguration.AddConfig(athena.Configure)
	ProviderConfiguration.AddConfig(autoscaling.Configure)
	ProviderConfiguration.AddConfig(backup.Configure)
	ProviderConfiguration.AddConfig(bedrock.Configure)
	ProviderConfiguration.AddConfig(cloudfront.Configure)
	ProviderConfiguration.AddConfig(cloudsearch.Configure)
	ProviderConfiguration.AddConfig(cloudwatch.Configure)
	ProviderConfiguration.AddConfig(cloudwatchlogs.Configure)
	ProviderConfiguration.AddConfig(cognitoidentity.Configure)
	ProviderConfiguration.AddConfig(cognitoidp.Configure)
	ProviderConfiguration.AddConfig(connect.Configure)
	ProviderConfiguration.AddConfig(cur.Configure)
	ProviderConfiguration.AddConfig(datasync.Configure)
	ProviderConfiguration.AddConfig(dax.Configure)
	ProviderConfiguration.AddConfig(devicefarm.Configure)
	ProviderConfiguration.AddConfig(dms.Configure)
	ProviderConfiguration.AddConfig(docdb.Configure)
	ProviderConfiguration.AddConfig(dynamodb.Configure)
	ProviderConfiguration.AddConfig(ebs.Configure)
	ProviderConfiguration.AddConfig(ec2.Configure)
	ProviderConfiguration.AddConfig(ecr.Configure)
	ProviderConfiguration.AddConfig(ecrpublic.Configure)
	ProviderConfiguration.AddConfig(ecs.Configure)
	ProviderConfiguration.AddConfig(efs.Configure)
	ProviderConfiguration.AddConfig(eks.Configure)
	ProviderConfiguration.AddConfig(elasticache.Configure)
	ProviderConfiguration.AddConfig(elb.Configure)
	ProviderConfiguration.AddConfig(elbv2.Configure)
	ProviderConfiguration.AddConfig(firehose.Configure)
	ProviderConfiguration.AddConfig(gamelift.Configure)
	ProviderConfiguration.AddConfig(globalaccelerator.Configure)
	ProviderConfiguration.AddConfig(glue.Configure)
	ProviderConfiguration.AddConfig(grafana.Configure)
	ProviderConfiguration.AddConfig(iam.Configure)
	ProviderConfiguration.AddConfig(kafka.Configure)
	ProviderConfiguration.AddConfig(kafkaconnect.Configure)
	ProviderConfiguration.AddConfig(kinesis.Configure)
	ProviderConfiguration.AddConfig(kinesisanalytics.Configure)
	ProviderConfiguration.AddConfig(kinesisanalytics2.Configure)
	ProviderConfiguration.AddConfig(kms.Configure)
	ProviderConfiguration.AddConfig(lakeformation.Configure)
	ProviderConfiguration.AddConfig(lambda.Configure)
	ProviderConfiguration.AddConfig(licensemanager.Configure)
	ProviderConfiguration.AddConfig(memorydb.Configure)
	ProviderConfiguration.AddConfig(mq.Configure)
	ProviderConfiguration.AddConfig(neptune.Configure)
	ProviderConfiguration.AddConfig(networkfirewall.Configure)
	ProviderConfiguration.AddConfig(opensearch.Configure)
	ProviderConfiguration.AddConfig(opensearchserverless.Configure)
	ProviderConfiguration.AddConfig(ram.Configure)
	ProviderConfiguration.AddConfig(rds.Configure)
	ProviderConfiguration.AddConfig(redshift.Configure)
	ProviderConfiguration.AddConfig(rolesanywhere.Configure)
	ProviderConfiguration.AddConfig(route53.Configure)
	ProviderConfiguration.AddConfig(route53resolver.Configure)
	ProviderConfiguration.AddConfig(route53recoverycontrolconfig.Configure)
	ProviderConfiguration.AddConfig(s3.Configure)
	ProviderConfiguration.AddConfig(secretsmanager.Configure)
	ProviderConfiguration.AddConfig(servicecatalog.Configure)
	ProviderConfiguration.AddConfig(organization.Configure)
	ProviderConfiguration.AddConfig(cloudwatchevents.Configure)
	ProviderConfiguration.AddConfig(budgets.Configure)
	ProviderConfiguration.AddConfig(servicediscovery.Configure)
	ProviderConfiguration.AddConfig(sfn.Configure)
	ProviderConfiguration.AddConfig(sns.Configure)
	ProviderConfiguration.AddConfig(sqs.Configure)
	ProviderConfiguration.AddConfig(transfer.Configure)
	ProviderConfiguration.AddConfig(directconnect.Configure)
	ProviderConfiguration.AddConfig(ds.Configure)
	ProviderConfiguration.AddConfig(qldb.Configure)
	ProviderConfiguration.AddConfig(fsx.Configure)
	ProviderConfiguration.AddConfig(networkmanager.Configure)
	ProviderConfiguration.AddConfig(opsworks.Configure)
	ProviderConfiguration.AddConfig(sagemaker.Configure)
	ProviderConfiguration.AddConfig(redshiftserverless.Configure)
	ProviderConfiguration.AddConfig(kendra.Configure)
	ProviderConfiguration.AddConfig(medialive.Configure)
	ProviderConfiguration.AddConfig(ssoadmin.Configure)
	ProviderConfiguration.AddConfig(identitystore.Configure)
	ProviderConfiguration.AddConfig(iot.Configure)
	ProviderConfiguration.AddConfig(cloudformation.Configure)
}
