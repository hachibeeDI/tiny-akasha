#!/usr/bin/env bash -v

cp -rf ../env/ ./

mkdir -p src/templates/static/{vendor,dist}
cp ../main src/
cp -rf ../template/index.html src/templates/
cp -rf ../template/static/dist/ src/templates/static/
cp -rf ../template/static/node_modules/react/dist/ src/templates/static/vendor/react
cp -rf ../template/static/node_modules/octicons/octicons/ src/templates/static/vendor/

