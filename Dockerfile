# Use Ubuntu 22.04 as the base image
FROM ubuntu:22.04

# Install app dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
#    apt-transport-https \
    ca-certificates \
    curl \
    bzip2
#    && rm -rf /var/lib/apt/lists/*

# Download and install sqlcmd
RUN curl -o sqlcmd.bz2 -fsSL https://github.com/microsoft/go-sqlcmd/releases/download/v1.6.0/sqlcmd-v1.6.0-linux-amd64.tar.bz2 \
    && tar -xf sqlcmd.bz2 \
    && chmod +x sqlcmd \
    && mv sqlcmd /bin/sqlcmd

## Download and install restic
#RUN curl -fsSL -o restic.bz2 https://github.com/restic/restic/releases/download/v0.16.4/restic_0.16.4_linux_amd64.bz2 \
#    && bzip2 -d restic.bz2 \
#    && chmod +x restic \
#    && mv restic /bin/restic

# Set the default shell to bash for interactive use
SHELL ["/bin/bash", "-c"]




#RUN set -x \
#  && apt-get update \
#  && apt-get install -y --no-install-recommends apt-transport-https ca-certificates curl wget gnupg lsb-release software-properties-common bzip2
#
#RUN set -x \
#    && curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > /etc/apt/trusted.gpg.d/microsoft.gpg \
#    && add-apt-repository "$(wget -qO- https://packages.microsoft.com/config/ubuntu/22.04/prod.list)" \
#    && apt-get update && apt-get install -y --no-install-recommends sqlcmd