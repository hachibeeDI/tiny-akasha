# -*- coding: utf-8 -*-

from fabric.api import lcd, local
from fabric.api import run, sudo, cd, env, hide, settings, put


env.use_ssh_config = True

# env.hosts = [TEST_SERVER]
env.key_filename = ['~/.ssh/github/id_rsa.github.org']


def prepare():
    with cd('~/var/http'):
        run('mkdir -p tiny-akasha/template/static/')


def build():
    local('GOOS=linux GOARCH=amd64 go build main.go')
    with lcd('./template/static/'):
        local('gulp build')


def deploy():
    prepare()
    build()
    put('./main', '~/var/http/tiny-akasha/')
    put('./template/index.html', '~/var/http/tiny-akasha/template/index.html')
    put('./template/static/dist/', '~/var/http/tiny-akasha/template/static/')
    put('./template/static/node_modules/react/dist/react-with-addons.js', '~/var/http/tiny-akasha/template/static/dist/')
    put('./template/static/node_modules/react/dist/react-with-addons.min.js', '~/var/http/tiny-akasha/template/static/dist/')


def clean():
    run('rm ./main')
