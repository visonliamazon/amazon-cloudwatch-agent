import React from "react";
import ReactDOM from "react-dom/client";
import { HashRouter, Route, Routes } from "react-router-dom";
import App from './containers/App'

//BrowserRouter,,Navigate
const root = ReactDOM.createRoot(document.getElementById("root"));

//This is where we setup our root component and the routes for the website

root.render(
  <App />,
);