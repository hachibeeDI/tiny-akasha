request = require 'superagent'

EachAnswerComponent = React.createClass(
  mixins: [Arda.mixin]

  delete: (ev) ->
    @dispatch 'answer:delete', @props.id

  render: () ->
    console.log 'each-answer render ', @
    template = require('./each-answer-view')
    template @
)



module.exports = EachAnswerComponent
