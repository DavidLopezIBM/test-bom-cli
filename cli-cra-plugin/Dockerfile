FROM registry.access.redhat.com/ubi8/nodejs-14


# AT tracker log path
RUN mkdir -p /opt/app-root/log/at

# switching user to root to install python for newrelic module
USER root
# RUN yum install python2 -y
RUN yum install --disablerepo=* --enablerepo=ubi-8-appstream --enablerepo=ubi-8-baseos python2 -y
# remove mysql package to avoid vulnerability issues
RUN yum remove mariadb-connector-c mariadb-connector-c-config mariadb-connector-c-devel -y
# removing mysql configuration to avoid vulnerabilities issues
RUN rm -fr /etc/mysql
# update packages
RUN subscription-manager register --username ${REDHAT_USERNAME} --password ${REDHAT_PASSWORD} --auto-attach
RUN yum update -y && yum clean all
# switching back to no root user
USER default
# copying app source
# source isn't copied anymore with original permissions
# COPY . /opt/app-root/src
# COPY --chown=default . /opt/app-root/src
# using user UID instead of its name as not existing on origin system
COPY --chown=1001 . /opt/app-root/src


RUN npm rebuild

RUN echo "I am $(whoami) user now - now we can start the microservice"

FROM registry.access.redhat.com/ubi8/nodejs-14


CMD echo "hello world"



