import Arda from 'arda'

let $c = React.createElement.bind(React);

import {Routers} from '../_router.js';
import QuestionComponent from './question-context';
import actions from './actions';


const IndexComponent = React.createClass({
  mixins: [Arda.mixin],

  render() {
    console.log('IndexComponent render', this);
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


import EachQuestionContext from '../each-question/context';


export default class IndexContext extends Arda.Context {
  get component() {
    return IndexComponent;
  }

  delegate(subscribe) {
    super.delegate();
    subscribe('show:questions', (questions) => {
      console.log('show:questions emitted');
      this.update((s) => {
        return {questions: questions};
      });
    });

    subscribe('question:loaded', (datas) => {
      console.log('question:loaded emitted');
      Routers.main.pushContext(
        EachQuestionContext,
        datas
      );
    });

    subscribe('question:reload', (datas) => {
      Routers.main.replaceContext(
        EachQuestionContext,
        datas
      );
    });

    subscribe('question:show', (id) => {
      Navigator.navigate(`/view/question/id/${id}`);
    });

    subscribe('question:delete', (id) => {
      this.update((s) => {
        let qs = s.questions.filter((q) => q.id !== id);
        return {questions: qs};
      });
    });
  }

  initState(props) {
    // TODO: これはAPIのクエリ側で適切にソートして対処するように修正
    let questions = props['questions'] || [];
    questions.reverse();
    return {questions: questions};
  }

  expandComponentProps(props, state) {
    return {questions: state['questions']};
  }
}

