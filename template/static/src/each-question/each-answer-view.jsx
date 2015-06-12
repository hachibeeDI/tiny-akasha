module.exports = function(component) {
  return (
    <li key={component.props.id} className="answer">
      <button className="answer__delete close__button" onClick={component.delete}></button>
      <h4 className="answer-user" > {component.props.username}</h4>
      <pre className="answer-content md-render-area" >
        {component.state.render}
      </pre>
    </li>
  );
};
