
import Arda from 'arda';
import Grapnel from 'grapnel';

import PostContext from './post-question/index.js';
import LoginContext from './login-link/index.jsx';


class _Routers {
  constructor() {
    this.post = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-post-question'));
    this.post.pushContext(PostContext, {});
    this.loginLink = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-login-panel'));
    this.loginLink.pushContext(LoginContext, {});


    this.main = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-main'));
  }
}


const Routers = new _Routers();
const Navigator = new Grapnel({pushState: true});


export {Routers, Navigator};
