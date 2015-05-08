#!/bin/bash

# Start to run zookeeper as background process
bin/zookeeper-server-start.sh config/zookeeper.properties &

# Start kafka server
bin/kafka-server-start.sh config/server.properties
