import React from 'react'
import ReactDOM from 'react-dom'

import { BrowserRouter, HashRouter } from 'react-router-dom'

import App from './components/App'

// Since we are using HtmlWebpackPlugin WITHOUT a template, we should create our own root node in the body element before rendering into it
let root = document.createElement('div')
root.id = "root"
document.body.appendChild( root )

ReactDOM.render((
  <HashRouter>
    <App />
  </HashRouter>
), document.getElementById("root"))
