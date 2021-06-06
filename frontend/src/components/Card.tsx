import React from 'react';
import '../App.css';

export type CardProps = {
  title: string,
  artist: string,
  citation: string,
}





export function Card (Props: CardProps) {
  return (
   
      <div className="card center rounded">
      <div style={{'position': "relative"}}>
      <img src="/images/canvas-oil-La-Grande-Jatte-Georges-Seurat-1884.jpg" alt="Work" className="rounded" style={{"width": "100%"}} />
      <div className="bottom-left rounded">
        <h4>
          <b className="rounded" >{Props.artist}</b>
        </h4>
        <p>{Props.title}</p>
        <a href={Props.citation}>Source</a>
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

  );
}
