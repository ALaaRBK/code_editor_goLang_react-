import React, { Component } from 'react';
import './App.css';
import NavbarHeader from './componants/navbar'
import NewCode from './containers/newCode';
class App extends Component {
  render() {
    return (
      <div className="App">
        <NavbarHeader />
        <NewCode/>
      </div>
    );
  }
}

export default App;
