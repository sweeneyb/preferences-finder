import * as functions from "firebase-functions";

// // Start writing Firebase Functions
// // https://firebase.google.com/docs/functions/typescript
//
export const helloWorld = functions.https.onRequest((request, response) => {
  functions.logger.info("Hello logs!", {structuredData: true});
  response.send("Hello from Firebase!");
});

// import * as admin from 'firebase-admin';
import * as express from "express";
import * as bodyParser from "body-parser";

// initialize firebase inorder to access its services
// admin.initializeApp(functions.config().firebase);

// initialize express server
const app = express();
const main = express();

interface Work {
  artist: string
  name: string,
  image: string,
  citation: string
}

const works: Work[] = [
  {
    artist: "Georges Seurat", name: "A Sunday on La Grande Jatte—1884",
    citation: "https://www.britannica.com/biography/Georges-Seurat/images-videos#/media/1/536352/94613",
    image: "/images/canvas-oil-La-Grande-Jatte-Georges-Seurat-1884.jpg",
  },
  {
    artist: "Vincent van Gogh",
    name: "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
    citation: "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
    image: "/images/Fishing-Boats-on-the-Beach-oil-canvas-1888.jpg",
  },
  {
    artist: "Adams, Ansel", name: "Canyon de Chelly",
    citation: "https://www.britannica.com/biography/Ansel-Adams-American-photographer#/media/1/5091/97348",
    image: "/images/photograph-Canyon-de-Chelly-Ansel-Adams.jpg",
  }];

function getRandomInt(max: number) {
  return Math.floor(Math.random() * max);
}

// curl http://localhost:5001/preference-finder/us-central1/api/api/v1/work/random
app.get("/work/random", async (req, res) =>
  res.send(works[getRandomInt(3)])
  // res.send("hi")
);

// add the path to receive request and
// set json as bodyParser to process the body
main.use("/api/v1", app);
main.use(bodyParser.json());
main.use(bodyParser.urlencoded({extended: false}));

// initialize the database and the collection
// const db = admin.firestore();
// const userCollection = 'users';

// define google cloud function name
export const api = functions.https.onRequest(main);
