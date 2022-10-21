# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT="1.19-bullseye"
FROM golang:${VARIANT}

ARG NON_ROOT_USER="www"
RUN groupadd -g 1000 ${NON_ROOT_USER}
RUN useradd -u 1000 -ms /bin/bash -g ${NON_ROOT_USER} ${NON_ROOT_USER}

ARG NODE_VERSION="none"
# [Choice] Node.js version: none, lts/*, 18, 16, 14
RUN if [ "${NODE_VERSION}" != "none" ]; then su ${NON_ROOT_USER} -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# Set destination for COPY
WORKDIR /var/www

COPY .env .

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . /var/www

# Build
RUN go build -o /build

COPY --chown=www:www . /var/www

USER www

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
ARG PORT="8080"
EXPOSE ${PORT}

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/build" ]
