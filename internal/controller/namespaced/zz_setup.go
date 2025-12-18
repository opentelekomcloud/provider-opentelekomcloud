// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	aclpolicyassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/aclpolicyassociatev2"
	aclpolicyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/aclpolicyv2"
	apipublishmentv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/apipublishmentv2"
	apiv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/apiv2"
	appcodev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/appcodev2"
	applicationauthorizationv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/applicationauthorizationv2"
	applicationv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/applicationv2"
	certificatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/certificatev2"
	customauthorizerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/customauthorizerv2"
	environmentv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/environmentv2"
	environmentvariablev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/environmentvariablev2"
	gatewayfeaturev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/gatewayfeaturev2"
	gatewayroutesv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/gatewayroutesv2"
	gatewayv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/gatewayv2"
	groupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/groupv2"
	responsev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/responsev2"
	signatureassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/signatureassociatev2"
	signaturev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/signaturev2"
	throttlingpolicyassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/throttlingpolicyassociatev2"
	throttlingpolicyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/throttlingpolicyv2"
	vpcchannelv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/apigw/vpcchannelv2"
	servicemeshv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/asm/servicemeshv1"
	volumev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/blockstorage/volumev2"
	policyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cbr/policyv3"
	vaultv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cbr/vaultv3"
	addonv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cce/addonv3"
	clusterv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cce/clusterv3"
	nodeattachv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cce/nodeattachv3"
	nodepoolv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cce/nodepoolv3"
	nodev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cce/nodev3"
	aclrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/aclrulev1"
	addressgroupmemberv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/addressgroupmemberv1"
	addressgroupv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/addressgroupv1"
	blacklistwhitelistrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/blacklistwhitelistrulev1"
	domainnamegroupv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/domainnamegroupv1"
	eipprotectionv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/eipprotectionv1"
	firewallv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/firewallv1"
	ipsprotectionv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/ipsprotectionv1"
	servicegroupmemberv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/servicegroupmemberv1"
	servicegroupv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cfw/servicegroupv1"
	floatingipassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/compute/floatingipassociatev2"
	floatingipv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/compute/floatingipv2"
	instancev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/compute/instancev2"
	keypairv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/compute/keypairv2"
	secgroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/compute/secgroupv2"
	servergroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/compute/servergroupv2"
	volumeattachv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/compute/volumeattachv2"
	clusterrestartv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/css/clusterrestartv1"
	clusterv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/css/clusterv1"
	configurationv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/css/configurationv1"
	snapshotconfigurationv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/css/snapshotconfigurationv1"
	eventnotificationv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cts/eventnotificationv3"
	trackerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/cts/trackerv3"
	instancev2dcs "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dcs/instancev2"
	instancev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/ddm/instancev1"
	schemav1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/ddm/schemav1"
	backupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dds/backupv3"
	instancev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dds/instancev3"
	ltslogv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dds/ltslogv3"
	hostv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/deh/hostv1"
	appv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dis/appv2"
	checkpointv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dis/checkpointv2"
	dumptaskv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dis/dumptaskv2"
	streamv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dis/streamv2"
	consumergroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/consumergroupv2"
	dedicatedinstancev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/dedicatedinstancev2"
	instancev1dms "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/instancev1"
	instancev2dms "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/instancev2"
	reassignpartitionsv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/reassignpartitionsv2"
	smartconnecttaskactionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/smartconnecttaskactionv2"
	smartconnecttaskv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/smartconnecttaskv2"
	smartconnectv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/smartconnectv2"
	topicv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/topicv1"
	userpermissionv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/userpermissionv1"
	userv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dms/userv2"
	ptrrecordv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dns/ptrrecordv2"
	recordsetv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dns/recordsetv2"
	zonev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dns/zonev2"
	taskv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/drs/taskv3"
	clusterv1dws "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/dws/clusterv1"
	instancev1ecs "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/ecs/instancev1"
	vpnconnectionmonitorv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/enterprisevpn/vpnconnectionmonitorv5"
	vpnconnectionv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/enterprisevpn/vpnconnectionv5"
	vpncustomergatewayv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/enterprisevpn/vpncustomergatewayv5"
	vpngatewayv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/enterprisevpn/vpngatewayv5"
	associationv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/er/associationv3"
	flowlogv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/er/flowlogv3"
	instancev3er "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/er/instancev3"
	propagationv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/er/propagationv3"
	routetablev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/er/routetablev3"
	staticroutev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/er/staticroutev3"
	vpcattachmentv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/er/vpcattachmentv3"
	volumev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/evs/volumev3"
	asyncinvokeconfigv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/fgs/asyncinvokeconfigv2"
	dependencyversionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/fgs/dependencyversionv2"
	eventv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/fgs/eventv2"
	functionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/fgs/functionv2"
	triggerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/fgs/triggerv2"
	firewallgroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/fw/firewallgroupv2"
	policyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/fw/policyv2"
	rulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/fw/rulev2"
	instancev3gemini "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/gemini/instancev3"
	hostgroupv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/hss/hostgroupv5"
	hostprotectionv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/hss/hostprotectionv5"
	agencyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/agencyv3"
	credentialv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/credentialv3"
	groupmembershipv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/groupmembershipv3"
	groupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/groupv3"
	loginpolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/loginpolicyv3"
	mappingv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/mappingv3"
	passwordpolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/passwordpolicyv3"
	projectv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/projectv3"
	protectionpolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/protectionpolicyv3"
	protocolv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/protocolv3"
	provider "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/provider"
	providerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/providerv3"
	roleassignmentv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/roleassignmentv3"
	rolev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/rolev3"
	usergroupmembershipv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/usergroupmembershipv3"
	userv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/identity/userv3"
	imageaccessacceptv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/images/imageaccessacceptv2"
	imageaccessv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/images/imageaccessv2"
	imagev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/images/imagev2"
	dataimagev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/ims/dataimagev2"
	imagev2ims "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/ims/imagev2"
	grantv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/kms/grantv1"
	keymaterialv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/kms/keymaterialv1"
	keyv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/kms/keyv1"
	certificatev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/certificatev3"
	ipgroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/ipgroupv3"
	listenerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/listenerv3"
	loadbalancerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/loadbalancerv3"
	ltslogv3lb "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/ltslogv3"
	memberv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/memberv3"
	monitorv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/monitorv3"
	policyv3lb "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/policyv3"
	poolv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/poolv3"
	rulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/rulev3"
	securitypolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lb/securitypolicyv3"
	groupv2logtank "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/logtank/groupv2"
	topicv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/logtank/topicv2"
	transferv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/logtank/transferv2"
	cceaccessv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/cceaccessv3"
	crossaccountaccessv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/crossaccountaccessv2"
	groupv2lts "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/groupv2"
	hostaccessv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/hostaccessv3"
	hostgroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/hostgroupv3"
	keywordsalarmrulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/keywordsalarmrulev2"
	notificationtemplatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/notificationtemplatev2"
	quicksearchcriteriav1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/quicksearchcriteriav1"
	streamv2lts "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/streamv2"
	transferv2lts "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/lts/transferv2"
	dnatrulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/nat/dnatrulev2"
	gatewayv2nat "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/nat/gatewayv2"
	snatrulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/nat/snatrulev2"
	floatingipassociatev2networking "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/floatingipassociatev2"
	floatingipv2networking "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/floatingipv2"
	networkv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/networkv2"
	portsecgroupassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/portsecgroupassociatev2"
	portv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/portv2"
	routerinterfacev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/routerinterfacev2"
	routerroutev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/routerroutev2"
	routerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/routerv2"
	secgrouprulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/secgrouprulev2"
	secgroupv2networking "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/secgroupv2"
	subnetv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/subnetv2"
	vipassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/vipassociatev2"
	vipv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/networking/vipv2"
	bucket "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/obs/bucket"
	bucketacl "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/obs/bucketacl"
	bucketinventory "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/obs/bucketinventory"
	bucketobject "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/obs/bucketobject"
	bucketobjectacl "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/obs/bucketobjectacl"
	bucketpolicy "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/obs/bucketpolicy"
	bucketreplication "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/obs/bucketreplication"
	natdnatrulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/privatenat/natdnatrulev3"
	natgatewayv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/privatenat/natgatewayv3"
	natsnatrulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/privatenat/natsnatrulev3"
	nattransitipv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/privatenat/nattransitipv3"
	providerconfig "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/providerconfig"
	backupv3rds "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rds/backupv3"
	instancev3rds "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rds/instancev3"
	maintenancev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rds/maintenancev3"
	parametergroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rds/parametergroupv3"
	publicipassociatev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rds/publicipassociatev3"
	readreplicav3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rds/readreplicav3"
	advancedqueryv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rms/advancedqueryv1"
	policyassignmentevaluatev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rms/policyassignmentevaluatev1"
	policyassignmentv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rms/policyassignmentv1"
	resourcerecorderv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rms/resourcerecorderv1"
	softwareconfigv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rts/softwareconfigv1"
	softwaredeploymentv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rts/softwaredeploymentv1"
	stackv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/rts/stackv1"
	buckets3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/s3/bucket"
	bucketobjects3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/s3/bucketobject"
	bucketpolicys3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/s3/bucketpolicy"
	filesystemv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/sfs/filesystemv2"
	shareaccessrulesv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/sfs/shareaccessrulesv2"
	turbosharev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/sfs/turbosharev1"
	subscriptionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/smn/subscriptionv2"
	topicattributev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/smn/topicattributev2"
	topicv2smn "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/smn/topicv2"
	domainv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/swr/domainv2"
	organizationpermissionsv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/swr/organizationpermissionsv2"
	organizationv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/swr/organizationv2"
	repositoryv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/swr/repositoryv2"
	mysqlbackupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/taurusdb/mysqlbackupv3"
	mysqlinstancev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/taurusdb/mysqlinstancev3"
	mysqlproxyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/taurusdb/mysqlproxyv3"
	mysqlquotav3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/taurusdb/mysqlquotav3"
	mysqlsqlcontrolrulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/taurusdb/mysqlsqlcontrolrulev3"
	resourcetagsv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/tms/resourcetagsv1"
	tagsv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/tms/tagsv1"
	bandwidthassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/bandwidthassociatev2"
	bandwidthv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/bandwidthv2"
	eipv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/eipv1"
	flowlogv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/flowlogv1"
	peeringconnectionaccepterv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/peeringconnectionaccepterv2"
	peeringconnectionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/peeringconnectionv2"
	routetablev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/routetablev1"
	routev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/routev2"
	secgrouprulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/secgrouprulev3"
	secgroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/secgroupv3"
	subnetv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/subnetv1"
	vpcv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpc/vpcv1"
	approvalv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpcep/approvalv1"
	endpointv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpcep/endpointv1"
	servicev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpcep/servicev1"
	endpointgroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpnaas/endpointgroupv2"
	ikepolicyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpnaas/ikepolicyv2"
	ipsecpolicyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpnaas/ipsecpolicyv2"
	servicev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpnaas/servicev2"
	siteconnectionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/vpnaas/siteconnectionv2"
	dedicatedalarmmaskingrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedalarmmaskingrulev1"
	dedicatedanticrawlerrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedanticrawlerrulev1"
	dedicatedantileakagerulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedantileakagerulev1"
	dedicatedblacklistrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedblacklistrulev1"
	dedicatedccrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedccrulev1"
	dedicatedcertificatev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedcertificatev1"
	dedicateddatamaskingrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicateddatamaskingrulev1"
	dedicateddomainv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicateddomainv1"
	dedicatedgeoiprulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedgeoiprulev1"
	dedicatedinstancev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedinstancev1"
	dedicatedknownattacksourcerulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedknownattacksourcerulev1"
	dedicatedpolicyv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedpolicyv1"
	dedicatedpreciseprotectionrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedpreciseprotectionrulev1"
	dedicatedreferencetablev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedreferencetablev1"
	dedicatedwebtamperrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/namespaced/waf/dedicatedwebtamperrulev1"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aclpolicyassociatev2.Setup,
		aclpolicyv2.Setup,
		apipublishmentv2.Setup,
		apiv2.Setup,
		appcodev2.Setup,
		applicationauthorizationv2.Setup,
		applicationv2.Setup,
		certificatev2.Setup,
		customauthorizerv2.Setup,
		environmentv2.Setup,
		environmentvariablev2.Setup,
		gatewayfeaturev2.Setup,
		gatewayroutesv2.Setup,
		gatewayv2.Setup,
		groupv2.Setup,
		responsev2.Setup,
		signatureassociatev2.Setup,
		signaturev2.Setup,
		throttlingpolicyassociatev2.Setup,
		throttlingpolicyv2.Setup,
		vpcchannelv2.Setup,
		servicemeshv1.Setup,
		volumev2.Setup,
		policyv3.Setup,
		vaultv3.Setup,
		addonv3.Setup,
		clusterv3.Setup,
		nodeattachv3.Setup,
		nodepoolv3.Setup,
		nodev3.Setup,
		aclrulev1.Setup,
		addressgroupmemberv1.Setup,
		addressgroupv1.Setup,
		blacklistwhitelistrulev1.Setup,
		domainnamegroupv1.Setup,
		eipprotectionv1.Setup,
		firewallv1.Setup,
		ipsprotectionv1.Setup,
		servicegroupmemberv1.Setup,
		servicegroupv1.Setup,
		floatingipassociatev2.Setup,
		floatingipv2.Setup,
		instancev2.Setup,
		keypairv2.Setup,
		secgroupv2.Setup,
		servergroupv2.Setup,
		volumeattachv2.Setup,
		clusterrestartv1.Setup,
		clusterv1.Setup,
		configurationv1.Setup,
		snapshotconfigurationv1.Setup,
		eventnotificationv3.Setup,
		trackerv3.Setup,
		instancev2dcs.Setup,
		instancev1.Setup,
		schemav1.Setup,
		backupv3.Setup,
		instancev3.Setup,
		ltslogv3.Setup,
		hostv1.Setup,
		appv2.Setup,
		checkpointv2.Setup,
		dumptaskv2.Setup,
		streamv2.Setup,
		consumergroupv2.Setup,
		dedicatedinstancev2.Setup,
		instancev1dms.Setup,
		instancev2dms.Setup,
		reassignpartitionsv2.Setup,
		smartconnecttaskactionv2.Setup,
		smartconnecttaskv2.Setup,
		smartconnectv2.Setup,
		topicv1.Setup,
		userpermissionv1.Setup,
		userv2.Setup,
		ptrrecordv2.Setup,
		recordsetv2.Setup,
		zonev2.Setup,
		taskv3.Setup,
		clusterv1dws.Setup,
		instancev1ecs.Setup,
		vpnconnectionmonitorv5.Setup,
		vpnconnectionv5.Setup,
		vpncustomergatewayv5.Setup,
		vpngatewayv5.Setup,
		associationv3.Setup,
		flowlogv3.Setup,
		instancev3er.Setup,
		propagationv3.Setup,
		routetablev3.Setup,
		staticroutev3.Setup,
		vpcattachmentv3.Setup,
		volumev3.Setup,
		asyncinvokeconfigv2.Setup,
		dependencyversionv2.Setup,
		eventv2.Setup,
		functionv2.Setup,
		triggerv2.Setup,
		firewallgroupv2.Setup,
		policyv2.Setup,
		rulev2.Setup,
		instancev3gemini.Setup,
		hostgroupv5.Setup,
		hostprotectionv5.Setup,
		agencyv3.Setup,
		credentialv3.Setup,
		groupmembershipv3.Setup,
		groupv3.Setup,
		loginpolicyv3.Setup,
		mappingv3.Setup,
		passwordpolicyv3.Setup,
		projectv3.Setup,
		protectionpolicyv3.Setup,
		protocolv3.Setup,
		provider.Setup,
		providerv3.Setup,
		roleassignmentv3.Setup,
		rolev3.Setup,
		usergroupmembershipv3.Setup,
		userv3.Setup,
		imageaccessacceptv2.Setup,
		imageaccessv2.Setup,
		imagev2.Setup,
		dataimagev2.Setup,
		imagev2ims.Setup,
		grantv1.Setup,
		keymaterialv1.Setup,
		keyv1.Setup,
		certificatev3.Setup,
		ipgroupv3.Setup,
		listenerv3.Setup,
		loadbalancerv3.Setup,
		ltslogv3lb.Setup,
		memberv3.Setup,
		monitorv3.Setup,
		policyv3lb.Setup,
		poolv3.Setup,
		rulev3.Setup,
		securitypolicyv3.Setup,
		groupv2logtank.Setup,
		topicv2.Setup,
		transferv2.Setup,
		cceaccessv3.Setup,
		crossaccountaccessv2.Setup,
		groupv2lts.Setup,
		hostaccessv3.Setup,
		hostgroupv3.Setup,
		keywordsalarmrulev2.Setup,
		notificationtemplatev2.Setup,
		quicksearchcriteriav1.Setup,
		streamv2lts.Setup,
		transferv2lts.Setup,
		dnatrulev2.Setup,
		gatewayv2nat.Setup,
		snatrulev2.Setup,
		floatingipassociatev2networking.Setup,
		floatingipv2networking.Setup,
		networkv2.Setup,
		portsecgroupassociatev2.Setup,
		portv2.Setup,
		routerinterfacev2.Setup,
		routerroutev2.Setup,
		routerv2.Setup,
		secgrouprulev2.Setup,
		secgroupv2networking.Setup,
		subnetv2.Setup,
		vipassociatev2.Setup,
		vipv2.Setup,
		bucket.Setup,
		bucketacl.Setup,
		bucketinventory.Setup,
		bucketobject.Setup,
		bucketobjectacl.Setup,
		bucketpolicy.Setup,
		bucketreplication.Setup,
		natdnatrulev3.Setup,
		natgatewayv3.Setup,
		natsnatrulev3.Setup,
		nattransitipv3.Setup,
		providerconfig.Setup,
		backupv3rds.Setup,
		instancev3rds.Setup,
		maintenancev3.Setup,
		parametergroupv3.Setup,
		publicipassociatev3.Setup,
		readreplicav3.Setup,
		advancedqueryv1.Setup,
		policyassignmentevaluatev1.Setup,
		policyassignmentv1.Setup,
		resourcerecorderv1.Setup,
		softwareconfigv1.Setup,
		softwaredeploymentv1.Setup,
		stackv1.Setup,
		buckets3.Setup,
		bucketobjects3.Setup,
		bucketpolicys3.Setup,
		filesystemv2.Setup,
		shareaccessrulesv2.Setup,
		turbosharev1.Setup,
		subscriptionv2.Setup,
		topicattributev2.Setup,
		topicv2smn.Setup,
		domainv2.Setup,
		organizationpermissionsv2.Setup,
		organizationv2.Setup,
		repositoryv2.Setup,
		mysqlbackupv3.Setup,
		mysqlinstancev3.Setup,
		mysqlproxyv3.Setup,
		mysqlquotav3.Setup,
		mysqlsqlcontrolrulev3.Setup,
		resourcetagsv1.Setup,
		tagsv1.Setup,
		bandwidthassociatev2.Setup,
		bandwidthv2.Setup,
		eipv1.Setup,
		flowlogv1.Setup,
		peeringconnectionaccepterv2.Setup,
		peeringconnectionv2.Setup,
		routetablev1.Setup,
		routev2.Setup,
		secgrouprulev3.Setup,
		secgroupv3.Setup,
		subnetv1.Setup,
		vpcv1.Setup,
		approvalv1.Setup,
		endpointv1.Setup,
		servicev1.Setup,
		endpointgroupv2.Setup,
		ikepolicyv2.Setup,
		ipsecpolicyv2.Setup,
		servicev2.Setup,
		siteconnectionv2.Setup,
		dedicatedalarmmaskingrulev1.Setup,
		dedicatedanticrawlerrulev1.Setup,
		dedicatedantileakagerulev1.Setup,
		dedicatedblacklistrulev1.Setup,
		dedicatedccrulev1.Setup,
		dedicatedcertificatev1.Setup,
		dedicateddatamaskingrulev1.Setup,
		dedicateddomainv1.Setup,
		dedicatedgeoiprulev1.Setup,
		dedicatedinstancev1.Setup,
		dedicatedknownattacksourcerulev1.Setup,
		dedicatedpolicyv1.Setup,
		dedicatedpreciseprotectionrulev1.Setup,
		dedicatedreferencetablev1.Setup,
		dedicatedwebtamperrulev1.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aclpolicyassociatev2.SetupGated,
		aclpolicyv2.SetupGated,
		apipublishmentv2.SetupGated,
		apiv2.SetupGated,
		appcodev2.SetupGated,
		applicationauthorizationv2.SetupGated,
		applicationv2.SetupGated,
		certificatev2.SetupGated,
		customauthorizerv2.SetupGated,
		environmentv2.SetupGated,
		environmentvariablev2.SetupGated,
		gatewayfeaturev2.SetupGated,
		gatewayroutesv2.SetupGated,
		gatewayv2.SetupGated,
		groupv2.SetupGated,
		responsev2.SetupGated,
		signatureassociatev2.SetupGated,
		signaturev2.SetupGated,
		throttlingpolicyassociatev2.SetupGated,
		throttlingpolicyv2.SetupGated,
		vpcchannelv2.SetupGated,
		servicemeshv1.SetupGated,
		volumev2.SetupGated,
		policyv3.SetupGated,
		vaultv3.SetupGated,
		addonv3.SetupGated,
		clusterv3.SetupGated,
		nodeattachv3.SetupGated,
		nodepoolv3.SetupGated,
		nodev3.SetupGated,
		aclrulev1.SetupGated,
		addressgroupmemberv1.SetupGated,
		addressgroupv1.SetupGated,
		blacklistwhitelistrulev1.SetupGated,
		domainnamegroupv1.SetupGated,
		eipprotectionv1.SetupGated,
		firewallv1.SetupGated,
		ipsprotectionv1.SetupGated,
		servicegroupmemberv1.SetupGated,
		servicegroupv1.SetupGated,
		floatingipassociatev2.SetupGated,
		floatingipv2.SetupGated,
		instancev2.SetupGated,
		keypairv2.SetupGated,
		secgroupv2.SetupGated,
		servergroupv2.SetupGated,
		volumeattachv2.SetupGated,
		clusterrestartv1.SetupGated,
		clusterv1.SetupGated,
		configurationv1.SetupGated,
		snapshotconfigurationv1.SetupGated,
		eventnotificationv3.SetupGated,
		trackerv3.SetupGated,
		instancev2dcs.SetupGated,
		instancev1.SetupGated,
		schemav1.SetupGated,
		backupv3.SetupGated,
		instancev3.SetupGated,
		ltslogv3.SetupGated,
		hostv1.SetupGated,
		appv2.SetupGated,
		checkpointv2.SetupGated,
		dumptaskv2.SetupGated,
		streamv2.SetupGated,
		consumergroupv2.SetupGated,
		dedicatedinstancev2.SetupGated,
		instancev1dms.SetupGated,
		instancev2dms.SetupGated,
		reassignpartitionsv2.SetupGated,
		smartconnecttaskactionv2.SetupGated,
		smartconnecttaskv2.SetupGated,
		smartconnectv2.SetupGated,
		topicv1.SetupGated,
		userpermissionv1.SetupGated,
		userv2.SetupGated,
		ptrrecordv2.SetupGated,
		recordsetv2.SetupGated,
		zonev2.SetupGated,
		taskv3.SetupGated,
		clusterv1dws.SetupGated,
		instancev1ecs.SetupGated,
		vpnconnectionmonitorv5.SetupGated,
		vpnconnectionv5.SetupGated,
		vpncustomergatewayv5.SetupGated,
		vpngatewayv5.SetupGated,
		associationv3.SetupGated,
		flowlogv3.SetupGated,
		instancev3er.SetupGated,
		propagationv3.SetupGated,
		routetablev3.SetupGated,
		staticroutev3.SetupGated,
		vpcattachmentv3.SetupGated,
		volumev3.SetupGated,
		asyncinvokeconfigv2.SetupGated,
		dependencyversionv2.SetupGated,
		eventv2.SetupGated,
		functionv2.SetupGated,
		triggerv2.SetupGated,
		firewallgroupv2.SetupGated,
		policyv2.SetupGated,
		rulev2.SetupGated,
		instancev3gemini.SetupGated,
		hostgroupv5.SetupGated,
		hostprotectionv5.SetupGated,
		agencyv3.SetupGated,
		credentialv3.SetupGated,
		groupmembershipv3.SetupGated,
		groupv3.SetupGated,
		loginpolicyv3.SetupGated,
		mappingv3.SetupGated,
		passwordpolicyv3.SetupGated,
		projectv3.SetupGated,
		protectionpolicyv3.SetupGated,
		protocolv3.SetupGated,
		provider.SetupGated,
		providerv3.SetupGated,
		roleassignmentv3.SetupGated,
		rolev3.SetupGated,
		usergroupmembershipv3.SetupGated,
		userv3.SetupGated,
		imageaccessacceptv2.SetupGated,
		imageaccessv2.SetupGated,
		imagev2.SetupGated,
		dataimagev2.SetupGated,
		imagev2ims.SetupGated,
		grantv1.SetupGated,
		keymaterialv1.SetupGated,
		keyv1.SetupGated,
		certificatev3.SetupGated,
		ipgroupv3.SetupGated,
		listenerv3.SetupGated,
		loadbalancerv3.SetupGated,
		ltslogv3lb.SetupGated,
		memberv3.SetupGated,
		monitorv3.SetupGated,
		policyv3lb.SetupGated,
		poolv3.SetupGated,
		rulev3.SetupGated,
		securitypolicyv3.SetupGated,
		groupv2logtank.SetupGated,
		topicv2.SetupGated,
		transferv2.SetupGated,
		cceaccessv3.SetupGated,
		crossaccountaccessv2.SetupGated,
		groupv2lts.SetupGated,
		hostaccessv3.SetupGated,
		hostgroupv3.SetupGated,
		keywordsalarmrulev2.SetupGated,
		notificationtemplatev2.SetupGated,
		quicksearchcriteriav1.SetupGated,
		streamv2lts.SetupGated,
		transferv2lts.SetupGated,
		dnatrulev2.SetupGated,
		gatewayv2nat.SetupGated,
		snatrulev2.SetupGated,
		floatingipassociatev2networking.SetupGated,
		floatingipv2networking.SetupGated,
		networkv2.SetupGated,
		portsecgroupassociatev2.SetupGated,
		portv2.SetupGated,
		routerinterfacev2.SetupGated,
		routerroutev2.SetupGated,
		routerv2.SetupGated,
		secgrouprulev2.SetupGated,
		secgroupv2networking.SetupGated,
		subnetv2.SetupGated,
		vipassociatev2.SetupGated,
		vipv2.SetupGated,
		bucket.SetupGated,
		bucketacl.SetupGated,
		bucketinventory.SetupGated,
		bucketobject.SetupGated,
		bucketobjectacl.SetupGated,
		bucketpolicy.SetupGated,
		bucketreplication.SetupGated,
		natdnatrulev3.SetupGated,
		natgatewayv3.SetupGated,
		natsnatrulev3.SetupGated,
		nattransitipv3.SetupGated,
		providerconfig.SetupGated,
		backupv3rds.SetupGated,
		instancev3rds.SetupGated,
		maintenancev3.SetupGated,
		parametergroupv3.SetupGated,
		publicipassociatev3.SetupGated,
		readreplicav3.SetupGated,
		advancedqueryv1.SetupGated,
		policyassignmentevaluatev1.SetupGated,
		policyassignmentv1.SetupGated,
		resourcerecorderv1.SetupGated,
		softwareconfigv1.SetupGated,
		softwaredeploymentv1.SetupGated,
		stackv1.SetupGated,
		buckets3.SetupGated,
		bucketobjects3.SetupGated,
		bucketpolicys3.SetupGated,
		filesystemv2.SetupGated,
		shareaccessrulesv2.SetupGated,
		turbosharev1.SetupGated,
		subscriptionv2.SetupGated,
		topicattributev2.SetupGated,
		topicv2smn.SetupGated,
		domainv2.SetupGated,
		organizationpermissionsv2.SetupGated,
		organizationv2.SetupGated,
		repositoryv2.SetupGated,
		mysqlbackupv3.SetupGated,
		mysqlinstancev3.SetupGated,
		mysqlproxyv3.SetupGated,
		mysqlquotav3.SetupGated,
		mysqlsqlcontrolrulev3.SetupGated,
		resourcetagsv1.SetupGated,
		tagsv1.SetupGated,
		bandwidthassociatev2.SetupGated,
		bandwidthv2.SetupGated,
		eipv1.SetupGated,
		flowlogv1.SetupGated,
		peeringconnectionaccepterv2.SetupGated,
		peeringconnectionv2.SetupGated,
		routetablev1.SetupGated,
		routev2.SetupGated,
		secgrouprulev3.SetupGated,
		secgroupv3.SetupGated,
		subnetv1.SetupGated,
		vpcv1.SetupGated,
		approvalv1.SetupGated,
		endpointv1.SetupGated,
		servicev1.SetupGated,
		endpointgroupv2.SetupGated,
		ikepolicyv2.SetupGated,
		ipsecpolicyv2.SetupGated,
		servicev2.SetupGated,
		siteconnectionv2.SetupGated,
		dedicatedalarmmaskingrulev1.SetupGated,
		dedicatedanticrawlerrulev1.SetupGated,
		dedicatedantileakagerulev1.SetupGated,
		dedicatedblacklistrulev1.SetupGated,
		dedicatedccrulev1.SetupGated,
		dedicatedcertificatev1.SetupGated,
		dedicateddatamaskingrulev1.SetupGated,
		dedicateddomainv1.SetupGated,
		dedicatedgeoiprulev1.SetupGated,
		dedicatedinstancev1.SetupGated,
		dedicatedknownattacksourcerulev1.SetupGated,
		dedicatedpolicyv1.SetupGated,
		dedicatedpreciseprotectionrulev1.SetupGated,
		dedicatedreferencetablev1.SetupGated,
		dedicatedwebtamperrulev1.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
