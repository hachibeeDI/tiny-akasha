'use strict';

import axios from 'axios';
import _ from 'lodash';

import {Routers, Navigator} from './_router.js';
global.Routers = Routers;
global.Navigator = Navigator;


document.addEventListener('DOMContentLoaded', () => {
  let IndexContext = require('./index/index.js');

  // routing
  (() => {
    var showRoot = () => {
      axios
        .get('/api/v1/question')
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
