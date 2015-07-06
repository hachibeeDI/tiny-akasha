request = require 'superagent'
md2react = require 'md2react'


EachAnswerComponent = React.createClass(
  mixins: [Arda.mixin]

  getInitialState: ->
    # 編集機能をつけるかもしれんのでstateに
    try
      renders = md2react @props.content,
          gfm: true
          breaks: true
          tables: true
      return render: renders
    catch e
      console.warn 'mark down parse error', e
      return render: []

  delete: (ev) ->
    @dispatch 'answer:delete', @props.id

  render: () ->
    console.log 'each-answer render ', @
    template = require('./each-answer-view')
    template @
)



module.exports = EachAnswerComponent
