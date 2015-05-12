
EachQuestionComponent = React.createClass(
  mixins: [Arda.mixin]

  render: () ->
    console.log 'each question component render', @props
    template = require('./view')
    template @
)


###
props:
  id: number
###
class EachQuestionContext extends Arda.Context
  component: EachQuestionComponent
  initState: (props) ->
    console.log 'each question init'
    return {question: [], }

  delegate: (subscribe) ->
    super
    subscribe 'context:created', =>
      $.get("/api/v1/question/id/#{@props.id}")
        .then((data) =>
          console.log 'each question context created', data
          @update((s) =>
            question: data
          )
        )

  expandComponentProps: (props, state) ->
    console.log 'each question expand', props, state
    return state['question']


module.exports = EachQuestionContext
