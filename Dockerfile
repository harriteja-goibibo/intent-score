# Get base centos image
FROM centos:7

#Install epel-release, debug tools, nginx, supervisor, boto packages.
#Create required directories
RUN \
  yum install -y epel-release bison python-setuptools bzip2 wget make gcc git &&  \
  easy_install supervisor && \
  mkdir -p /logs/intent-score /etc/supervisord.d  && \
  yum clean all && \
  rm -f /etc/localtime && \
  ln -s /usr/share/zoneinfo/Asia/Kolkata /etc/localtime

#Install Go
RUN \
  cd /tmp && \
  wget https://storage.googleapis.com/golang/go1.12.6.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz && \
  ln -s /usr/local/go/bin/go /bin/go && \
  ln -s /usr/local/go/bin/gofmt /bin/gofmt

#Set environment variables
ENV PATH=$PATH:/usr/local/go/bin:/usr/local/goibibo/intent-score/bin
ENV GO111MODULE=on CGO_ENABLED=0

EXPOSE 80
#Set argument env - to receive input
ARG env

WORKDIR /usr/local/goibibo/intent-score

#Add supervisord and functions in startup scripts
COPY ./init/supervisord /etc/rc.d/init.d/
COPY ./init/services/* /etc/supervisord.d/

COPY ./pkg ./pkg
COPY ./go.mod .

RUN go mod download

#Copy source directory
COPY ./ .

#Build
RUN \
    make && \
    chmod 644 ./deployments/intent-score.sh && \
    chmod 755 /etc/rc.d/init.d/supervisord
