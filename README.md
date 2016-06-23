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

Following is the markdown output

|        ORG |                          SPACE |                    APPLICATION |      STATE |                                                                                                                                                                                                                                                  BUILDPACK |
|      ----- |                          ----- |                          ----- |      ----- |                                                                                                                                                                                                                                                      ----- |
|     system |                   apps-manager |             apps-manager-green |    STARTED |                                                                                                                                                                                                                                             ruby_buildpack |
```

```
$ cf buildpack-usage --csv

Following is the csv output

ORG,SPACE,APPLICATION,STATE,BUILDPACK
system,apps-manager,apps-manager-green,STARTED,ruby_buildpack

```

## Uninstall

```
$ cf uninstall-plugin buildpack-usage
```
