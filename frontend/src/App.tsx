import React, { useState } from 'react';
import './App.css';
import { Card } from './components/Card'


const works = [
  {
    title: "A Sunday on La Grande Jatte-1884",
    artist: "Georges Seurat",
    citation: "https://www.britannica.com/biography/Georges-Seurat/images-videos#/media/1/536352/94613",
    image: "/images/canvas-oil-La-Grande-Jatte-Georges-Seurat-1884.jpg"
  },
  {
    title: "Canyon de Chelly",
    artist: "Ansel Adams",
    citation: "https://www.britannica.com/biography/Ansel-Adams-American-photographer/images-videos#/media/1/5091/97348",
    image: "/images/photograph-Canyon-de-Chelly-Ansel-Adams.jpg"
  },
  {
    title: "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
    artist: "Vincent van Gogh",
    citation: "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
    image: "/images/Fishing-Boats-on-the-Beach-oil-canvas-1888.jpg"
  }

]

function getRandomInt(max: number) {
  return Math.floor(Math.random() * max);
}

function sendMessage(rating: number, user: string|null, image: string) {
  if (user === null) {
    console.log("no user; not logging")
    return;
  }
  fetch('/preference-finder/us-central1/api/api/v1/works/rate', {
    method: 'POST',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      "user": user,
      "rating": rating,
      "image": image
    })
  });
}



function App() {
  const [work, setWork] = useState(works[getRandomInt(3)])

  const params = new URLSearchParams(window.location.search)
  const user = params.get('user')

  function dislike() {
    console.log("dislike")
    sendMessage( 0, user, work.image)
    setWork(works[getRandomInt(3)])
  }

  function like() {
    console.log("like")
    sendMessage(5, user, work.image)
    setWork(works[getRandomInt(3)])
  }

  



  return (
    <div className="App">
      <h2 style={{ 'textAlign': "center" }}>Company Logo</h2>

      <Card cardProps={work} cardCallbacks={{ onDislike: dislike, onLike: like }} />

    </div>
  );
}

export default App;
