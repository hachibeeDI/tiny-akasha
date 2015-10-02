
import Arda from 'arda';

import ArdaActionCreator from '../utils/action.js';


class QuestionAction extends ArdaActionCreator {

  showQuestion(id) {
    this.dispatch('question:show', id);
  }

  deleteQuestion(id) {
    this.dispatch('question:delete', id);
    request
      .del(`/api/v1/question/id/${id}`)
      .end((err, res) => {
        if (err) { console.error(err); return; }
        this.dispatch('question:delete', id);
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
  componentDidMount: () => {
    this.action = new QuestionAction(this);
  },

  onInnerClick: () => {
    this.action.showQuestion(this.props.id);
  },

  onDeleteClick: () => {
    this.action.deleteQuestion(this.props.id);
  },

  render: () => {
    // TODO: 削除ボタンのデザインや仕様は要検討
    return (
      <li className="question" key={this.props.id}>
        <div className='question__inner' onClick={this.onInnerClick.bind(this)}>
          <h3 className='question__title'>this.props.title</h3>
          <p className='question__digest'>this.props.content</p>
        </div>
        <button className='question__delete-button' onClick={this.onDeleteClick.bind(this)}> 'この質問を削除する'</button>
      </li>
    );
  }
});

export default QuestionComponent;
