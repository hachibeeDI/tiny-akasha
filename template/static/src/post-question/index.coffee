$c = React.createElement.bind(React)


PostComponent = React.createClass(
  mixins: [Arda.mixin, React.addons.LinkedStateMixin]
  getInitialState: ()  ->
    return name: '', title: '', content: ''

  postQuestion: (ev) ->
    ev.preventDefault()
    $.post('/api/v1/question', {
      'title': @state.title
      'name': @state.name
      'content': @state.content
    })

  render: () ->
    $c('form', {className: 'post__panel', onSubmit: @postQuestion}, [
        $c('label', {className: 'label--row'},
            ['タイトル', $c('input', {type: 'text', name: 'title', valueLink: @linkState('title')})]),
        $c('label', {className: 'label--row'},
            ['名前', $c('input', {type: 'text', name: 'name', valueLink: @linkState('name')})]),
        $c('label', {className: 'label--row'},
            ['内容', $c('textarea', {name: 'content', valueLink: @linkState('content')})]),
        $c('input', {type: 'submit', value: '投稿', }),
      ]
    )
)


class PostContext extends Arda.Context
  component: PostComponent
  # delegate: (subscribe) ->
  #   super
  #   subscribe 'context:created', =>
  #     $.get('/api/v1/question')
  #       .then((data) =>
  #         console.log data
  #         @update((s) =>
  #           questions: data['questions']
  #         )
  #       )


  initState: (props) ->
    return {}

  expandComponentProps: (props, state) ->
    return {}


module.exports = PostContext
