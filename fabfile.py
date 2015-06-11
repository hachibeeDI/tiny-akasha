# -*- coding: utf-8 -*-

from fabric.api import lcd, local
from fabric.api import run, sudo, cd, env, hide, settings, put
from fabric.contrib.files import exists


env.use_ssh_config = True

# env.hosts = [TEST_SERVER]


def prepare():
    with cd('~/var/http'):
        run('mkdir -p tiny-akasha/template/static/dist/')
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


def deploy_vendors():
    ''' サードパーティのデプロイ '''
    root = '~/var/http/tiny-akasha/template/static/'
    run('mkdir -p {0}vendor/{{react,font/octicons}}'.format(root))
    path_to_react = '{0}vendor/react/'.format(root)
    if not exists('{0}react-with-addons.js'.format(path_to_react)):
        put('./template/static/node_modules/react/dist/*.js', path_to_react)

    path_to_octicon = '{0}vendor/font/octicons/'.format(root)
    if not exists(path_to_octicon + 'octicons.ttf'):
        put('./template/static/node_modules/octicons/octicons/*', path_to_octicon)


def deploy_assets():
    build_assets()
    deploy_vendors()
    put('./template/static/dist/', '~/var/http/tiny-akasha/template/static/')


def deploy():
    prepare()
    deploy_assets()
    deploy_server()


def stop():
    sudo('supervisorctl stop akasha')


def start():
    sudo('supervisorctl start akasha')


def clean():
    run('rm ./main')
