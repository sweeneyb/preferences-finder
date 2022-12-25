import * as functions from "firebase-functions";
import * as cors from "cors";

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
// Automatically allow cross-origin requests
app.use(cors({origin: "https://preference-finder.web.app"}));
main.use(cors({origin: "https://preference-finder.web.app"}));


export type Work = {
  title: string,
  artist: string,
  image: string,
  citation: string,
  id: string,
  collection: string
}

type Rating = {
  work: Work,
  user: string,
  rating: number
}

function validateNumber(input: string): number {
  return parseInt(input);
}

function validateRating(input: number): number {
  if ( 0 <= input && input <= 5) {
    return input;
  }
  throw new Error("Parameter is not within acceptable ranges");
}

function validateString(input: string, fieldName: string): string {
  const pattern = /^[/\-\w.]*$/;
  if (input.match(pattern)) {
    return input;
  }
  throw new Error("Parameter "+fieldName+" has unexpected characters");
}

// curl -id @request.json -H "Content-Type: application/json" http://localhost:5001/preference-finder/us-central1/api/api/v1/works/rate
app.post("/works/rate/", async (req, res) => {
  try {
    const rating: Rating = {
      rating: validateRating(validateNumber(req.body.rating)),
      user: validateString(req.body.user, "user"),
      work: req.body.work,

    };
    console.log("past validation", JSON.stringify(rating));
    console.log("rating: ", rating.user, rating.work.collection,
        rating.work.id);
    res.sendStatus(200);
  } catch (err) {
    console.log(err);
    res.sendStatus(400);
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
