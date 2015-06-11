
loadQuestionData = (id) ->
  Promise.all([
    $.get("/api/v1/question/id/#{id}"),
    $.get("/api/v1/question/id/#{id}/answer"),
  ])


module.exports = {
  showQuestion: (id) ->
    loadQuestionData(id)
      .then (data) =>
        console.log 'question:show occurd', data
        Routers.main.pushContext(
          require('../each-question/context'),
          _.merge(data[0], data[1])
        )
      .catch (error) ->
        console.error 'each question', error

  reloadQuestion: (id) ->
    loadQuestionData(id)
      .then (data) =>
        console.log 'question:reload occurd', data
        Routers.main.replaceContext(
          require('../each-question/context'),
          _.merge(data[0], data[1])
        )
      .catch (error) ->
        console.error 'each question', error
}
