window.React = require 'react'
window.Promise = require 'bluebird'
window.Arda = require 'arda'
window.$ = require 'jquery'
window._ = require 'lodash'
window.Routers = {}


document.addEventListener 'DOMContentLoaded', () ->
  Routers.post = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-post-question'))
  PostContext = require './post-question/index'
  Routers.post.pushContext(PostContext, {})

  Routers.main = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-main'))
  IndexContext = require './index/index'
  $.get('/api/v1/question')
    .then((data) =>
      console.log 'start applications!', data
      unless data.error?
        Routers.main.pushContext(IndexContext, data)
    )
