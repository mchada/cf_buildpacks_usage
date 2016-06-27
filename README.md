# Cloud Foundry Buildpack Usage CLI Plugin

Cloud Foundry plugin extension to view all buildpacks used in across a Cloud Foundry or in specific organizations and spaces.

## Install

```
$ go get github.com/rahul-kj/cf_buildpacks_usage
$ cf install-plugin $GOPATH/bin/cf_buildpacks_usage
```

## Usage

**SAMPLE OUTPUT**

```
$ cf buildpack-usage

Following is the markdown output

|        ORG |                          SPACE |                    APPLICATION |      STATE |                                                                                                                                                                                                                                                  BUILDPACK |
|      ----- |                          ----- |                          ----- |      ----- |                                                                                                                                                                                                                                                      ----- |
|     system |                         system |                apps-manager-js |    STARTED |                                                                                                                                                                                                                                       staticfile_buildpack |
|     system |                         system |               app-usage-server |    STARTED |                                                                                                                                                                                                                                             ruby_buildpack |
|     system |                         system |                        console |    STARTED |                                                                                                                                                                                                                                             ruby_buildpack |
|     system |                         system |            app-usage-scheduler |    STARTED |                                                                                                                                                                                                                                             ruby_buildpack |
|     system |                         system |               app-usage-worker |    STARTED |                                                                                                                                                                                                                                             ruby_buildpack |
|     system |          notifications-with-ui |               notifications-ui |    STARTED |                                                                                                                                                                                                                                                         Go |
|     system |                    autoscaling |                      autoscale |    STARTED |                                                                                                                                                                                                                                                         Go |
```

```
$ cf buildpack-usage --csv

Following is the csv output

ORG,SPACE,APPLICATION,STATE,BUILDPACK
system,system,apps-manager-js,STARTED,staticfile_buildpack
system,system,app-usage-server,STARTED,ruby_buildpack
system,system,console,STARTED,ruby_buildpack
system,system,app-usage-scheduler,STARTED,ruby_buildpack
system,system,app-usage-worker,STARTED,ruby_buildpack
system,notifications-with-ui,notifications-ui,STARTED,Go
system,autoscaling,autoscale,STARTED,Go

```

```
$ cf buildpack-usage --verbose

Following is the csv output

ORG,SPACE,APPLICATION,STATE,INSTANCES,MEMORY,DISK
system,system,apps-manager-js,STARTED,6,64 MB,1024 MB
system,system,app-usage-server,STARTED,1,128 MB,1024 MB
system,system,console,STARTED,6,1024 MB,1024 MB
system,system,app-usage-scheduler,STARTED,1,128 MB,1024 MB
system,system,app-usage-worker,STARTED,1,1024 MB,1024 MB
system,notifications-with-ui,notifications-ui,STARTED,1,64 MB,1024 MB
system,autoscaling,autoscale,STARTED,1,256 MB,1024 MB

```

## Uninstall

```
$ cf uninstall-plugin buildpack-usage
```
