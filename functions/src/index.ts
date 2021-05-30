import * as functions from "firebase-functions";

// // Start writing Firebase Functions
// // https://firebase.google.com/docs/functions/typescript
//
export const helloWorld = functions.https.onRequest((request, response) => {
  functions.logger.info("Hello logs!", {structuredData: true});
  response.send("Hello from Firebase!");
});

// import * as admin from 'firebase-admin';
import * as express from 'express';
import * as bodyParser from "body-parser";

//initialize firebase inorder to access its services
// admin.initializeApp(functions.config().firebase);

//initialize express server
const app = express();
const main = express();

//add the path to receive request and set json as bodyParser to process the body 
main.use('/api/v1', app);
main.use(bodyParser.json());
main.use(bodyParser.urlencoded({ extended: false }));

//initialize the database and the collection 
// const db = admin.firestore();
// const userCollection = 'users';

//define google cloud function name
export const api = functions.https.onRequest(main);

interface Work {
  name: String,
  link: String
}

var works : Work[] = [
  {name: "name1", link: "https://www.gooogle.com"}, 
  {name: "name2", link: "https://www.google.com"},
  {name: "name3", link: "https://www.google.com"} ]

function getRandomInt(max: number) {
    return Math.floor(Math.random() * max);
}  

// curl http://localhost:5001/preference-finder/us-central1/api/api/v1/work/random
app.get("/work/random", async (req, res) =>
  res.send(works[getRandomInt(3)])
)