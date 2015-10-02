import Arda from 'arda'

import Actions from './action.js';

import mdPreview from '../markdown-previewer/component.js';


let EachAnswerComponent = React.createClass({
  mixins: [Arda.mixin],

  getInitialState() {
    // 編集機能をつけるかもしれんのでstateに
    return {render: this.props.content};
  },

  componentDidMount() {
    this.action = new Actions(this);
  },

  delete(ev) {
    this.action.deleteAnswer(this.props.id);
  },

  render() {
    return (
      <li key={this.props.id} className="answer">
        <div className="answer__content" >
          <h4 className="answer__user" >
            {this.props.username}
            <button className="answer__delete close__button" onClick={this.delete.bind(this)}></button>
          </h4>
          <mdPreview content={this.state.render} addtionalClass="answer__text"/ >
        </div>
      </li>
    );
  }
});


export default EachAnswerComponent;
