
import md2react from 'md2react';


/*
* this.param {string} txt
* this.return {Array}
*/
var _renderMd = (txt) => {
  try {
    return md2react(txt, {
      gfm: true,
      breaks: true,
      tables: true
    });
  }
  catch (e) {
    console.warn('mark down parse error', e);
    return [];
  }
};



export default MarkdownPreviewerComponent = React.createClass({
    mixins: [Arda.mixin],

    render: () => {
      console.log('MarkdownPreviewerComponent render', this.props);
      var template = require('./view.jsx');
      return template(_renderMd(this.props.content), this.props.addtionalClass);
    }
});
