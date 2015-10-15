'use strict';

import {EventEmitter} from 'events';

import axios from 'axios';
import _ from 'lodash';

import {Routers, Navigator} from './_router.js';
global.Routers = Routers;
global.Navigator = Navigator;


import IndexContext from './index/index.js';
import EachQuestionContext from './each-question/context.js';


document.addEventListener('DOMContentLoaded', () => {

  // routing
  (() => {
    var showRoot = () => {
      axios
        .get('/api/v1/question')
        .then((res) => {
          console.log('init index', res);
          let data = res.data;
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
      // TODO: このdispatcherどうしよう
      const dispatcher = new EventEmitter();
      dispatcher.on('question:loaded', (data) => {
        Routers.main.pushContext(EachQuestionContext, data);
      });
      dispatcher.dispatch = dispatcher.emit.bind(dispatcher);
      const IndexActions = require('./index/actions');
      const act = new IndexActions(dispatcher);
      act.showQuestion(req.params.id);
    });
  })();
});
