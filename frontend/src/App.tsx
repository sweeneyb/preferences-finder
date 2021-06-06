import React from 'react';
// import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <h2 style={ {'textAlign': "center"}}>Company Logo</h2>

      <div className="card center rounded">
      <div style={{'position': "relative"}}>
      <img src="/images/canvas-oil-La-Grande-Jatte-Georges-Seurat-1884.jpg" alt="Work" className="rounded" style={{"width": "100%"}} />
      <div className="bottom-left rounded">
        <h4>
          <b className="rounded" >Georges Seurat</b>
        </h4>
        <p>A Sunday on La Grande Jatte-1884</p>
        <a href="https://www.britannica.com/biography/Georges-Seurat/images-videos#/media/1/536352/94613">Source</a>
      </div>
      </div>

      <div className="center" style={{"textAlign": "center"}}>
        <button className="button rounded no" style={{"backgroundColor": "red"}} >
          <div className="buttonText">Not for me</div>
        </button>
        <button className="button rounded yes" style={{"backgroundColor": "green"}}>
          <div className="buttonText">I like this</div>
        </button>
      </div>
    </div>

    </div>
  );
}

export default App;
