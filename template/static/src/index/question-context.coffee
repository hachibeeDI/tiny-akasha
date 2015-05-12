
$c = React.createElement.bind(React)


###
props:
  id: number
  title: string
  username: string
  content: string
###
QuestionComponent = React.createClass(
  mixins: [Arda.mixin]
  showQuestion: () ->
    EachQuestionContext = require '../each-question/context'
    Router.pushContext(EachQuestionContext, {id: @props.id})

  render: () ->
    console.log 'question!', @props
    $c('li', {key: @props.id, onClick: @showQuestion}, @props.title)
)


module.exports = QuestionComponent
