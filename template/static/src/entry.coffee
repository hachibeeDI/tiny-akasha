window.React = require 'react'
window.Promise = require 'bluebird'
window.Arda = require 'arda'
window.$ = require 'jquery'
window._ = require 'lodash'


document.addEventListener 'DOMContentLoaded', () ->
  window.PostRouter = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-post-question'))
  PostContext = require './post-question/index'
  PostRouter.pushContext(PostContext, {})

  window.Router = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-main'))
  IndexContext = require './index/index'
  Router.pushContext(IndexContext, {})
