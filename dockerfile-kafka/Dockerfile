FROM java:openjdk-8-jre

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get install -y \
                              zookeeper \
                              wget \
                              dnsutils \
                              vim \
            && rm -rf /var/lib/apt/lists/*

ENV KAFKA_VERSION 0.8.2.1
ENV SCALA_VERSION 2.10
RUN wget -q \
    http://apache.mirrors.spacedump.net/kafka/${KAFKA_VERSION}/kafka_${SCALA_VERSION}-${KAFKA_VERSION}.tgz \
    -O /tmp/kafka.tgz \
    && tar xfz /tmp/kafka.tgz -C /opt \
    && rm /tmp/kafka.tgz \
    && mv /opt/kafka_${SCALA_VERSION}-${KAFKA_VERSION} /opt/kafka


ADD usr/local/bin/run.sh /usr/local/bin/run.sh

WORKDIR /opt/kafka
CMD ["/usr/local/bin/run.sh"]

