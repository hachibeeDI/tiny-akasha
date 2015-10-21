
import axios from 'axios';

import Arda from 'arda';


const Component = React.createClass({
  mixins: [Arda.mixin, React.addons.LinkedStateMixin],
  getInitialState() {
    return {name: '', title: '', content: ''};
  },

  close(ev) {
    this.dispatch('questions:close');
  },

  reloadQuestion() {
    axios
      .get('/api/v1/question')
      .then((res) => {
        if (res.data.error) {
          console.error(res.data.error);
          return Promise.reject(res.data.error);
        }
        this.dispatch('questions:reload', res.data);
      })
      .catch((err) => {
        // TODO: emit error
        console.error('questions:reload occurd', err.message);
      });
  },

  postQuestion(ev) {
    ev.preventDefault();
    axios
      .post('/api/v1/question', {
        'title': this.state.title,
        'name': this.state.name,
        'content': this.state.content
      })
      .then((data) => {
        console.log('/api/v1/question returns', data);
        this.reloadQuestion();
      });
  },

  render() {
    // console.log('post component render');
    let MarkdownPreviewerComponent = require('../markdown-previewer/component');
    return (
      <div className='post-panel--dark__cover'>
        <div className='post-panel'>
          <button onClick={this.close} className="close__button"></button>
          <form className='post-panel__form' onSubmit={this.postQuestion}>
            <label className='label--row'>
              タイトル
              <input type='text' name='title' valueLink={this.linkState('title')} />
            </label>
            <label className='label--row'>
              名前
              <input type='text' name='name' valueLink={this.linkState('name')} />
            </label>
            <label className='label--row'>
              内容
              <textarea
                className='post-panel__form__content' name='content'
                valueLink={this.linkState('content')}
              />
            </label>
            <input type='submit' value='投稿' />
            <MarkdownPreviewerComponent addtionalClass='post-panel__preview' content={this.state.content} />
          </form>
        </div>
      </div>
    );
  }
});
export default Component;
