$c = React.createElement.bind(React)
IndexContext = require '../index/index'


###
投稿用パネルの各項目などを管理する
###
class PostPanelContext extends Arda.Context
  delegate: (subscribe) ->
    super
    subscribe 'questions:reload', () =>
      $.get('/api/v1/question')
        .then (data) =>
          console.log 'questions:reload occurd', data
          unless data.error?
            Routers.main.replaceContext(IndexContext, data)

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
        .then (data) =>
          console.log '/api/v1/question returns', data
          @dispatch 'questions:reload'

        .then (data) ->
          Routers.post.popContext()

    render: () ->
      post__panel = $c('div', {className: 'post__panel', }, [
        $c('button', {onClick: @close, className: "close__button"}, ''),
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
  getInitialState: () ->
    return searchWord: ''

  showPostPanel: () ->
    Routers.post.pushContext(PostPanelContext, {})

  seachQuestionsByWord: (ev) ->
    # TODO: submitイベントを捕まえたほうが良いかも
    return unless ev.keyCode == 13
    $.post('/api/v1/question/search', {'word': @state.searchWord})
     .done (data) =>
       console.log '/api/v1/question/search', data
       unless data.error?
         @dispatch 'search:questions', data['questions']

  render: () ->
    $c('nav', {className: 'controll-panel'},
      $c('button', {onClick: @showPostPanel, className: 'button__open-post octicon-pencil'}),
      $c('input', {
        className: 'search-box',
        type: 'text',
        placeholder: 'search',
        onKeyDown: @seachQuestionsByWord,
        valueLink: @linkState('searchWord'),
        }
      ),
    )
)


class PostContext extends Arda.Context
  delegate: (subscribe) ->
    super
    subscribe 'search:questions', (questions) =>
      Routers.main.pushContext(IndexContext, {questions: questions})


  component: PostFrontComponent

  initState: (props) ->
    return {}

  expandComponentProps: (props, state) ->
    return {}


module.exports = PostContext
