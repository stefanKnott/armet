import logo from './armet.png';
import './App.css';
import React, {useState, setState} from 'react';
import axios from 'axios';
import Layout from './components/Layout.jsx'
import { ProSidebarProvider } from 'react-pro-sidebar';

function App() {
  const [currentContext, setCurrentContext] = useState("");

  return (
    <div className="App">
      <header className="App-header">
      <img className="logo" src={logo} alt={"logo"}/> 
      <h2 className="title">armet </h2>
        {currentContext}
      </header>
      <ProSidebarProvider>
      <Layout />
      </ProSidebarProvider>
    </div>
  );
}

export default App;
