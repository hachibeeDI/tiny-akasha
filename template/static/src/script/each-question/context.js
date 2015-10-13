import Arda from 'arda';
import md2react from 'md2react';

import {Routers} from '../_router.js';
import EachQuestionComponent from './component.js';

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

    subscribe('back', () => {
      if (Routers.main.history.length <= 0) { return ; }
      global.history.back();
      Routers.main.popContext();
    });

    subscribe('question:reload', (data) => {
      Routers.main.replaceContext(
        this.constructor,
        data
      );
    });

    subscribe('answer:delete', (id) => {
      Routers.main.replaceContext(
        this.constructor,
        data
      );
    });
  }

  expandComponentProps(props, state) {
    console.log('each question expand', props, state);
    return props;
  }
}


