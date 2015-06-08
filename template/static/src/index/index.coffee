
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
    subscribe 'show:questions', (questions) =>
      @update((s) => questions: questions)

    subscribe 'question:show', (id) =>
      Promise.all([
        $.get("/api/v1/question/id/#{id}"),
        $.get("/api/v1/question/id/#{id}/answer"),
      ])
        .then (data) =>
          console.log 'question:show occurd', data
          Routers.main.pushContext(
            require('../each-question/context'),
            _.merge(data[0], data[1])
          )
        .catch (error) ->
          console.error 'each question', error

  initState: (props) ->
    return {questions: props['questions'] or [], }

  expandComponentProps: (props, state) ->
    return {questions: state['questions']}


module.exports = IndexContext
