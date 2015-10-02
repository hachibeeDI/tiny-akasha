
import request from 'superagent';

import Arda from 'arda';

const Component = React.createClass({
  mixins: [Arda.mixin, React.addons.LinkedStateMixin],
  getInitialState: ()  => {
    return {name: '', title: '', content: ''};
  },

  close: (ev) => {
    this.dispatch('questions:close');
  },

  reloadQuestion: () => {
    request
      .get('/api/v1/question')
      .end((err, res) => {
        console.log('questions:reload occurd', res.text);
        if (data.error) { console.error(data.error); return; }
        this.dispatch('questions:reload');
      });
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
      this.reloadQuestion();
    });
  },

  render: () => {
    let mdPreview = require('../markdown-previewer/component');
    let post__panel = (
      <div className='post-panel'>
        <button onClick={this.close.bind(this)} className="close__button"></button>
        <form className='post-panel__form' onSubmit={this.postQuestion.bind(this)}>
          <label className='label--row'>
            'タイトル'
            <input type='text' name='title' valueLink={this.linkState('title')} />
          </label>
          <label className='label--row'>
            '名前'
            <input type='text' name='name' valueLink={this.linkState('name')} />
          </label>
          <label className='label--row'>
            '内容'
            <textarea className='post-panel__form__content' name='content' valueLink={this.linkState('content')} />
          </label>
          <input type='submit' value='投稿' />
          <mdPreview addtionalClass='post-panel__preview' content={this.state.content} />
        </form>
      </div>
    );
    return (<div className='post-panel--dark__cover'><post__panel /></div>);
  }
});
