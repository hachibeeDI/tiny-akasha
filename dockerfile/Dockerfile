FROM debian:jessie
MAINTAINER OGURA_Daiki 8hachibee125@gmail.com


RUN apt-get update -yq && apt-get upgrade -yq && \
    apt-get install -yq --no-install-recommends supervisor && \
    apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

ADD ./env/supervisord/goji.conf /etc/supervisor/conf.d/

RUN mkdir -p /var/http/tiny-akasha/template/static/dist/; \
    mkdir -p /var/http/tiny-akasha/template/static/vendor/{react,font/octicons}

# deploy server side files
ADD ./src/main /var/http/tiny-akasha/
ADD ./src/templates/index.html /var/http/tiny-akasha/template/

# deploy frontend files
ADD ./src/templates/static/vendor/react/react-with-addons.min.js /var/http/tiny-akasha/template/static/vendor/react/
ADD ./src/templates/static/vendor/octicons/* /var/http/tiny-akasha/template/static/vendor/octicons/

# ENTRYPOINT []
