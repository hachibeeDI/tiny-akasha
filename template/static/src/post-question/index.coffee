$c = React.createElement.bind(React)


###
投稿用パネルの各項目などを管理する
###
class PostPanelContext extends Arda.Context
  component: React.createClass(
    mixins: [Arda.mixin, React.addons.LinkedStateMixin]
    getInitialState: ()  ->
      return name: '', title: '', content: ''

    close: (ev) ->
      Routers.post.popContext()

    postQuestion: (ev) ->
      ev.preventDefault()
      $.post('/api/v1/question', {
        'title': @state.title
        'name': @state.name
        'content': @state.content
      })
      Routers.post.popContext()

    render: () ->
      post__panel = $c('div', {className: 'post__panel', }, [
        $c('button', {onClick: @close, className: "icon-font close__button"}, ''),
        $c('form', {className: 'post__panel__form', onSubmit: @postQuestion}, [
            $c('label', {className: 'label--row'},
                ['タイトル', $c('input', {type: 'text', name: 'title', valueLink: @linkState('title')})]),
            $c('label', {className: 'label--row'},
                ['名前', $c('input', {type: 'text', name: 'name', valueLink: @linkState('name')})]),
            $c('label', {className: 'label--row'},
                ['内容', $c('textarea', {name: 'content', valueLink: @linkState('content')})]),
            $c('input', {type: 'submit', value: '投稿', }),
          ]
        )]
      )
      return $c('div', {
        className: 'post__panel--dark__cover'
        onClick: @close
      }, post__panel)
  )

  initState: (props) ->
    return {}

  expandComponentProps: (props, state) ->
    return {}


###
投稿用パネルを開くためのボタン的な意義を持つ
###
PostFrontComponent = React.createClass(
  mixins: [Arda.mixin, React.addons.LinkedStateMixin]
  showPostPanel: () ->
    Routers.post.pushContext(PostPanelContext, {})

  render: () ->
    $c('button', {
        onClick: @showPostPanel,
        className: 'icon-font octicon-pencil button__open-post'
      }, '')
)


class PostContext extends Arda.Context
  delegate: (subscribe) ->
    super
    subscribe 'search:questions', (questions) =>
      IndexContext = require '../index/index'
      Routers.main.pushContext(IndexContext, {questions: questions})

  component: PostFrontComponent

  initState: (props) ->
    return {}

  expandComponentProps: (props, state) ->
    return {}


module.exports = PostContext
