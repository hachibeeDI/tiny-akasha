
$c = React.createElement.bind(React)

QuestionComponent = require './question-context'


IndexComponent = React.createClass(
  mixins: [Arda.mixin]
  # componentWillMount: () ->

  render: () ->
    console.log 'index component'
    $c('div', {},
      $c('ul', {},
        @props.questions.map (q) ->
          q['key'] = q['id']
          $c(QuestionComponent, q)
      ),
    )
)


class IndexContext extends Arda.Context
  component: IndexComponent
  delegate: (subscribe) ->
    super
    subscribe 'show:questions', (questions) =>
      @update((s) => questions: questions)

    subscribe 'question:show', (id) =>
      Routers.main.pushContext(require('../each-question/context'), {id: id})

  initState: (props) ->
    return {questions: props['questions'] or [], }

  expandComponentProps: (props, state) ->
    return {questions: state['questions']}


module.exports = IndexContext
