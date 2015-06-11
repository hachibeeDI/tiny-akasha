
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

  deleteQuestion: () ->
    @dispatch 'question:delete', @props.id

  render: () ->
    # TODO: 削除ボタンのデザインや仕様は要検討
    $c('li', {className: 'question', key: @props.id},
      $c('div', {className: 'question__inner', onClick: @showQuestion},
        $c('h3', {className: 'question__title'}, @props.title),
        $c('p', {className: 'question__digest'}, @props.content),
      ),
      $c('button', {className: 'question__delete-button', onClick: @deleteQuestion}, 'この質問を削除する'),
    )


module.exports = QuestionComponent
