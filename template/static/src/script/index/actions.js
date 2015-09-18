import request from 'superagent-bluebird-promise';


let loadQuestionData = (id) => {
  return Promise.all([
    $.get("/api/v1/question/id/#{id}"),
    $.get("/api/v1/question/id/#{id}/answer"),
  ]);
};


export default Actions = {
  showQuestion: (id) => {
    loadQuestionData(id)
      .then((data) => {
        console.log('question:show occurd', data);
        Routers.main.pushContext(
          require('../each-question/context'),
          _.merge(data[0], data[1])
        );
      })
      .catch((error) => {
        console.error('each question', error);
      });
  },

  reloadQuestion: (id) => {
    loadQuestionData(id)
      .then((data) => {
        console.log('question:reload occurd', data);
        Routers.main.replaceContext(
          require('../each-question/context'),
          _.merge(data[0], data[1])
        );
      })
      .catch((error) => {
        console.error('each question', error);
      });
  },

  deleteQuestion: (id) => {
    return request
      .del("/api/v1/question/id/#{id}")
      .promise();
  }
}
