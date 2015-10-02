import Arda from 'arda'

let $c = React.createElement.bind(React);

import QuestionComponent from './question-context';
import actions from './actions';


var IndexComponent = React.createClass({
  mixins: [Arda.mixin],

  render: () => {
    return $c('div', {},
      $c('ul', {className: 'questions__ul'},
        this.props.questions.map((q) => {
          q['key'] = q['id'];
          return $c(QuestionComponent, q);
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
    subscribe('show:questions', (questions) => {
      this.update((s) => {questions: questions});
    });

    subscribe('question:loaded', (datas) => {
      Routers.main.pushContext(
        require('../each-question/context'),
        datas
      );
    });

    subscribe('question:reload', (datas) => {
      Routers.main.replaceContext(
        require('../each-question/context'),
        datas
      );
    });

    subscribe('question:show', (id) => {
      Navigator.navigate(`/view/question/id/${id}`);
    });

    subscribe('question:delete', (id) => {
      this.update((s) => {
        let qs = s.questions.filter((q) => { q.id != id});
        return {questions: qs};
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

