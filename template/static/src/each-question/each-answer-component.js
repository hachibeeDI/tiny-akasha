import md2react from 'md2react';


export default EachAnswerComponent = React.createClass({
  mixins: [Arda.mixin],

  getInitialState: () => {
    // 編集機能をつけるかもしれんのでstateに
    try {
      var renders = md2react(this.props.content, {
        gfm: true,
        breaks: true,
        tables: true
      });
      return {render: renders};
    }
    catch (e) {
      console.warn('mark down parse error', e);
      return {render: []};
    }
  },

  delete: (ev) => {
    this.dispatch('answer:delete', this.props.id);
  },

  render: () => {
    console.log('each-answer render ', this);
    var template = require('./each-answer-view.jsx');
    return template(this);
  }
});



