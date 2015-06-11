module.exports = function(component) {
  return (
    <li className="answer">
      <button className="answer__delete close__button" onClick={component.delete}></button>
      <h4 className="answer-user" > {component.props.username}</h4>
      <pre className="answer-content" >
        {component.props.content}
      </pre>
    </li>
  );
};
