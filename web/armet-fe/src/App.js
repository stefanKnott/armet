import logo from './logo.svg';
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
        {currentContext}
      </header>
      <ProSidebarProvider>
      <Layout />
      </ProSidebarProvider>
    </div>
  );
}

export default App;
