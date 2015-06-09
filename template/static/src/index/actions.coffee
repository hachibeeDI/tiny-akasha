

module.exports = {
  showQuestion: (id) ->
    Promise.all([
      $.get("/api/v1/question/id/#{id}"),
      $.get("/api/v1/question/id/#{id}/answer"),
    ])
      .then (data) =>
        console.log 'question:show occurd', data
        Routers.main.pushContext(
          require('../each-question/context'),
          _.merge(data[0], data[1])
        )
      .catch (error) ->
        console.error 'each question', error
}
