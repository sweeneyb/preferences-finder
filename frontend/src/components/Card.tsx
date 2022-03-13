import React from 'react';
import '../App.css';

export type CardProps = {
  title: string,
  artist: string,
  image: string,
  citation: string,
}


export type CardCallbacks = {
  onLike: Function,
  onDislike: Function
}


export function Card (props: {cardProps: CardProps|undefined, cardCallbacks: CardCallbacks}) {
  if (props.cardProps === undefined) {
    return <div className="card center rounded">Card is undefined</div>
  } else {
  return (
      
      <div className="card center rounded">
      <div style={{'position': "relative"}}>
      <img src={props.cardProps.image} alt="Work" className="rounded" style={{"width": "100%"}} />
      <div className="bottom-left rounded">
        <h4>
          <b className="rounded" >{props.cardProps.artist}</b>
        </h4>
        <p>{props.cardProps.title}</p>
        <a href={props.cardProps.citation}>Source</a>
      </div>
      </div>

      <div className="center" style={{"textAlign": "center"}}>
        <button className="button rounded no" style={{"backgroundColor": "#af4154"}} onClick={ () => props.cardCallbacks.onDislike()}>
          <div className="buttonText" >Not for me</div>
        </button>
        <button className="button rounded yes" style={{"backgroundColor": "#019875"}} onClick={ () => props.cardCallbacks.onLike()} >
          <div className="buttonText">I like this</div>
        </button>
      </div>
    </div>

  );
}
}
