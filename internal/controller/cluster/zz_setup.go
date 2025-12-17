// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	aclpolicyassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/aclpolicyassociatev2"
	aclpolicyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/aclpolicyv2"
	apipublishmentv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/apipublishmentv2"
	apiv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/apiv2"
	appcodev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/appcodev2"
	applicationauthorizationv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/applicationauthorizationv2"
	applicationv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/applicationv2"
	certificatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/certificatev2"
	customauthorizerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/customauthorizerv2"
	environmentv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/environmentv2"
	environmentvariablev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/environmentvariablev2"
	gatewayfeaturev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/gatewayfeaturev2"
	gatewayroutesv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/gatewayroutesv2"
	gatewayv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/gatewayv2"
	groupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/groupv2"
	responsev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/responsev2"
	signatureassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/signatureassociatev2"
	signaturev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/signaturev2"
	throttlingpolicyassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/throttlingpolicyassociatev2"
	throttlingpolicyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/throttlingpolicyv2"
	vpcchannelv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/apigw/vpcchannelv2"
	servicemeshv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/asm/servicemeshv1"
	volumev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/blockstorage/volumev2"
	policyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cbr/policyv3"
	vaultv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cbr/vaultv3"
	addonv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/addonv3"
	clusterv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/clusterv3"
	nodeattachv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/nodeattachv3"
	nodepoolv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/nodepoolv3"
	nodev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/nodev3"
	aclrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/aclrulev1"
	addressgroupmemberv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/addressgroupmemberv1"
	addressgroupv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/addressgroupv1"
	blacklistwhitelistrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/blacklistwhitelistrulev1"
	domainnamegroupv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/domainnamegroupv1"
	eipprotectionv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/eipprotectionv1"
	firewallv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/firewallv1"
	ipsprotectionv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/ipsprotectionv1"
	servicegroupmemberv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/servicegroupmemberv1"
	servicegroupv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cfw/servicegroupv1"
	floatingipassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/floatingipassociatev2"
	floatingipv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/floatingipv2"
	instancev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/instancev2"
	keypairv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/keypairv2"
	secgroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/secgroupv2"
	servergroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/servergroupv2"
	volumeattachv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/volumeattachv2"
	clusterrestartv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/css/clusterrestartv1"
	clusterv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/css/clusterv1"
	configurationv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/css/configurationv1"
	snapshotconfigurationv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/css/snapshotconfigurationv1"
	eventnotificationv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cts/eventnotificationv3"
	trackerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cts/trackerv3"
	instancev2dcs "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dcs/instancev2"
	instancev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/ddm/instancev1"
	schemav1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/ddm/schemav1"
	backupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dds/backupv3"
	instancev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dds/instancev3"
	ltslogv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dds/ltslogv3"
	hostv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/deh/hostv1"
	appv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dis/appv2"
	checkpointv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dis/checkpointv2"
	dumptaskv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dis/dumptaskv2"
	streamv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dis/streamv2"
	consumergroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/consumergroupv2"
	dedicatedinstancev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/dedicatedinstancev2"
	instancev1dms "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/instancev1"
	instancev2dms "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/instancev2"
	reassignpartitionsv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/reassignpartitionsv2"
	smartconnecttaskactionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/smartconnecttaskactionv2"
	smartconnecttaskv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/smartconnecttaskv2"
	smartconnectv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/smartconnectv2"
	topicv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/topicv1"
	userpermissionv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/userpermissionv1"
	userv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/userv2"
	ptrrecordv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dns/ptrrecordv2"
	recordsetv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dns/recordsetv2"
	zonev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dns/zonev2"
	taskv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/drs/taskv3"
	clusterv1dws "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dws/clusterv1"
	instancev1ecs "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/ecs/instancev1"
	vpnconnectionmonitorv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/enterprisevpn/vpnconnectionmonitorv5"
	vpnconnectionv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/enterprisevpn/vpnconnectionv5"
	vpncustomergatewayv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/enterprisevpn/vpncustomergatewayv5"
	vpngatewayv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/enterprisevpn/vpngatewayv5"
	associationv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/er/associationv3"
	flowlogv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/er/flowlogv3"
	instancev3er "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/er/instancev3"
	propagationv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/er/propagationv3"
	routetablev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/er/routetablev3"
	staticroutev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/er/staticroutev3"
	vpcattachmentv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/er/vpcattachmentv3"
	volumev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/evs/volumev3"
	asyncinvokeconfigv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/asyncinvokeconfigv2"
	dependencyversionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/dependencyversionv2"
	eventv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/eventv2"
	functionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/functionv2"
	triggerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/triggerv2"
	firewallgroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fw/firewallgroupv2"
	policyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fw/policyv2"
	rulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fw/rulev2"
	instancev3gemini "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/gemini/instancev3"
	hostgroupv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/hss/hostgroupv5"
	hostprotectionv5 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/hss/hostprotectionv5"
	agencyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/agencyv3"
	credentialv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/credentialv3"
	groupmembershipv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/groupmembershipv3"
	groupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/groupv3"
	loginpolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/loginpolicyv3"
	mappingv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/mappingv3"
	passwordpolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/passwordpolicyv3"
	projectv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/projectv3"
	protectionpolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/protectionpolicyv3"
	protocolv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/protocolv3"
	provider "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/provider"
	providerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/providerv3"
	roleassignmentv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/roleassignmentv3"
	rolev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/rolev3"
	usergroupmembershipv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/usergroupmembershipv3"
	userv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/identity/userv3"
	imageaccessacceptv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/images/imageaccessacceptv2"
	imageaccessv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/images/imageaccessv2"
	imagev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/images/imagev2"
	dataimagev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/ims/dataimagev2"
	imagev2ims "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/ims/imagev2"
	grantv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/kms/grantv1"
	keymaterialv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/kms/keymaterialv1"
	keyv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/kms/keyv1"
	certificatev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/certificatev3"
	ipgroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/ipgroupv3"
	listenerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/listenerv3"
	loadbalancerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/loadbalancerv3"
	ltslogv3lb "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/ltslogv3"
	memberv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/memberv3"
	monitorv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/monitorv3"
	policyv3lb "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/policyv3"
	poolv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/poolv3"
	rulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/rulev3"
	securitypolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/securitypolicyv3"
	groupv2logtank "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/logtank/groupv2"
	topicv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/logtank/topicv2"
	transferv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/logtank/transferv2"
	cceaccessv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/cceaccessv3"
	crossaccountaccessv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/crossaccountaccessv2"
	groupv2lts "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/groupv2"
	hostaccessv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/hostaccessv3"
	hostgroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/hostgroupv3"
	keywordsalarmrulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/keywordsalarmrulev2"
	notificationtemplatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/notificationtemplatev2"
	quicksearchcriteriav1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/quicksearchcriteriav1"
	streamv2lts "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/streamv2"
	transferv2lts "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lts/transferv2"
	dnatrulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/nat/dnatrulev2"
	gatewayv2nat "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/nat/gatewayv2"
	snatrulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/nat/snatrulev2"
	floatingipassociatev2networking "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/floatingipassociatev2"
	floatingipv2networking "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/floatingipv2"
	networkv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/networkv2"
	portsecgroupassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/portsecgroupassociatev2"
	portv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/portv2"
	routerinterfacev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/routerinterfacev2"
	routerroutev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/routerroutev2"
	routerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/routerv2"
	secgrouprulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/secgrouprulev2"
	secgroupv2networking "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/secgroupv2"
	subnetv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/subnetv2"
	vipassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/vipassociatev2"
	vipv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/networking/vipv2"
	bucket "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucket"
	bucketacl "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketacl"
	bucketinventory "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketinventory"
	bucketobject "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketobject"
	bucketobjectacl "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketobjectacl"
	bucketpolicy "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketpolicy"
	bucketreplication "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketreplication"
	natdnatrulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/privatenat/natdnatrulev3"
	natgatewayv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/privatenat/natgatewayv3"
	natsnatrulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/privatenat/natsnatrulev3"
	nattransitipv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/privatenat/nattransitipv3"
	providerconfig "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/providerconfig"
	backupv3rds "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/backupv3"
	instancev3rds "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/instancev3"
	maintenancev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/maintenancev3"
	parametergroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/parametergroupv3"
	publicipassociatev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/publicipassociatev3"
	readreplicav3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/readreplicav3"
	advancedqueryv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rms/advancedqueryv1"
	policyassignmentevaluatev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rms/policyassignmentevaluatev1"
	policyassignmentv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rms/policyassignmentv1"
	resourcerecorderv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rms/resourcerecorderv1"
	softwareconfigv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rts/softwareconfigv1"
	softwaredeploymentv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rts/softwaredeploymentv1"
	stackv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rts/stackv1"
	buckets3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/s3/bucket"
	bucketobjects3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/s3/bucketobject"
	bucketpolicys3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/s3/bucketpolicy"
	filesystemv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/sfs/filesystemv2"
	shareaccessrulesv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/sfs/shareaccessrulesv2"
	turbosharev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/sfs/turbosharev1"
	subscriptionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/smn/subscriptionv2"
	topicattributev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/smn/topicattributev2"
	topicv2smn "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/smn/topicv2"
	domainv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/swr/domainv2"
	organizationpermissionsv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/swr/organizationpermissionsv2"
	organizationv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/swr/organizationv2"
	repositoryv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/swr/repositoryv2"
	mysqlbackupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/taurusdb/mysqlbackupv3"
	mysqlinstancev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/taurusdb/mysqlinstancev3"
	mysqlproxyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/taurusdb/mysqlproxyv3"
	mysqlquotav3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/taurusdb/mysqlquotav3"
	mysqlsqlcontrolrulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/taurusdb/mysqlsqlcontrolrulev3"
	resourcetagsv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/tms/resourcetagsv1"
	tagsv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/tms/tagsv1"
	bandwidthassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/bandwidthassociatev2"
	bandwidthv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/bandwidthv2"
	eipv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/eipv1"
	flowlogv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/flowlogv1"
	peeringconnectionaccepterv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/peeringconnectionaccepterv2"
	peeringconnectionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/peeringconnectionv2"
	routetablev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/routetablev1"
	routev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/routev2"
	secgrouprulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/secgrouprulev3"
	secgroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/secgroupv3"
	subnetv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/subnetv1"
	vpcv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/vpcv1"
	approvalv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpcep/approvalv1"
	endpointv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpcep/endpointv1"
	servicev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpcep/servicev1"
	endpointgroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpnaas/endpointgroupv2"
	ikepolicyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpnaas/ikepolicyv2"
	ipsecpolicyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpnaas/ipsecpolicyv2"
	servicev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpnaas/servicev2"
	siteconnectionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpnaas/siteconnectionv2"
	dedicatedalarmmaskingrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedalarmmaskingrulev1"
	dedicatedanticrawlerrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedanticrawlerrulev1"
	dedicatedantileakagerulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedantileakagerulev1"
	dedicatedblacklistrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedblacklistrulev1"
	dedicatedccrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedccrulev1"
	dedicatedcertificatev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedcertificatev1"
	dedicateddatamaskingrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicateddatamaskingrulev1"
	dedicateddomainv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicateddomainv1"
	dedicatedgeoiprulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedgeoiprulev1"
	dedicatedinstancev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedinstancev1"
	dedicatedknownattacksourcerulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedknownattacksourcerulev1"
	dedicatedpolicyv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedpolicyv1"
	dedicatedpreciseprotectionrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedpreciseprotectionrulev1"
	dedicatedreferencetablev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedreferencetablev1"
	dedicatedwebtamperrulev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/waf/dedicatedwebtamperrulev1"
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
