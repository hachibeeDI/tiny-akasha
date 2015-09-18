import Arda from 'arda'

import IndexContext from '../index/index'

var $c = React.createElement.bind(React)

/*
* 投稿用パネルの各項目などを管理する
*/
class PostPanelContext extends Arda.Context {
  delegate(subscribe) {
    super.delegate();
    subscribe('questions:reload', () => {
      $.get('/api/v1/question')
        .then((data) => {
          console.log('questions:reload occurd', data);
          if (data.error) {
            Routers.main.replaceContext(IndexContext, data);
          }
        });
    });
  }

  get component() {
    return React.createClass({
      mixins: [Arda.mixin, React.addons.LinkedStateMixin],
      getInitialState: ()  => {
        return {name: '', title: '', content: ''};
      },

      close: (ev) => {
        Routers.post.popContext()
      },

      postQuestion: (ev) => {
        ev.preventDefault();
        $.post('/api/v1/question', {
          'title': this.state.title,
          'name': this.state.name,
          'content': this.state.content
        })
          .then((data) => {
            console.log('/api/v1/question returns', data);
            this.dispatch('questions:reload');
          })
          .then((data) => {
            Routers.post.popContext();
          });
        },
      render: () => {
        let mdPreview = require('../markdown-previewer/component')
        let post__panel =
          $c('div', {className: 'post-panel', },
            $c('button', {onClick: this.close, className: "close__button"}, ''),
            $c('form', {className: 'post-panel__form', onSubmit: this.postQuestion},
              $c('label', {className: 'label--row'},
                  'タイトル', $c('input', {type: 'text', name: 'title', valueLink: this.linkState('title')})),
              $c('label', {className: 'label--row'},
                  '名前', $c('input', {type: 'text', name: 'name', valueLink: this.linkState('name')})),
              $c('label', {className: 'label--row'},
                  '内容', $c('textarea', {className: 'post-panel__form__content', name: 'content', valueLink: this.linkState('content')})),
              $c('input', {type: 'submit', value: '投稿', }),
              $c(mdPreview, {addtionalClass: 'post-panel__preview', content: this.state.content})
            )
        );
        return $c('div', { className: 'post-panel--dark__cover' }, post__panel);
      }
    });
  }

  initState(props) {
    return {};
  }

  expandComponentProps(props, state) {
    return {};
  }
}


/*
 * 投稿用パネルを開くためのボタン的な意義を持つ
*/
PostFrontComponent = React.createClass({
  mixins: [Arda.mixin, React.addons.LinkedStateMixin],
  getInitialState: () => {
    return {searchWord: ''};
  },

  showPostPanel: () => {
    Routers.post.pushContext(PostPanelContext, {});
  },

  seachQuestionsByWord: (ev) => {
    // TODO: submitイベントを捕まえたほうが良いかも
    if (ev.keyCode !== 13) {
      return;
    }
    $.post('/api/v1/question/search', {'word': this.state.searchWord})
     .done((data) => {
       console.log('/api/v1/question/search', data);
       if (!data.error) {
         this.dispatch('search:questions', data['questions']);
       }
      });
  },

  render: () => {
    return $c('nav', {className: 'controll-panel'},
      $c('button', {onClick: this.showPostPanel, className: 'button__open-post octicon-pencil'}),
      $c('input', {
        className: 'search-box',
        type: 'text',
        placeholder: 'search',
        onKeyDown: this.seachQuestionsByWord,
        valueLink: this.linkState('searchWord'),
      })
    )
  }
});


export default class PostContext extends Arda.Context {
  delegate(subscribe) {
    super.delegate();
    subscribe('search:questions', (questions) => {
      Routers.main.pushContext(IndexContext, {questions: questions});
    });
  }

  get component() {return PostFrontComponent; }

  initState(props) {
    return {};
  }

  expandComponentProps(props, state) {
    return {};
  }
}

