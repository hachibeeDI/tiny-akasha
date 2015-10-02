'use strict';

import Arda from 'arda';
import $ from 'jquery';
import _ from 'lodash';
import Grapnel from 'grapnel';

import PostContext from './post-question/index.js';

window.Routers = {};
window.Navigator = new Grapnel({pushState: true});

document.addEventListener('DOMContentLoaded', () => {
  Routers.post = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-post-question'));
  Routers.post.pushContext(PostContext, {});

  Routers.main = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-main'));
  let IndexContext = require('./index/index.js');

  // routing
  (() => {
    var showRoot = () => {
      $.get('/api/v1/question')
       .then((data) => {
          if (!data.error) {
            Routers.main.pushContext(IndexContext, data);
          }
          else if (data.error === 'no data') {
            Routers.main.pushContext(IndexContext, {questions: []});
          }
          else {
            console.error(data.error);
          }
        });
    };

    Navigator.get('/', (req) => {
      showRoot();
    });
    Navigator.get('/view', (req) => {
      showRoot();
    });
    Navigator.get('/view/question/id/:id', (req) => {
      const Actions = require('./index/actions');
      Actions.showQuestion(req.params.id);
    });
  })();
});
