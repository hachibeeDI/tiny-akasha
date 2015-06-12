window.React = require 'react'
window.Promise = require 'bluebird'
window.Arda = require 'arda'
window.$ = require 'jquery'
window._ = require 'lodash'
window.Routers = {}

Grapnel = require 'grapnel'
window.Navigator = new Grapnel(pushState: true)


document.addEventListener 'DOMContentLoaded', () ->
  Routers.post = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-post-question'))
  PostContext = require './post-question/index'
  Routers.post.pushContext(PostContext, {})

  Routers.main = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-main'))
  IndexContext = require './index/index'

  # routing
  do ->
    showRoot = ->
      $.get('/api/v1/question')
       .then (data) =>
          unless data.error?
            Routers.main.pushContext(IndexContext, data)

    Navigator.get '/', (req) ->
      showRoot()
    Navigator.get '/view', (req) ->
      showRoot()

    Navigator.get '/view/question/id/:id', (req) ->
      Actions = require './index/actions'
      Actions.showQuestion(req.params.id)
