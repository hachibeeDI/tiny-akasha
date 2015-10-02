import request from 'superagent';

import ArdaActionCreator from '../utils/action.js';


let loadQuestionData = (id) => {
  return Promise.all([
    $.get(`/api/v1/question/id/${id}`),
    $.get(`/api/v1/question/id/${id}/answer`),
  ]);
};


export default class Actions extends ArdaActionCreator {
  showQuestion(id) {
    loadQuestionData(id)
      .then((data) => {
        console.log('question:show occurd', data);
        this.dispatch('question:loaded', _.merge(data[0], data[1]));
      })
      .catch((error) => {
        console.error('each question', error);
      });
  }

  reloadQuestion(id) {
    loadQuestionData(id)
      .then((data) => {
        console.log('question:reload occurd', data);
        this.dispatch('question:reload', _.merge(data[0], data[1]));
      })
      .catch((error) => {
        console.error('each question', error);
      });
  }

  deleteQuestion(id) {
    request
      .del(`/api/v1/question/id/${id}`)
      .end((err, res) => {
        if (err) { console.error(err); return; }
        this.dispatch('question:delete', id);
      });
  }
}
