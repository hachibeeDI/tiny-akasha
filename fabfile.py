# -*- coding: utf-8 -*-

from fabric.api import lcd, local
from fabric.api import run, sudo, cd, env, hide, settings, put
from fabric.contrib.files import exists


env.use_ssh_config = True

# env.hosts = [TEST_SERVER]


def prepare():
    with cd('~/var/http'):
        run('mkdir -p tiny-akasha/template/static/')
    put('./env/supervisord/goji.conf', '/etc/supervisor/conf.d/', use_sudo=True)


def build_server():
    local('GOOS=linux GOARCH=amd64 gom exec go build main.go')


def build_assets():
    with lcd('./template/static/'):
        local('gulp build')


def build():
    build_server()
    build_assets()


def deploy_server():
    build_server()
    put('./main', '~/var/http/tiny-akasha/')
    put('./template/index.html', '~/var/http/tiny-akasha/template/index.html')


def deploy_assets():
    build_assets()
    put('./template/static/dist/', '~/var/http/tiny-akasha/template/static/')
    put('./template/static/node_modules/react/dist/react-with-addons.js', '~/var/http/tiny-akasha/template/static/dist/')
    put('./template/static/node_modules/react/dist/react-with-addons.min.js', '~/var/http/tiny-akasha/template/static/dist/')


def deploy():
    prepare()
    deploy_assets()
    stop()
    deploy_server()
    start()


def stop():
    run('supervisorctl stop')


def start():
    run('supervisorctl start')


def clean():
    run('rm ./main')
