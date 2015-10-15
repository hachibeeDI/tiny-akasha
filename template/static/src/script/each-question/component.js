
import Arda from 'arda'

import Actions from './action.js';

let EachQuestionComponent = React.createClass({
  mixins: [Arda.mixin],

  getInitialState(){
    return {preview: ''};
  },

  componentDidMount() {
    this.action = new Actions(this);
  },

  goBack(ev) {
    this.action.goBack();
  },

  onHandleAnswerFormSubmit(ev) {
    ev.preventDefault();
    let username = React.findDOMNode(this.refs.form__user);
    let content = React.findDOMNode(this.refs.form__content);
    this.action.sendAnswer(username, content);
  },

  renderPreviewMd(ev) {
    this.setState({preview: this.refs.form__content.getDOMNode().value});
  },

  render() {
    console.log('each question component render', this.props);
    let template = require('./view.jsx');
    return template(this);
  }
});


export default EachQuestionComponent;
