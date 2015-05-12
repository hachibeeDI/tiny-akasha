window.React = require 'react'
window.Promise = require 'bluebird'
window.Arda = require 'arda'
window.$ = require 'jquery'
window._ = require 'lodash'


document.addEventListener 'DOMContentLoaded', () ->
  IndexContext = require './index/index'
  window.Router = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-main'))
  Router.pushContext(IndexContext, {})
