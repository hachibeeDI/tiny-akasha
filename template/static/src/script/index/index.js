
let $c = React.createElement.bind(React);

import QuestionContext from './question-context';


var IndexComponent = React.createClass({
  mixins: [Arda.mixin],

  render: () => {
    return $c('div', {},
      $c('ul', {className: 'questions__ul'},
        this.props.questions.map((q) => {
          q['key'] = q['id'];
          return $c(QuestionContext, q);
        })
      )
    );
  }
});


export default class IndexContext extends Arda.Context {
  get component() {
    return IndexComponent;
  }

  delegate(subscribe) {
    super.delegate();
    var actions = require('./actions');
    subscribe('show:questions', (questions) => {
      this.update((s) => {questions: questions});
    });

    subscribe('question:show', (id) => {
      Navigator.navigate("/view/question/id/#{id}");
    });

    subscribe('question:delete', (id) => {
      actions.deleteQuestion(id)
        .then((data) => {
          this.update((s) => {
            let qs = s.questions.filter((q) => { q.id != id});
            return {questions: qs};
          });
        })
        .catch((err) => {
          console.log('question:delete has error')
          console.error(err)
        });
    });
  }

  initState(props) {
    // TODO: これはAPIのクエリ側で適切にソートして対処するように修正
    var questions = props['questions'] || [];
    questions.reverse();
    return {questions: questions};
  }

  expandComponentProps(props, state) {
    return {questions: state['questions']};
  }
}

