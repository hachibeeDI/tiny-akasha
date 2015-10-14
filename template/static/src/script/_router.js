
import Arda from 'arda';
import Grapnel from 'grapnel';

import PostContext from './post-question/index.js';


class _Routers {
  constructor() {
    this.post = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-post-question'));
    this.post.pushContext(PostContext, {});

    this.main = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-main'));
  }
}




var Routers = new _Routers();
var Navigator = new Grapnel({pushState: true});


export {Routers, Navigator};
