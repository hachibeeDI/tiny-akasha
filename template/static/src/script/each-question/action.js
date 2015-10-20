
import axios from 'axios';
import merge from 'lodash/object/merge';

import ArdaActionCreator from '../utils/action.js';


let loadQuestionData = (id) => {
  return axios.all([
    axios.get(`/api/v1/question/id/${id}`),
    axios.get(`/api/v1/question/id/${id}/answer`),
  ])
  .then((res) => merge(res[0].data, res[1].data));
};


export default class Actions extends ArdaActionCreator {
  goBack(ev) {
    this.dispatch('back');
  }

  sendAnswer(id, username, content) {
    if (username == '' || content == '') {
      return ;
    }
    return axios
      .post(`/api/v1/question/id/${id}/answer`, {name: username, content: content})
      .then((res) => {
        console.log('question created');
        return res;
      });
  }

  reloadQuestion(id) {
    loadQuestionData(id)
      .then((data) => {
        console.log('question:reload occurd', data);
        this.dispatch('question:reload', data);
      })
      .catch((error) => {
        console.error('each question', error);
      });
  }

  deleteAnswer(answerId) {
    axios
      .delete(`/api/v1/answer/id/${id}`)
      .then((deleteRes) => {
        console.log(deleteRes);
        return loadQuestionData(id)
          .then((loadData) => {
            console.log('question:reload occurd', loadData);
            this.dispatch('answer:delete', loadData);
          });
      })
      .catch((err) => {
        console.error(err);
      });
  }
}


