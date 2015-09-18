import Arda from 'arda'
import request from 'superagent-bluebird-promise';
import md2react from 'md2react';

/* 
 * this.param {string} txt
 * this.return {Array}
 **/
_renderMd = (txt) => {
  try {
    return md2react(txt, {
      gfm: true,
      breaks: true,
      tables: true
    });
  }
  catch (e) {
    console.warn('mark down parse error', e);
    return [];
  }
}


EachQuestionComponent = React.createClass({
  mixins: [Arda.mixin],

  getInitialState: () => {
    preview: []
  },

  goBack: (ev) => {
    if (Routers.main.history.length <= 0) {
      return ;
    }
    global.history.back();
    Routers.main.popContext();
  },

  onHandleAnswerFormSubmit: (ev) => {
    ev.preventDefault();

    let username = React.findDOMNode(this.refs.form__user)
    let content = React.findDOMNode(this.refs.form__content)
    if (username == '' || content == '') { return ; }
    request
      .post("/api/v1/question/id///{this.props.id}/answer")
      .send({name: username.value, content: content.value})
      .set('Accept', 'application/json')
      .then((data) => {
        console.log('question created');
        username.value = '';
        content.value = '';
        this.setState({preview: []});
        this.dispatch('question:reload', this.props.id);
      })
      .catch((err) => {
        console.error(err);
      });
  },

  renderPreviewMd: (ev) => {
    this.setState({preview: _renderMd(this.refs.form__content.getDOMNode().value)});
  },

  render: () => {
    console.log('each question component render', this.props);
    let template = require('./view.jsx');
    return template(this);
  }
});


const Actions = require('../index/actions');

/*
props:
  id: number
  title: string
  username: string
  content: string
  answers: Array{string}
*/
export default class EachQuestionContext extends Arda.Context {
  get component() {
    return EachQuestionComponent;
  }
  // initState: (props) => {
  //   console.log 'each question init'
  //   return {question: [], }

  delegate(subscribe) {
    super.delegate();

    subscribe('question:reload', (id) => {
      Actions.reloadQuestion(id);
    });

    subscribe('answer:delete', (id) => {
      request
        .del("/api/v1/answer/id///{id}")
        .end((err, res) => {
          if (err) {
            console.error(err);
            return ;
          }
          console.log(res);
          Actions.reloadQuestion(this.props.id);
        });
    });
  }

  expandComponentProps(props, state) {
    console.log('each question expand', props, state);
    props.content = _renderMd(props.content);
    return props;
  }
}


