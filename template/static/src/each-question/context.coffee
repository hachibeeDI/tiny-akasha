request = require 'superagent'

EachQuestionComponent = React.createClass(
  mixins: [Arda.mixin]

  goBack: (ev) ->
    Routers.main.popContext()

  onHandleAnswerFormSubmit: (ev) ->
    ev.preventDefault()
    form = ev.target
    username = form.elements['user'].value
    content = form.elements['content'].value
    return if username == '' || content == ''
    request
      .post "/api/v1/question/id/#{@props.id}/answer"
      .send name: username, content: content
      .set 'Accept', 'application/json'
      .end (err, res) =>
        console.log 'question created'
        @dispatch 'question:show', @props.id

  render: () ->
    console.log 'each question component render', @props
    template = require('./view')
    template @
)


###
props:
  id: number
  title: string
  username: string
  content: string
  answers: Array{string}
###
class EachQuestionContext extends Arda.Context
  component: EachQuestionComponent
  # initState: (props) ->
  #   console.log 'each question init'
  #   return {question: [], }

  delegate: (subscribe) ->
    super
    subscribe 'answer:delete', (id) =>
      request
        .del "/api/v1/answer/id/#{id}"
        .end (err, res) =>
          return console.error err if err?
          console.log res
          actions = require '../index/actions'
          actions.showQuestion(@props.id)


  expandComponentProps: (props, state) ->
    console.log 'each question expand', props, state
    return props


module.exports = EachQuestionContext
