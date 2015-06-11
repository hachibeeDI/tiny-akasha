
$c = React.createElement.bind(React)


###
props:
  id: number
  title: string
  username: string
  content: string
###
QuestionComponent = React.createClass
  mixins: [Arda.mixin]
  showQuestion: () ->
    @dispatch 'question:show', @props.id

  render: () ->
    $c('li', {className: 'question', key: @props.id, onClick: @showQuestion},
      $c('div', {className: 'question__inner'},
        $c('h3', {className: 'question__title'}, @props.title),
        $c('p', {className: 'question__digest'}, @props.content),
        $c('button', {className: 'question__delete-button'}, 'この質問を削除する'),
      )
    )


module.exports = QuestionComponent
