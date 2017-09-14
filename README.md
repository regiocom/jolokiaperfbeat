# Jolokiaperfbeat

Welcome to Jolokiaperfbeat.

This [beat](https://www.elastic.co/de/products/beats) collects web service
performance metrics from a Java application server (e. g. [WildFly](http://wildfly.org/)). 

## Using Jolokiaperfbeat

* Enable [JMX integration for CXF](http://cxf.apache.org/docs/jmx-management.html).
* Install [Jolokia](https://jolokia.org/) to the application server, or 
  install [Hawtio](http://hawt.io/) which already contains Jolokia.  
* Configure `jolokiaperfbeat.yml`
  * `period`: The pause interval between data collection calls
  * `baseurl`: The path to access the Jolokia artifact
  * `output`: see [output options](https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-configuration-options.html)
* Install jolokiaperfbeat as a service, or
  run jolokiaperfbeat in a console
 

## Building Jolokiaperfbeat

Ensure that this folder is at the following location:
`${GOPATH}/github.com/regiocom/jolokiaperfbeat`

### Requirements

* [Golang](https://golang.org/dl/) 1.7
* `${GOPATH}/github.com/elastic/beats` on branch 5.5

### Init Project
To get running with Jolokiaperfbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Jolokiaperfbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/regiocom/jolokiaperfbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Jolokiaperfbeat run the command below. This will generate a binary
in the same directory with the name jolokiaperfbeat.

```
make
```


### Run

To run Jolokiaperfbeat with debugging output enabled, run:

```
./jolokiaperfbeat -c jolokiaperfbeat.yml -e -d "*"
```


### Test

To test Jolokiaperfbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/jolokiaperfbeat.template.json and etc/jolokiaperfbeat.asciidoc

```
make update
```


### Cleanup

To clean  Jolokiaperfbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Jolokiaperfbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/regiocom/jolokiaperfbeat
cd ${GOPATH}/github.com/regiocom/jolokiaperfbeat
git clone https://github.com/regiocom/jolokiaperfbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
