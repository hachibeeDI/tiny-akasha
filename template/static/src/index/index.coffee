
$c = React.createElement.bind(React)

QuestionComponent = require './question-context'


IndexComponent = React.createClass(
  mixins: [Arda.mixin]
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
    subscribe 'context:created', =>
      $.get('/api/v1/question')
        .then((data) =>
          console.log data
          @update((s) =>
            questions: data['questions']
          )
        )

    subscribe 'question:show', (id) =>
      Router.pushContext(require('../each-question/context'), {id: id})

  initState: (props) ->
    return {questions: [], }

  expandComponentProps: (props, state) ->
    return {questions: state['questions']}


module.exports = IndexContext