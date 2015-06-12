
module.exports = function (render, addtionalClass) {
  var classNames = 'md-render-area ' + addtionalClass;
  return (
    <div className={classNames}>
      {render}
    </div>
  );
};
