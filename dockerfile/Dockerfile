FROM debian:jessie
MAINTAINER OGURA_Daiki 8hachibee125@gmail.com

ENV MYSQL_USER admin
ENV MYSQL_PASS qwerty15
ENV DB_HOST localhost
ENV DB_PORT 3306


RUN apt-get update -yq && apt-get upgrade -yq && \
    apt-get install -yq --no-install-recommends supervisor && \
    apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

ADD ./env/supervisord/supervisord.conf /etc/supervisor/
ADD ./env/supervisord/goji.conf /etc/supervisor/conf.d/

RUN mkdir -p /var/http/tiny-akasha/template/static/dist/; \
    mkdir -p /var/http/tiny-akasha/template/static/vendor/react; \
    mkdir -p /var/http/tiny-akasha/template/static/vendor/font/octicons

# deploy server side files
ADD ./src/main /var/http/tiny-akasha/
ADD ./src/templates/index.html /var/http/tiny-akasha/template/

# deploy frontend files
ADD ./src/templates/static/dist/bundle.js /var/http/tiny-akasha/template/static/dist/
ADD ./src/templates/static/dist/style/main.css /var/http/tiny-akasha/template/static/dist/style/

ADD ./src/templates/static/vendor/react/react-with-addons.min.js /var/http/tiny-akasha/template/static/vendor/react/
ADD ./src/templates/static/vendor/react/react-with-addons.js /var/http/tiny-akasha/template/static/vendor/react/
ADD ./src/templates/static/vendor/octicons/* /var/http/tiny-akasha/template/static/vendor/font/octicons/

EXPOSE 80


ENTRYPOINT ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]

