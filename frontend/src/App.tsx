import React, { useEffect, useState } from 'react';
import './App.css';
import { Card } from './components/Card'

import { initializeApp } from "firebase/app";
import { collectionGroup, query, where, collection, doc, getDoc, getDocs, setDoc, getFirestore } from "firebase/firestore";
import { getStorage, ref, getDownloadURL } from "firebase/storage"

const firebaseConfig = {
  apiKey: process.env.REACT_APP_FIREBASE_API_KEY,
  authDomain: process.env.REACT_APP_FIREBASE_AUTH_DOMAIN,
  databaseURL: process.env.REACT_APP_FIREBASE_DATABASE_URL,
  projectId: process.env.REACT_APP_FIREBASE_PROJECT_ID,
  storageBucket: process.env.REACT_APP_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.REACT_APP_FIREBASE_SENDER_ID,
  appId: process.env.REACT_APP_FIREBASE_APP_ID
}
// Initialize Firebase
const app = initializeApp(firebaseConfig);
// Initialize Cloud Firestore and get a reference to the service
const db = getFirestore(app);
const storage = getStorage();

export type Work = {
  title: string,
  artist: string,
  image: string,
  citation: string,
  id: string,
  collection: string
}


const works: Work[] = [
  {
    title: "A Sunday on La Grande Jatte-1884",
    artist: "Georges Seurat",
    citation: "https://www.britannica.com/biography/Georges-Seurat/images-videos#/media/1/536352/94613",
    image: "/images/canvas-oil-La-Grande-Jatte-Georges-Seurat-1884.jpg",
    id: "1",
    collection: "default"
  },
  {
    title: "Canyon de Chelly",
    artist: "Ansel Adams",
    citation: "https://www.britannica.com/biography/Ansel-Adams-American-photographer/images-videos#/media/1/5091/97348",
    image: "/images/photograph-Canyon-de-Chelly-Ansel-Adams.jpg",
    id: "2",
    collection: "default"
  },
  {
    title: "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
    artist: "Vincent van Gogh",
    citation: "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
    image: "/images/Fishing-Boats-on-the-Beach-oil-canvas-1888.jpg",
    id: "3",
    collection: "default"
  }

]


function getRandomInt(max: number) {
  return Math.floor(Math.random() * max);
}

function sendMessage(rating: number, user: string | null, work: Work | undefined) {
  if (user === null) {
    console.log("no user; not logging")
    return;
  }
  if (work === undefined ) {
    console.log("undefined work; returning")
    return;
  }
  fetch("https://us-central1-preference-finder.cloudfunctions.net/api/api/v1/works/rate", {
    // fetch('/preference-finder/us-central1/api/api/v1/works/rate', {
    method: 'POST',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      "user": user,
      "rating": rating,
      "work": work
    })
  });
}

function getConverter(name: string) {

  return {
    toFirestore: (collection: any) => {
      console.log(collection)
      return { "name": "foo" }
    },
    fromFirestore: (snapshot: any, options: any): Work => {
      const data = snapshot.data(options);
      var work: Work = {
        title: data.name,
        artist: data.artist,
        citation: data.citation,
        // TODO convert this go dynamically grab the download URL https://firebase.google.com/docs/storage/web/download-files
        image: "https://firebasestorage.googleapis.com/v0/b/preference-finder.appspot.com/o" + data.imageURL + "?alt=media",
        id: snapshot.id,
        collection: name
      }
      return work
    }
  }
}


async function getCollection(): Promise<Work[]> {
  const collection: Work[] = []
  const collectionName = 'third'
  const works = query(collectionGroup(db, 'works').withConverter(getConverter(collectionName)), where('collections', 'array-contains-any', [collectionName]));
  const querySnapshot = await getDocs(works);
  querySnapshot.forEach((doc) => {
    // console.log("doc: ", doc.data())
    collection.push(doc.data())
  });
  return Promise.resolve(collection)
}



function App() {
  const params = new URLSearchParams(window.location.search)
  const user = params.get('user')
  const [work, setWork] = useState<Work>()
  const [workList, setWorkList] = useState<Work[]>([])

  useEffect(() => {
    console.log("in an effect.  Setting work.  workList.length: ", workList.length)
    let combined = workList;
    while (combined.length < 3) {
      const which = getRandomInt(works.length)
      const workToPush = works[which]
      console.log("adding to stack: ", work?.title)
      combined.push(workToPush)
    }
    setWork(workList[0])
    setWorkList(combined)
    console.log("in an effect.  Setting work...  workList.length: ", workList.length)
  }, [workList])

  useEffect(() => {
    const performAsync = async () => {
      const collection = await getCollection()
      console.log("collection: ", collection)
      works.push(...collection)
    }
    performAsync()
      .catch(console.error)
  }, [])

  return (
    <div className="App">
      <h2 style={{ 'textAlign': "center" }}>Find the Art you <em>Love</em></h2>
      {<Card cardProps={work} cardCallbacks={{ onDislike: dislike, onLike: like }} />}
    </div>
  );

  function dislike() {
    console.log("dislike")
    sendMessage(0, user, work)
    setWorkList(workList.slice(1))
  }

  function like() {
    console.log("like")
    sendMessage(5, user, work)
    setWorkList(workList.slice(1))
  }


}

export default App;
