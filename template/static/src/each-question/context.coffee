request = require 'superagent-bluebird-promise'
md2react = require 'md2react'

###*
* @param {string} txt
* @return {Array}
###
_renderMd = (txt) ->
  try
    md2react txt,
      gfm: true
      breaks: true
      tables: true
  catch e
    console.warn 'mark down parse error', e
    []


EachQuestionComponent = React.createClass(
  mixins: [Arda.mixin]

  getInitialState: ->
    preview: []

  goBack: (ev) ->
    return if Routers.main.history.length <= 0
    global.history.back()
    Routers.main.popContext()

  onHandleAnswerFormSubmit: (ev) ->
    ev.preventDefault()
    username = React.findDOMNode(@refs.form__user)
    content = React.findDOMNode(@refs.form__content)
    return if username == '' || content == ''
    request
      .post "/api/v1/question/id/#{@props.id}/answer"
      .send name: username.value, content: content.value
      .set 'Accept', 'application/json'
      .then (data) =>
        console.log 'question created'
        username.value = ''
        content.value = ''
        @dispatch 'question:reload', @props.id

      .catch (err) ->
        console.error err

  renderPreviewMd: (ev) ->
    @setState preview: _renderMd @refs.form__content.getDOMNode().value

  render: () ->
    console.log 'each question component render', @props
    template = require('./view')
    template @
)

Actions = require '../index/actions'

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

    subscribe 'question:reload', (id) =>
      Actions.reloadQuestion(id)

    subscribe 'answer:delete', (id) =>
      request
        .del "/api/v1/answer/id/#{id}"
        .end (err, res) =>
          return console.error err if err?
          console.log res
          Actions.reloadQuestion(@props.id)


  expandComponentProps: (props, state) ->
    console.log 'each question expand', props, state
    props.content = _renderMd props.content
    return props


module.exports = EachQuestionContext
