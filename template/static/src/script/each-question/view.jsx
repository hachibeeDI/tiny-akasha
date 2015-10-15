import throttle from 'lodash/function/throttle';

const eachAnswer = require("./each-answer-component");
import mdPreview from '../markdown-previewer/component.js';

module.exports = function(component) {
  const answers = component.props.answers.map(function(ans) {
    return React.createElement(eachAnswer, ans);
  });
  return (
    <div className="each-question">
      <button className="each-question__back-button" onClick={component.goBack.bind(component)}>戻る</button>
      <h2 className="each-question__title">{component.props.title}</h2>
      <div className="each-question__content" >
        <h4 className="each-question__posted-user">{component.props.username}</h4>
        <pre className="each-question__text md-render-area" >
          {component.props.content}
        </pre>
      </div>
      <h3 className="answers-section__header">回答</h3>
      <ul className="answers">
        {answers}
      </ul>
      <form className="answer-form" onSubmit={component.onHandleAnswerFormSubmit}>
        <input type='text' placeholder='your name' ref='form__user' className="answer-form__name" />
        <textarea
          ref='form__content'
          className="answer-form__content"
          onChange={throttle(component.renderPreviewMd.bind(component), 200)}>
        </textarea>
        <mdPreview content={component.state.preview} addtionalClass='post-panel__preview' />
        <button type='submit'>投稿</button>
      </form>
    </div>
  );
};
