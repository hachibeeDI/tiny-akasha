module.exports = function(component) {
  return (
    <li key={component.props.id} className="answer">
      <div className="answer__content" >
        <h4 className="answer__user" >
          {component.props.username}
          <button className="answer__delete close__button" onClick={component.delete}></button>
        </h4>
        <pre className="answer__text md-render-area" >
          {component.state.render}
        </pre>
      </div>
    </li>
  );
};
