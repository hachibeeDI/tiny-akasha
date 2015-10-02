
import request from 'superagent';

import ArdaActionCreator from '../utils/action.js';


let loadQuestionData = (id) => {
  return Promise.all([
    $.get(`/api/v1/question/id/${id}`),
    $.get(`/api/v1/question/id/${id}/answer`),
  ]);
};


export default class Actions extends ArdaActionCreator {
  goBack(ev) {
    this.dispatch('back');
  }

  sendAnswer(username, content) {
    // FIXME: DOMいじってるのイクない
    if (username == '' || content == '') { return ; }
    request
      .post(`/api/v1/question/id/${this.props.id}/answer`)
      .send({name: username.value, content: content.value})
      .set('Accept', 'application/json')
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
    request
      .del(`/api/v1/answer/id/${id}`)
      .end((err, res) => {
        if (err) { console.error(err); return ; }

        console.log(res);
        loadQuestionData(id)
          .then((data) => {
            console.log('question:reload occurd', data);
            this.dispatch('answer:delete', _.merge(data[0], data[1]));
          })
          .catch((error) => {
            console.error('delete anser in load question', error);
          });
      });
  }
}


