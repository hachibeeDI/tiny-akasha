import axios from 'axios';
import merge from 'lodash/object/merge';

import ArdaActionCreator from '../utils/action.js';


let loadQuestionData = (id) => {
  return Promise.all([
    axios.get(`/api/v1/question/id/${id}`),
    axios.get(`/api/v1/question/id/${id}/answer`),
  ]);
};


export default class Actions extends ArdaActionCreator {
  showQuestion(id) {
    loadQuestionData(id)
      .then((data) => {
        console.log('question:show occurd', data);
        this.dispatch('question:loaded', merge(data[0], data[1]));
      })
      .catch((error) => {
        console.error('each question', error);
      });
  }

  reloadQuestion(id) {
    loadQuestionData(id)
      .then((data) => {
        console.log('question:reload occurd', data);
        this.dispatch('question:reload', merge(data[0], data[1]));
      })
      .catch((error) => {
        console.error('each question', error);
      });
  }

  deleteQuestion(id) {
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
