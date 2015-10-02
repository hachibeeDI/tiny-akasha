
import axios from 'axios';

import ArdaActionCreator from '../utils/action.js';


let loadQuestionData = (id) => {
  return axios.all([
    axios.get(`/api/v1/question/id/${id}`),
    axios.get(`/api/v1/question/id/${id}/answer`),
  ]);
};


export default class Actions extends ArdaActionCreator {
  goBack(ev) {
    this.dispatch('back');
  }

  sendAnswer(username, content) {
    // FIXME: DOMいじってるのイクない
    if (username == '' || content == '') { return ; }
    axios
      .post(`/api/v1/question/id/${this.props.id}/answer`, {name: username.value, content: content.value})
      .then((data) => {
        console.log('question created');
        username.value = '';
        content.value = '';
        this.setState({preview: ''});
        this.reloadQuestion();
      })
      .catch((err) => {
        console.error(err);
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

  deleteAnswer(answerId) {
    axios
      .delete(`/api/v1/answer/id/${id}`)
      .then((res) => {
        console.log(res);
        return loadQuestionData(id)
          .then((data) => {
            console.log('question:reload occurd', data);
            this.dispatch('answer:delete', _.merge(data[0], data[1]));
          });
      })
      .catch((err) => {
        console.error(err);
      });
  }
}


