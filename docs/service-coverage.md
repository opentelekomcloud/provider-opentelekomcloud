# T Cloud Public Service Coverage

We are continuously onboarding more services. If you have a use case for a service that is not yet available, you can request it's implementation by creating a GitHub issue with a brief description of your use case.

### State of services

🟢 Verified - Service has been configured and examples are e2e tested.  

🔵 Available - CRDs are available and resources can be used, with limited configuration and validation. These services might change during implementation.

🔴 Unavailable - No CRDs available yet.


| SERVICE NAME | SHORT NAME | API GROUP(s) | STATE |
|---|---|---|---|
| Elastic Cloud Server | [ecs](../examples/namespaced/ecs) | ecs.opentelekomcloud.m.crossplane.io<br>compute.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Identity and Access Management | [iam](../examples/namespaced/identity) | identity.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Log Tank Service | [lts](../examples/namespaced/lts) | lts.opentelekomcloud.m.crossplane.io<br>logtank.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Cloud Container Engine | [cce](../examples/namespaced/cce) | cce.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Software Repository for Container | [swr](../examples/namespaced/swr) | swr.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Key Management Service | [kms](../examples/namespaced/kms) | kms.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Elastic Volume Service | [evs](../examples/namespaced/evs) | evs.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Object Storage Service | [obs](../examples/namespaced/obs) | obs.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Scalable File Service & SFS Turbo | [sfs](../examples/namespaced/sfs) | sfs.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Virtual Private Cloud | [vpc](../examples/namespaced/vpc) | vpc.opentelekomcloud.m.crossplane.io<br>networking.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Elastic Load Balancing | [elb](../examples/namespaced/lb) | lb.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Domain Name Service | [dns](../examples/namespaced/dns) | dns.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| NAT Gateway | [nat](../examples/namespaced/nat) | nat.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Elastic IP | [eip](../examples/namespaced/vpc) | vpc.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Relational Database Service | [rds](../examples/namespaced/rds) | rds.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Distributed Cache Service | [dcs](../examples/namespaced/dcs) | dcs.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Document Database Service | [dds](../examples/namespaced/dds) | dds.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Distributed Message Service (for Kafka) | [dms](../examples/namespaced/dms) | dms.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Web Application Firewall | [waf](../examples/namespaced/waf) | waf.opentelekomcloud.m.crossplane.io | 🟢 Verified |
| Block Storage Service | blockstorage | blockstorage.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Dedicated Load Balancer | dlb | lb.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Image Management Service | ims | ims.opentelekomcloud.m.crossplane.io<br>images.opentelekomcloud.m.crossplane.io | 🔵 Available |
| FunctionGraph | fgs | fgs.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Dedicated Host | deh | deh.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Resource Formation Service | rfs | rts.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Cloud Trace Service | cts | cts.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Tag Management Service | tms | tms.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Config | rms | rms.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Application Service Mesh | asm | asm.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Host Security Service | hss | hss.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Web Application Firewall (Classic) | waf | waf.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Cloud Firewall | cfw | cfw.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Data Warehouse Service | dws | dws.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Cloud Search Service | css | css.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Cloud Backup and Recovery | cbr | cbr.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Distributed Message Service (for RocketMQ) | dms | dms.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Distributed Message Service | dms | dms.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Simple Message Notification | smn | smn.opentelekomcloud.m.crossplane.io | 🔵 Available |
| API Gateway | apigw | apigw.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Enterprise Router | er | er.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Virtual Private Network | vpn | vpnaas.opentelekomcloud.m.crossplane.io<br>enterprisevpn.opentelekomcloud.m.crossplane.io | 🔵 Available |
| VPC Endpoint | vpcep | vpcep.opentelekomcloud.m.crossplane.io | 🔵 Available |
| TaurusDB | taurusdb | taurusdb.opentelekomcloud.m.crossplane.io | 🔵 Available |
| GeminiDB | gemini | gemini.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Distributed Database Middleware | ddm | ddm.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Data Replication Service | drs | drs.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Private NAT | privatenat | privatenat.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Virtual Private Cloud Firewall | fw | fw.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Object Storage Service (S3) | s3 | s3.opentelekomcloud.m.crossplane.io | 🔵 Available |
| Auto Scaling | as | - | 🔴 Unavailable |
| Enterprise Project Service | eps | - | 🔴 Unavailable |
| Cloud Eye | ces | - | 🔴 Unavailable |
| Cloud Container Instance | cci | - | 🔴 Unavailable |
| Anti-DDoS | antiddos | - | 🔴 Unavailable |
| Database Security Service | dbss | - | 🔴 Unavailable |
| DataArts Studio | dataarts | - | 🔴 Unavailable |
| ModelArts | modelarts | - | 🔴 Unavailable |
| MapReduce Service | mrs | - | 🔴 Unavailable |
| Data Lake Insight | dli | - | 🔴 Unavailable |
| Optical Character Recognition | ocr | - | 🔴 Unavailable |
| Cloud Server Backup Service | csbs | - | 🔴 Unavailable |
| Storage Disaster Recovery Service | sdrs | - | 🔴 Unavailable |
| Volume Backup Service | vbs | - | 🔴 Unavailable |
| Application Performance Management | apm | - | 🔴 Unavailable |
| Application Operations Management | aom | - | 🔴 Unavailable |
| Direct Connect | dc | - | 🔴 Unavailable |
| Enterprise Switch | esw | - | 🔴 Unavailable |
| Cloud Connect | cc | - | 🔴 Unavailable |
| General Purpose File System | gpfs | - | 🔴 Unavailable |
| GaussDB (for MySQL) | gaussdb | - | 🔴 Unavailable |
