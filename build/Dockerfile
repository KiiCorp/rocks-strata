FROM golang:latest AS golang

# GOPATH is `/go`
RUN export GOPATH='/go'

# Configure GOOS and GOARCH
RUN export GOOS='linux'
RUN export GOARCH='amd64'

# go get libraries which are needed for building strata and mongoq command
RUN go get -v github.com/facebookgo/rocks-strata/strata \
  github.com/AdRoll/goamz/aws \
  github.com/AdRoll/goamz/s3 \
  github.com/AdRoll/goamz/s3/s3test \
  github.com/facebookgo/mgotest \
  gopkg.in/mgo.v2 \
  gopkg.in/mgo.v2/bson \
  golang.org/x/crypto/ssh/terminal \
  github.com/kr/pty \
  github.com/Azure/azure-sdk-for-go/storage \
  github.com/minio/minio-go

# Build strata command and mongoq command
RUN cd $GOPATH/src/github.com/facebookgo/rocks-strata/strata/cmd/mongo/lreplica_drivers/strata && go install
RUN cd $GOPATH/src/github.com/facebookgo/rocks-strata/strata/cmd/mongo/lreplica_drivers/mongoq && go install

FROM centos:7
# Define arguments
ARG GOPATH="/go"
ARG TARGET_PATH="/usr/bin"
ARG RPM_PACKAGE_NAME="rocks-strata"
ARG RPM_TARGET_PATH="/rpm"
ARG ARCHITECTURE="x86_64"
ARG CIRCLE_BUILD_NUM="local"

# Copy strata and mongoq command
COPY --from=golang ${GOPATH}/bin/* ${TARGET_PATH}/

# Install fpm
RUN yum install -y ruby-devel gcc make rpm-build rubygems
RUN gem install --no-ri --no-rdoc fpm

# Build rpm package
WORKDIR ${RPM_TARGET_PATH}
RUN fpm -s dir -t rpm \
  -n "${RPM_PACKAGE_NAME}" \
  -a "${ARCHITECTURE}" \
  --iteration "${CIRCLE_BUILD_NUM}.el7" \
  --url "https://github.com/facebookgo/rocks-strata" \
  --description "rocks-strata is a framework for managing incremental backups of databases that use the RocksDB storage engine" \
  ${TARGET_PATH}/strata ${TARGET_PATH}/mongoq
