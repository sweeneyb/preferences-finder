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

// curl http://localhost:5001/preference-finder/us-central1/api/api/v1/works/random
app.get("/works/random", async (req, res) =>
  res.send(works[getRandomInt(3)])
  // res.send("hi")
);

app.get("/works/all", async (req, res) =>
  res.send(works)
  // res.send("hi")
);

type Rating = {
  image: string,
  user: string,
  rating: number
}

function validateNumber(input: string): number {
  return parseInt(input)
}

function validateRating(input: number): number {
  if ( 0 <= input && input <= 5) {
    return input
  }
  throw 'Parameter is not within acceptable ranges'
}

function validateString (input: string, fieldName: string): string {
  const pattern = /^[\/\-\w\.]*$/
  if( input.match(pattern)) {
    return input
  }
  throw 'Parameter '+fieldName+' has unexpected characters'
}

// curl -id @request.json -H "Content-Type: application/json" http://localhost:5001/preference-finder/us-central1/api/api/v1/works/rate
app.post("/works/rate/", async(req, res) => {
  try {
    var rating: Rating = { 
      rating : validateRating(validateNumber(req.body.rating)),
      user: validateString(req.body.user, "user"),
      image: validateString( req.body.image, "image")
    
    }
    console.log("past validation", rating)
    res.sendStatus(200)
  } catch (err) {
    console.log(err)
    res.sendStatus(400)
  }
});

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
