
import Arda from 'arda'


/*
  props:
    id: number
    title: string
    username: string
    content: string
 */
var QuestionComponent = React.createClass({
  mixins: [Arda.mixin],
  showQuestion: () => {
    this.dispatch('question:show', this.props.id);
  },

  deleteQuestion: () => {
    this.dispatch('question:delete', this.props.id);
  },

  render: () => {
    // TODO: 削除ボタンのデザインや仕様は要検討
    return (
      <li className="question" key={this.props.id}>
        <div className='question__inner' onClick={this.showQuestion}>
          <h3 className='question__title'>this.props.title</h3>
          <p className='question__digest'>this.props.content</p>
        </div>
        <button className='question__delete-button' onClick={this.deleteQuestion}> 'この質問を削除する'</button>
      </li>
    );
  }
});

export default QuestionComponent;
