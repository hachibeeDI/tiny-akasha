'use strict';

import React from 'react'
import Promise from 'bluebird'
import Arda from 'arda'
import $ from 'jquery'
import _ from 'lodash'
import Grapnel from 'grapnel'

import PostContext from './post-question/index'

var Routers = {}


window.Navigator = new Grapnel({pushState: true})

document.addEventListener('DOMContentLoaded', () => {
  Routers.post = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-post-question'))
  Routers.post.pushContext(PostContext, {})

  Routers.main = new Arda.Router(Arda.DefaultLayout, document.getElementById('app-main'))
  let IndexContext = require('./index/index')

  // routing
  (() => {
    var showRoot = () => {
      $.get('/api/v1/question')
       .then((data) => {
          if (!data.error) {
            Routers.main.pushContext(IndexContext, data);
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
      Actions = require('./index/actions');
      Actions.showQuestion(req.params.id);
    });
  })();
});
