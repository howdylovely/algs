// package upkg

// require (
// 	cloud.google.com/go v0.57.0 // indirect
// 	git.jd.com/jcloud-api-gateway/jdcloud-apigateway-signature-go v0.0.1
// 	git.jd.com/mypass/helm-utils v0.0.10
// 	git.jd.com/mypass/map2properties v0.0.5
// 	git.jd.com/mypass/mypass-auth-middleware v0.1.1
// 	git.jd.com/mypass/mypass-client v0.0.1
// 	github.com/Masterminds/semver/v3 v3.1.0
// 	github.com/avast/retry-go v3.0.0+incompatible
// 	github.com/caarlos0/env v3.5.0+incompatible
// 	github.com/chartmuseum/auth v0.4.1
// 	github.com/dgrijalva/jwt-go v3.2.0+incompatible
// 	github.com/docker/distribution v2.7.1+incompatible
// 	github.com/docker/go-connections v0.4.0
// 	github.com/docker/go-metrics v0.0.1 // indirect
// 	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7
// 	github.com/fatih/structs v1.1.0
// 	github.com/gbrlsnchs/jwt/v3 v3.0.0-rc.2
// 	github.com/ghodss/yaml v1.0.0
// 	github.com/gin-contrib/zap v0.0.1
// 	github.com/gin-gonic/gin v1.6.3
// 	github.com/go-logr/logr v0.1.0
// 	github.com/go-logr/zapr v0.1.0
// 	github.com/go-oauth2/oauth2/v4 v4.1.2
// 	github.com/go-playground/validator/v10 v10.2.0
// 	github.com/go-sql-driver/mysql v1.5.0
// 	github.com/gofrs/uuid v3.3.0+incompatible
// 	github.com/google/uuid v1.1.1
// 	github.com/gorilla/mux v1.7.4
// 	github.com/gorilla/websocket v1.4.2
// 	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
// 	github.com/imdario/mergo v0.3.9
// 	github.com/influxdata/influxdb-client-go v1.4.0
// 	github.com/jinzhu/gorm v1.9.15
// 	github.com/json-iterator/go v1.1.10
// 	github.com/labstack/gommon v0.3.0
// 	github.com/olivere/elastic v6.2.35+incompatible
// 	github.com/olivere/elastic/v7 v7.0.19
// 	github.com/opencontainers/go-digest v1.0.0
// 	github.com/pkg/errors v0.9.1
// 	github.com/rancher/norman v0.0.0-20200930000340-693d65aaffe3
// 	github.com/rancher/remotedialer v0.0.1
// 	github.com/sirupsen/logrus v1.6.0
// 	github.com/spf13/pflag v1.0.5
// 	github.com/spf13/viper v1.7.0
// 	github.com/thedevsaddam/gojsonq v2.3.0+incompatible
// 	github.com/yvasiyarov/go-metrics v0.0.0-20150112132944-c25f46c4b940 // indirect
// 	github.com/yvasiyarov/gorelic v0.0.7 // indirect
// 	github.com/yvasiyarov/newrelic_platform_go v0.0.0-20160601141957-9c099fbc30e9 // indirect
// 	go.uber.org/zap v1.16.0
// 	golang.org/x/net v0.0.0-20201006153459-a7d1128ccaa0 // indirect
// 	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
// 	golang.org/x/tools v0.0.0-20200512131952-2bc93b1c0c88 // indirect
// 	google.golang.org/genproto v0.0.0-20200511104702-f5ebc3bea380 // indirect
// 	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
// 	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
// 	gopkg.in/natefinch/lumberjack.v2 v2.0.0
// 	gopkg.in/yaml.v2 v2.3.0
// 	helm.sh/helm/v3 v3.3.1
// 	k8s.io/api v0.18.8
// 	k8s.io/apiextensions-apiserver v0.18.8
// 	k8s.io/apimachinery v0.18.8
// 	k8s.io/cli-runtime v0.18.8
// 	k8s.io/client-go v0.18.8
// 	k8s.io/klog v1.0.0
// 	k8s.io/kubectl v0.18.8
// 	sigs.k8s.io/controller-runtime v0.6.3
// 	sigs.k8s.io/kubefed v0.4.1
// 	sigs.k8s.io/yaml v1.2.0
// )

// replace (
// 	git.jd.com/jcloud-api-gateway/jdcloud-apigateway-signature-go => ../jdcloud-apigateway-signature-go
// 	git.jd.com/mypass/helm-utils => ../helm-utils
// 	git.jd.com/mypass/map2properties => ../map2properties
// 	git.jd.com/mypass/mypass-auth-middleware => ../mypass-auth-middleware
// 	git.jd.com/mypass/mypass-client => ../mypass-client
// 	//github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
// 	github.com/opencontainers/runc => github.com/opencontainers/runc v1.0.0-rc10
// 	github.com/rancher/remotedialer => ../remotedialer
// 	golang.org/x/sys => golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6
// 	google.golang.org/grpc => google.golang.org/grpc v1.26.0
// )

// go 1.14

// require (
// 	cloud.google.com/go v0.48.0 // indirect
// 	github.com/aws/aws-sdk-go v1.17.7
// 	github.com/go-logr/logr v0.1.0
// 	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
// 	github.com/golang/protobuf v1.3.3 // indirect
// 	github.com/google/uuid v1.1.1
// 	github.com/hashicorp/consul/api v1.4.0
// 	github.com/hashicorp/go-uuid v1.0.1
// 	github.com/onsi/ginkgo v1.12.0 // indirect
// 	github.com/onsi/gomega v1.9.0 // indirect
// 	github.com/operator-framework/operator-sdk v0.15.1-0.20200224212241-e8424d5f3def
// 	github.com/pkg/errors v0.9.1 // indirect
// 	github.com/prometheus/client_golang v1.4.1
// 	github.com/prometheus/procfs v0.0.10 // indirect
// 	github.com/spf13/cobra v0.0.6 // indirect
// 	github.com/spf13/pflag v1.0.5
// 	github.com/stretchr/testify v1.5.1 // indirect
// 	go.uber.org/multierr v1.5.0 // indirect
// 	go.uber.org/zap v1.14.0 // indirect
// 	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d // indirect
// 	golang.org/x/lint v0.0.0-20200130185559-910be7a94367 // indirect
// 	golang.org/x/mod v0.2.0 // indirect
// 	golang.org/x/net v0.0.0-20200225223329-5d076fcf07a8 // indirect
// 	golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae // indirect
// 	golang.org/x/tools v0.0.0-20200225230052-807dcd883420 // indirect
// 	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
// 	k8s.io/api v0.0.0
// 	k8s.io/apimachinery v0.0.0
// 	k8s.io/client-go v12.0.0+incompatible
// 	sigs.k8s.io/controller-runtime v0.4.0
// 	sigs.k8s.io/yaml v1.2.0 // indirect
// )

// // Pinned to kubernetes-1.16.2
// replace (
// 	k8s.io/api => k8s.io/api v0.0.0-20191016110408-35e52d86657a
// 	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191016113550-5357c4baaf65
// 	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
// 	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191016112112-5190913f932d
// 	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191016114015-74ad18325ed5
// 	k8s.io/client-go => k8s.io/client-go v0.0.0-20191016111102-bec269661e48
// 	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191016115326-20453efc2458
// 	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191016115129-c07a134afb42
// 	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20191004115455-8e001e5d1894
// 	k8s.io/component-base => k8s.io/component-base v0.0.0-20191016111319-039242c015a9
// 	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190828162817-608eb1dad4ac
// 	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191016115521-756ffa5af0bd
// 	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191016112429-9587704a8ad4
// 	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191016114939-2b2b218dc1df
// 	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191016114407-2e83b6f20229
// 	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191016114748-65049c67a58b
// 	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191016120415-2ed914427d51
// 	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191016114556-7841ed97f1b2
// 	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191016115753-cf0698c3a16b
// 	k8s.io/metrics => k8s.io/metrics v0.0.0-20191016113814-3b1a734dba6e
// 	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191016112829-06bb3c9d77c9

// )

// replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309 // Required by Helm

// replace github.com/openshift/api => github.com/openshift/api v0.0.0-20190924102528-32369d4db2ad // Required until https://github.com/operator-framework/operator-lifecycle-manager/pull/1241 is resolved

// module git.jd.com/mypass/mypass-monitor-operator

// go 1.13

// require (
// 	git.jd.com/mypass/mypass-auth-middleware v0.0.0-00010101000000-000000000000
// 	github.com/cloudflare/cfssl v0.0.0-20180726162950-56268a613adf
// 	github.com/coreos/prometheus-operator v0.34.0
// 	github.com/emicklei/go-restful v2.11.1+incompatible
// 	github.com/fatih/structs v1.1.0
// 	github.com/gin-gonic/gin v1.6.3
// 	github.com/olivere/elastic v6.2.30+incompatible
// 	github.com/operator-framework/operator-sdk v0.15.2
// 	github.com/pkg/errors v0.8.1
// 	github.com/prometheus/alertmanager v0.20.0
// 	github.com/prometheus/client_golang v1.5.1 // indirect
// 	github.com/prometheus/common v0.9.1
// 	github.com/prometheus/prometheus v2.3.2+incompatible
// 	github.com/satori/go.uuid v1.2.0
// 	github.com/spf13/pflag v1.0.5
// 	github.com/stretchr/testify v1.4.0
// 	github.com/tidwall/gjson v1.6.0
// 	github.com/tidwall/pretty v1.0.1 // indirect
// 	golang.org/x/sys v0.0.0-20200413165638-669c56c373c4 // indirect
// 	golang.org/x/tools v0.0.0-20191127201027-ecd32218bd7f // indirect
// 	gopkg.in/yaml.v2 v2.2.8
// 	k8s.io/api v0.0.0
// 	k8s.io/apimachinery v0.0.0
// 	k8s.io/client-go v12.0.0+incompatible
// 	k8s.io/klog v1.0.0
// 	sigs.k8s.io/controller-runtime v0.4.0
// 	sigs.k8s.io/yaml v1.1.0
// )

// // Pinned to kubernetes-1.16.2
// replace (
// 	k8s.io/api => k8s.io/api v0.0.0-20191016110408-35e52d86657a
// 	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191016113550-5357c4baaf65
// 	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
// 	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191016112112-5190913f932d
// 	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191016114015-74ad18325ed5
// 	k8s.io/client-go => k8s.io/client-go v0.0.0-20191016111102-bec269661e48
// 	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191016115326-20453efc2458
// 	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191016115129-c07a134afb42
// 	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20191004115455-8e001e5d1894
// 	k8s.io/component-base => k8s.io/component-base v0.0.0-20191016111319-039242c015a9
// 	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190828162817-608eb1dad4ac
// 	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191016115521-756ffa5af0bd
// 	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191016112429-9587704a8ad4
// 	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191016114939-2b2b218dc1df
// 	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191016114407-2e83b6f20229
// 	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191016114748-65049c67a58b
// 	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191016120415-2ed914427d51
// 	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191016114556-7841ed97f1b2
// 	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191016115753-cf0698c3a16b
// 	k8s.io/metrics => k8s.io/metrics v0.0.0-20191016113814-3b1a734dba6e
// 	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191016112829-06bb3c9d77c9
// )

// replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309 // Required by Helm

// replace github.com/openshift/api => github.com/openshift/api v0.0.0-20190924102528-32369d4db2ad // Required until https://github.com/operator-framework/operator-lifecycle-manager/pull/1241 is resolved

// replace github.com/prometheus/tsdb => github.com/prometheus/tsdb v0.3.1

// // auth middleware
// replace git.jd.com/mypass/mypass-auth-middleware => ../mypass-auth-middleware

// replace git.jd.com/jcloud-api-gateway/jdcloud-apigateway-signature-go => ../../jcloud-api-gateway/jdcloud-apigateway-signature-go

// module git.jd.com/mypass/mypass-auth-middleware

// go 1.14

// require (
// 	git.jd.com/jcloud-api-gateway/jdcloud-apigateway-signature-go v0.0.1
// 	github.com/caarlos0/env v3.5.0+incompatible
// 	github.com/gbrlsnchs/jwt/v3 v3.0.0-rc.2
// 	github.com/gin-gonic/gin v1.6.3
// 	github.com/gofrs/uuid v3.3.0+incompatible
// 	gopkg.in/yaml.v2 v2.2.8
// )

// module git.jd.com/mypass/mypass-client

// go 1.13

// require (
// 	github.com/fsnotify/fsnotify v1.4.9 // indirect
// 	github.com/go-openapi/spec v0.20.3 // indirect
// 	github.com/go-openapi/swag v0.19.15 // indirect
// 	github.com/golang/protobuf v1.5.2 // indirect
// 	github.com/hashicorp/golang-lru v0.5.4 // indirect
// 	github.com/imdario/mergo v0.3.9 // indirect
// 	github.com/json-iterator/go v1.1.10 // indirect
// 	github.com/kr/pretty v0.2.1 // indirect
// 	github.com/mailru/easyjson v0.7.7 // indirect
// 	github.com/onsi/gomega v1.10.1 // indirect
// 	github.com/prometheus/procfs v0.0.11 // indirect
// 	github.com/stretchr/testify v1.7.0 // indirect
// 	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
// 	golang.org/x/mod v0.4.2 // indirect
// 	golang.org/x/net v0.0.0-20210331060903-cb1fcc7394e5 // indirect
// 	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558 // indirect
// 	golang.org/x/tools v0.1.0 // indirect
// 	google.golang.org/appengine v1.6.7 // indirect
// 	gopkg.in/yaml.v2 v2.4.0 // indirect
// 	helm.sh/helm/v3 v3.3.1
// 	k8s.io/api v0.18.8
// 	k8s.io/apimachinery v0.18.8
// 	k8s.io/client-go v0.18.8
// 	k8s.io/code-generator v0.18.8
// 	k8s.io/klog v1.0.0
// 	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29 // indirect
// 	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451 // indirect
// 	sigs.k8s.io/controller-runtime v0.6.0
// )

// replace (
// 	github.com/go-openapi/analysis => github.com/go-openapi/analysis v0.19.10
// 	github.com/go-openapi/errors => github.com/go-openapi/errors v0.19.4
// 	github.com/go-openapi/jsonpointer => github.com/go-openapi/jsonpointer v0.19.3
// 	github.com/go-openapi/jsonreference => github.com/go-openapi/jsonreference v0.19.3
// 	github.com/go-openapi/loads => github.com/go-openapi/loads v0.19.5
// 	github.com/go-openapi/runtime => github.com/go-openapi/runtime v0.19.15
// 	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.7
// 	github.com/go-openapi/strfmt => github.com/go-openapi/strfmt v0.19.5
// 	github.com/go-openapi/swag => github.com/go-openapi/swag v0.19.9
// 	github.com/go-openapi/validate => github.com/go-openapi/validate v0.19.8
// )
// module git.jd.com/mypass/mypass-openapi-operator

// go 1.13

// require (
// 	git.jd.com/mypass/helm-utils v0.0.10
// 	git.jd.com/mypass/mypass-client v0.0.1
// 	github.com/caarlos0/env v3.5.0+incompatible
// 	github.com/ghodss/yaml v1.0.0
// 	github.com/go-git/go-git/v5 v5.1.0
// 	github.com/go-logr/logr v0.1.0
// 	github.com/gofrs/flock v0.8.0
// 	github.com/golang/protobuf v1.5.2
// 	github.com/google/uuid v1.1.2
// 	github.com/gorilla/websocket v1.4.2
// 	github.com/onsi/ginkgo v1.14.1
// 	github.com/onsi/gomega v1.10.2
// 	github.com/pkg/errors v0.9.1
// 	github.com/rancher/remotedialer v0.0.1
// 	github.com/sirupsen/logrus v1.6.0
// 	github.com/spf13/pflag v1.0.5
// 	go.uber.org/zap v1.16.0
// 	golang.org/x/net v0.0.0-20210331060903-cb1fcc7394e5
// 	google.golang.org/grpc v1.31.0
// 	google.golang.org/protobuf v1.26.0
// 	gopkg.in/natefinch/lumberjack.v2 v2.0.0
// 	gopkg.in/yaml.v2 v2.4.0
// 	helm.sh/helm/v3 v3.3.1
// 	k8s.io/api v0.18.8
// 	k8s.io/apiextensions-apiserver v0.18.8
// 	k8s.io/apimachinery v0.18.8
// 	k8s.io/cli-runtime v0.18.8
// 	k8s.io/client-go v0.18.8
// 	k8s.io/klog v1.0.0
// 	k8s.io/klog/v2 v2.0.0
// 	k8s.io/kubectl v0.18.8
// 	sigs.k8s.io/controller-runtime v0.6.2
// 	sigs.k8s.io/yaml v1.2.0

// )

// replace (
// 	git.jd.com/mypass/helm-utils => ../helm-utils
// 	git.jd.com/mypass/mypass-client => ../mypass-client
// 	github.com/go-openapi/analysis => github.com/go-openapi/analysis v0.19.10
// 	github.com/go-openapi/errors => github.com/go-openapi/errors v0.19.4
// 	github.com/go-openapi/jsonpointer => github.com/go-openapi/jsonpointer v0.19.3
// 	github.com/go-openapi/jsonreference => github.com/go-openapi/jsonreference v0.19.3
// 	github.com/go-openapi/loads => github.com/go-openapi/loads v0.19.5
// 	github.com/go-openapi/runtime => github.com/go-openapi/runtime v0.19.15
// 	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.7
// 	github.com/go-openapi/strfmt => github.com/go-openapi/strfmt v0.19.5
// 	github.com/go-openapi/swag => github.com/go-openapi/swag v0.19.9
// 	github.com/go-openapi/validate => github.com/go-openapi/validate v0.19.8
// 	github.com/rancher/remotedialer => ../remotedialer
// )

// module git.jd.com/mypass/vessel-api-test

// go 1.15

// require (
// 	git.jd.com/jvessel/openapi-go-sdk v0.0.0-20201215082858-5c1df1d86e07
// 	github.com/gofrs/uuid v4.0.0+incompatible // indirect
// 	github.com/gorilla/mux v1.8.0
// 	github.com/imdario/mergo v0.3.11 // indirect
// 	github.com/satori/go.uuid v1.2.0
// 	github.com/stretchr/testify v1.7.0
// 	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
// 	gopkg.in/yaml.v2 v2.3.0
// 	k8s.io/api v0.18.8
// 	k8s.io/apimachinery v0.18.8
// 	k8s.io/client-go v0.18.8
// 	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009 // indirect
// )

// module git.jd.com/mypass/mypass-kube-proxy

// go 1.12

// require (
// 	git.jd.com/mypass/mypass-client v0.0.3-0.20210108130801-9231f9dc3f18
// 	github.com/Masterminds/semver/v3 v3.1.0 // indirect
// 	github.com/avast/retry-go v3.0.0+incompatible // indirect
// 	github.com/caarlos0/env v3.5.0+incompatible // indirect
// 	github.com/chartmuseum/auth v0.4.1 // indirect
// 	github.com/dgrijalva/jwt-go v3.2.0+incompatible
// 	github.com/docker/distribution v2.7.1+incompatible // indirect
// 	github.com/docker/docker v1.4.2-0.20200203170920-46ec8731fbce // indirect
// 	github.com/docker/go-connections v0.4.0 // indirect
// 	github.com/docker/go-metrics v0.0.1 // indirect
// 	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
// 	github.com/emicklei/go-restful v2.9.6+incompatible // indirect
// 	github.com/emicklei/go-restful-openapi v1.4.1 // indirect
// 	github.com/fatih/structs v1.1.0 // indirect
// 	github.com/fsnotify/fsnotify v1.4.9 // indirect
// 	github.com/gbrlsnchs/jwt/v3 v3.0.0-rc.2 // indirect
// 	github.com/ghodss/yaml v1.0.0 // indirect
// 	github.com/gin-contrib/zap v0.0.1 // indirect
// 	github.com/gin-gonic/gin v1.6.3 // indirect
// 	github.com/go-logr/logr v0.1.0 // indirect
// 	github.com/go-logr/zapr v0.1.0 // indirect
// 	github.com/go-oauth2/oauth2/v4 v4.1.2 // indirect
// 	github.com/go-playground/validator/v10 v10.2.0 // indirect
// 	github.com/go-sql-driver/mysql v1.5.0 // indirect
// 	github.com/gofrs/uuid v3.3.0+incompatible // indirect
// 	github.com/google/uuid v1.1.1 // indirect
// 	github.com/gorilla/mux v1.7.4
// 	github.com/gorilla/websocket v1.4.2 // indirect
// 	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
// 	github.com/imdario/mergo v0.3.9 // indirect
// 	github.com/influxdata/influxdb-client-go v1.4.0 // indirect
// 	github.com/jinzhu/gorm v1.9.15 // indirect
// 	github.com/json-iterator/go v1.1.10 // indirect
// 	github.com/labstack/gommon v0.3.0 // indirect
// 	github.com/olivere/elastic v6.2.35+incompatible // indirect
// 	github.com/olivere/elastic/v7 v7.0.19 // indirect
// 	github.com/opencontainers/go-digest v1.0.0 // indirect
// 	github.com/pkg/errors v0.9.1
// 	github.com/rancher/norman v0.0.0-20200930000340-693d65aaffe3
// 	github.com/rancher/remotedialer v0.0.1
// 	github.com/sirupsen/logrus v1.6.0
// 	github.com/spf13/cobra v1.1.3
// 	github.com/spf13/pflag v1.0.5
// 	github.com/spf13/viper v1.7.1
// 	github.com/yvasiyarov/go-metrics v0.0.0-20150112132944-c25f46c4b940 // indirect
// 	github.com/yvasiyarov/gorelic v0.0.7 // indirect
// 	github.com/yvasiyarov/newrelic_platform_go v0.0.0-20160601141957-9c099fbc30e9 // indirect
// 	go.uber.org/zap v1.16.0 // indirect
// 	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
// 	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
// 	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
// 	gopkg.in/yaml.v2 v2.4.0
// 	helm.sh/helm/v3 v3.3.1 // indirect
// 	k8s.io/api v0.18.8
// 	k8s.io/apiextensions-apiserver v0.18.8 // indirect
// 	k8s.io/apimachinery v0.18.8
// 	k8s.io/apiserver v0.18.8 // indirect
// 	k8s.io/cli-runtime v0.18.8 // indirect
// 	k8s.io/client-go v0.18.8
// 	k8s.io/klog v1.0.0
// 	k8s.io/kubectl v0.18.8 // indirect
// 	//sigs.k8s.io/cluster-api v0.3.14 // indirect
// 	sigs.k8s.io/controller-runtime v0.6.3
// 	sigs.k8s.io/kubefed v0.4.1 // indirect
// 	sigs.k8s.io/yaml v1.2.0 // indirect

// )

// replace (
// 	github.com/spf13/afero => github.com/spf13/afero v1.2.2
// 	github.com/spf13/pflag => github.com/spf13/pflag v1.0.5
// 	github.com/spf13/viper => github.com/spf13/viper v1.4.0

// 	git.jd.com/mypass/mypass-client => ../mypass-client
// 	github.com/rancher/remotedialer => ../remotedialer
// )

// module github.com/opsgenie/kubernetes-event-exporter

// go 1.14

// require (
// 	cloud.google.com/go v0.60.0 // indirect
// 	cloud.google.com/go/bigquery v1.9.0
// 	cloud.google.com/go/pubsub v1.3.1
// 	github.com/Masterminds/goutils v1.1.0 // indirect
// 	github.com/Masterminds/semver v1.5.0 // indirect
// 	github.com/Masterminds/sprig v2.22.0+incompatible
// 	github.com/Shopify/sarama v1.24.1
// 	github.com/aws/aws-sdk-go v1.30.10
// 	github.com/elastic/go-elasticsearch/v6 v6.8.10
// 	github.com/elastic/go-elasticsearch/v7 v7.4.1
// 	github.com/gin-contrib/pprof v1.3.0
// 	github.com/gin-gonic/gin v1.6.3
// 	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
// 	github.com/gophercloud/gophercloud v0.0.0-20190126172459-c818fa66e4c8 // indirect
// 	github.com/hashicorp/golang-lru v0.5.3
// 	github.com/huandu/xstrings v1.2.0 // indirect
// 	github.com/imdario/mergo v0.3.8 // indirect
// 	github.com/json-iterator/go v1.1.10
// 	github.com/klauspost/cpuid v1.2.3 // indirect
// 	github.com/mitchellh/copystructure v1.0.0 // indirect
// 	github.com/nlopes/slack v0.6.0
// 	github.com/opsgenie/opsgenie-go-sdk-v2 v1.0.3
// 	github.com/prometheus/client_golang v1.9.0
// 	github.com/rs/zerolog v1.16.0
// 	github.com/sirupsen/logrus v1.6.0
// 	github.com/spf13/pflag v1.0.5 // indirect
// 	github.com/stretchr/objx v0.2.0 // indirect
// 	github.com/stretchr/testify v1.6.1
// 	github.com/zsais/go-gin-prometheus v0.1.0
// 	google.golang.org/api v0.28.0
// 	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
// 	gopkg.in/inf.v0 v0.9.1 // indirect
// 	gopkg.in/natefinch/lumberjack.v2 v2.0.0
// 	gopkg.in/yaml.v2 v2.3.0
// 	k8s.io/api v0.20.1
// 	k8s.io/apimachinery v0.20.1
// 	k8s.io/apiserver v0.20.1 // indirect
// 	k8s.io/client-go v0.20.1
// 	k8s.io/klog v1.0.0
// )

// module mypass-modify

// go 1.15

// require (
// 	github.com/pkg/errors v0.9.1
// 	github.com/spf13/cobra v1.1.1
// 	github.com/spf13/pflag v1.0.5
// 	k8s.io/klog v1.0.0
// )

// module git.jd.com/lishijun10/mypass-jindowin-api

// go 1.15

// require (
// 	git.jd.com/lishijun10/map2properties v0.0.1
// 	github.com/Masterminds/sprig/v3 v3.1.0
// 	github.com/go-logr/logr v0.2.1
// 	github.com/huandu/xstrings v1.3.2 // indirect
// 	github.com/mitchellh/go-homedir v1.1.0
// 	github.com/spf13/cobra v1.1.1
// 	github.com/spf13/viper v1.7.1
// 	github.com/thedevsaddam/gojsonq v2.3.0+incompatible
// 	go.uber.org/zap v1.13.0
// 	gopkg.in/natefinch/lumberjack.v2 v2.0.0
// 	k8s.io/api v0.19.3
// 	k8s.io/apiextensions-apiserver v0.19.3
// 	k8s.io/apimachinery v0.19.3
// 	k8s.io/cli-runtime v0.19.3
// 	k8s.io/client-go v0.19.3
// 	sigs.k8s.io/controller-runtime v0.6.3
// 	sigs.k8s.io/kubefed v0.5.0
// 	sigs.k8s.io/yaml v1.2.0
// 	git.jd.com/mypass/mypass-client v0.0.1
// )

// replace(
//     git.jd.com/mypass/mypass-client => ../mypass-client
// )

// module git.jd.com/lishijun10/mypass-jindowin-api

// go 1.15

// require (
// 	git.jd.com/lishijun10/map2properties v0.0.1
// 	github.com/Masterminds/sprig/v3 v3.1.0
// 	github.com/go-logr/logr v0.2.1
// 	github.com/huandu/xstrings v1.3.2 // indirect
// 	github.com/mitchellh/go-homedir v1.1.0
// 	github.com/spf13/cobra v1.1.1
// 	github.com/spf13/viper v1.7.1
// 	github.com/thedevsaddam/gojsonq v2.3.0+incompatible
// 	go.uber.org/zap v1.13.0
// 	gopkg.in/natefinch/lumberjack.v2 v2.0.0
// 	k8s.io/api v0.19.3
// 	k8s.io/apiextensions-apiserver v0.19.3
// 	k8s.io/apimachinery v0.19.3
// 	k8s.io/cli-runtime v0.19.3
// 	k8s.io/client-go v0.19.3
// 	sigs.k8s.io/controller-runtime v0.6.3
// 	sigs.k8s.io/kubefed v0.5.0
// 	sigs.k8s.io/yaml v1.2.0
// 	git.jd.com/mypass/mypass-client v0.0.1
// )

// replace(
//     git.jd.com/mypass/mypass-client => ../mypass-client
// )

// module git.jd.com/mypass/mypass-controller-manager

// go 1.13

// require (
// 	git.jd.com/mypass/mypass-client v0.1.1-0.20210524070449-2e242abe7382
// 	github.com/go-git/go-git/v5 v5.1.0
// 	github.com/go-logr/logr v0.1.0
// 	github.com/onsi/ginkgo v1.14.0
// 	github.com/onsi/gomega v1.10.1
// 	github.com/pkg/errors v0.9.1
// 	github.com/sirupsen/logrus v1.6.0
// 	github.com/spf13/cobra v1.1.3
// 	github.com/spf13/pflag v1.0.5
// 	helm.sh/helm/v3 v3.3.4
// 	k8s.io/api v0.18.8
// 	k8s.io/apiextensions-apiserver v0.18.8
// 	k8s.io/apimachinery v0.18.8
// 	k8s.io/client-go v0.18.8
// 	k8s.io/klog v1.0.0
// 	k8s.io/klog/v2 v2.0.0
// 	sigs.k8s.io/controller-runtime v0.6.2
// 	sigs.k8s.io/kubefed v0.4.1
// 	sigs.k8s.io/yaml v1.2.0
// )

// replace (
// 	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.8
//  	git.jd.com/mypass/mypass-client => ../mypass-client
// )

// module git.jd.com/mypass/map2properties

// go 1.15

// require (
// 	github.com/fatih/structs v1.1.0
// 	github.com/pkg/errors v0.9.1
// 	gopkg.in/yaml.v2 v2.3.0 // indirect
// 	sigs.k8s.io/yaml v1.2.0
// )

// module git.jd.com/mypass/mypass-client

// go 1.13

// require (
// 	github.com/fsnotify/fsnotify v1.4.9 // indirect
// 	github.com/go-openapi/spec v0.20.3 // indirect
// 	github.com/go-openapi/swag v0.19.15 // indirect
// 	github.com/golang/protobuf v1.5.2 // indirect
// 	github.com/hashicorp/golang-lru v0.5.4 // indirect
// 	github.com/imdario/mergo v0.3.9 // indirect
// 	github.com/json-iterator/go v1.1.10 // indirect
// 	github.com/kr/pretty v0.2.1 // indirect
// 	github.com/mailru/easyjson v0.7.7 // indirect
// 	github.com/onsi/gomega v1.10.1 // indirect
// 	github.com/prometheus/procfs v0.0.11 // indirect
// 	github.com/stretchr/testify v1.7.0 // indirect
// 	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
// 	golang.org/x/mod v0.4.2 // indirect
// 	golang.org/x/net v0.0.0-20210331060903-cb1fcc7394e5 // indirect
// 	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558 // indirect
// 	golang.org/x/tools v0.1.0 // indirect
// 	google.golang.org/appengine v1.6.7 // indirect
// 	gopkg.in/yaml.v2 v2.4.0 // indirect
// 	helm.sh/helm/v3 v3.3.1
// 	k8s.io/api v0.18.8
// 	k8s.io/apimachinery v0.18.8
// 	k8s.io/client-go v0.18.8
// 	k8s.io/code-generator v0.18.8
// 	k8s.io/klog v1.0.0
// 	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29 // indirect
// 	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451 // indirect
// 	sigs.k8s.io/controller-runtime v0.6.0
// )

// replace (
// 	github.com/go-openapi/analysis => github.com/go-openapi/analysis v0.19.10
// 	github.com/go-openapi/errors => github.com/go-openapi/errors v0.19.4
// 	github.com/go-openapi/jsonpointer => github.com/go-openapi/jsonpointer v0.19.3
// 	github.com/go-openapi/jsonreference => github.com/go-openapi/jsonreference v0.19.3
// 	github.com/go-openapi/loads => github.com/go-openapi/loads v0.19.5
// 	github.com/go-openapi/runtime => github.com/go-openapi/runtime v0.19.15
// 	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.7
// 	github.com/go-openapi/strfmt => github.com/go-openapi/strfmt v0.19.5
// 	github.com/go-openapi/swag => github.com/go-openapi/swag v0.19.9
// 	github.com/go-openapi/validate => github.com/go-openapi/validate v0.19.8
// )
