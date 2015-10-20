
import Arda from 'arda'

import Actions from './action.js';

let EachQuestionComponent = React.createClass({
  mixins: [Arda.mixin],

  getInitialState() {
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
    // TODO: error用のアクションを作るべし
    this.action.sendAnswer(this.props.id, username.value, content.value)
      .then(() => {
        username.value = '';
        content.value = '';
        this.setState({preview: ''});
        this.action.reloadQuestion(this.props.id);
      })
      .catch((err) => console.error(err));
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
