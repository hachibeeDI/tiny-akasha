import Arda from 'arda';
import axios from 'axios';

import IndexContext from '../index/index';
import Component from './component.jsx';

import {Routers} from '../_router.js';


/*
* 投稿用パネルの各項目などを管理する
*/
class PostPanelContext extends Arda.Context {
  get component() {
    return Component;
  }

  delegate(subscribe) {
    super.delegate();
    subscribe('questions:reload', (data) => {
      Routers.main.replaceContext(IndexContext, data)
        .then((ctx) => {
          Routers.post.popContext();
        });
    });
    subscribe('questions:close', () => {
      Routers.post.popContext();
    });
  }

  initState(props) {
    return {};
  }

  expandComponentProps(props, state) {
    return {};
  }
}


const $c = React.createElement.bind(React);

/*
 * 投稿用パネルを開くためのボタン的な意義を持つ
*/
let PostFrontComponent = React.createClass({
  mixins: [Arda.mixin, React.addons.LinkedStateMixin],
  getInitialState() {
    return {searchWord: ''};
  },

  showPostPanel() {
    console.log('showPostPanel exec');
    Routers.post.pushContext(PostPanelContext, {});
  },

  seachQuestionsByWord(ev) {
    // TODO: submitイベントを捕まえたほうが良いかも
    if (ev.keyCode !== 13) {
      return;
    }
    axios
      .post('/api/v1/question/search', {'word': this.state.searchWord})
      .then((data) => {
        console.log('/api/v1/question/search', data);
        if (!data.error) {
          this.dispatch('search:questions', data['questions']);
        }
      })
      .catch((err) => {
        console.error(err);
      });
  },

  render() {
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
  get component() {
    return PostFrontComponent;
  }

  delegate(subscribe) {
    super.delegate();
    subscribe('search:questions', (questions) => {
      Routers.main.pushContext(IndexContext, {questions: questions});
    });
  }

  initState(props) {
    return {};
  }

  expandComponentProps(props, state) {
    return {};
  }
}

