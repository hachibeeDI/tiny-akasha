
EachQuestionComponent = React.createClass(
  mixins: [Arda.mixin]

  goBack: (ev) ->
    Routers.main.popContext()

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
###
class EachQuestionContext extends Arda.Context
  component: EachQuestionComponent
  # initState: (props) ->
  #   console.log 'each question init'
  #   return {question: [], }

  delegate: (subscribe) ->
    super


  expandComponentProps: (props, state) ->
    console.log 'each question expand', props, state
    return props


module.exports = EachQuestionContext
