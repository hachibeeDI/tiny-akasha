module.exports = function(component) {
  return (
    <li key={component.props.id} className="answer">
      <button className="answer__delete close__button" onClick={component.delete}></button>
      <h4 className="answer-user" > {component.props.username}</h4>
      <pre className="answer-content" >
        {component.props.content}
      </pre>
    </li>
  );
};
