$c = React.createElement.bind(React)


PostComponent = React.createClass(
  mixins: [Arda.mixin]
  render: () ->
    $c('form', {className: 'post__panel', }, [
        $c('label', {}, $c('input', {type: 'text', name: 'title'})),
        $c('label', {}, $c('input', {type: 'text', name: 'name'})),
        $c('label', {}, $c('input', {type: 'text', name: 'content'})),
        $c('input', {type: 'submit'}, '投稿'),
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
