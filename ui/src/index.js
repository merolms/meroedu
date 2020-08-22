import React from 'react';
import ReactDOM from 'react-dom';
import * as serviceWorker from './serviceWorker';

import { Provider } from 'react-redux';
import { applyMiddleware, createStore } from "redux";
import reducers from './redux/reducers';
import App from "./App";
import { composeWithDevTools } from "redux-devtools-extension/developmentOnly";
import thunk from 'redux-thunk';
import 'semantic-ui-css/semantic.min.css';
import './index.scss';

const store = createStore(reducers, composeWithDevTools(applyMiddleware(thunk)));

ReactDOM.render(
    <Provider store={store}>
        <App/>
    </Provider>,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
