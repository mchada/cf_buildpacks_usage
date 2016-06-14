# Cloud Foundry Buildpack Usage CLI Plugin

Cloud Foundry plugin extension to view all buildpacks used in across a Cloud Foundry or in specific organizations and spaces.

## Install

```
$ go get github.com/rahul-kj/cf_buildpacks_usage
$ cf install-plugin $GOPATH/bin/cf_buildpacks_usage
```

## Usage

```
$ cf buildpack-usage

Following is the table of apps and buildpacks app deployments

-------------------------------
| console - ruby_buildpack |
| app-usage-server - ruby_buildpack |
| apps-manager-js - staticfile_buildpack |
| app-usage-scheduler - ruby_buildpack |
| app-usage-worker - ruby_buildpack |
| notifications - Go |
| notifications-ui - Go |
| autoscale - Go |
| spring-cloud-broker - java-buildpack=v3.6-offline-https://github.com/cloudfoundry/java-buildpack.git#5194155 java-main open-jdk-like-jre=1.8.0_71 open-jdk-like-memory-calculator=2.0.1_RELEASE spring-auto-reconfiguration=1.10.0_RELEASE |
| spring-cloud-broker-worker - java-buildpack=v3.6-offline-https://github.com/cloudfoundry/java-buildpack.git#5194155 java-main open-jdk-like-jre=1.8.0_71 open-jdk-like-memory-calculator=2.0.1_RELEASE spring-auto-reconfiguration=1.10.0_RELEASE |
| metrics - java_buildpack_offline |
| metrics-ui - staticfile_buildpack |
-------------------------------

## Uninstall

```
$ cf uninstall-plugin buildpack-usage
```
