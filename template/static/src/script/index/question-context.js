
import Arda from 'arda';
import axios from 'axios';

import ArdaActionCreator from '../utils/action.js';


class QuestionAction extends ArdaActionCreator {

  showQuestion(id) {
    this.dispatch('question:show', id);
  }

  deleteQuestion(id) {
    this.dispatch('question:delete', id);
    axios
      .delete(`/api/v1/question/id/${id}`)
      .then((res) => {
        this.dispatch('question:delete', id);
      })
      .catch((err) => {
        console.error(err);
      });
  }
}


/*
  props:
    id: number
    title: string
    username: string
    content: string
 */
const QuestionComponent = React.createClass({
  mixins: [Arda.mixin],
  componentWillMount() {
    this.action = new QuestionAction(this);
    console.log(this);
  },

  onInnerClick() {
    console.log(this);
    this.action.showQuestion(this.props.id);
  },

  onDeleteClick() {
    this.action.deleteQuestion(this.props.id);
  },

  render() {
    console.log('QuestionComponent render', this);
    // TODO: 削除ボタンのデザインや仕様は要検討
    return (
      <li className="question" key={this.props.id}>
        <div className='question__inner' onClick={this.onInnerClick}>
          <h3 className='question__title'>{this.props.title}</h3>
          <p className='question__digest'>{this.props.content}</p>
        </div>
        <button className='question__delete-button' onClick={this.onDeleteClick}>
          この質問を削除する
        </button>
      </li>
    );
  }
});

export default QuestionComponent;
