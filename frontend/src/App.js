import React from "react";
import "./App.css"
import HomePage from "./components/Home/HomePage";
import LoginForm from "./components/Login/LoginForm"
import { BrowserRouter, Route, Switch } from "react-router-dom";

const App = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={HomePage}></Route>
        <Route path="/login" component={LoginForm}></Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
