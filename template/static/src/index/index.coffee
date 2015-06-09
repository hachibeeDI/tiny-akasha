
$c = React.createElement.bind(React)

QuestionContext = require './question-context'


IndexComponent = React.createClass(
  mixins: [Arda.mixin]
  # componentWillMount: () ->

  render: () ->
    console.log 'index component'
    $c('div', {},
      $c('ul', {className: 'questions__ul'},
        @props.questions.map (q) ->
          q['key'] = q['id']
          $c(QuestionContext, q)
      ),
    )
)


class IndexContext extends Arda.Context
  component: IndexComponent
  delegate: (subscribe) ->
    super
    actions = require './actions'
    subscribe 'show:questions', (questions) =>
      @update((s) => questions: questions)

    subscribe 'question:show', (id) =>
      actions.showQuestion(id)

  initState: (props) ->
    return {questions: props['questions'] or [], }

  expandComponentProps: (props, state) ->
    return {questions: state['questions']}


module.exports = IndexContext
