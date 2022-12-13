### builder stage, used to compile the application
FROM openfunction/maven:3-jdk-8-openj9 AS builder

# copy the sources into the builder stage
COPY . /app
WORKDIR /app

# build the application
RUN mvn package

### final stage, used to setup the open liberty runtime
FROM openfunction/open-liberty:kernel-java8-openj9-ubi

# copy liberty config
COPY --chown=1001:0 src/main/liberty/config /config/

# copy compiled war file from the builder stage
COPY --chown=1001:0 --from=builder /app/target/*.war /config/apps/

# configure liberty based on the server.xml
RUN configure.sh
