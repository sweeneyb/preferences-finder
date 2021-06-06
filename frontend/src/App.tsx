import React from 'react';
// import logo from './logo.svg';
import './App.css';
import {Card} from './components/Card'


const work = {title: "A Sunday on La Grande Jatte-1884",
              artist: "Georges Seurat", 
              citation: "https://www.britannica.com/biography/Georges-Seurat/images-videos#/media/1/536352/94613"}
function App() {
  return (
    <div className="App">
      <h2 style={ {'textAlign': "center"}}>Company Logo</h2>

      <Card {...work} />

    </div>
  );
}

export default App;
