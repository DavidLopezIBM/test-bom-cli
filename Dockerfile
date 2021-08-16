# UBI8 comes with yum package manager
FROM registry.access.redhat.com/ubi8/ubi:latest

# Install required dependencies
RUN yum -y upgrade-minimal && yum install -y wget gnupg ca-certificates python3-pip curl \
    nodejs npm wget sudo java unzip @python27 golang git
# Upgrade Openssl 
RUN dnf install -y openssl-libs 

WORKDIR /

# Install JQ & ibmcloud & Maven & Gradle
RUN wget -O jq https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 &&\
    chmod +x ./jq &&\
    cp jq /usr/bin &&\
    jq --version &&\
    wget https://download.clis.cloud.ibm.com/ibm-cloud-cli/1.2.0/IBM_Cloud_CLI_1.2.0_amd64.tar.gz -O ibmcloud-cli.tar.gz &&\
    tar xvf ibmcloud-cli.tar.gz &&\
    rm ibmcloud-cli.tar.gz &&\
    cp Bluemix_CLI/bin/ibmcloud /usr/local/bin/ &&\
    ibmcloud --version && \
    wget https://mirrors.ocf.berkeley.edu/apache/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.tar.gz -O maven.tar.gz && \
    tar -zxvf maven.tar.gz && \
    mv apache-maven-3.6.3 /opt/maven && \
    rm maven.tar.gz && \
    ln -s apache-maven-3.6.3 maven && \
    wget https://services.gradle.org/distributions/gradle-6.8-bin.zip && \
    unzip gradle-*.zip && \
    mkdir /opt/gradle && \
    cp -pr gradle-*/* /opt/gradle

# Add required Environmental Variables for Gradle
ENV PATH=/opt/gradle/bin:${PATH}
RUN gradle -v
# Add required Java Environmental Variables for Maven
ENV M2_HOME=/opt/maven
ENV PATH=${M2_HOME}/bin:${PATH}
RUN mvn -version


RUN curl -L https://download.docker.com/linux/static/stable/x86_64/docker-19.03.0.tgz --output /tmp/docker-19.03.0.tgz && \
    tar -C /usr/share/ -xf /tmp/docker-19.03.0.tgz && \
    cp /usr/share/docker/docker /usr/local/bin/

RUN curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl" &&\
    cp kubectl /usr/bin && \
    chmod +x /usr/bin/kubectl

RUN rm -rf /etc/apache2

RUN curl https://bootstrap.pypa.io/pip/2.7/get-pip.py -o get-pip.py && \
    python2 get-pip.py

# Policies and scripts
COPY cli-cra-plugin/ /tmp/cli-cra-plugin


CMD echo "hello world"
