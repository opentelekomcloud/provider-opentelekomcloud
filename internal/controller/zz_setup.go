/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	volumev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/blockstorage/volumev2"
	addonv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/addonv3"
	clusterv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/clusterv3"
	nodepoolv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/nodepoolv3"
	nodev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/cce/nodev3"
	floatingipassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/floatingipassociatev2"
	floatingipv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/floatingipv2"
	instancev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/instancev2"
	keypairv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/keypairv2"
	secgroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/secgroupv2"
	servergroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/servergroupv2"
	volumeattachv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/compute/volumeattachv2"
	instancev2dcs "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dcs/instancev2"
	instancev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dds/instancev3"
	hostv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/deh/hostv1"
	appv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dis/appv2"
	checkpointv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dis/checkpointv2"
	dumptaskv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dis/dumptaskv2"
	streamv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dis/streamv2"
	instancev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/instancev1"
	instancev2dms "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/instancev2"
	topicv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/topicv1"
	userpermissionv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/userpermissionv1"
	userv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dms/userv2"
	ptrrecordv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dns/ptrrecordv2"
	recordsetv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dns/recordsetv2"
	zonev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/dns/zonev2"
	instancev1ecs "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/ecs/instancev1"
	volumev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/evs/volumev3"
	asyncinvokeconfigv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/asyncinvokeconfigv2"
	eventv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/eventv2"
	functionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/functionv2"
	triggerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fgs/triggerv2"
	firewallgroupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fw/firewallgroupv2"
	policyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fw/policyv2"
	rulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/fw/rulev2"
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
	keyv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/kms/keyv1"
	certificatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/certificatev2"
	certificatev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/certificatev3"
	ipgroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/ipgroupv3"
	l7policyv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/l7policyv2"
	l7rulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/l7rulev2"
	listenerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/listenerv2"
	listenerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/listenerv3"
	loadbalancerv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/loadbalancerv2"
	loadbalancerv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/loadbalancerv3"
	memberv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/memberv2"
	memberv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/memberv3"
	monitorv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/monitorv2"
	monitorv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/monitorv3"
	policyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/policyv3"
	poolv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/poolv2"
	poolv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/poolv3"
	rulev3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/rulev3"
	securitypolicyv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/securitypolicyv3"
	whitelistv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/lb/whitelistv2"
	groupv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/logtank/groupv2"
	topicv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/logtank/topicv2"
	transferv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/logtank/transferv2"
	dnatrulev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/nat/dnatrulev2"
	gatewayv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/nat/gatewayv2"
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
	bucketinventory "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketinventory"
	bucketobject "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketobject"
	bucketpolicy "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketpolicy"
	bucketreplication "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/obs/bucketreplication"
	providerconfig "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/providerconfig"
	backupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/backupv3"
	instancev3rds "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/instancev3"
	parametergroupv3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/parametergroupv3"
	readreplicav3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/rds/readreplicav3"
	buckets3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/s3/bucket"
	bucketobjects3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/s3/bucketobject"
	bucketpolicys3 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/s3/bucketpolicy"
	filesystemv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/sfs/filesystemv2"
	shareaccessrulesv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/sfs/shareaccessrulesv2"
	turbosharev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/sfs/turbosharev1"
	subscriptionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/smn/subscriptionv2"
	topicattributev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/smn/topicattributev2"
	topicv2smn "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/smn/topicv2"
	bandwidthassociatev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/bandwidthassociatev2"
	bandwidthv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/bandwidthv2"
	eipv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/eipv1"
	flowlogv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/flowlogv1"
	peeringconnectionaccepterv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/peeringconnectionaccepterv2"
	peeringconnectionv2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/peeringconnectionv2"
	routetablev1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/routetablev1"
	routev2 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/routev2"
	subnetv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/subnetv1"
	vpcv1 "github.com/opentelekomcloud/provider-opentelekomcloud/internal/controller/vpc/vpcv1"
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
		volumev2.Setup,
		addonv3.Setup,
		clusterv3.Setup,
		nodepoolv3.Setup,
		nodev3.Setup,
		floatingipassociatev2.Setup,
		floatingipv2.Setup,
		instancev2.Setup,
		keypairv2.Setup,
		secgroupv2.Setup,
		servergroupv2.Setup,
		volumeattachv2.Setup,
		instancev2dcs.Setup,
		instancev3.Setup,
		hostv1.Setup,
		appv2.Setup,
		checkpointv2.Setup,
		dumptaskv2.Setup,
		streamv2.Setup,
		instancev1.Setup,
		instancev2dms.Setup,
		topicv1.Setup,
		userpermissionv1.Setup,
		userv2.Setup,
		ptrrecordv2.Setup,
		recordsetv2.Setup,
		zonev2.Setup,
		instancev1ecs.Setup,
		volumev3.Setup,
		asyncinvokeconfigv2.Setup,
		eventv2.Setup,
		functionv2.Setup,
		triggerv2.Setup,
		firewallgroupv2.Setup,
		policyv2.Setup,
		rulev2.Setup,
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
		keyv1.Setup,
		certificatev2.Setup,
		certificatev3.Setup,
		ipgroupv3.Setup,
		l7policyv2.Setup,
		l7rulev2.Setup,
		listenerv2.Setup,
		listenerv3.Setup,
		loadbalancerv2.Setup,
		loadbalancerv3.Setup,
		memberv2.Setup,
		memberv3.Setup,
		monitorv2.Setup,
		monitorv3.Setup,
		policyv3.Setup,
		poolv2.Setup,
		poolv3.Setup,
		rulev3.Setup,
		securitypolicyv3.Setup,
		whitelistv2.Setup,
		groupv2.Setup,
		topicv2.Setup,
		transferv2.Setup,
		dnatrulev2.Setup,
		gatewayv2.Setup,
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
		bucketinventory.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		bucketreplication.Setup,
		providerconfig.Setup,
		backupv3.Setup,
		instancev3rds.Setup,
		parametergroupv3.Setup,
		readreplicav3.Setup,
		buckets3.Setup,
		bucketobjects3.Setup,
		bucketpolicys3.Setup,
		filesystemv2.Setup,
		shareaccessrulesv2.Setup,
		turbosharev1.Setup,
		subscriptionv2.Setup,
		topicattributev2.Setup,
		topicv2smn.Setup,
		bandwidthassociatev2.Setup,
		bandwidthv2.Setup,
		eipv1.Setup,
		flowlogv1.Setup,
		peeringconnectionaccepterv2.Setup,
		peeringconnectionv2.Setup,
		routetablev1.Setup,
		routev2.Setup,
		subnetv1.Setup,
		vpcv1.Setup,
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
