FROM mysql:latest

# initdbはdocker-compose.ymlでやるとPermission DeniedになるのでDockerfileで
COPY db/initdb/ /docker-entrypoint-initdb.d/
RUN chmod 777 /docker-entrypoint-initdb.d/*.sh
