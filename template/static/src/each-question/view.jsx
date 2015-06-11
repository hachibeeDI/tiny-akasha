var eachAnswer = require("./each-answer-component");

module.exports = function(component) {
  var answers = component.props.answers.map(function(ans) {
    return React.createElement(eachAnswer, ans);
  });
  return (
    <div className="each-question">
      <button className="each-question__back-button" onClick={component.goBack}>戻る</button>
      <h2 className="each-question__title">{component.props.title}</h2>
      <p className="each-question__posted-user">{component.props.username}</p>
      <p className="each-question__content" >{component.props.content}</p>
      <h3 className="answers-section__header">回答</h3>
      <ul className="answers"> {answers} </ul>
      <form className="answer-form" onSubmit={component.onHandleAnswerFormSubmit}>
        <input type='text' placeholder='your name' name='user' className="answer-form__name" />
        <textarea name='content' className="answer-form__content"></textarea>
        <button type='submit'>投稿</button>
      </form>
    </div>
  );
};
